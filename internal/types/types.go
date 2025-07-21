package types

type Todo struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
}
