package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LLMRequest struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

type StreamResponse struct {
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
	Err     string `json:"err,omitempty"`
}

func main() {
	url := "http://localhost:8080/ai/v1/stream"

	// 构造请求体
	reqBody := LLMRequest{
		Content: "你好，请帮我翻译成英文。",
		Type:    "translate_zh2en",
	}
	jsonData, _ := json.Marshal(reqBody)

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("开始接收流式响应：")

	// 按行读取流式响应
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// fmt.Printf("收到原始数据: %s\n", line)

		// 解析为 JSON
		var event StreamResponse
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			fmt.Printf("JSON解析失败: %v, 原始数据: %s\n", err, line)
			continue
		}

		switch event.Type {
		case "event_message":
			fmt.Print(event.Content)
		case "event_err":
			fmt.Printf("\n错误: %s\n", event.Err)
			return
		case "event_done":
			fmt.Println("\n流结束")
			return
		default:
			fmt.Printf("\n未知事件类型: %s\n", event.Type)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取错误: %v\n", err)
	}
}
