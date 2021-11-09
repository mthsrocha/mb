package configs

import (
	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("Host", "0.0.0.0")
	viper.SetDefault("Port", 5012)

	viper.SetDefault("Database.Driver", "mysql")
	viper.SetDefault("Database.Host", "root:secret@/mercadobitcoin")

}

func InitConfig() {
	setDefaults()

	viper.BindEnv("Enviromnent", "ENV_NAME")
	viper.BindEnv("Host", "HOST")
	viper.BindEnv("Port", "PORT")
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.BindEnv("readTimeout", "READ_TIMEOUT")
	viper.BindEnv("writeTimeout", "WRITE_TIMEOUT")

	viper.BindEnv("Database.Driver", "DATABASE_DRIVER")
	viper.BindEnv("Database.Host", "DATABASE_HOST")

}
