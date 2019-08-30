package vo

// OrderProduct struct
type OrderProduct struct {
	OrderProductID        string              `json:"orderProductId,omitempty"`
	DistributionProductVo DistributionProduct `json:"distributionProductObj,omitempty"`
}
