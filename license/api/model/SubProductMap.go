package model

// SubProductMap struct
type SubProductMap struct {
	Model
	SubProductMapID string  `gorm:"type:varchar(50);primary_key" json:"subProductMapId,omitempty"`
	Product         Product `gorm:"foreignkey:ParentProductID"`
	//Product         Product `gorm:"foreignkey:SubProductID"`
	ParentProductID string `json:"parentProductId,omitempty"`
	SubProductID    string `json:"subProductId,omitempty"`
}

func (s *SubProductMap) TableName() string {
	return "SubProductMap"
}
