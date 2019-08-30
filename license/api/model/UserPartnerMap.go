package model

// UserPartnerMap struct
type UserPartnerMap struct {
	Model
	UserPartnerID string  `gorm:"type:varchar(50);primary_key" json:"userPartnerId,omitempty"`
	Partner       Partner `gorm:"foreignkey:PartnerId"`
	PartnerID     string  `json:"partnerId,omitempty"`
	AppUserID     string  `gorm:"type:varchar(50);not null" json:"appUserId,omitempty"`
	UserRole      int     `gorm:"type:int" json:"userRole,omitempty"`
	Status        int     `gorm:"type:int" json:"status,omitempty"`
}

func (u *UserPartnerMap) TableName() string {
	return "UserPartnerMap"
}
