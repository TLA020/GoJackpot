package models

type Gambler struct {
	Conn *Connection
}

func NewGambler(conn *Connection) *Gambler {
	return &Gambler{
		Conn: conn,
	}
}
