package vo

//PartnerUser properties
type PartnerUser struct {
	UpID       string `json:"userId,omitempty"`
	PartnerID  string `json:"partnerId,omitempty"`
	AuthUserID string `json:"authUserId,omitempty"`
	UserRole   string `json:"userRole,omitempty"`
	Status     string `json:"status,omitempty"`
}
