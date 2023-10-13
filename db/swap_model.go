package db

type InputCoin struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Amount int    `json:"amount"`
}
