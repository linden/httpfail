package httpfail

import (
	"errors"
	"net/http"
)

type Reason struct {
	Message string
	Code    int
}

func NewReason(message string, code int) *Reason {
	return &Reason{
		Message: message,
		Code:    code,
	}
}

type Failure struct {
	err    error
	Reason *Reason
}

func (f *Failure) Error() string {
	if f.err == nil {
		return ""
	}

	return f.err.Error()
}

func (f *Failure) Is(err error) bool {
	return errors.Is(f.err, err)
}

func (f *Failure) Write(w http.ResponseWriter) {
	http.Error(w, f.Reason.Message, f.Reason.Code)
}

func New(err error, reason *Reason) *Failure {
	return &Failure{
		err:    err,
		Reason: reason,
	}
}
