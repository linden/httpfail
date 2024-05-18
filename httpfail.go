package httpfail

import (
	"errors"
	"net/http"
)

type Reason struct {
	message string
	code    int
}

func NewReason(message string, code int) *Reason {
	return &Reason{
		message: message,
		code:    code,
	}
}

type Failure struct {
	err    error
	reason *Reason
}

func (f *Failure) Reason() *Reason {
	return f.reason
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
	http.Error(w, f.reason.message, f.reason.code)
}

func New(err error, reason *Reason) *Failure {
	return &Failure{
		err:    err,
		reason: reason,
	}
}
