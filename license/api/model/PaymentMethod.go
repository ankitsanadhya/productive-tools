package model

// PaymentMethod struct
type PaymentMethod struct {
	Model
	PaymentID string  `gorm:"type:varchar(50);primary_key" json:"paymentId,omitempty"`
	Partner   Partner `gorm:"foreignkey:PartnerID"`
	PartnerID string  `json:"partnerId,omitempty"`
	Name      string  `gorm:"type:varchar(100)" json:"name,omitempty"`
	Detail    string  `gorm:"type:text" json:"detail,omitempty"`
}

func (p *PaymentMethod) TableName() string {
	return "PaymentMethod"
}
