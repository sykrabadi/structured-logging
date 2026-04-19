package server

import (
	"net/http"
)

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) error {
	err := h.service.Home(r.Context())
	if err != nil {
		return err
	}

	return nil
}
