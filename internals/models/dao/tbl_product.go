package dao

import "time"

type Product struct {
	ID        string    `json:"id" gorm:"type:varchar(50)" db_connector:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null" db_connector:"name"`
	Amount    int       `json:"amount" gorm:"not null" db_connector:"amount"`
	Price     float64   `json:"price" gorm:"type:numeric(10,2);not null" db_connector:"price"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp" db_connector:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp" db_connector:"updated_at"`
}

func (table *Product) TableName() string {
	return "tbl_product"
}
