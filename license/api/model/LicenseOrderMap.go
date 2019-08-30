package model

// LicenseOrderMap struct
type LicenseOrderMap struct {
	Model
	LicenseOrderID string     `gorm:"type:varchar(50);primary_key" json:"licenseOrderId,omitempty"`
	Order          Order      `gorm:"foreignkey:OrderID"`
	OrderID        string     `json:"orderId,omitempty"`
	LicenseKey     LicenseKey `gorm:"foreignkey:LicenseID"`
	LicenseID      string     `json:"licenseId,omitempty"`
}

func (l *LicenseOrderMap) TableName() string {
	return "LicenseOrderMap"
}
