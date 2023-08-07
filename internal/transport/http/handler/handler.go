package handler

import (
	"diploma/internal/transport/http/dto"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Service interface {
	GetResultData() (dto.ResultT, error)
}
type Cacher interface {
	Set(value dto.ResultT)
	Get() (dto.ResultT, bool)
	Delete() error
}
type Handler struct {
	cache  Cacher
	l      *logrus.Logger
	Router *mux.Router
	svc    Service
}

func New(l *logrus.Logger, s Service, cache Cacher) *Handler {
	router := mux.NewRouter()
	return &Handler{
		cache:  cache,
		l:      l,
		Router: router,
		svc:    s,
	}
}
func (h *Handler) RegisterRoutes() {

	h.Router.HandleFunc("/", h.GetResultData)
}
func (h *Handler) GetResultData(w http.ResponseWriter, r *http.Request) {
	response, found := h.cache.Get()
	if !found {
		resp, err := h.svc.GetResultData()
		if err != nil {
			h.l.Info(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
		}
		h.cache.Set(resp)
		response = resp
	}

	resp, err := json.Marshal(response)
	if err != nil {
		h.l.Info(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
