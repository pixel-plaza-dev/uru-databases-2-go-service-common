package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// UserLogInAttempt is the struct for the user log in attempt entity
type UserLogInAttempt struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	IPv4Address  string             `json:"ipv4_address" bson:"ipv4_address"`
	AttemptedAt  time.Time          `json:"attempted_at" bson:"attempted_at"`
	IsSuccessful bool               `json:"is_successful" bson:"is_successful"`
}

// JwtRefreshToken is the struct for the JWT refresh token entity
type JwtRefreshToken struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	UserID               primitive.ObjectID `json:"user_id" bson:"user_id"`
	UserLogInAttemptID   primitive.ObjectID `json:"user_log_in_attempt_id,omitempty" bson:"user_log_in_attempt_id,omitempty"`
	ParentRefreshTokenID primitive.ObjectID `json:"parent_refresh_token_id,omitempty" bson:"parent_refresh_token_id,omitempty"`
	IssuedAt             time.Time          `json:"issued_at" bson:"issued_at"`
	ExpiresAt            time.Time          `json:"expires_at" bson:"expires_at"`
	RevokedAt            time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// JwtAccessToken is the struct for the JWT access token entity
type JwtAccessToken struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	UserID            primitive.ObjectID `json:"user_id" bson:"user_id"`
	JwtRefreshTokenID primitive.ObjectID `json:"jwt_refresh_token_id" bson:"jwt_refresh_token_id"`
	IssuedAt          time.Time          `json:"issued_at" bson:"issued_at"`
	ExpiresAt         time.Time          `json:"expires_at" bson:"expires_at"`
	RevokedAt         time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// UserRole is the struct for the user role entity
type UserRole struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	RoleID           primitive.ObjectID `json:"role_id" bson:"role_id"`
	AssignedByUserID primitive.ObjectID `json:"assigned_by_user_id" bson:"assigned_by_user_id"`
	RevokedByUserID  primitive.ObjectID `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	AssignedAt       time.Time          `json:"assigned_at" bson:"assigned_at"`
	RevokedAt        time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// Role is the struct for the role entity
type Role struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	CreatedByUserID primitive.ObjectID `json:"created_by_user_id" bson:"created_by_user_id"`
	RevokedByUserID primitive.ObjectID `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	RevokedAt       time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// RolePermission is the struct for the role permission entity
type RolePermission struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	AssignedByUserID primitive.ObjectID `json:"assigned_by_user_id" bson:"assigned_by_user_id"`
	RevokedByUserID  primitive.ObjectID `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	RoleID           primitive.ObjectID `json:"role_id" bson:"role_id"`
	PermissionID     primitive.ObjectID `json:"permission_id" bson:"permission_id"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	RevokedAt        time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// Permission is the struct for the permission entity
type Permission struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	CreatedByUserID primitive.ObjectID `json:"created_by_user_id" bson:"created_by_user_id"`
	RevokedByUserID primitive.ObjectID `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	Action          string             `json:"action" bson:"action"`
	Resource        string             `json:"resource" bson:"resource"`
	Description     string             `json:"description" bson:"description"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	RevokedAt       time.Time          `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}
