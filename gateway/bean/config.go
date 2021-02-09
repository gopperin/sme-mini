package bean

import ()

// ServerConfig ServerConfig Struct
type ServerConfig struct {
	Port         string
	ViewLimit    int64
	APIAppendKey string
	APIMd5Key    string
	Release      string
	Version      string
	GrpcURI      string
	DesKey       string
}

// APIConfig APIConfig Struct
type APIConfig struct {
	AppendKey   string
	Md5Key      string
	IPWhiteList map[string]interface{}
}

// RedisConfig RedisConfig Struct
type RedisConfig struct {
	Addr     string
	Pwd      string
	PoolSize int
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
