package models

type Gambler struct {
	Account  Account	`json:"Account"`
	Conn    *Connection	`json:"-"`
}

func NewUser(account Account,  conn *Connection) *Gambler {
	return &Gambler{
		Account: account,
		Conn: conn,
	}
}

