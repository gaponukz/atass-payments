package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	RabbitmqUrl string
	Port        string
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
		RabbitmqUrl: os.Getenv("rabbitmqUrl"),
		Port:        os.Getenv("port"),
	}
}
