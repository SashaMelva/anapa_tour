package hendler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPinsVewHendler(t *testing.T) {
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
			name:       "add pins",
			method:     http.MethodGet,
			path:       "/pins_vew/",
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

			serever.PinsVewHendler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
