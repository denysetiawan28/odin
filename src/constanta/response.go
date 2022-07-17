package constanta

import "errors"

var (
	ErrorUserNotFound         = ErrMsg{Code: "VY2_001", Msg: "User not found"}
	ErrorDuplicateUser        = ErrMsg{Code: "VY2_002", Msg: "Username Already Taken"}
	ErrorDuplicateEmailUser   = ErrMsg{Code: "VY2_003", Msg: "Email Already Taken"}
	ErrorHashingPassword      = ErrMsg{Code: "VY2_004", Msg: "Password Hashing Failed"}
	ErrorInvalidFormatDate    = ErrMsg{Code: "VY2_005", Msg: "Format Date Must Be DD-MM-YYYY"}
	ErrorFailedToSaveUser     = ErrMsg{Code: "VY2_006", Msg: "Failed to Create User"}
	ErrorFailedToUpdateUser   = ErrMsg{Code: "VY2_007", Msg: "Failed to Update User"}
	ErrorFailedAuthenticate   = ErrMsg{Code: "VY2_008", Msg: "Failed to Authenticate User, Check Username / Password"}
	ErrorInvalidAccessToken   = ErrMsg{Code: "VY2_009", Msg: "Invalid Access Token"}
	ErrorExpiredAccessToken   = ErrMsg{Code: "VY2_010", Msg: "Expired Access Token"}
	ErrorChangePasswordFailed = ErrMsg{Code: "VY2_010", Msg: "Failed to change Password"}

	ErrorFailedToLovHeader        = ErrMsg{Code: "VY2_014", Msg: "Failed to Create LOV Header"}
	ErrorFailedToLovDetail        = ErrMsg{Code: "VY2_015", Msg: "Failed to Create LOV Details"}
	ErrorLovHeaderIdNotFound      = ErrMsg{Code: "VY2_016", Msg: "LOV Header Id Not Found"}
	ErrorLovHeaderNotFound        = ErrMsg{Code: "VY2_017", Msg: "Lov Header Is Empty"}
	ErrorLovDetailRequestNotFound = ErrMsg{Code: "VY2_018", Msg: "Lov Detail Request Is Empty"}
	ErrorLovDetailNotFound        = ErrMsg{Code: "VY2_019", Msg: "Lov Detail Data Not Found"}
	ErrorLovDetailFailedUpdate    = ErrMsg{Code: "VY2_020", Msg: "Failed to Update Lov Details"}

	ErrorRoleDataNotFound       = ErrMsg{Code: "VY2_021", Msg: "Role Data Not Found"}
	ErrorFailedToCreateRole     = ErrMsg{Code: "VY2_022", Msg: "Failed to Create Role"}
	ErrorFailedToUpdateRole     = ErrMsg{Code: "VY2_023", Msg: "Failed to Update Role"}
	ErrorPermissionDataNotFound = ErrMsg{Code: "VY2_024", Msg: "Permission Data Not Found"}
	ErrorAddPrivFailedToAdd     = ErrMsg{Code: "VY2_025", Msg: "Failed to Mapping Additional Privilege to User"}
	ErrorAddPrivFailedToRemove  = ErrMsg{Code: "VY2_026", Msg: "Failed to Remove Additional Privilege From User"}

	ErrorModuleDataNotFound       = ErrMsg{Code: "VY2_031", Msg: "Module Data Not Found"}
	ErrorMenuDataNotFound         = ErrMsg{Code: "VY2_032", Msg: "Menu Data Not Found"}
	ErrorFailedToCreateModule     = ErrMsg{Code: "VY2_033", Msg: "Failed to Create Module"}
	ErrorFailedToCreateMenu       = ErrMsg{Code: "VY2_034", Msg: "Failed to Create Menu"}
	ErrorFailedToUpdateModule     = ErrMsg{Code: "VY2_035", Msg: "Failed to Update Module"}
	ErrorFailedToUpdateMenu       = ErrMsg{Code: "VY2_036", Msg: "Failed to Update Menu"}
	ErrorMenuPartialDataNotFound  = ErrMsg{Code: "VY2_037", Msg: "Some Menu Data Not Found"}
	ErrorMenuParentIdBelongTo     = ErrMsg{Code: "VY2_038", Msg: "Parent belong to others module"}
	ErrorRoleMenuStateNotValid    = ErrMsg{Code: "VY2_039", Msg: "State not valid for creating new role"}
	ErrorMenuIdDuplicate          = ErrMsg{Code: "VY2_040", Msg: "Menu id duplicated on request"}
	ErrorMenuNotFoundAddPrivilege = ErrMsg{Code: "VY2_041", Msg: "Some menu not found for additional privileges"}
	ErrorPermissionNotFound       = ErrMsg{Code: "VY2_042", Msg: "Permission not found"}
	ErrorFailedToRemovePermission = ErrMsg{Code: "VY2_043", Msg: "Failed to remove permission"}
	ErrorFailedToAddPermission    = ErrMsg{Code: "VY2_043", Msg: "Failed to add permission"}

	ErrorFailedExecuteQuery = ErrMsg{Code: "VY2_050", Msg: "Failed to Execute Query"}

	ErrorBindingRequest           = ErrMsg{Code: "VY2_090", Msg: "Invalid Binding"}
	ErrorInvalidRequest           = ErrMsg{Code: "VY2_091", Msg: "Invalid Request"}
	ErrorInvalidDate              = ErrMsg{Code: "VY2_092", Msg: "Invalid Date"}
	ErrorInvalidDateFormat        = ErrMsg{Code: "VY2_093", Msg: "Format Date Is Not ISO-8601"}
	ErrorComparisonBetween2Date   = ErrMsg{Code: "VY2_094", Msg: "Invalid Compare Between 2 Dates"}
	ErrorInvalidDefaultStatus     = ErrMsg{Code: "VY2_095", Msg: "Invalid Default Status (active / inactive)"}
	ErrorNewAndOldSamePassword    = ErrMsg{Code: "VY2_100", Msg: "Old Password and New Password is same, please change new password"}
	ErrorNewAndReconfirmDifferent = ErrMsg{Code: "VY2_101", Msg: "New Password And Reconfirm Is different, please check your password"}
	ErrorInvalidOldPassword       = ErrMsg{Code: "VY2_102", Msg: "Invalid old password"}

	ErrorGeneral = ErrMsg{Code: "VY3_999", Msg: "General Error"}
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
