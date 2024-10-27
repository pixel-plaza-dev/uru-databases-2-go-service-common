package auth_service

// UserLogInAttempt is the struct for the user log in attempt
type UserLogInAttempt struct {
	ID           string `json:"id" bson:"_id"`
	UserID       string `json:"user_id" bson:"user_id"`
	IPAddress    string `json:"ip_address" bson:"ip_address"`
	AttemptedAt  string `json:"attempted_at" bson:"attempted_at"`
	IsSuccessful bool   `json:"is_successful" bson:"is_successful"`
}

// JwtRefreshToken is the struct for the JWT refresh token
type JwtRefreshToken struct {
	ID                   string `json:"id" bson:"_id"`
	UserID               string `json:"user_id" bson:"user_id"`
	UserLogInAttemptID   string `json:"user_log_in_attempt_id" bson:"user_log_in_attempt_id"`
	ParentRefreshTokenID string `json:"parent_refresh_token_id" bson:"parent_refresh_token_id"`
	CreatedAt            string `json:"created_at" bson:"created_at"`
	ExpiresAt            string `json:"expires_at" bson:"expires_at"`
	RevokedAt            string `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// JwtRefreshTokenLog is the struct for the JWT refresh token log
type JwtRefreshTokenLog struct {
	ID                string `json:"id" bson:"_id"`
	JwtRefreshTokenID string `json:"jwt_refresh_token_id" bson:"jwt_refresh_token_id"`
	IPAddress         string `json:"ip_address" bson:"ip_address"`
	UsedAt            string `json:"used_at" bson:"used_at"`
}

// JwtAccessToken is the struct for the JWT access token
type JwtAccessToken struct {
	ID                string `json:"id" bson:"_id"`
	UserID            string `json:"user_id" bson:"user_id"`
	JwtRefreshTokenID string `json:"jwt_refresh_token_id" bson:"jwt_refresh_token_id"`
	CreatedAt         string `json:"created_at" bson:"created_at"`
	ExpiresAt         string `json:"expires_at" bson:"expires_at"`
	RevokedAt         string `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// JwtAccessTokenLog is the struct for the JWT access token log
type JwtAccessTokenLog struct {
	ID               string `json:"id" bson:"_id"`
	JwtAccessTokenID string `json:"jwt_access_token_id" bson:"jwt_access_token_id"`
	IPAddress        string `json:"ip_address" bson:"ip_address"`
	UsedAt           string `json:"used_at" bson:"used_at"`
}

// UserRole is the struct for the user role
type UserRole struct {
	ID               string `json:"id" bson:"_id"`
	UserID           string `json:"user_id" bson:"user_id"`
	RoleID           string `json:"role_id" bson:"role_id"`
	AssignedByUserID string `json:"assigned_by_user_id" bson:"assigned_by_user_id"`
	RevokedByUserID  string `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	AssignedAt       string `json:"assigned_at" bson:"assigned_at"`
	RevokedAt        string `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// Role is the struct for the role
type Role struct {
	ID              string `json:"id" bson:"_id"`
	CreatedByUserID string `json:"created_by_user_id" bson:"created_by_user_id"`
	RevokedByUserID string `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	Name            string `json:"name" bson:"name"`
	Description     string `json:"description" bson:"description"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
	RevokedAt       string `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// RolePermission is the struct for the role permission
type RolePermission struct {
	ID               string `json:"id" bson:"_id"`
	AssignedByUserID string `json:"assigned_by_user_id" bson:"assigned_by_user_id"`
	RevokedByUserID  string `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	RoleID           string `json:"role_id" bson:"role_id"`
	PermissionID     string `json:"permission_id" bson:"permission_id"`
	CreatedAt        string `json:"created_at" bson:"created_at"`
	RevokedAt        string `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
}

// Permission is the struct for the permission
type Permission struct {
	ID              string `json:"id" bson:"_id"`
	CreatedByUserID string `json:"created_by_user_id" bson:"created_by_user_id"`
	RevokedByUserID string `json:"revoked_by_user_id,omitempty" bson:"revoked_by_user_id,omitempty"`
	Action          string `json:"action" bson:"action"`
	Resource        string `json:"resource" bson:"resource"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
	RevokedAt       string `json:"revoked_at,omitempty" bson:"revoked_at,omitempty"`
	Description     string `json:"description" bson:"description"`
}
