package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {

	trace := errors.New("cassandra connection error")

	err := NewInternalServerError("internal server error", trace)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "internal server error", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Trace)
	assert.EqualValues(t, 1, len(err.Trace))
	assert.Same(t, trace, err.Trace[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("it is a bad request")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "it is a bad request", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewFoundError(t *testing.T) {
	err := NewNotFoundError("it was not found")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "it was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewError(t *testing.T) {
	err := NewError("new error")
	assert.NotNil(t, err)
	assert.EqualValues(t, "new error", err.Error())
}
