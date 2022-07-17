package server

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"time"
	"voyager2/src/handler"
	"voyager2/src/server/container"
	"voyager2/src/server/middleware"
	"voyager2/src/server/router"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func StartHttpServer(container *container.DefaultContainer) {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	validate := validator.New()
	validate.RegisterValidation("ISO8601date", IsISO8601Date)

	e.Validator = &CustomValidator{validator: validate}

	middleware.SetupMiddleware(e)
	router.InitializeRouter(e, handler.InitializeHandler(container))

	port := fmt.Sprintf("%s%s", ":", container.Config.Server.Port)
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func IsISO8601Date(fl validator.FieldLevel) bool {

	if fl.Field().String() == "" {
		return true
	}

	ISO8601DateRegexString := "^\\d{4}(-\\d\\d(-\\d\\d(T\\d\\d:\\d\\d(:\\d\\d)?(\\.\\d+)?(([+-]\\d\\d:\\d\\d)|Z)?)?)?)?$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}
