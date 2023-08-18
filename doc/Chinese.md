# g-JTT: Golang JTT 框架

[中文版本](doc/Chinese.md) | [English Version](README.md)

**g-JTT** 是一个开源的 Golang 框架，旨在简化使用 JTT 系列协议进行通信的应用程序开发。该框架可以用于解析和处理 JTT808、JTT1078 和 JTT809 等不同的协议。

## 特性
- 连接管理：高效地管理 TCP 连接和连接池。
- 消息解码：使用 JTT 协议规范解码传入的消息。
- 通讯服务：设置通讯服务器以侦听传入连接。
- 可扩展性：设计为易于扩展以支持不同的 JTT 协议。

## 安装
使用 go get 来安装 g-JTT：
```bash
go get -u github.com/mingkid/g-JTT
```

## 使用示例
```go
package main

import (
	"fmt"
	
	"github.com/mingkid/g-JTT/protocol/codec"
	"github.com/mingkid/g-JTT/protocol/msg"
)

func main() {
	// Create a connection pool and engine
	engine := jtt.Default()

	// Register message handlers
	engine.RegisterHandler(msg.MsgID(0x0200), handleMessage)

	// Start the communication server
	err := engine.Serve(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}

func handleMessage(ctx *jtt.Context) {
	var (
		msg     msg.M0200
		decoder codec.Decoder
	)

	decoder.Decode(&msg, ctx.Body())
	fmt.Printf("Hello, %s", msg.Head.Phone)
}
```

## 许可证
g-JTT 使用 [Apache 2.0](LICENSE) 许可证。