package publish

import "github.com/DataInsightHub/Event-Service/topic"

type (
	Publisher struct {
		topic.Topic `json:",inline" bson:",inline"`
		BodySchema string `json:"bodyschema" bson:"bodyschema"`
	}
)