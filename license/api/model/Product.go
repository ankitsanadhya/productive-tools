package model

//Product struct
type Product struct {
	Model
	ProductID     string  `gorm:"type:varchar(50);primary_key" json:"productId,omitempty"`
	Partner       Partner `gorm:"foreignkey:PartnerID"`
	PartnerID     string  `gorm:"type:varchar(50)" json:"partnerId,omitempty"`
	Name          string  `gorm:"type:varchar(100)" json:"name,omitempty"`
	Description   string  `gorm:"type:text" json:"description,omitempty"`
	DefaultPrice  float64 `gorm:"type:double;default null" json:"defaultPrice,omitempty"`
	IsDefault     bool    `gorm:"type:tinyint(1);not null;default:0" json:"isDefault,omitempty"`
	Specification string  `gorm:"type:text" json:"specifications,omitempty"`
}

func (p *Product) TableName() string {
	return "Product"
}
