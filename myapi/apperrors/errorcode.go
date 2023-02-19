package apperrors

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
