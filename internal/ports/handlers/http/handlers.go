package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fakovacic/ports/internal/ports"
)

func New(c *ports.Config, service ports.Service) *Handler {
	return &Handler{
		config:  c,
		service: service,
	}
}

type Handler struct {
	config  *ports.Config
	service ports.Service
}

func (h *Handler) writeResponse(w http.ResponseWriter, r *http.Request, resp interface{}) {
	reqID := ports.GetCtxStringVal(r.Context(), ports.ContextKeyRequestID)

	l := h.config.Log.With().
		Str("reqID", reqID).
		Interface("response", resp).
		Logger()
	l.Info().Msg("http response")

	bytes, err := json.Marshal(resp)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (h *Handler) writeError(w http.ResponseWriter, r *http.Request, er ErrorResponse) {
	reqID := ports.GetCtxStringVal(r.Context(), ports.ContextKeyRequestID)

	l := h.config.Log.With().
		Str("reqID", reqID).
		Interface("response", er).
		Logger()
	l.Error().Msg("http response")

	bytes, err := json.Marshal(er)
	if err != nil {
		log.Println("marshal", err)

		return
	}

	// write header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.Status)

	// write return body
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}
