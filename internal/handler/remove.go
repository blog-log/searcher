package handler

import (
	"encoding/json"
	"net/http"

	"bloglog.com/search/v1/pkg/api/v1/remove"
	"github.com/gorilla/mux"
)

func (h *SearchHandler) Remove(w http.ResponseWriter, r *http.Request) {
	// extract
	vars := mux.Vars(r)
	id := vars["id"]

	// do actual work
	if err := h.remover(r.Context(), id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := remove.Response{
		Id: id,
	}
	json.NewEncoder(w).Encode(response)
}
