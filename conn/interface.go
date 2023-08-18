package conn

type Pool interface {
	Add(terminalID string, conn *Connection)
	Get(terminalID string) *Connection
}
