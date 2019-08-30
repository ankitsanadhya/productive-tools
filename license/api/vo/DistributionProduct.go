package vo

//DistributionProduct struct
type DistributionProduct struct {
	DistributionProductID string  `json:"distributionproductId,omitempty"`
	Product               Product `json:"product,omitempty"`
	Price                 float64 `json:"price,omitempty"`
}
