package models

import (
	"fmt"
	"net/http"
)

type Err interface {
	Code() int
	Error() string
}

type err struct {
	StatusCode int
	Text       string
}

func (f err) Code() int {
	return f.StatusCode
}

func (f err) Error() string {
	return f.Text
}

func InternalError(text string) Err {
	return err{
		StatusCode: http.StatusInternalServerError,
		Text:       text,
	}
}

func Unauthorized(text string) Err {
	return err{
		StatusCode: http.StatusUnauthorized,
		Text:       text,
	}
}

func BadRequest(text string) Err {
	return err{
		StatusCode: http.StatusBadRequest,
		Text:       text,
	}
}

func Forbidden(text string) Err {
	return err{
		StatusCode: http.StatusForbidden,
		Text:       text,
	}
}

func ErrorWithSuccess(text string) Err {
	return err{
		StatusCode: http.StatusOK,
		Text:       text,
	}
}

func Concat(errs ...Err) Err {
	var (
		oErr err
		ln   int
	)
	for _, err := range errs {
		if err == nil {
			continue
		}
		oErr.StatusCode = err.Code()
		ln = len(oErr.Text)
		if ln == 0 {
			oErr.Text = err.Error()
		} else {
			oErr.Text = fmt.Sprintf("%s\n%s", oErr.Text, err.Error())
		}
	}
	if ln == 0 {
		return nil
	}
	return oErr
}
