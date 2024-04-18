package event

import (
	"time"

	"github.com/DataInsightHub/Event-Service/topic"
)

type (
	Event struct {
		topic.Topic `json:",inline" bson:",inline"`
		ContextID   string    `json:"context_id" bson:"context_id"`
		Time        time.Time `json:"-" bson:"time"`
		Body        string    `json:"body" bson:"body"`
	}
)
