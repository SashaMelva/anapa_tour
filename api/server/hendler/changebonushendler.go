package hendler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"
)

func (s *Service) ChangeBonusHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if req.Method == http.MethodPost {
		s.changeBalance(w, req, ctx)
	} else {
		args := req.URL.Query()
		orgId := args.Get("orgCategoryid")

		if len(orgId) > 0 {
			strOrgId, err := strconv.Atoi(orgId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			s.changeBalanceByCategory(strOrgId, w, req, ctx)
		}
	}
}

func (s *Service) changeBalance(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var balance loyaltymodel.ChangeBalance

	body, err := io.ReadAll(req.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &balance)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	err = s.app.ChangeBalance(&balance)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func (s *Service) changeBalanceByCategory(orgId int, w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var balance loyaltymodel.ChangeBalanceForOrganization

	body, err := io.ReadAll(req.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &balance)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	err = s.app.ChangeBalanceByCategory(&balance)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
