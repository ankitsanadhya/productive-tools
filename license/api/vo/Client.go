package vo

// Client struct
type Client struct {
	ClientID    string `json:"clientId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PublicKey   string `json:"-"`
	URL         string `json:"url,omitempty"`
}
