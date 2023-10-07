# g-JTT: Golang JTT 框架

[中文版本](doc/Chinese.md) | [English Version](README.md)

**g-JTT** 是一个开源的 Golang 框架，旨在简化使用 JTT 系列协议进行通信的应用程序开发。该框架可以用于解析和处理 JT/T 808、JT/T 1078 和 JT/T 809 等不同的协议。

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

	jtt "github.com/mingkid/g-jtt"
	"github.com/mingkid/g-JTT/protocol/codec"
	"github.com/mingkid/g-JTT/protocol/msg"
)

func main() {
	// Create a connection pool and engine
	engine := jtt.Default()

	// Register message handlers
	engine.RegisterHandler(msg.MsgID(0x0200), handleMessage)

	// Start the communication server
	_ = engine.Serve("", 9300)
}

func handleMessage(ctx *jtt.Context) {
	var (
		m msg.Msg[msg.M0200]
		d codec.Decoder
	)

	d.Decode(m, ctx.Data())
	fmt.Printf("Hello, %s", msg.Head.Phone)
	_ = ctx.Generic(msg.M8001ResultSuccess)
}
```

## 解码和编码
当涉及到数据传输和通信协议时，解码和编码是关键的步骤，它们用于将复杂的数据结构转换为字节序列以进行传输，并将接收到的字节序列重新转换为原始的数据结构。此时，解码器（`Decoder`）和编码器（`Encoder`）扮演着这一重要角色。

### 解码原理

解码是将接收到的字节数据转换为高级数据结构的过程。项目中的 `Decoder` 对象负责实现解码逻辑。它通过递归遍历数据结构的字段，根据标签规则，从字节序列中逐步解析出数据，并将其设置到相应的字段中。

### 编码原理

编码是将高级数据结构转换为字节序列的过程，以便进行传输。项目中的 `Encoder` 对象负责实现编码逻辑。它同样通过递归遍历数据结构的字段，根据标签规则，将字段的数据按照一定的格式写入字节缓冲区中，最终生成待传输的字节序列。

### 标签使用规则表

标签规则是项目中用于指导解码和编码的关键元素。下面是一张表格，展示了标签规则的使用情况、适用的字段类型以及规则的效果：

| 标签规则               | 适用字段类型                    | 解码器支持 | 编码器支持 | 效果              |
|------------------------|---------------------------|----------|----------|-----------------|
| `jtt[版本号]:"-"`       | 任意类型                      | 支持     | 支持     | 跳过字段            |
| `jtt[版本号]:"raw,长度"` | `[]byte` 类型               | 支持     | 支持     | 按长度读取指定位数数据     |
| `jtt[版本号]:"bcd,长度"` | `string` 类型               | 支持     | 支持     | 读取 BCD 编码数据     |
| `jtt[版本号]:"长度"`     | `string` 类型               | 支持     | 支持     | 按长度读取数据         |
| 无标签                 | `string` 类型               | 支持     | 支持     | 按照 GBK 编码读取剩余数据 |
| 无标签                 | `[]byte` 类型               | 支持     | 支持     | 读取剩余数据          |
| 无标签                 | `uint8`、`uint16`、`uint32` | 支持     | 支持     | 按类型读取数据         |
| 无标签                    | `map[uint8][]byte` 类型     | 支持     | 支持     | 读取和解码 map 键值对   |

在使用标签规则时，请注意适用的字段类型和规则的效果。这些标签将指导解码和编码过程，确保数据能够正确地传输和转换。

### JTT 中的原始数据类型和 Go 类型映射表

在 JTT 系列协议中，原始数据类型与 Go 语言的数据类型之间存在一定的映射关系。以下是一些常见的映射关系示例：

|   JTT 数据类型   | Go 数据类型 |         描述及要求         |
| --------------- |---------| ----------------------- |
| BYTE            | uint8   | 无符号单字节整型（字节，8 位 |
| WORD            | uint16  | 无符号双字节整型（字，16 位） |
| DWORD           | uint32  | 无符号四字节整型（双字，32 位）|
| BYTE[n]         | [n]byte | n 字节 |
| BCD(n)          | string  | 8421 码，n 字节 |
| STRING          | string  | GBK 编码，若无数据，置空 |

这些映射关系可以指导在解码和编码过程中使用正确的 Go 数据类型来处理不同的字段。

通过以上的解码和编码原理、标签使用规则表以及 JTT 数据类型与 Go 数据类型映射表，您可以更好地理解和使用 Decoder 和 Encoder 模块，以便处理 JTT 系列协议相关的通信数据。

## 许可证
g-JTT 使用 [Apache 2.0](LICENSE) 许可证。