package model

//Distribution struct
type Distribution struct {
	Model
	DistributionID string `gorm:"type:varchar(50);primary_key" json:"distributionId,omitempty"`
	//Tenant         Tenant `gorm:"foreignkey:TenantId"`
	TenantID string `json:"tenantId,omitempty"`

	Name                string `gorm:"type:varchar(100)" json:"name,omitempty"`
	Description         string `gorm:"type:text" json:"description,omitempty"`
	LicenseType         *int   `gorm:"default:0" json:"licensetype,omitempty"`
	Validity            string `json:"validity,omitempty"`
	MaxActivation       int64  `json:"maxActivation,omitempty"`
	ActivationPerDevice int64  `json:"activationPerDevice,omitempty"`
}

func (d *Distribution) TableName() string {
	return "Distribution"
}
