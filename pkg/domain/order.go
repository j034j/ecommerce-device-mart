package domain

import "time"

type OrderStatus struct {
	ID     uint   `gorm:"not null;primaryKey"`
	Status string `gorm:"not null;unique"`
}

type OrderLine struct {
	ID              uint          `gorm:"not null;primaryKey"`
	UserID          uint          `gorm:"not null"`
	User            User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AddressesID     uint          `gorm:"not null"`
	Addresses       Addresses     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID       uint          `gorm:"not null"`
	Product         Product       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PaymentMethodID int           `gorm:"not null"`
	PaymentMethod   PaymentMethod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrderStatusId   int           `gorm:"not null"`
	Qty             int           `gorm:"not null"`
	Price           float32       `gorm:"not null"`
	CouponID        uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

/*

Processing
    */
