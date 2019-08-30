package model

import "time"

//Partner struct
type Activation struct {
	Model
	ActivationID string     `gorm:"type:varchar(50);primary_key" json:"activationId,omitempty"`
	Operation    *int       `gorm:"default:0" json:"operation,omitempty"`
	LicenseKey   LicenseKey `gorm:"foreignkey:LicenseID"`
	LicenseID    string     `json:"licenseId,omitempty"`
	IP           string     `json:"Ip,omitempty"`
	DeviceInfo   string     `json:"deviceInfo,omitempty"`
	DeviceID     string     `json:"deviceId,omitempty"`
	Date         time.Time  `json:"date,omitempty"`
}

func (a *Activation) TableName() string {
	return "Activation"
}
