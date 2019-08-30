package model

// Client struct
type Client struct {
	Model
	ClientID    string `gorm:"type:varchar(50);primary_key" json:"clientId,omitempty"`
	Name        string `gorm:"type:varchar(100)" json:"name,omitempty"`
	PublicKey   string `json:"publicKey,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}

func (c *Client) TableName() string {
	return "Client"
}
