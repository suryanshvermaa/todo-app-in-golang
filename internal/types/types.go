package types

type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo" validate:"required"`
	Completed bool   `json:"completed" validate:"required"`
}
