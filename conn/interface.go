package conn

type Pool interface {
	Add(termID string, conn *Connection)
	Get(termID string) (*Connection, bool)
}
