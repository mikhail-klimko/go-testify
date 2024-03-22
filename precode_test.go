package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 5
	target := fmt.Sprintf("/cafe?city=moscow&count=%d", totalCount)
	req := httptest.NewRequest("GET", target, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, strings.Join(cafeList["moscow"], ","), responseRecorder.Body.String())
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	target := "/cafe?city=minsk&count=1"
	errorMsg := "wrong city value"
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, errorMsg, responseRecorder.Body.String())
}

func TestMainHandlerBaseRequest(t *testing.T) {
	target := "/cafe?city=moscow&count=1"
	answer := []string{"Мир кофе"}
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
	assert.Len(t, answer, len(strings.Split(responseRecorder.Body.String(), ",")))
}
