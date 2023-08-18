# g-JTT: Golang JTT Framework

[中文版本](doc/Chinese.md) | [English Version](README.md)

**g-JTT** is an open-source Golang framework designed to simplify the development of applications that communicate using various JTT protocols. This framework can be used to parse and handle different protocols such as JTT808, JTT1078, and JTT809.

## Features
- Connection Management: Efficiently manage TCP connections and connection pools.
- Message Decoding: Decode incoming messages using JTT protocol specifications.
- Communication Service: Set up a communication server to listen for incoming connections.
- Extendable: Designed to be easily extended to support different JTT protocols.

## Installation
Use go get to install g-JTT:

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

## License
g-JTT is licensed under the [Apache 2.0 License](LICENSE).