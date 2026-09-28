package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handler "code-quality-practice/02-mocks"
	"code-quality-practice/02-mocks/mocks"

	"go.uber.org/mock/gomock"
)

func TestHandler(t *testing.T) {
	service := mocks.NewMockGreeter(gomock.NewController(t))
	request := httptest.NewRequest(http.MethodGet, "/?name=Alice", nil)
	service.EXPECT().Greet(request.Context(), "Alice").Return("Hello, Alice!", nil)
	response := httptest.NewRecorder()

	handler.NewHandler(service).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if got := response.Body.String(); got != "Hello, Alice!" {
		t.Errorf("body = %q, want %q", got, "Hello, Alice!")
	}
}
