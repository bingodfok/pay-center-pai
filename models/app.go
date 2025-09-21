package models

import "time"

// AppStatus 应用状态
type AppStatus string

const (
	AppStatusActive    AppStatus = "active"    // 激活
	AppStatusInactive  AppStatus = "inactive"  // 未激活
	AppStatusSuspended AppStatus = "suspended" // 暂停
)

// App 应用信息
type App struct {
	ID          string    `json:"id" gorm:"primaryKey"`         // 应用ID
	AppName     string    `json:"app_name" gorm:"unique"`       // 应用名称
	AppSecret   string    `json:"app_secret"`                   // 应用密钥
	Description string    `json:"description"`                  // 应用描述
	Status      AppStatus `json:"status" gorm:"index"`          // 应用状态
	Permissions string    `json:"permissions" gorm:"type:text"` // 权限列表(JSON)
	QuotaLimit  int64     `json:"quota_limit"`                  // 配额限制
	QuotaUsed   int64     `json:"quota_used"`                   // 已使用配额
	CreatedAt   time.Time `json:"created_at"`                   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`                   // 更新时间
}

// AppPermission 应用权限
type AppPermission struct {
	ID         string    `json:"id" gorm:"primaryKey"` // 权限ID
	AppID      string    `json:"app_id" gorm:"index"`  // 应用ID
	Permission string    `json:"permission"`           // 权限名称
	Resource   string    `json:"resource"`             // 资源
	Action     string    `json:"action"`               // 操作
	CreatedAt  time.Time `json:"created_at"`           // 创建时间
}

// AppQuota 应用配额使用记录
type AppQuota struct {
	ID        string    `json:"id" gorm:"primaryKey"` // 记录ID
	AppID     string    `json:"app_id" gorm:"index"`  // 应用ID
	Date      string    `json:"date" gorm:"index"`    // 日期 (YYYY-MM-DD)
	QuotaUsed int64     `json:"quota_used"`           // 当日使用配额
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"updated_at"`           // 更新时间
}
