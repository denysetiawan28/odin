package constanta

import "errors"

//Application Error Code
var (
	ErrorCustomerNotFound = ErrMsg{Code: "ODIN_001", Msg: "Customer not found"}
	ErrorBadRequest       = ErrMsg{Code: "ODIN_002", Msg: "Bad request"}

	ErrorGeneral = ErrMsg{Code: "ODIN_999", Msg: "General Error"}
)

var (
	Success = ErrMsg{Code: "000", Msg: "Success"}
)

var (
	DbFailedToExecuteQuery = errors.New("failed to execute query from db")
	DbFailedToInsertData   = errors.New("failed to insert data to db")
	DbFailedToUpdateData   = errors.New("failed to update data to db")
	DbFailedToDeleteData   = errors.New("failed to delete data from db")
)

var (
	DateFormatInvalid = errors.New("format date invalid")
	DateEmpty         = errors.New("date field empty")
	Date1AfterDate2   = errors.New("date1 after date2")
	Date2BeforeDate1  = errors.New("date2 before date1")
)

var (
	SysDatabaseFailedInit   = "failed to init database"
	SysConfigFailedRead     = "failed to read config"
	SysConfigUnmarshall     = "failed to unmarshall config"
	SysRepositoryFailedInit = "failed to init repository"
	SysServiceFailedInit    = "failed to init service"
)

type ErrMsg struct {
	Code string
	Msg  string
}
