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

func (s *Service) ActionHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if req.URL.Path == "/action/" {
		switch req.Method {
		case http.MethodPost:
			s.cerateAction(w, req, ctx)
		case http.MethodPut:
			s.updateAction(w, req, ctx)
		default:
			s.getAllActions(w, req, ctx)
		}
	} else {
		args := req.URL.Query()
		id := args.Get("id")
		typeAction := args.Get("type")
		active := args.Get("active")

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
				s.deletePinByID(w, intId)
			// case http.MethodGet:
			// 	s.getEventHandlerById(w, ctx, intId)
			default:
				s.Logger.Error(fmt.Sprintf("expect method GET or DELETE at /event?=<id>, got %v", req.Method))
				return
			}
			return
		}
		if typeAction != "" {
			if active != "" {
				flag, _ := strconv.ParseBool(active)
				s.getActiveAndTypeAction(w, flag, typeAction, ctx)
			}
		}
		if active != "" {
			flag, _ := strconv.ParseBool(active)
			s.getActiveAction(w, flag, ctx)
		}
	}

}

func (s *Service) cerateAction(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var action loyaltymodel.Action

	body, err := io.ReadAll(req.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &action)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	id, err := s.app.CerateAction(&action)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", id)))
}

func (s *Service) deleteActionByID(w http.ResponseWriter, id int) {
	s.Logger.Info("delet event by id %v", id)

	err := s.app.DeleteActionByID(uint32(id))

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func (s *Service) getAllActions(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	s.Logger.Info("handling get all events at %s\n", req.URL.Path)

	allPins, err := s.app.GetAllAction()

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

func (s *Service) getActiveAction(w http.ResponseWriter, flag bool, ctx context.Context) {
	allPins, err := s.app.GetActiveAction(flag)

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

func (s *Service) getActiveAndTypeAction(w http.ResponseWriter, flag bool, typeAction string, ctx context.Context) {
	allPins, err := s.app.GetActiveAndTypeAction(flag, typeAction)

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

func (s *Service) updateAction(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	s.Logger.Info("edit event at %v\n", req.URL.Path)
	fmt.Println("qweqewqe")

	pin := loyaltymodel.Action{}
	body, err := io.ReadAll(req.Body)

	if err != nil {
		s.Logger.Error(w, "err %q\n", err.Error())
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &pin)
		if err != nil {
			s.Logger.Error(w, "can't unmarshal: ", err.Error())
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("can't unmarshal: " + err.Error()))
		}
	}

	if pin.Id == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id event empty"))
		return
	}

	s.Logger.Info(pin)

	errEdit := s.app.EditAction(ctx, &pin)

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
