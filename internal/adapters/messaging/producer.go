package messaging

import (
	"encoding/json"

	"github.com/NavidKalashi/twitter/internal/core/domain/events"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	channel *amqp.Channel
}

func NewRabbitMQProducer(channel *amqp.Channel) ports.Producer {
	return &RabbitMQProducer{channel: channel}
}

func (rp *RabbitMQProducer) ProducerFeedEvents(feed events.Feed) error {
	payload, err := json.Marshal(feed)
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
			Body:        payload,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (rp *RabbitMQProducer) ProducerGestureEvents(createdGesture events.Gesture) error {
	body, err := json.Marshal(createdGesture)
	if err != nil {
		return err
	}

	err = rp.channel.Publish(
		"likes_exchange",
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