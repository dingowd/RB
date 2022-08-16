package config

type Config struct {
	Logger  LoggerConf
	DSN     string
	HTTPSrv string
}

type LoggerConf struct {
	Level   string
	LogFile string
}

func NewConfig() *Config {
	return &Config{}
}

func Default() *Config {
	return &Config{
		Logger: LoggerConf{
			Level:   "INFO",
			LogFile: "./log.txt",
		},
		DSN:     "user=postgres dbname=WB sslmode=disable password=masterkey",
		HTTPSrv: "127.0.0.1:3541",
	}
}
