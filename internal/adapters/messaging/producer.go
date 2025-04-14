package messaging

import (
	"encoding/json"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	channel *amqp.Channel
}

func NewRabbitMQProducer(channel *amqp.Channel) ports.Producer {
	return &RabbitMQProducer{channel: channel}
}

func (rp *RabbitMQProducer) ProducerFeedEvents(createdTweet *models.Tweet) error {
	body, err := json.Marshal(createdTweet)
	if err != nil {
		return err
	}

	err = rp.channel.Publish(
		"feed_exchange",
		"", // routing key for fanout
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}