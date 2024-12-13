package order

import (
	"database/sql"
	"gorm.io/gorm"
)

// User is the PostgreSQL user model
type User struct {
	gorm.Model
	UserIDHex string
}

// CartProduct is the PostgreSQL cart product model
type CartProduct struct {
	gorm.Model
	CartID             uint64
	Cart               Cart `gorm:"foreignKey:CartID"`
	BranchProductIDHex string
	Quantity           uint64
	ProductPrice       float64
	TotalPrice         float64
	Discount           sql.NullFloat64
}

// Cart is the PostgreSQL cart model
type Cart struct {
	gorm.Model
	UserID  uint64
	User    User `gorm:"foreignKey:UserID"`
	OrderID sql.NullInt64
	Order   Order `gorm:"foreignKey:OrderID"`
}

// Order is the PostgreSQL order model
type Order struct {
	CashierID             sql.NullInt64
	Cashier               User `gorm:"foreignKey:CashierID"`
	EstimatedDeliveryDate sql.NullTime
}
