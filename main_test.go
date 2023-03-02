package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	// 3rd party modules
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

//
// Writing unit tests for all the silly handlers
//

// test GET method
func TestSillyGet(t *testing.T) {
	router := gin.Default()
	router.GET("/", SillyGet)

	mockResponse := `{"object":"Hi! I'm a silly object"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, mockResponse, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

// test POST method
func TestSillyPost(t *testing.T) {
	router := gin.Default()
	router.POST("/", SillyPost)

	bodyReader := bytes.NewReader([]byte(`"object":"Hi! I'm a silly object"}`))
	mockResponse := `{"message":"Silly object created!"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, mockResponse, w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}

// test PUT method
func TestSillyPut(t *testing.T) {
	router := gin.Default()
	router.PUT("/", SillyPut)

	bodyReader := bytes.NewReader([]byte(`{"object":"Hi! I'm a silly object"}`))
	mockResponse := `{"message":"Silly object modified!"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, mockResponse, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

// test DELETE method
func TestSillyDelete(t *testing.T) {
	router := gin.Default()
	router.DELETE("/", SillyDelete)

	mockResponse := `{"message":"Silly object deleted!"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, mockResponse, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}