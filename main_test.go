package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
)

func TestAdd(t *testing.T) {
	total := Add(1, 3)
	assert.NotNil(t, total, "The `total` should not be nil")
	assert.Equal(t, 4, total, "Expecting 4")
}

func TestSubstract(t *testing.T) {
	result := Substract(3, 2)
	assert.NotNil(t, result, "The `result` should not be nil")
	assert.Equal(t, 1, result, "Expecting 1")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Response expected")
	assert.Equal(t, "Hello world", response.Body.String(), "Incorrect response body")
}

