package types

type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *BizError) Error() string {
	return e.Message
}

func NewBizError(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}
