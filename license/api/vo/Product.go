package vo

// Product struct
type Product struct {
	ProductID     string       `json:"productId,omitempty"`
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	IsDefault     bool         `json:"isDefault,omitempty"`
	DefaultPrice  float64      `json:"defaultPrice,omitempty"`
	Specification string       `json:"spec,omitempty"`
	SubProducts   []SubProduct `json:"subProducts,omitempty"`
}
