package payment

type Payment struct {
	ID     int
	Amount int
	UserID int `gorm:"foreignkey"`
}
