package vo

// Partner struct
type Partner struct {
	PartnerID     string          `json:"partnerId,omitempty"`
	Name          string          `json:"name,omitempty"`
	Users         []PartnerUser   `json:"users,omitempty"`
	Description   string          `json:"description,omitempty"`
	Tenants       []Tenant        `json:"tenants,omitempty"`
	Product       []Product       `json:"product,omitempty"`
	Currency      string          `json:"currency,omitempty"`
	PaymentMethod []PaymentMethod `json:"patmentMethods,omitempty"`
}
