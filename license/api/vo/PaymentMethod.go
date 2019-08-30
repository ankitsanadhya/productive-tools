package vo

// PaymentMethod struct
type PaymentMethod struct {
	PaymentID string `json:"paymentId,omitempty"`
	Name      string `json:"name,omitempty"`
	Detail    string `json:"detail,omitempty"`
}
