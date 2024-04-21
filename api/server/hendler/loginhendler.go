package hendler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	autenficationmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/autenfication"
)

func (s *Service) LoginHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if req.Method == http.MethodPost {
		s.loginAccount(w, req, ctx)
	}

}

func (s *Service) loginAccount(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var account autenficationmodel.Account

	body, err := io.ReadAll(req.Body)

	if err != nil {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &account)
		if err != nil {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	reqAuth, err := s.app.LoginAccout(&account)

	if err != nil {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	js, err := json.Marshal(reqAuth)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
