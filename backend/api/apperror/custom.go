package apperror

import "fmt"

type CustomError struct {
	Err        error  `json:"-"`
	Msg        string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewError(err error, msg string) CustomError {
	return CustomError{
		Err:        err,
		Msg:        msg,
		StatusCode: StatusCode[msg],
	}
}

func (c CustomError) Error() string {
	result := fmt.Sprintf("Error: %v \nMessage: %v \nStatus Code: %v\n", c.Err, c.Msg, c.StatusCode)
	return result
}
