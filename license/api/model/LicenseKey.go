package model

//LicenseKey struct
type LicenseKey struct {
	Model
	LicenseID  string `gorm:"type:varchar(50);primary_key" json:"licenseID,omitempty"`
	LicenseKey string `gorm:"type:varchar(100);unique_index" json:"licenseKey,omitempty"`
	Status     int    `gorm:"type:int" json:"status,omitempty"`
}

func (l *LicenseKey) TableName() string {
	return "LicenseKey"
}
