package handler

import (
	"bytes"
	"encoding/json"
	"github.com/ecodeclub/ginx"
	"github.com/gin-gonic/gin"
	"github.com/yumosx/poc/internal/domain"
	"github.com/yumosx/poc/internal/service"
	"github.com/yumosx/poc/internal/utils/logger"
)

type Handler struct {
	svc *service.AIService
}

func NewHandler(svc *service.AIService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Route(engine *gin.Engine) {
	g := engine.Group("/ai/v1")
	g.GET("/list", ginx.B(h.List))
	g.POST("/run", ginx.B(h.Run))
	g.POST("/stream", h.Stream)
	g.GET("/task/:id", ginx.W(h.GetTask))
}

// List 获取功能列表接口, 未来考虑存储在数据库中
func (h *Handler) List(ctx *ginx.Context, req ListRequest) (ginx.Result, error) {
	functions := []Function{
		{Name: "中译英", Desc: "中文翻译成为英文", Type: "translate_zh2en"},
		{Name: "英译中", Desc: "英文翻译成为中文", Type: "translate_en2zh"},
		{Name: "总结功能", Desc: "对文字进行总结", Type: "summarize"},
	}

	return ginx.Result{
		Code: 200,
		Msg:  "success",
		Data: ListResponse{
			Total:     3,
			Functions: functions,
		},
	}, nil
}

// Run 提交对应的任务, 并且异步执行
func (h *Handler) Run(ctx *ginx.Context, req SubmitTaskRequest) (ginx.Result, error) {
	id, err := h.svc.RunTask(ctx, domain.Task{UUID: req.Id, Content: req.Content, Type: req.Type})
	if err != nil {
		return ginx.Result{Code: 500, Data: "内部错误"}, err
	}
	return ginx.Result{Code: 200, Data: id}, nil
}

// GetTask 获取任务的执行结果和状态
func (h *Handler) GetTask(ctx *ginx.Context) (ginx.Result, error) {
	id := ctx.Param("id")
	uuid, err := id.AsString()
	if err != nil {
		return ginx.Result{Code: 500, Data: "内部错误"}, err
	}
	task, err := h.svc.GetTask(ctx, uuid)
	if err != nil {
		return ginx.Result{Code: 500, Data: "内部错误"}, err
	}
	return ginx.Result{Code: 200, Data: TaskResponse{Id: task.UUID, State: task.State, Type: task.Type, Result: task.Result}}, nil
}

// Stream 调用对应的大模型, 并以 stream 的方式返回
func (h *Handler) Stream(ctx *gin.Context) {
	var req LLMRequest
	if err := ctx.Bind(&req); err != nil {
		logger.Errorf("stream| bind| %v", err)
		return
	}

	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	stream, err := h.svc.Stream(ctx, domain.LLMRequest{Content: req.Content, Type: req.Type})
	if err != nil {
		h.eventErr(ctx, err)
		return
	}
	h.stream(ctx, stream)
}

func (h *Handler) stream(ctx *gin.Context, ch chan domain.StreamResponse) {
	for {
		select {
		case event, ok := <-ch:
			if !ok || event.Done {
				h.eventDone(ctx)
				return
			}
			if event.Err != nil {
				h.eventErr(ctx, event.Err)
				return
			}
			h.eventMessage(ctx, event.Content)
		case <-ctx.Request.Context().Done():
			return
		}
	}
}

func (h *Handler) sendEvent(ctx *gin.Context, data string) {
	buf := bytes.Buffer{}
	buf.WriteString(data)
	buf.WriteByte('\n')
	_, _ = ctx.Writer.Write(buf.Bytes())
	ctx.Writer.Flush()
}

func (h *Handler) eventMessage(ctx *gin.Context, content string) {
	event := StreamResponse{
		Type:    EventMessage,
		Content: content,
	}
	eventStr, _ := json.Marshal(event)
	h.sendEvent(ctx, string(eventStr))
}

func (h *Handler) eventErr(ctx *gin.Context, err error) {
	event := StreamResponse{
		Type: EventErr,
		Err:  err.Error(),
	}
	eventStr, _ := json.Marshal(event)
	h.sendEvent(ctx, string(eventStr))
}

func (h *Handler) eventDone(ctx *gin.Context) {
	event := StreamResponse{
		Type: EventDone,
	}
	eventStr, _ := json.Marshal(event)
	h.sendEvent(ctx, string(eventStr))
}
