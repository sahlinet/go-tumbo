package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestServer(t *testing.T) {

	tests := []struct {
		url                string
		method             string
		expectedHTTPStatus int
		expectedMessage    string
	}{
		{
			url:                "/huhu/start",
			method:             "POST",
			expectedHTTPStatus: http.StatusOK,
			expectedMessage:    `{"Message":"Worker huhu started"}`,
		},
		{
			url:                "/huhu/start",
			method:             "GET",
			expectedHTTPStatus: http.StatusMethodNotAllowed,
			expectedMessage:    ``,
		},
	}

	for _, tt := range tests {

		w := httptest.NewRecorder()
		r := mux.NewRouter()
		Worker().AddRoute(r)
		r.ServeHTTP(w, httptest.NewRequest(tt.method, tt.url, nil))

		if w.Code != tt.expectedHTTPStatus {
			t.Error("Did not get expected HTTP status code, got", w.Code)
		}

		if strings.TrimRight(w.Body.String(), "\n") != tt.expectedMessage {
			t.Error("Did not get expected message, got", w.Body.String())
		}

	}
}
