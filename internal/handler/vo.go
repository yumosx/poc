package handler

type ListRequest struct {
	limit  int
	offset int
}

type Function struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

type ListResponse struct {
	Total     int        `json:"total"`
	Functions []Function `json:"functions"`
}

type SubmitTaskRequest struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	Content string `json:"content"`
}

type TaskResponse struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	State  string `json:"state"`
	Result string `json:"result"`
}
