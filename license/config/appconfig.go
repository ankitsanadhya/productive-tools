package config

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	HostName   string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

type Messages struct {
	Data string `json:"data"`
}
