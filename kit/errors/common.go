package errors

type CustomError struct {
	Msg        string
	StatusCode int
}

func NewCustomError(msg string, statusCode int) CustomError {
	return CustomError{Msg: msg, StatusCode: statusCode}
}

func (err CustomError) Error() string {
	return err.Msg
}
