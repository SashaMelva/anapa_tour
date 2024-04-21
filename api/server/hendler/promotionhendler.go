package hendler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"
)

func (s *Service) PromotionHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.URL.Path == "/promotion/" {
		switch req.Method {
		case http.MethodPost:
			s.ceratePromotion(w, req, ctx)
		case http.MethodPut:
			s.updatePromotion(w, req, ctx)
		}
	} else {
		args := req.URL.Query()
		id := args.Get("id")
		idCategory := args.Get("idCategory")

		if len(id) > 0 {
			s.Logger.Info("id event " + id)
			intId, err := strconv.Atoi(id)
			if err != nil {
				s.Logger.Error(fmt.Sprintf("is not valid if event id, got %v", id))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			switch req.Method {
			case http.MethodDelete:
				s.deletePromotionByID(w, intId)
			case http.MethodGet:
				s.getPromotionByID(intId, w, ctx)
			default:
				s.Logger.Error(fmt.Sprintf("expect method GET or DELETE at /event?=<id>, got %v", req.Method))
				return
			}
			return
		}

		if idCategory != "" {
			s.Logger.Info("id event " + id)
			intIdCategory, err := strconv.Atoi(id)
			if err != nil {
				s.Logger.Error(fmt.Sprintf("is not valid if event id, got %v", id))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if req.Method == http.MethodGet {
				s.getPromotionByOrganizationID(intIdCategory, w, ctx)
			}
			return
		}
	}
}

func (s *Service) ceratePromotion(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var promotion loyaltymodel.Promotion

	body, err := io.ReadAll(req.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &promotion)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	id, err := s.app.CeratePromotion(&promotion)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", id)))
}

func (s *Service) deletePromotionByID(w http.ResponseWriter, id int) {
	s.Logger.Info("delet event by id %v", id)

	err := s.app.DeletePromotionByID(uint32(id))

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func (s *Service) getPromotionByOrganizationID(orgId int, w http.ResponseWriter, ctx context.Context) {
	allPins, err := s.app.GetPromotionByOrganizationID(orgId)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(allPins)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (s *Service) getPromotionByID(id int, w http.ResponseWriter, ctx context.Context) {
	allPins, err := s.app.GetPromotionByID(id)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(allPins)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (s *Service) updatePromotion(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	prompt := loyaltymodel.Promotion{}
	body, err := io.ReadAll(req.Body)

	if err != nil {
		s.Logger.Error(w, "err %q\n", err.Error())
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &prompt)
		if err != nil {
			s.Logger.Error(w, "can't unmarshal: ", err.Error())
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("can't unmarshal: " + err.Error()))
		}
	}

	if prompt.Id == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id event empty"))
		return
	}

	s.Logger.Info(prompt)

	errEdit := s.app.EditPromotion(ctx, &prompt)

	if errEdit != nil {
		s.Logger.Error(w, errEdit.Error(), http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("dont update"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
