package hendler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOrganizationHendler(t *testing.T) {
	serever := testService()
	testCase := []struct {
		name       string
		method     string
		path       string
		body       []byte
		want       string
		statusCode int
	}{
		{
			name:       "get by id admin",
			method:     http.MethodGet,
			path:       "/organization?adminId=1",
			body:       []byte(``),
			want:       ``,
			statusCode: http.StatusOK,
		},
		{
			name:       "get by id admin",
			method:     http.MethodGet,
			path:       "/organization?category=1",
			body:       []byte(``),
			want:       ``,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			bodyReader := bytes.NewReader(tc.body)
			request := httptest.NewRequest(tc.method, tc.path, bodyReader)
			responseRecorder := httptest.NewRecorder()

			serever.OrganizationHendler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
