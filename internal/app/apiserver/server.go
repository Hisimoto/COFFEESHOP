package apiserver

import (
	"encoding/json"
	"github.com/Hisimoto/COFFEESHOP/internal/app/model"
	"github.com/Hisimoto/COFFEESHOP/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/users", s.handleUsersGet()).Methods("GET")
	s.router.HandleFunc("/buycoffe", s.buyCoffee()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email          string `json:"email"`
		MembershipType int    `json:"membershipType"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:          req.Email,
			MembershipType: req.MembershipType,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleUsersGet() http.HandlerFunc {
	type request struct {
		ID             int    `json:"ID"`
		Email          string `json:"email"`
		MembershipType int    `json:"membershipType"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			ID:             req.ID,
			Email:          req.Email,
			MembershipType: req.MembershipType,
		}
		u, err := s.store.User().FindByEmail(u.Email)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) buyCoffee() http.HandlerFunc {
	type request struct {
		ID         int `json:"ID"`
		CoffeeType int `json:"coffeeType"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.Order{
			ID:         req.ID,
			CoffeeType: req.CoffeeType,
		}

		MaxCoffeeCnt, err := s.store.User().GetNumbersOfCoffeeByTypeAndUserId(u.ID, u.CoffeeType)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		ActualCoffeeCnt, err := s.store.Order().CountCoffeByTypeAndUserId(u.ID, u.CoffeeType)
		//s.respond(w, r, http.StatusCreated, actualCoffeeCnt) // tmp
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if ActualCoffeeCnt < MaxCoffeeCnt {

			//insert 1 order, response 200 ok
			if err = s.store.Order().CreateOrder(u, u.CoffeeType); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			s.respondCode(w, r, http.StatusOK)
		} else {
			var Timeleft string
			Timeleft, err = s.store.Order().CheckRemainingTime(u.ID, u.CoffeeType)
			if err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusTooManyRequests, Timeleft) // insert how much time should wait for next coffe
		}

	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
func (s *server) respondCode(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
}
