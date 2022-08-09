package accounts

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) FindUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		httperror.BadRequest(w, fmt.Errorf("username must not be an empty string"))
		return
	}

	account, err := h.app.Dao.Account().FindByUsername(r.Context(), username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	if account == nil {
		httperror.BadRequest(w, fmt.Errorf("no such user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
