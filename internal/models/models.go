package models

type Token struct {
	Id         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	NetworkId  uint32 `json:"network_id" db:"network_id"`
	CurrencyId uint32 `json:"currency_id" db:"currency_id"`
	IsActive   bool   `json:"is_active" db:"is_active"`
}
