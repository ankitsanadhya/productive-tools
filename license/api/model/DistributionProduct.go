package model

//DistributionProduct struct
type DistributionProduct struct {
	Model
	DistributionProductID string       `gorm:"type:varchar(50);primary_key" json:"distributionProductId,omitempty"`
	Distribution          Distribution `gorm:"foreignkey:DistributionID"`
	DistributionID        string       `json:"distributionID,omitempty"`

	Product   Product `gorm:"foreignkey:ProductID"`
	ProductID string  `json:"productId,omitempty"`

	Price float32 `json:"price,omitempty"`
}

func (d *DistributionProduct) TableName() string {
	return "DistributionProduct"
}
