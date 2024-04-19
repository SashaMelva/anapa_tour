package hendler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPinHendler(t *testing.T) {
	serever := testService()
	testCase := []struct {
		name       string
		method     string
		path       string
		body       []byte
		want       string
		statusCode int
	}{
		// {
		// 	name:       "reg user",
		// 	method:     http.MethodPost,
		// 	path:       "/registration/",
		// 	body:       []byte(`{"id":0,"login":"login","password":"password","role":"admin"}`),
		// 	want:       ``,
		// 	statusCode: http.StatusBadGateway,
		// },
		{
			name:       "add pins",
			method:     http.MethodPost,
			path:       "/pin/",
			body:       []byte(`{"name":"login23","description":"password","coordinat":"admin","activ":"true"}`),
			want:       ``,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			bodyReader := bytes.NewReader(tc.body)
			request := httptest.NewRequest(tc.method, tc.path, bodyReader)
			responseRecorder := httptest.NewRecorder()

			serever.PinHendler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
