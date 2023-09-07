package config

type Config struct {
	Db       *DB `yaml:"DB"`
	GrpcPort int `yaml:"GrpcPort"`
}

type DB struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Db       string `yaml:"Db"`
	Timeout  int    `yaml:"Timeout"`
}
