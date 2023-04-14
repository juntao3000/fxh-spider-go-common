package baseCommon

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
)

func GetAppConfig(cfg any) error {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("/app/")
	v.AddConfigPath("/app/config/")
	v.AddConfigPath("/etc/spider/")
	v.AddConfigPath("d:\\temp\\")
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	decoderCfgOption := func(cfg *mapstructure.DecoderConfig) {
		cfg.TagName = "json"
	}

	return v.Unmarshal(cfg, decoderCfgOption)
}

func GetBaseConfig(cfg any) error {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("/app/")
	v.AddConfigPath("/app/config/")
	v.AddConfigPath("/etc/spider/")
	v.AddConfigPath("d:\\temp\\")
	v.SetConfigName("base")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	decoderCfgOption := func(cfg *mapstructure.DecoderConfig) {
		cfg.TagName = "json"
	}

	return v.Unmarshal(cfg, decoderCfgOption)
}

func IsDebug() bool {
	_, ok := os.LookupEnv("DEBUG")
	return ok
}
