package domain

type LLMRequest struct {
	Content string
}

type StreamResponse struct {
	Content string
}

const (
	Padding = "padding"
	Running = "running"
	Failed  = "failed"
	Success = "success"
)

type Task struct {
	UUID    string
	State   string // 任务的执行状态
	Type    string // 任务的类型
	Content string // 任务内容
	Result  string // 任务的执行结果
}
