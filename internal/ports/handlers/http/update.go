package http

import (
	"encoding/json"
	goErrors "errors"
	"io"
	"net/http"

	"github.com/fakovacic/ports/internal/ports"
	"github.com/fakovacic/ports/internal/ports/errors"
	"github.com/julienschmidt/httprouter"
)

type UpdateRequest struct {
	Port *ports.Port `json:"port"`
}

type UpdateResponse struct {
	Port *ports.Port `json:"port"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, par httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	var req UpdateRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	port, err := h.service.Update(r.Context(), par.ByName("id"), req.Port)
	if err != nil {
		var e errors.Error

		ok := goErrors.As(err, &e)
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

	h.writeResponse(w, r, &UpdateResponse{
		Port: port,
	})
}
