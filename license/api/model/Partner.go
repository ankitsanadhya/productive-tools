package model

//Partner struct
type Partner struct {
	Model
	PartnerID   string `gorm:"type:varchar(50);primary_key" json:"partnerId,omitempty"`
	Name        string `gorm:"type:varchar(100)" json:"name,omitempty"`
	Description string `gorm:"type:text" json:"description,omitempty"`
	Currency    string `gorm:"type:varchar(5);not null;default:'USD'" json:"currency,omitempty"`
	// Product     []Product
}

func (p *Partner) TableName() string {
	return "Partner"
}
