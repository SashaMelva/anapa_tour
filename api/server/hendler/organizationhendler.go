package hendler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *Service) OrganizationHendler(w http.ResponseWriter, req *http.Request) {
	if http.MethodGet == req.Method {
		args := req.URL.Query()
		category := args.Get("category")
		adminId := args.Get("adminId")

		if category != "" {
			categoryInt, err := strconv.Atoi(category)

			if err != nil {
				s.Logger.Error(fmt.Sprintf("is not valid if categoryId id, got %v", categoryInt))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			s.getOrganizationByCategory(w, categoryInt)
			return

		}
		if adminId != "" {
			adminIdInt, err := strconv.Atoi(adminId)

			if err != nil {
				s.Logger.Error(fmt.Sprintf("is not valid if adminId id, got %v", adminId))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			s.getOrganizationByAdminId(w, adminIdInt)
			return
		}
	}
}
func (s *Service) getOrganizationByCategory(w http.ResponseWriter, categoryId int) {

	allOrganizations, err := s.app.GetOrganizationByCategory(categoryId)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(allOrganizations)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (s *Service) getOrganizationByAdminId(w http.ResponseWriter, id int) {

	allOrganizations, err := s.app.GetOrganizationByAdminId(id)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(allOrganizations)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
