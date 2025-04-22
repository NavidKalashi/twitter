package messaging

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/NavidKalashi/twitter/internal/core/domain/events"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	channel *amqp.Channel
}

func NewRabbitMQConsumer(channel *amqp.Channel) ports.Consume {
	return &RabbitMQConsumer{channel: channel}
}


func (rc *RabbitMQConsumer) ConsumeFeedEvents(handler func(events.Feed)) error {
	msgs, err := rc.channel.Consume(
		"feed_events",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume feed events: %w", err)
	}

	go func() {
		for msg := range msgs {
			var e events.Feed
			if err := json.Unmarshal(msg.Body, &e); err != nil {
				log.Printf("error decoding feed event: %v", err)
				continue
			}
			log.Println("e:", e)
			handler(e)
		}
	}()

	return nil
}

func (rc *RabbitMQConsumer) ConsumeGestureEvents(handler func(events.Gesture)) error {
	msgs, err := rc.channel.Consume(
		"like_events", // queue
		"",            // consumer
		true,          // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume like events: %w", err)
	}

	go func() {
		for msg := range msgs {
			var e events.Gesture
			if err := json.Unmarshal(msg.Body, &e); err != nil {
				log.Printf("error decoding like event: %v", err)
				continue
			}
			handler(e)
		}
	}()

	return nil
}
