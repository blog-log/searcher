package create

type Response struct {
	Id      string `json:"id"`
	Repo    string `json:"repo"`
	Branch  string `json:"branch"`
	Path    string `json:"path"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
