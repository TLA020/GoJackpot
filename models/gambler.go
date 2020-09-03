package models

type Gambler struct {
	Conn *Client
}

func NewGambler(conn *Client) *Gambler {
	return &Gambler{
		Conn: conn,
	}
}
