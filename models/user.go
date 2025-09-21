package models

import (
	"time"
)

// UserStatus 用户状态
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"   // 激活
	UserStatusInactive UserStatus = "inactive" // 停用
	UserStatusLocked   UserStatus = "locked"   // 锁定
)

// UserRole 用户角色
type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"    // 管理员
	UserRoleOperator UserRole = "operator" // 操作员
	UserRoleViewer   UserRole = "viewer"   // 查看者
)

// User 用户模型
type User struct {
	ID           string     `json:"id" gorm:"primaryKey"`
	Username     string     `json:"username" gorm:"unique;not null"`
	Email        string     `json:"email" gorm:"unique;not null"`
	PasswordHash string     `json:"-" gorm:"not null"` // 密码哈希，不返回给前端
	Salt         string     `json:"-" gorm:"not null"` // 密码盐值
	RealName     string     `json:"realName" gorm:"column:real_name"`
	Phone        string     `json:"phone"`
	Avatar       string     `json:"avatar"`
	Role         UserRole   `json:"role" gorm:"not null;default:'viewer'"`
	Status       UserStatus `json:"status" gorm:"not null;default:'active'"`
	LastLoginAt  *time.Time `json:"lastLoginAt" gorm:"column:last_login_at"`
	LastLoginIP  string     `json:"lastLoginIP" gorm:"column:last_login_ip"`
	LoginCount   int64      `json:"loginCount" gorm:"column:login_count;default:0"`
	FailedCount  int64      `json:"failedCount" gorm:"column:failed_count;default:0"` // 连续失败次数
	LockedAt     *time.Time `json:"lockedAt" gorm:"column:locked_at"`                 // 锁定时间
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=6,max=50"`
	Captcha   string `json:"captcha" validate:"required,len=4"`
	CaptchaID string `json:"captchaId" validate:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string    `json:"token"`
	User      UserInfo  `json:"user"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// UserInfo 用户信息（不包含敏感信息）
type UserInfo struct {
	ID          string     `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	RealName    string     `json:"realName"`
	Phone       string     `json:"phone"`
	Avatar      string     `json:"avatar"`
	Role        UserRole   `json:"role"`
	Status      UserStatus `json:"status"`
	LastLoginAt *time.Time `json:"lastLoginAt"`
	LastLoginIP string     `json:"lastLoginIP"`
	LoginCount  int64      `json:"loginCount"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string   `json:"username" validate:"required,min=3,max=20"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=6,max=50"`
	RealName string   `json:"realName" validate:"required,min=2,max=20"`
	Phone    string   `json:"phone" validate:"omitempty,len=11"`
	Role     UserRole `json:"role" validate:"required,oneof=admin operator viewer"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Email    string     `json:"email" validate:"omitempty,email"`
	RealName string     `json:"realName" validate:"omitempty,min=2,max=20"`
	Phone    string     `json:"phone" validate:"omitempty,len=11"`
	Avatar   string     `json:"avatar"`
	Role     UserRole   `json:"role" validate:"omitempty,oneof=admin operator viewer"`
	Status   UserStatus `json:"status" validate:"omitempty,oneof=active inactive locked"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=50"`
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	UserID      string `json:"userId" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=50"`
}

// CaptchaResponse 验证码响应
type CaptchaResponse struct {
	CaptchaID string `json:"captchaId"`
	ImageData string `json:"imageData"` // Base64编码的图片数据
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Users []UserInfo `json:"users"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}

// UserFilters 用户查询过滤器
type UserFilters struct {
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Role     UserRole   `json:"role"`
	Status   UserStatus `json:"status"`
	Page     int        `json:"page"`
	Size     int        `json:"size"`
}
