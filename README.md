# g-JTT: Golang JTT Framework

[中文版本](doc/Chinese.md) | [English Version](README.md)

**g-JTT** is an open-source Golang framework designed to simplify the development of applications using JTT series protocols for communication. This framework can be used to parse and process various protocols such as JT/T 808, JT/T 1078, and JT/T 809.

## Features
- Connection Management: Efficiently manage TCP connections and connection pools.
- Message Decoding: Decode incoming messages using the JTT protocol specification.
- Communication Server: Set up a communication server to listen for incoming connections.
- Extensibility: Designed to be easily extended to support different JTT protocols.

## Installation
Use `go get` to install g-JTT:
```bash
go get -u github.com/mingkid/g-JTT
```

## Usage
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
	_ = engine.Serve(":8080")
}

func handleMessage(ctx *jtt.Context) {
	var decoder codec.Decoder
	m := msg.New[msg.M0200]()

	decoder.Decode(m, ctx.Data())
	fmt.Printf("Hello, %s", msg.Head.Phone)
}
```

## Decoding and Encoding
When it comes to data transmission and communication protocols, decoding and encoding are crucial steps. They are used to convert complex data structures into byte sequences for transmission, and to reconvert received byte sequences back into the original data structures. At this point, the Decoder and Encoder play this important role.

### Decoding Principle

Decoding is the process of converting received byte data into high-level data structures. The `Decoder` object in the project is responsible for implementing the decoding logic. It recursively traverses the fields of the data structure, parses data from the byte sequence step by step according to the tag rules, and sets it into the corresponding fields.

### Encoding Principle

Encoding is the process of converting high-level data structures into byte sequences for transmission. The `Encoder` object in the project is responsible for implementing the encoding logic. It similarly traverses the fields of the data structure recursively, writes data of the fields into the byte buffer according to the tag rules and a certain format, and finally generates the byte sequence for transmission.

### Tag Use Rule Table

Tag rules are key elements in guiding decoding and encoding in the project. Below is a table that shows the usage, applicable field types, and effects of tag rules:

| Tag Rule                | Applicable Field Types             | Decoder Support | Encoder Support | Effect                              |
|-------------------------|------------------------------------|-----------------|-----------------|-------------------------------------|
| `jtt[Version]:"-"`       | Any Type                           | Supported       | Supported       | Skip Field                          |
| `jtt[Version]:"raw,Len"` | `[]byte` Type                      | Supported       | Supported       | Read Data by Length                 |
| `jtt[Version]:"bcd,Len"` | `string` Type                      | Supported       | Supported       | Read BCD Encoded Data               |
| `jtt[Version]:"Len"`     | `string` Type                      | Supported       | Supported       | Read Data by Length                 |
| No Tag                  | `string` Type                      | Supported       | Supported       | Read Remaining Data                 |
| No Tag                  | `[]byte` Type                      | Supported       | Supported       | Read Remaining Data                 |
| No Tag                  | `uint8`、`uint16`、`uint32`、`string` | Supported       | Supported       | Read Data by Type                   |
| No Tag                    | `map[uint8][]byte` type                    | Supported       | Supported       | Read and decode map key-value pairs |

When using tag rules, pay attention to the applicable field types and the effects of the rules. These tags will guide the decoding and encoding process, ensuring that data can be correctly transmitted and converted.

### Mapping of JTT Raw Data Types to Go Types

In the JTT series protocols, there is a mapping relationship between raw data types and Go data types. Here are some examples of the mapping relationships:

|   JTT Data Type   | Go Data Type |            Description and Requirements           |
| ----------------- |--------------| --------------------------------------------- |
| BYTE              | uint8        | Unsigned 8-bit integer (byte, 8 bits)         |
| WORD              | uint16       | Unsigned 16-bit integer (word, 16 bits)       |
| DWORD             | uint32       | Unsigned 32-bit integer (dword, 32 bits)      |
| BYTE[n]           | [n]byte      | n bytes                                      |
| BCD(n)            | string       | BCD encoding, n bytes                        |
| STRING            | string       | GBK encoded, empty if no data                |

These mapping relationships can guide the use of the correct Go data types to handle different fields during the decoding and encoding processes.

Through the decoding and encoding principles, the tag use rule table, and the mapping of JTT data types to Go data types, you can better understand and use the Decoder and Encoder modules to handle communication data related to the JTT series protocols.

## License
g-JTT is licensed under the [Apache 2.0 License](LICENSE).