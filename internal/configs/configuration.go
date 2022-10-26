package configs

import "github.com/spf13/viper"

func LoadConfig(profile string) error {
	viper.SetConfigName(profile)
	viper.SetConfigType("json")

	if profile == "local" {
		viper.AddConfigPath("./configs")
	} else {
		viper.AutomaticEnv()
	}

	return viper.ReadInConfig()
}
