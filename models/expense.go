package models

type Expense struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	Description string `json:"description"`
}
