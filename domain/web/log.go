package web

type Log struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Path   string `json:"path"`
	Body   string `json:"body"`
}
