package model

// TanentPaymentMap struct
type TanentPaymentMap struct {
	Model
	TanentPaymentID string        `gorm:"type:varchar(50);primary_key" json:"tanentPaymentId,omitempty"`
	Tanent          Tanent        `gorm:"foreignkey:TanentID"`
	TanentID        string        `json:"tanentId,omitempty"`
	PaymentMethod   PaymentMethod `gorm:"foreignkey:PaymentID"`
	PaymentID       string        `json:"paymentId,omitempty"`
}

func (t *TanentPaymentMap) TableName() string {
	return "TanentPaymentMap"
}
