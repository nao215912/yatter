package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

type CreateRequest struct {
	Username string
	Password string
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}
	if req.Username == "" {
		httperror.BadRequest(w, fmt.Errorf("username must not be an empty string"))
		return
	}
	if req.Password == "" {
		httperror.BadRequest(w, fmt.Errorf("password must not be an empty string"))
		return
	}

	account := &object.Account{Username: req.Username}
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	ar := h.app.Dao.Account()
	ctx := r.Context()
	if err := ar.Create(ctx, account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	account, err := ar.FindByUsername(ctx, account.Username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
