package config

// Configurations exported
type Configurations struct {
	Server          ServerConfigurations
	Database        DatabaseConfigurations
	JWT_SECRET      string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBHost     string
	DBDriver   string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}
