package client

import (
	"errors"
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code int
	Err  error
}

func HTTPErrorFromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr
	} else {
		return &HTTPError{http.StatusInternalServerError, err}
	}
}

func IsNotFound(err error) bool {
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return httpErr.Code == http.StatusNotFound
	}
	return false
}

func HTTPErrorf(code int, format string, a ...interface{}) *HTTPError {
	return &HTTPError{code, fmt.Errorf(format, a...)}
}

func (err *HTTPError) Error() string {
	s := fmt.Sprintf("%v %v", err.Code, http.StatusText(err.Code))
	if err.Err != nil {
		return fmt.Sprintf("%v: %v", s, err.Err)
	} else {
		return s
	}
}

func (err *HTTPError) Unwrap() error {
	return err.Err
}
