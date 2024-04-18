package topic

type (
	Topic struct {
		ID string `json:"id" bson:"id"`
		Topic string `json:"topic" bson:"topic"`
	}
)