package data

const (
	CodeSuccess = 200
	CodeFailed  = 500
)

type BaseResponse struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func ResponseSuccess(data interface{}) BaseResponse {
	return BaseResponse{
		Code:   CodeSuccess,
		ErrMsg: "ok",
		Data:   data,
	}
}

func ResponseFailed(errMsg string) BaseResponse {
	return BaseResponse{
		Code:   CodeFailed,
		ErrMsg: errMsg,
	}
}
