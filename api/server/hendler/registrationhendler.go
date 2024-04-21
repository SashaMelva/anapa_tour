package hendler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	autenficationmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/autenfication"
)

func (s *Service) RegistrationHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if req.Method == http.MethodPost {
		s.regAccount(w, req, ctx)
	}

}

func (s *Service) regAccount(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var account autenficationmodel.Account

	body, err := io.ReadAll(req.Body)
	s.Logger.Info("parse")
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		err = json.Unmarshal(body, &account)
		if err != nil {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
	s.Logger.Info(account.Login)
	err = s.app.CheckUniqueLogin(account.Login)

	if err != nil {
		// s.Logger.Fatal(err.Error())
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	s.Logger.Info("Плдбзователя с таким логин не существует")
	s.Logger.Info("Старт регистрации")
	id, errNew := s.app.RegisterAccount(&account)

	if errNew != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.Logger.Info("Ура регистрации")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id : %v", id)))
}
