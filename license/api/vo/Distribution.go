package vo

// Distribution struct
type Distribution struct {
	DistributionID      string                `json:"patnerId,omitempty"`
	Name                string                `json:"name,omitempty"`
	DistributionProduct []DistributionProduct `json:"distributionProduct,omitempty"`
	Description         string                `json:"description,omitempty"`
	LicenseType         string                `json:"licenseType,omitempty"`
	LicenseKey          []string              `json:"licenseKey,omitempty"`
	Validity            string                `json:"validity,omitempty"`
	MaxActivation       int                   `json:"maxActivation,omitempty"`
	ActivationPerDevice int                   `json:"validity.activationPerDevice,omitempty"`
}
