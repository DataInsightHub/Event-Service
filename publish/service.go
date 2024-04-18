package publish

import (
	"fmt"

	"github.com/DataInsightHub/Event-Service/logger"
	"github.com/DataInsightHub/Event-Service/topic"
)

type (
	Repo interface {
		InsertOrUpdatePublisher(Publisher) (Publisher, error)
	}

	TopicService interface {
		SaveTopic(topic.Topic) error
	}
)

type (
	Service struct {
		logger           logger.Logger
		repo             Repo
		topicService     TopicService
	}
)

func NewService(logger logger.Logger,
	repo Repo,
	topicService TopicService) *Service {
	return &Service{
		logger:           logger,
		repo:             repo,
		topicService:     topicService,
	}
}

func (s *Service) SignUpPublisher(publisher Publisher) error {
	if publisher.ID == "" {
		msg := "ID is empty. Please provide an ID for the publisher"
		s.logger.Error(msg)
		return fmt.Errorf("ID is empty. Please provide an ID for the publisher")
	}

	if publisher.Topic.Topic == "" {
		msg := "topic is empty. Please provide a topic for the publisher"
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	if publisher.BodySchema == "" {
		msg := "bodyschema is empty. Please provide a bodyschema for the publisher"
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	newPublisher, err := s.repo.InsertOrUpdatePublisher(publisher)
	if err != nil {
		msg := fmt.Sprintf("could not insert or update publisher. Error: %v", err)
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	err = s.topicService.SaveTopic(newPublisher.Topic)
	if err != nil {
		msg := fmt.Sprintf("could not insert topic. Error: %v", err)
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	return nil
}