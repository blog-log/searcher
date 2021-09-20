package model

type AlgoliaPage struct {
	Id      string `json:"objectID"` // important json is objectID to set id in algolia
	Repo    string `json:"repo"`
	Branch  string `json:"branch"`
	Path    string `json:"path"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
