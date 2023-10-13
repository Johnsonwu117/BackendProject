package db

type InputCoin struct {
	Source string  `json:"source"`
	Target string  `json:"target"`
	Amount float64 `json:"amount"`
}

type OutputCoin struct {
	Amount float64 `json:"amount"`
}
