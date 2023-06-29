package cmd

import "time"

type Config struct {
	Authorize struct {
		UserId         int       `mapstructure:"user_id"`
		OrganizationId int       `mapstructure:"organization_id"`
		AccessToken    int       `mapstructure:"access_token"`
		RefreshToken   int       `mapstructure:"refresh_token"`
		Expiry         time.Time `mapstructure:"expiry"`
	}
	Datastore struct {
		Driver string `mapstructure:"driver"`
		Source string `mapstructure:"source"`
	}
}
