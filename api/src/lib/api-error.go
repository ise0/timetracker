package lib

type apiErrCode int

const (
	ApiErrCode400 apiErrCode = 400
	ApiErrCode500 apiErrCode = 500
)

type ApiError struct {
	Code apiErrCode
	s    string
}

func (e ApiError) Error() string {
	return e.s
}

func ApiError400(text string) ApiError {
	if text == "" {
		text = "Bad request"
	}
	return ApiError{
		s:    text,
		Code: ApiErrCode400,
	}
}

func ApiError500(text string) ApiError {
	if text == "" {
		text = "Internal error"
	}
	return ApiError{
		s:    text,
		Code: ApiErrCode500,
	}
}
