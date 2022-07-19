package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/honyshyota/l0-wb-test/internal/service"
	"github.com/sirupsen/logrus"
)

func newServer(config *config.Config, service service.Service) *http.Server {
	srv := &http.Server{
		Addr:           config.App.Port,
		Handler:        newRouter(service),
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: 1 << 20,
	}

	return srv
}

func newRouter(service service.Service) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	handler := newHanler(service)

	router.Get("/orders/{id}", handler.ordersSearch)
	router.Get("/bad/{id}", handler.badMessagesSearch)

	return router
}

type handler struct {
	service service.Service
}

func newHanler(service service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ordersSearch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	orderId, err := strconv.Atoi(id)
	if err != nil {
		logrus.Println("failed convert string to int")
		h.error(w, r, http.StatusBadRequest, err)
	}

	model := h.service.FindById(orderId)
	if model == nil {
		logrus.Println("Not found")
		h.error(w, r, http.StatusNotFound, errors.New("not found"))
	}

	res, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		h.error(w, r, http.StatusInternalServerError, err)
	}

	h.respond(w, r, http.StatusOK, res)
}

func (h *handler) badMessagesSearch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	badMes, err := strconv.Atoi(id)
	if err != nil {
		logrus.Println("failed convert string to int")
		h.error(w, r, http.StatusBadRequest, err)
	}

	model := h.service.FindByIdBadMessage(badMes)
	if model == nil {
		logrus.Println("Not found")
		h.error(w, r, http.StatusNotFound, errors.New("not found"))
	}

	h.respond(w, r, http.StatusOK, model)
}

func (h *handler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (h *handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		switch v := data.(type) {
		case []byte:
			w.Write(v)
		default:
			json.NewEncoder(w).Encode(data)
		}
	}
}
