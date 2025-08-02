## poc

一个 简单的 AI 服务， 用来演示效果

整个项目分成了下面几层:
1. handler 用于处理对应的
2. service 处理对应的业务逻辑
3. llm: 大模型服务, 目前有两个接口, handle 和 stream,
    handle 是完整的调用，stream 是流式调用
4. repo: 是整个持久化层