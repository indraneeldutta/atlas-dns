package common

import (
	"github.com/spf13/viper"
)

func InitialiseApplication() {

}

// SetupEnvironment sets up the configs and environment for the application to start
func SetupEnvironment() error {
	err := viper.BindEnv("gopath")
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return err
}
