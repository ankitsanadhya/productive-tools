package vo

//Tenant properties
type Tenant struct {
	TenantID       string          `json:"tenantId,omitempty"`
	TenantName     string          `json:"tenantName,omitempty"`
	Description    string          `json:"description,omitempty"`
	Products       []Product       `json:"products,omitempty"`
	Distributions  []Distribution  `json:"distributions,omitempty"`
	Users          []PartnerUser   `json:"users,omitempty"`
	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`
}
