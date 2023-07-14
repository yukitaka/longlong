package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Authorize struct {
		UserId         int       `mapstructure:"user_id" yaml:"user_id"`
		OrganizationId int       `mapstructure:"organization_id" yaml:"organization_id"`
		AccessToken    string    `mapstructure:"access_token" yaml:"access_token"`
		RefreshToken   string    `mapstructure:"refresh_token" yaml:"refresh_token"`
		Expiry         time.Time `mapstructure:"expiry" yaml:"expiry"`
	} `mapstructure:"authorize" yaml:"authorize"`
	Datastore struct {
		Driver string `mapstructure:"driver"`
		Source string `mapstructure:"source"`
	} `mapstructure:"datastore" yaml:"datastore"`
}

func LoadFromFile(name, fileType, path string) (Config, error) {
	var conf Config
	viper.SetConfigName(name)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		}
	}
	if err := viper.Unmarshal(&conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}

func (c *Config) StoreAuth(accessToken, refreshToken string, expiry time.Time) {
	c.Authorize.AccessToken = accessToken
	c.Authorize.RefreshToken = refreshToken
	c.Authorize.Expiry = expiry

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/llctl")
	viper.Set("authorize", c.Authorize)
	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}
}
