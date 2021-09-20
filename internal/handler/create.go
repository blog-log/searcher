package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"bloglog.com/search/v1/internal/markdown"
	"bloglog.com/search/v1/internal/model"
	"bloglog.com/search/v1/pkg/api/v1/create"
)

func (h *SearchHandler) Create(w http.ResponseWriter, r *http.Request) {
	// extract
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request create.Request
	json.Unmarshal(reqBody, &request)

	// request to internal model
	page := model.Page{
		Id:      request.Id,
		Repo:    request.Repo,
		Branch:  request.Branch,
		Path:    request.Path,
		Title:   request.Title,
		Content: request.Content,
	}

	// do actual work
	// get content
	content, err := markdown.Get(&page)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	page.Content = content

	if err := h.adder(r.Context(), &page); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(page)
}
