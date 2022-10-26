package config

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	AuthSvcUrl   string `mapstructure:"AUTH_SVC_URL"`
	AdminSvcUrl  string `mapstructure:"ADMIN_SVC_URL"`
	ItemSvcUrl   string `mapstructure:"ITEM_SVC_URL"`
	WalletSvcUrl string `mapstructure:"WALLET_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
