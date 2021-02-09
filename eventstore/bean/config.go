package bean

import ()

// ServerConfig ServerConfig Struct
type ServerConfig struct {
	GrpcPort string
}

// DBConfig DBConfig Struct
type DBConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}

// LiftConfig LiftConfig
type LiftConfig struct {
	Addrs     []string
	Partition int32
	Subjects  map[string]interface{}
}
