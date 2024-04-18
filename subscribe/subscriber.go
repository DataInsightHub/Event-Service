package subscribe

import "github.com/DataInsightHub/Event-Service/topic"

type (
	Subscriber struct {
		topic.Topic `json:",inline" bson:",inline"`
		URL         string    `json:"url" bson:"url"`
		BasicAuth   BasicAuth `json:"auth" bson:"basic_auth"`
	}

	BasicAuth struct {
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
	}
)
