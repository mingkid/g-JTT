package conn

type Pool interface {
	// Add 添加终端 ID 对应的连接对象
	Add(termID string, conn *Connection)
	// Get 返回终端 ID 对应的连接对象
	Get(termID string) (*Connection, bool)
}
