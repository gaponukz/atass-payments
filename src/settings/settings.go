package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	RabbitmqUrl      string
	Port             string
	PostgresHost     string `json:"postgresHost"`
	PostgresUser     string `json:"postgresUser"`
	PostgresPassword string `json:"postgresPassword"`
	PostgresDbname   string `json:"postgresDbname"`
	PostgresPort     string `json:"postgresPort"`
	PostgresSslmode  string `json:"postgresSslmode"`
}

type dotEnvSettings struct{}

func NewDotEnvSettings() dotEnvSettings {
	return dotEnvSettings{}
}

func (sts dotEnvSettings) Load() Settings {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: can not load dot env: %v\n", err)
	}

	return Settings{
		RabbitmqUrl:      os.Getenv("rabbitmqUrl"),
		Port:             os.Getenv("port"),
		PostgresHost:     os.Getenv("postgresHost"),
		PostgresUser:     os.Getenv("postgresUser"),
		PostgresPassword: os.Getenv("postgresPassword"),
		PostgresDbname:   os.Getenv("postgresDbname"),
		PostgresPort:     os.Getenv("postgresPort"),
		PostgresSslmode:  os.Getenv("postgresSslmode"),
	}
}
