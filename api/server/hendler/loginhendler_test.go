package hendler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHendler(t *testing.T) {
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
			name:       "reg user",
			method:     http.MethodPost,
			path:       "/login/",
			body:       []byte(`{"login":"login23","password":"password","role":"admin"}`),
			want:       ``,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			bodyReader := bytes.NewReader(tc.body)
			request := httptest.NewRequest(tc.method, tc.path, bodyReader)
			responseRecorder := httptest.NewRecorder()

			serever.LoginHendler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

// func testService() *Service {

// 	log := log.NewLogger(&config.ConfigLogger{
// 		Level:       zapcore.InfoLevel,
// 		LogEncoding: "console",
// 	}, "./")

// 	connection := sql.New(&config.ConfigDB{
// 		NameDB:   "anapa_db",
// 		Host:     "127.0.0.1",
// 		Port:     "5432",
// 		User:     "postgres",
// 		Password: "123456",
// 	}, log)

// 	memstorage := memory.New(connection.StorageDb)
// 	calendar := app.New(log, memstorage, "secret")

// 	return &Service{
// 		Logger: *log,
// 		app:    *calendar,
// 	}
// }
