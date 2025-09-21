package models

import "time"

// PlatformStatus 平台状态
type PlatformStatus string

const (
	PlatformStatusActive      PlatformStatus = "active"      // 激活
	PlatformStatusInactive    PlatformStatus = "inactive"    // 未激活
	PlatformStatusMaintenance PlatformStatus = "maintenance" // 维护中
)

// Platform 支付平台类型
type Platform struct {
	ID           string         `json:"id" gorm:"primaryKey"`        // 平台ID
	PlatformName string         `json:"platform_name" gorm:"unique"` // 平台名称
	DisplayName  string         `json:"display_name"`                // 显示名称
	Status       PlatformStatus `json:"status" gorm:"index"`         // 平台状态
	CreatedAt    time.Time      `json:"created_at"`                  // 创建时间
	UpdatedAt    time.Time      `json:"updated_at"`                  // 更新时间
}

// PlatformConfig 平台配置实例
type PlatformConfig struct {
	ID          string         `json:"id" gorm:"primaryKey"`     // 配置ID
	ConfigName  string         `json:"config_name"`              // 配置名称
	PlatformID  string         `json:"platform_id" gorm:"index"` // 平台ID
	Environment string         `json:"environment"`              // 环境(production/test/sandbox)
	Status      PlatformStatus `json:"status" gorm:"index"`      // 配置状态
	Config      string         `json:"config" gorm:"type:text"`  // 配置参数(JSON)
	CreatedAt   time.Time      `json:"created_at"`               // 创建时间
	UpdatedAt   time.Time      `json:"updated_at"`               // 更新时间
}

// PaymentMethodStatus 支付方式状态
type PaymentMethodStatus string

const (
	PaymentMethodStatusActive   PaymentMethodStatus = "active"   // 激活
	PaymentMethodStatusInactive PaymentMethodStatus = "inactive" // 未激活
)

// PaymentMethod 支付方式
type PaymentMethod struct {
	ID          string              `json:"id" gorm:"primaryKey"`     // 方式ID
	PlatformID  string              `json:"platform_id" gorm:"index"` // 平台ID
	MethodName  string              `json:"method_name"`              // 方式名称
	DisplayName string              `json:"display_name"`             // 显示名称
	Status      PaymentMethodStatus `json:"status" gorm:"index"`      // 状态
	Config      string              `json:"config" gorm:"type:text"`  // 配置参数(JSON)
	CreatedAt   time.Time           `json:"created_at"`               // 创建时间
	UpdatedAt   time.Time           `json:"updated_at"`               // 更新时间
}

// PlatformConfigParams 平台配置参数结构
type PlatformConfigParams struct {
	// 微信支付配置
	WechatAppID    string `json:"wechat_app_id,omitempty"`    // 微信AppID
	WechatMchID    string `json:"wechat_mch_id,omitempty"`    // 微信商户号
	WechatAPIKey   string `json:"wechat_api_key,omitempty"`   // 微信API密钥
	WechatCertPath string `json:"wechat_cert_path,omitempty"` // 微信证书路径
	WechatKeyPath  string `json:"wechat_key_path,omitempty"`  // 微信私钥路径

	// 支付宝配置
	AlipayAppID      string `json:"alipay_app_id,omitempty"`      // 支付宝应用ID
	AlipayPrivateKey string `json:"alipay_private_key,omitempty"` // 支付宝私钥
	AlipayPublicKey  string `json:"alipay_public_key,omitempty"`  // 支付宝公钥
	AlipayGateway    string `json:"alipay_gateway,omitempty"`     // 支付宝网关

	// Apple Pay配置
	AppleMerchantID string `json:"apple_merchant_id,omitempty"` // Apple商户ID
	AppleCertPath   string `json:"apple_cert_path,omitempty"`   // Apple证书路径
	AppleKeyPath    string `json:"apple_key_path,omitempty"`    // Apple私钥路径

	// 通用配置
	NotifyURL string `json:"notify_url,omitempty"` // 回调地址
	ReturnURL string `json:"return_url,omitempty"` // 返回地址
	Sandbox   bool   `json:"sandbox,omitempty"`    // 是否沙箱环境
}
