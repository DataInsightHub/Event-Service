package logger

import (
	"fmt"
	"os"
)

const (
	LOGGER = "LOGGER"
)

func BuildLogger() (Logger, error) {
	logger := os.Getenv(LOGGER)
	if logger == "" {
		return NewLogrusLogger(), fmt.Errorf("could not get %v from the env file", LOGGER)
	}

	switch logger {
	case "logrus":
		return NewLogrusLogger(), nil
	default:
		return NewLogrusLogger(), nil
	}
}
