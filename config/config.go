package config

var Conf Config

type Config struct {
	Redmine Redmine
}

type Redmine struct {
	Servers []Server `toml:"server"`
}

type Server struct {
	Endpoint   string `toml:"endpoint"`
	Apikey     string `toml:"apikey"`
	ProjectIDs []int  `toml:"project"`
	UserID     string `toml:"user_id"`
}
