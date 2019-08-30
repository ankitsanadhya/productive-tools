package model

//Order struct
type Order struct {
	Model
	OrderID        string        `gorm:"type:varchar(50);primary_key" json:"productId,omitempty"`
	Distribution   Distribution  `gorm:"foreignkey:DistributionID"`
	DistributionID string        `json:"distributionId,omitempty"`
	Price          float64       `gorm:"type:double;default null" json:"price,omitempty"`
	TotalPrice     float64       `gorm:"type:double;default null" json:"totalPrice,omitempty"`
	TransactionID  string        `gorm:"type:varchar(50)" json:"transactionId,omitempty"`
	PaymentMethod  PaymentMethod `gorm:"foreignkey:PaymentID"`
	PaymentID      string        `json:"paymentId,omitempty"`
	LicenseCount   int           `gorm:"type:int;not null;default:0" json:"licenseCount,omitempty"`
}

func (o *Order) TableName() string {
	return "Order"
}
