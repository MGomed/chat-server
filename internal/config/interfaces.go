package config

// APIConfig is grpc config interface
type APIConfig interface {
	Address() string
}

// AccessConfig is access service config interface
type AccessConfig interface {
	Address() string
}

// PgConfig is postgres config interface
type PgConfig interface {
	DSN() string
}
