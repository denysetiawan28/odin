package util

import (
	"net/http"
	"odin/src/constanta"
	"odin/src/payload/response"
)

func CreateBadRequestResponse(dt constanta.ErrMsg, data interface{}) (code int, resp response.GlobalResponse) {
	code = http.StatusBadRequest
	resp = response.GlobalResponse{
		Status:  dt.Code,
		Message: dt.Msg,
		Data:    data,
	}
	return
}

func CreateFailedResponse(dt constanta.ErrMsg, data interface{}) (code int, resp response.GlobalResponse) {
	code = http.StatusInternalServerError
	resp = response.GlobalResponse{
		Status:  dt.Code,
		Message: dt.Msg,
		Data:    data,
	}
	return
}

func CreateSuccessResponse(data interface{}) (code int, resp response.GlobalResponse) {
	code = http.StatusOK
	resp = response.GlobalResponse{
		Status:  constanta.Success.Code,
		Message: constanta.Success.Msg,
		Data:    data,
	}
	return
}
