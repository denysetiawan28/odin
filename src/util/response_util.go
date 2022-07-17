package util

import (
	"voyager2/src/constanta"
	"voyager2/src/payload/response"
)

func CreateGlobalResponse(dt constanta.ErrMsg, data interface{}) (resp response.GlobalResponse) {
	resp = response.GlobalResponse{
		Status:  dt.Code,
		Message: dt.Msg,
		Data:    data,
	}
	return
}
