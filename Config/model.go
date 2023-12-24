package config

type config struct {
	Postgres postgres `mapstructure:",squash"`
}

type postgres struct {
	Port     string `mapstructure:"ENV_PORT"`
	HostName string `mapstructure:"ENV_HOST_NAME"`
	User     string `mapstructure:"ENV_USER"`
	Password string `mapstructure:"ENV_PASSWORD"`
	Dbname   string `mapstructure:"ENV_DBNAME"`
}
