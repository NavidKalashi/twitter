package service

import (
	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type ProduceService struct {
	produce ports.Producer
}

func NewProduceService(produce ports.Producer) *ProduceService {
	return &ProduceService{produce: produce}
}

func (ps *ProduceService) Produce(CreatedTweet *models.Tweet) error {
	return ps.produce.ProducerFeedEvents(CreatedTweet)
}