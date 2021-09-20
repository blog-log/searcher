package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"bloglog.com/search/v1/pkg/api/v1/search"
)

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	// extract
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request search.Request
	json.Unmarshal(reqBody, &request)

	// do actual work
	pages, err := h.searcher(r.Context(), request.Query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(pages)
}
