package vo

import "time"

//Activation properties
type Activation struct {
	Date       time.Time `json:"date,omitempty"`
	Operation  string    `json:"operation,omitempty"`
	LicenceKey string    `json:"licenceKey,omitempty"`
	DeviceInfo string    `json:"deviceInfo,omitempty"`
	DeviceID   string    `json:"deviceId,omitempty"`
}
