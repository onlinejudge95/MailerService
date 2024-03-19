package mailer

import "github.com/spf13/viper"

type Config struct {
	Environment string `mapstructure:"ENV"`
	BindAddress string `mapstructure:"BINDADDRESS"`
	Identity    string `mapstructure:"IDENTITY"`
	Sender      string `mapstructure:"SENDER"`
	Password    string `mapstructure:"PASSWORD"`
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
	Recipient   string `mapstructure:"RECIPIENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)

	return
}
