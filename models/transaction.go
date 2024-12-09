package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Payment model
type Transaction struct {
	gorm.Model
	UserId    int       `gorm:"not null" json:"user_id"`
	Method    string    `gorm:"type:varchar(50);not null" json:"method"` // Payment method (e.g., credit card, PayPal)
	Amount    float64   `gorm:"not null" json:"amount"`                  // Total payment amount
	Status    string    `gorm:"type:varchar(50);not null" json:"status"` // Payment status (e.g., pending, completed)
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	// Relationships
	TransactionDetails []TransactionDetail `gorm:"foreignKey:PaymentID" json:"transaction_details"`
}

// TransactionDetail model
type TransactionDetail struct {
	gorm.Model
	PaymentID   uint      `gorm:"not null" json:"payment_id"`                     // Foreign key to Payment
	ProductID   int       `gorm:"type:int(11);not null" json:"product_id"`        // Name of the product or service
	ProductName string    `gorm:"type:varchar(100);not null" json:"product_name"` // Name of the product or service
	Quantity    int       `gorm:"not null" json:"quantity"`                       // Quantity of items
	Price       float64   `gorm:"not null" json:"price"`                          // Price per unit
	Discount    float64   `gorm:"not null" json:"discount"`                       // Price per unit
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
