package model

// TanentProductMap struct
type TanentProductMap struct {
	Model
	TanentProductID string  `gorm:"type:varchar(50);primary_key" json:"tanentProductId,omitempty"`
	Tanent          Tanent  `gorm:"foreignkey:TanentID"`
	TanentID        string  `json:"tanentId,omitempty"`
	Product         Product `gorm:"foreignkey:ProductID"`
	ProductID       string  `json:"productID,omitempty"`
}

func (t *TanentProductMap) TableName() string {
	return "TanentPaymentMap"
}
