package models

import "time"

// PaymentStatus 支付状态
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"   // 待支付
	PaymentStatusPaid      PaymentStatus = "paid"      // 已支付
	PaymentStatusFailed    PaymentStatus = "failed"    // 支付失败
	PaymentStatusCancelled PaymentStatus = "cancelled" // 已取消
	PaymentStatusRefunded  PaymentStatus = "refunded"  // 已退款
)

// PaymentPlatform 支付平台
type PaymentPlatform string

const (
	PaymentPlatformWechat PaymentPlatform = "wechat" // 微信支付
	PaymentPlatformAlipay PaymentPlatform = "alipay" // 支付宝
	PaymentPlatformApple  PaymentPlatform = "apple"  // Apple Pay
)

// Payment 支付订单
type Payment struct {
	ID            string          `json:"id" gorm:"primaryKey"`   // 支付ID
	OrderID       string          `json:"order_id" gorm:"index"`  // 订单ID
	AppID         string          `json:"app_id" gorm:"index"`    // 应用ID
	Platform      PaymentPlatform `json:"platform" gorm:"index"`  // 支付平台
	Amount        int64           `json:"amount"`                 // 支付金额(分)
	Currency      string          `json:"currency"`               // 货币类型
	Subject       string          `json:"subject"`                // 订单标题
	Body          string          `json:"body"`                   // 订单描述
	Status        PaymentStatus   `json:"status" gorm:"index"`    // 支付状态
	PaymentURL    string          `json:"payment_url"`            // 支付链接
	QRCode        string          `json:"qr_code"`                // 二维码
	NotifyURL     string          `json:"notify_url"`             // 回调地址
	ReturnURL     string          `json:"return_url"`             // 返回地址
	Extra         string          `json:"extra" gorm:"type:text"` // 扩展参数(JSON)
	PlatformTxnID string          `json:"platform_txn_id"`        // 平台交易号
	PaidAt        *time.Time      `json:"paid_at"`                // 支付时间
	CreatedAt     time.Time       `json:"created_at"`             // 创建时间
	UpdatedAt     time.Time       `json:"updated_at"`             // 更新时间
}

// RefundStatus 退款状态
type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"   // 待退款
	RefundStatusSuccess   RefundStatus = "success"   // 退款成功
	RefundStatusFailed    RefundStatus = "failed"    // 退款失败
	RefundStatusCancelled RefundStatus = "cancelled" // 已取消
)

// Refund 退款订单
type Refund struct {
	ID               string          `json:"id" gorm:"primaryKey"`    // 退款ID
	PaymentID        string          `json:"payment_id" gorm:"index"` // 支付ID
	OrderID          string          `json:"order_id" gorm:"index"`   // 订单ID
	AppID            string          `json:"app_id" gorm:"index"`     // 应用ID
	Platform         PaymentPlatform `json:"platform" gorm:"index"`   // 支付平台
	RefundAmount     int64           `json:"refund_amount"`           // 退款金额(分)
	RefundReason     string          `json:"refund_reason"`           // 退款原因
	Status           RefundStatus    `json:"status" gorm:"index"`     // 退款状态
	PlatformRefundID string          `json:"platform_refund_id"`      // 平台退款号
	RefundedAt       *time.Time      `json:"refunded_at"`             // 退款时间
	CreatedAt        time.Time       `json:"created_at"`              // 创建时间
	UpdatedAt        time.Time       `json:"updated_at"`              // 更新时间
}
