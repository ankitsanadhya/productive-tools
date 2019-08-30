package vo

//SubProduct struct
type SubProduct struct {
	SubProductID  string  `json:"subProductId,omitempty"`
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	IsDefault     bool    `json:"isDefault,omitempty"`
	DefaultPrice  float64 `json:"defaultPrice,omitempty"`
	Specification string  `json:"spec,omitempty"`
}
