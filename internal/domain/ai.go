package domain

type LLMRequest struct {
	Type    string //当前这个请求是做什么的
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

// 任务的类型
const (
	ZH2EN     = "translate_zh2en"
	EN2ZH     = "translate_en2zh"
	SUMMARIZE = "summarize"
)

type Task struct {
	UUID    string
	State   string // 任务的执行状态
	Type    string // 任务的类型
	Content string // 任务内容
	Result  string // 任务的执行结果
}
