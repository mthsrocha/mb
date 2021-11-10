package configs

import (
	"github.com/spf13/viper"
)

type configuration struct {
	Enviromnent string

	Database struct {
		Host    string
		Dbname  string
		User    string
		Passw   string
		Disable string
	}
	Server struct {
		Httport string
	}
}

var Configuration configuration

func setDefaults() {
	viper.SetDefault("Server.Httport", 5012)

	viper.SetDefault("Database.Host", "127.0.0.1")
	viper.SetDefault("Database.Dbname", "izanami")
	viper.SetDefault("Database.User", "root")
	viper.SetDefault("Database.Passw", "root")
	viper.SetDefault("Database.Disable", "disable")

	viper.SetDefault("Enviromnent", "Dev")

}

func InitConfig() {
	setDefaults()
	viper.AutomaticEnv()

	viper.BindEnv("Enviromnent", "ENV_NAME")
	viper.BindEnv("Host", "HOST")
	viper.BindEnv("Port", "PORT")
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.BindEnv("readTimeout", "READ_TIMEOUT")
	viper.BindEnv("writeTimeout", "WRITE_TIMEOUT")

	viper.BindEnv("Database.Driver", "DATABASE_DRIVER")
	viper.BindEnv("Database.Host", "DATABASE_HOST")

}
