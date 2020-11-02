package storage

// Storage ...
type Storage struct {
	config *Config
}

// New ...
func New(config *Config) *Storage {
	return &Config{
		config: config
	}
}