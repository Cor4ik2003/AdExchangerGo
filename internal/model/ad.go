package model

type Ad struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsAuction   bool   `json:"is_auction"`
	Price       uint64 `json:"price"`
}
