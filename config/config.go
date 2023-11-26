package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	APP struct {
		Name         string `envconfig:"APP_NAME" default:":dk-notification-service"`
		Port         string `envconfig:"APP_PORT" default:"9000"`
		Env          string `envconfig:"APP_ENV" default:"development"`
		RouteService string `envconfig:"APP_ROUTE_SERVICE" default:"/"`
	}
	DB struct {
		Host string `envconfig:"DB_HOST" required:"true" default:"hms.db"`
		Port string `envconfig:"DB_PORT"`
		User string `envconfig:"DB_USER"`
		Pass string `envconfig:"DB_PASS"`
		Name string `envconfig:"DB_NAME"`
	}
}

var envFile = ".env"

func loadEnv(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

func LoadConfig() (*Configuration, error) {
	if err := loadEnv(envFile); err != nil {
		return nil, err
	}

	config := new(Configuration)
	if err := envconfig.Process("APP", config); err != nil {
		return nil, err
	}
	return config, nil
}

func GetEnvByKey(key string) string {
	err := loadEnv(envFile)
	if err != nil {
		// log.Fatal("Error loading .env file", err)
	}

	return os.Getenv(key)
}
