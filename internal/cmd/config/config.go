package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Authorize struct {
		UserId         int       `mapstructure:"user_id"`
		OrganizationId int       `mapstructure:"organization_id"`
		AccessToken    string    `mapstructure:"access_token"`
		RefreshToken   string    `mapstructure:"refresh_token"`
		Expiry         time.Time `mapstructure:"expiry"`
	}
	Datastore struct {
		Driver string `mapstructure:"driver"`
		Source string `mapstructure:"source"`
	}
}

func (c *Config) Store(accessToken, refreshToken string, expiry time.Time) {
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
