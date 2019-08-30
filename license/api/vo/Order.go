package vo

// Order struct
type Order struct {
	OrderID        string         `json:"orderId,omitempty"`
	Price          float64        `json:"price,omitempty"`
	TotalPrice     float64        `json:"totalPrice,omitempty"`
	CouponID       string         `json:"couponId,omitempty"`
	DistributionVo Distribution   `json:"distributionObj,omitempty"`
	PaymentVo      PaymentMethod  `json:"paymentObj,omitempty"`
	OrderProducts  []OrderProduct `json:"orderProducts,omitempty"`
	TransactionID  string         `json:"transactionId,omitempty"`
	LicenceCount   int            `json:"licenceCount,omitempty"`
}
