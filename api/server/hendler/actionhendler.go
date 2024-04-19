package hendler

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"time"
// )

// func (s *Service) ActionHendler(w http.ResponseWriter, req *http.Request) {
// 	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
// 	defer cancel()
// 	if req.URL.Path == "/action/" {
// 		switch req.Method {
// 		case http.MethodPost:
// 			// s.cerateAction(w, req, ctx)
// 			// case http.MethodDelete:
// 			// 	s.app.DeletePin(w, req)
// 			// case http.MethodPut:
// 			// 	s.updateAction(w, req, ctx)
// 			// default:
// 			// 	s.getAllAction(w, req, ctx)
// 		}
// 	} else {
// 		args := req.URL.Query()
// 		id := args.Get("id")

// 		if len(id) > 0 {
// 			s.Logger.Info("id event " + id)
// 			intId, err := strconv.Atoi(id)
// 			if err != nil {
// 				s.Logger.Error(fmt.Sprintf("is not valid if event id, got %v", id))
// 				http.Error(w, err.Error(), http.StatusBadRequest)
// 				return
// 			}

// 			switch req.Method {
// 			case http.MethodDelete:
// 				s.deletePinByID(w, intId)
// 			// case http.MethodGet:
// 			// 	s.getEventHandlerById(w, ctx, intId)
// 			default:
// 				s.Logger.Error(fmt.Sprintf("expect method GET or DELETE at /event?=<id>, got %v", req.Method))
// 				return
// 			}
// 		}
// 	}

// }

// // func (s *Service) cerateAction(w http.ResponseWriter, req *http.Request, ctx context.Context) {
// // 	var pin loyaltymodel.Pin

// // 	body, err := io.ReadAll(req.Body)

// // 	if err != nil {
// // 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// // 		w.WriteHeader(http.StatusBadRequest)
// // 		w.Write([]byte(err.Error()))
// // 	} else {
// // 		err = json.Unmarshal(body, &pin)
// // 		if err != nil {
// // 			w.Header().Set("Content-Type", "application/json; charset=utf-8")
// // 			w.WriteHeader(http.StatusInternalServerError)
// // 			w.Write([]byte(err.Error()))
// // 		}
// // 	}

// // 	id, err := s.app.CerateAction(&pin)

// // 	if err != nil {
// // 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// // 		w.WriteHeader(http.StatusInternalServerError)
// // 		w.Write([]byte(err.Error()))
// // 	}

// // 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// // 	w.WriteHeader(http.StatusOK)
// // 	w.Write([]byte(fmt.Sprintf("%v", id)))
// // }
