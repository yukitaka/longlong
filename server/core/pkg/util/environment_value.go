package util

import (
	"github.com/joho/godotenv"
	"os"
)

type EnvironmentValue struct{}

func NewEnvironmentValue() (*EnvironmentValue, error) {
	return &EnvironmentValue{}, godotenv.Load("../.env")
}

func GetEnvironmentValue(key string) (string, error) {
	ev, err := NewEnvironmentValue()
	if err != nil {
		return "", err
	}

	return ev.Get(key), nil
}

func (ev *EnvironmentValue) Get(key string) string {
	return os.Getenv(key)
}
