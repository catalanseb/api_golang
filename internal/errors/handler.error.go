package errors

func ErrResp(code int, status string, message string) ErrorResponse {
	e := ErrorResponse{
		Code:    code,
		Status:  status,
		Message: message,
	}

	return e
}
