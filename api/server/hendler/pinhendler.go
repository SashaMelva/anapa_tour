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

func (s *Service) PinHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if req.URL.Path == "/pin/" {
		switch req.Method {
		case http.MethodPost:
			s.ceratePin(w, req, ctx)
		// case http.MethodDelete:
		// 	s.app.DeletePin(w, req)
		case http.MethodPut:
			s.updatePin(w, req, ctx)
		default:
			s.getAllPins(w, req, ctx)
		}
	} else {
		args := req.URL.Query()
		id := args.Get("id")

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
		}
	}

}

func (s *Service) ceratePin(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var pin loyaltymodel.Pin

	body, err := io.ReadAll(req.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &pin)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	id, err := s.app.CeratePin(&pin)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", id)))
}

func (s *Service) deletePinByID(w http.ResponseWriter, id int) {
	s.Logger.Info("delet event by id %v", id)

	err := s.app.DeletePinByID(uint32(id))

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func (s *Service) getAllPins(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	s.Logger.Info("handling get all events at %s\n", req.URL.Path)

	allPins, err := s.app.GetAllPins(ctx)

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

func (s *Service) updatePin(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	s.Logger.Info("edit event at %v\n", req.URL.Path)
	fmt.Println("qweqewqe")

	pin := loyaltymodel.Pin{}
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

	errEdit := s.app.EditPin(ctx, &pin)

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
