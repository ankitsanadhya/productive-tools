package vo

//License properties
type License struct {
	LicenceKey    string       `json:"licenceKey,omitempty"`
	Acitations    []Activation `json:"acitations,omitempty"`
	LicenceType   string       `json:"licenceType,omitempty"`
	Products      []Product    `json:"products,omitempty"`
	Validity      int          `json:"validity,omitempty"`
	MaxActivation int          `json:"maxActivation,omitempty"`
	OrderID       string       `json:"orderId,omitempty"`
}
