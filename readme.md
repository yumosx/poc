# poc

一个 简单的 AI 服务， 用来演示效果

一、服务启动方式

1. 设置环境变量

在启动服务前，请先设置你的 AI Token：

```shell
export AI_TOKEN="your_token_here"
```
2. 启动服务

Docker Compose 启动服务：
```
cd .script
docker compose up
```


## 架构介绍

1. Handler 层
    负责接收并处理 Web 请求，将请求转发至对应的服务层。
2. Service 层
    实现核心业务逻辑，包括：
        - 创建任务
        - 查询任务状态
        - 异步执行任务
        - LLM 模块（大模型服务） 提供两种调用接口：
            - handle: 同步调用，返回完整结果
            - stream: 流式调用，逐步返回响应
3. Repo 层（持久化层） 负责数据的持久化存储，当前主要用于存储任务信息（Task）。


## 接口:

1. 获取功能
```curl
curl -X GET 'http://localhost:8080/ai/v1/list'
```
响应:
```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "total": 3,
        "functions": [
            {
                "name": "中译英",
                "desc": "中文翻译成为英文",
                "type": "translate_zh2en"
            },
            {
                "name": "英译中",
                "desc": "英文翻译成为中文",
                "type": "translate_en2zh"
            },
            {
                "name": "总结功能",
                "desc": "对文字进行总结",
                "type": "summarize"
            }
        ]
    }
}
```

2. 创建对应的任务
```curl
curl -X POST 'http://localhost:8080/ai/v1/run' \
    -H 'Content-Type: application/json' \
    -d '{
        "content":"请把这句话翻译成英文：你好阿哥",
        "type":"translate_zh2en"
    }'
```

返回的结果
```shell
{"code":200,"msg":"","data":"5365c0c8-daed-4894-95ec-69d18ee70506"}
```

3. 获取对应的任务, id 需要替换
```curl
curl -X GET 'http://localhost:8080/ai/v1/task/5365c0c8-daed-4894-95ec-69d18ee70506'
```

```shell
{
    "code": 200,
    "msg": "",
    "data": {
        "id": "5365c0c8-daed-4894-95ec-69d18ee70506",
        "type": "",
        "state": "success",
        "result": "Hello, brother."
    }
}
```
4. 流式接口

流式接口的测试，需要使用对应的 client.go

- 服务端打印:
```go
[GIN] 2025/08/02 - 23:52:02 | 200 | 22.282327583s |       127.0.0.1 | POST     "/ai/v1/stream"
2025/08/02 23:53:59.216272 DEBUG: eventMessage|
2025/08/02 23:53:59.216531 DEBUG: eventMessage|Hello
2025/08/02 23:53:59.748328 DEBUG: eventMessage|,
2025/08/02 23:53:59.848696 DEBUG: eventMessage| please
2025/08/02 23:53:59.848796 DEBUG: eventMessage| help
2025/08/02 23:53:59.944297 DEBUG: eventMessage| me
2025/08/02 23:53:59.944433 DEBUG: eventMessage| translate
2025/08/02 23:54:00.047413 DEBUG: eventMessage| this
2025/08/02 23:54:00.143595 DEBUG: eventMessage| into
2025/08/02 23:54:00.143665 DEBUG: eventMessage| English
2025/08/02 23:54:00.241179 DEBUG: eventMessage|.
2025/08/02 23:54:00.241687 DEBUG: eventMessage|
```

- 客户端打印

**使用的是拼接之后的**
```go
开始接收流式响应：
Hello, please help me translate this into English.
流结束
```
