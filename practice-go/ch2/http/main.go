package main

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	db *sqlx.DB
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var v interface{}
	err := h.db.Get(&v, "SELECT * FROM persion WHERE id=$1", r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(v)

}
