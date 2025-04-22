package service

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/events"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type ProduceService struct {
	produce ports.Producer
}

func NewProduceService(produce ports.Producer) *ProduceService {
	return &ProduceService{produce: produce}
}

func (ps *ProduceService) ProduceFeed(feed events.Feed) error {
	return ps.produce.ProducerFeedEvents(feed)
}

func (ps *ProduceService) ProducerGesture(createdGesture events.Gesture) error {
	return ps.produce.ProducerGestureEvents(createdGesture)
}