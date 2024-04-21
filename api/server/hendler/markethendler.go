package hendler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	marcketmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/marcket"
)

func (s *Service) MarketHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if req.Method == http.MethodPost {
		s.getAllProducts(w, req, ctx)
	} else if req.Method == http.MethodPost {
		s.byProducts(w, req, ctx)
	}

}

func (s *Service) getAllProducts(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	allProductions, err := s.app.GetAiiProduction()

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(allProductions)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
func (s *Service) byProducts(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var prodBy marcketmodel.ByProduction

	body, err := io.ReadAll(req.Body)
	s.Logger.Info("parse")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		err = json.Unmarshal(body, &prodBy)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	err = s.app.ByProductions(prodBy)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
