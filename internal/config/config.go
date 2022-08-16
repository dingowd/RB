package config

type Config struct {
	HTTPSrv   string
	CacheTick int
	Logger    LoggerConf
	DB        Mongo
}

type LoggerConf struct {
	Level   string
	LogFile string
}

type Mongo struct {
	DSN        string
	DB         string
	Collection string
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
		HTTPSrv:   "127.0.0.1:9091",
		CacheTick: 10,
		DB: Mongo{
			DSN:        "mongodb://127.0.0.1",
			DB:         "RB",
			Collection: "example",
		},
	}
}
