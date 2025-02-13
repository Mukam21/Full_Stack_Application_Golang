package entity

type Transaction struct {
	ID     int     `db:"id" json:"ID"`
	Type   string  `db:"type" json:"type" binding:"required, gte=2, lte=5"`
	Ticker string  `db:"ticker" json:"ticker" binding:"required, gte=1, lte=30"`
	Volume float32 `db:"volume" json:"volume" binding:"required"`
	Price  float32 `db:"price" json:"price" binding:"required"`
	Date   string  `db:"date" json:"date" binding:"required"`
}
