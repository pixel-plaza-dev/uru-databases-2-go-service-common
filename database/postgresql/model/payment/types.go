package payment

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// User is the PostgreSQL user model
type User struct {
	gorm.Model
	UserIDHex string
}

// OrderPayments is the PostgreSQL order payments model
type OrderPayments struct {
	gorm.Model
	OrderID uint64
}

// OrderPayment is the PostgreSQL order payment model
type OrderPayment struct {
	gorm.Model
	OrderPaymentsID uint64
	OrderPayments   OrderPayments `gorm:"foreignKey:OrderPaymentsID"`
	PaymentID       uint64
	Payment         Payment `gorm:"foreignKey:PaymentID"`
}

// BranchRentPayments is the PostgreSQL branch rent payments model
type BranchRentPayments struct {
	gorm.Model
	BranchRentIDHex string
}

// BranchRentPayment is the PostgreSQL branch rent payment model
type BranchRentPayment struct {
	gorm.Model
	BranchRentPaymentsID uint64
	BranchRentPayments   BranchRentPayments `gorm:"foreignKey:BranchRentPaymentsID"`
	PaymentID            uint64
	Payment              Payment `gorm:"foreignKey:PaymentID"`
}

// Payment is the PostgreSQL payment model
type Payment struct {
	gorm.Model
	PaymentAccountID uint64
	PaymentAccount   PaymentAccount `gorm:"foreignKey:PaymentAccountID"`
	Amount           float64
	PaymentDate      time.Time
	RefundedAt       sql.NullTime
	VerifiedByUserID sql.NullInt64
	VerifiedByUser   User `gorm:"foreignKey:VerifiedByUserID"`
}

// PaymentAccount is the PostgreSQL payment account model
type PaymentAccount struct {
	gorm.Model
	BranchIDHex        sql.NullString
	AccountPlatform    string
	AccountNumber      sql.NullString
	AccountName        sql.NullString
	AccountEmail       sql.NullString
	AccountPhoneNumber sql.NullString
	IsActive           bool
}
