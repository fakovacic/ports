package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/errors"
	"github.com/julienschmidt/httprouter"
)

type CreateRequest struct {
	Port *ports.Port `json:"port"`
}

type CreateResponse struct {
	Port *ports.Port `json:"port"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, par httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	var req CreateRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	port, err := h.service.Create(r.Context(), req.Port)
	if err != nil {
		e, ok := err.(errors.Error)
		if ok {
			h.writeError(w, r, ErrorResponse{
				Message: e.Error(),
				Status:  e.Status,
			})

			return
		}

		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	h.writeResponse(w, r, &CreateResponse{
		Port: port,
	})
}