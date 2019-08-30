package model

// Tanent struct
type Tanent struct {
	Model
	TanentID    string  `gorm:"type:varchar(50);primary_key" json:"tanentId,omitempty"`
	Partner     Partner `gorm:"foreignkey:PartnerID"`
	PartnerID   string  `json:"partnerId,omitempty"`
	Name        string  `gorm:"type:varchar(100)" json:"name,omitempty"`
	Description string  `gorm:"type:text" json:"description,omitempty"`
}

func (t *Tanent) TableName() string {
	return "Tanent"
}
