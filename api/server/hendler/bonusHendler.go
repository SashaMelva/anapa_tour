package hendler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (s *Service) BonusHendler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if req.URL.Path == "/bonusbalance/" {
		// switch req.Method {
		// // case http.MethodPost:
		// // 	s.changeSumm(w, req, ctx)
		// case http.MethodPut:
		// 	s.changeSumm(w, req, ctx)
		// default:

		// }
	} else {
		args := req.URL.Query()
		id := args.Get("accountId")

		if len(id) > 0 {
			intId, err := strconv.Atoi(id)
			if err != nil {
				s.Logger.Error(fmt.Sprintf("is not valid if event id, got %v", id))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			switch req.Method {
			case http.MethodGet:
				s.getBalanceByAccountId(intId, w, ctx)
			default:
				s.Logger.Error(fmt.Sprintf("expect method GET or DELETE at /event?=<id>, got %v", req.Method))
				return
			}
			return
		}
	}
}

func (s *Service) getBalanceByAccountId(orgId int, w http.ResponseWriter, ctx context.Context) {
	balanse, err := s.app.GetBalanceByAccountId(orgId)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(balanse)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
