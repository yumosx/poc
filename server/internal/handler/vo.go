package handler

type ListRequest struct {
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

type LLMRequest struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

const (
	EventErr     = "event_err"
	EventMessage = "event_message"
	EventDone    = "event_done"
)

type StreamResponse struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Err     string `json:"err"`
}
