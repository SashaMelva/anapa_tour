package hendler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestActionHendler(t *testing.T) {
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
		// 	name:       "Creat action",
		// 	method:     http.MethodPost,
		// 	path:       "/action/",
		// 	body:       []byte(`{"id":0,"login":"login","password":"password","role":"admin"}`),
		// 	want:       ``,
		// 	statusCode: http.StatusOK,
		// },
		{
			name:       "reg user",
			method:     http.MethodGet,
			path:       "/action?orgId=1",
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

			serever.ActionHendler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
