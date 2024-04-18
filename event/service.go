package event

import (
	"fmt"
	"time"

	"github.com/DataInsightHub/Event-Service/logger"
	"github.com/DataInsightHub/Event-Service/subscribe"
	"github.com/DataInsightHub/Go-Helper/fp"
)

type (
	Repo interface {
		InsertEvent(Event) (Event, error)
	}

	SubscribeService interface {
		GetAllSubscriberForTopic(string) ([]subscribe.Subscriber, error)
		SendEvent(subscribe.Subscriber, string) error
	}

	TopicService interface {
		GetAllUniqueTopics() ([]string, error)
	}
)

type (
	Service struct {
		logger           logger.Logger
		repo             Repo
		subscribeService SubscribeService
		topicService     TopicService
	}
)

func NewService(logger logger.Logger,
	repo Repo,
	subscribeService SubscribeService,
	topicService TopicService) *Service {
		return &Service{
			logger: logger,
			repo: repo,
			subscribeService: subscribeService,
			topicService: topicService,
		}
	}

func (s *Service) PublishEvent(event Event) error {
	topics, err := s.topicService.GetAllUniqueTopics()
	if err != nil {
		msg := fmt.Sprintf("could not get all unique topics. Error: %v", err)
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	if !fp.Contains(topics, event.Topic.Topic) {
		msg := fmt.Sprintf("the topic: %v doesnt exists, please sign up first to publish events for this topic", event.Topic.Topic)
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	if event.ContextID == "" {
		msg := "please provide a context id for the event"
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	if event.Body == "" {
		msg := "please provide a body for the event"
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	event.Time = time.Now()
	savedEvent, err := s.repo.InsertEvent(event)
	if err != nil {
		msg := fmt.Sprintf("could not insert event. Error: %v", err)
		s.logger.Error(msg)
		return fmt.Errorf(msg)
	}

	go s.sendEventToSubscriber(savedEvent)

	return nil
}

func (s *Service) sendEventToSubscriber(event Event) {
	subscribers, err := s.subscribeService.GetAllSubscriberForTopic(event.Topic.Topic)
	if err != nil {
		msg := fmt.Sprintf("could not get subscriber for event: %v. Error: %v", event.Topic.Topic, err)
		s.logger.Error(msg)
	}

	fp.ForEachParallel(subscribers, func(_ int, subscriber subscribe.Subscriber) {
		err := s.subscribeService.SendEvent(subscriber, event.Body)
		if err != nil {
			msg := fmt.Sprintf("could not send event to subscriber: %v. Error: %v", event.ID, err)
			s.logger.Error(msg)
		}
	})
}
