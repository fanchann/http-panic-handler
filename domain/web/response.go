package web

type WebResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
