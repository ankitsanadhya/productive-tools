package model

// TanentUserMap struct
type TanentUserMap struct {
	Model
	TanentUserID   string  `gorm:"type:varchar(50);primary_key" json:"tanentUserId,omitempty"`
	Tanent         Tanent  `gorm:"foreignkey:TanentID"`
	TanentID       string  `json:"tanentId,omitempty"`
	UserPartnerMap Product `gorm:"foreignkey:UserPartnerID"`
	UserPartnerID  string  `json:"userPartnerId,omitempty"`
}

func (t *TanentUserMap) TableName() string {
	return "TanentUserMap"
}
