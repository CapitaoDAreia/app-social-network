package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type ResponsesMock struct {
	mock.Mock
}

func NewResponsesMock() *ResponsesMock {
	return &ResponsesMock{}
}

func (responses *ResponsesMock) FormatResponseToJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	responses.Called(w, statusCode, data)
}

func (responses *ResponsesMock) FormatResponseToCustomError(w http.ResponseWriter, statusCode int, err error) {
	responses.Called(w, statusCode, err)
}
