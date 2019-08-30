package model

// OrderProductMap struct
type OrderProductMap struct {
	Model
	OrderProductID        string              `gorm:"type:varchar(50);primary_key" json:"orderProductId,omitempty"`
	Order                 Order               `gorm:"foreignkey:OrderID"`
	OrderID               string              `json:"orderID,omitempty"`
	DistributionProduct   DistributionProduct `gorm:"foreignkey:DistributionProductID"`
	DistributionProductID string              `json:"distributionProductID,omitempty"`
}

func (o *OrderProductMap) TableName() string {
	return "OrderProductMap"
}
