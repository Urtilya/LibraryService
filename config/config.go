package config

type Config struct {
	Db       *DB
	GrpcPort int
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     int
	Db       string
	Timeout  int
}
