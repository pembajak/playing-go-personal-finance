package http

import (
	"encoding/json"
	"net/http"

	"github.com/pembajak/personal-finance/internal/pkg/errors"
)

type Rest interface {
	GetHTTPMethod() string
	GetPattern() string
	GetHandler() Handler
}

type rest struct {
	HTTPMethod string
	Pattern    string
	Handler    Handler
}

func NewRest(httpMethod string, pattern string, handler Handler) Rest {
	return &rest{
		HTTPMethod: httpMethod,
		Pattern:    pattern,
		Handler:    handler,
	}
}

type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := handler(w, r)
	if err != nil {
		WriteError(w, r, err, 500)
		return
	}

	WriteSuccess(w, r, data)
}

func WriteError(w http.ResponseWriter, r *http.Request, err error, errorCode int) {
	response := NewFailed(
		"0001",
		NewResponseDesc(err.Error()),
		errorCode,
		NewMeta("1.0", "healthy", "development"),
	)
	if _, ok := err.(*errors.InternalError); ok {
		response = NewFailed(
			"0001",
			NewResponseDesc(err.Error()),
			http.StatusInternalServerError,
			NewMeta("1.0", "healthy", "development"),
		)
	}

	if _, ok := err.(*errors.Unauthorized); ok {
		response = NewFailed(
			"0002",
			NewResponseDesc(err.Error()),
			http.StatusUnauthorized,
			NewMeta("1.0", "healthy", "development"),
		)
	}

	if _, ok := err.(*errors.BadRequest); ok {
		response = NewFailed(
			"0003",
			NewResponseDesc(err.Error()),
			http.StatusBadRequest,
			NewMeta("1.0", "healthy", "development"),
		)
	}

	compose(w, r, response, response.HttpStatus)
}

func WriteSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	response := NewSuccess(
		"0000",
		NewResponseDesc("-"),
		data,
		http.StatusOK,
		NewMeta("1.0", "healthy", "development"),
	)

	compose(w, r, response, response.HttpStatus)
}

func compose(w http.ResponseWriter, r *http.Request, response interface{}, httpStatus int) {
	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to unmarshal"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(res)
}

func (rest *rest) GetHandler() Handler {
	return rest.Handler
}

func (rest *rest) GetHTTPMethod() string {
	return rest.HTTPMethod
}

func (rest *rest) GetPattern() string {
	return rest.Pattern
}
