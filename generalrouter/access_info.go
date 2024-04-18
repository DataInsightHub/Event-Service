package generalrouter

import (
	"fmt"
	"os"
)

const (
	PUB_USER = "PUB_USER"
	PUB_PASS = "PUB_PASS"
	SUB_USER = "SUB_USER"
	SUB_PASS = "SUB_PASS"
)

type (
	AccessInfo struct {
		PublisherUsername  string
		PablisherPassword  string
		SubscriberUsername string
		SubscriberPassword string
	}
)

func LoadAccessInfo() (AccessInfo, error) {
	pubUser := os.Getenv(PUB_USER)
	if pubUser == "" {
		return AccessInfo{}, fmt.Errorf("could not get %v from the env file", PUB_USER)
	}

	pubPass := os.Getenv(PUB_PASS)
	if pubPass == "" {
		return AccessInfo{}, fmt.Errorf("could not get %v from the env file", PUB_PASS)
	}

	subUser := os.Getenv(SUB_USER)
	if subUser == "" {
		return AccessInfo{}, fmt.Errorf("could not get %v from the env file", SUB_USER)
	}

	subPass := os.Getenv(SUB_PASS)
	if subPass == "" {
		return AccessInfo{}, fmt.Errorf("could not get %v from the env file", SUB_USER)
	}

	return AccessInfo{
		PublisherUsername: pubUser,
		PablisherPassword: pubPass,
		SubscriberUsername: subUser,
		SubscriberPassword: subPass,
	}, nil
}
