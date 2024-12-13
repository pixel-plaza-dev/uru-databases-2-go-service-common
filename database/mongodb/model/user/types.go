package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User is the MongoDB user model
type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	FirstName      string             `json:"first_name" bson:"first_name"`
	LastName       string             `json:"last_name" bson:"last_name"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	DeletedAt      time.Time          `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Birthdate      time.Time          `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	JoinedAt       time.Time          `json:"joined_at" bson:"joined_at"`
	UUID           string             `json:"uuid" bson:"uuid"`
}

// UserEmail is the MongoDB user email model
type UserEmail struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Email      string             `json:"email" bson:"email"`
	AssignedAt time.Time          `json:"assigned_at" bson:"assigned_at"`
	VerifiedAt time.Time          `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
	RevokedAt  time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
	IsPrimary  bool               `json:"is_primary" bson:"is_primary"`
}

// UserPhoneNumber is the MongoDB user phone number model
type UserPhoneNumber struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	AssignedAt  time.Time          `json:"assigned_at" bson:"assigned_at"`
	VerifiedAt  time.Time          `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
	RevokedAt   time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// UserEmailVerification is the MongoDB user email verification model
type UserEmailVerification struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserEmailID primitive.ObjectID `json:"user_email_id" bson:"user_email_id"`
	UUID        string             `json:"uuid" bson:"uuid"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	ExpiresAt   time.Time          `json:"expires_at" bson:"expires_at"`
	VerifiedAt  time.Time          `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
	RevokedAt   time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// UserPhoneNumberVerification is the MongoDB user phone number verification model
type UserPhoneNumberVerification struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	UserPhoneNumberID primitive.ObjectID `json:"user_phone_number_id" bson:"user_phone_number_id"`
	VerificationCode  string             `json:"verification_code" bson:"verification_code"`
	CreatedAt         time.Time          `json:"created_at" bson:"created_at"`
	ExpiresAt         time.Time          `json:"expires_at" bson:"expires_at"`
	VerifiedAt        time.Time          `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
	RevokedAt         time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// UserResetPassword is the MongoDB user password reset model
type UserResetPassword struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"`
	UsedAt    time.Time          `json:"used_at,omitempty" bson:"used_at,omitempty"`
	RevokedAt time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// UserUsernameLog is the MongoDB user username log model
type UserUsernameLog struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Username   string             `json:"username" bson:"username"`
	AssignedAt time.Time          `json:"assigned_at" bson:"assigned_at"`
}

// UserHashedPasswordLog is the MongoDB user hashed password log model
type UserHashedPasswordLog struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	AssignedAt     time.Time          `json:"assigned_at" bson:"assigned_at"`
}
