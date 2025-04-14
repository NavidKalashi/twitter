package rabbitmq

import (
	"fmt"

	"github.com/NavidKalashi/twitter/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func InitRabbitMQ(cfg *config.Config) (*RabbitMQ, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	if err := declareExchangeAndQueue(ch, "likes_exchange", "like_events"); err != nil {
		return nil, err
	}

	if err := declareExchangeAndQueue(ch, "feed_exchange", "feed_events"); err != nil {
		return nil, err
	}

	fmt.Println("RabbitMQ connected. Exchanges and queues declared and bound.")
	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func declareExchangeAndQueue(ch *amqp.Channel, exchangeName, queueName string) error {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange %s: %w", exchangeName, err)
	}

	queue, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", queueName, err)
	}

	err = ch.QueueBind(
		queue.Name,   // queue name
		"",           // routing key (ignored for fanout)
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue %s to exchange %s: %w", queueName, exchangeName, err)
	}

	return nil
}

func (r *RabbitMQ) Close() error {
    if err := r.channel.Close(); err != nil {
        return fmt.Errorf("failed to close channel: %w", err)
    }
    if err := r.conn.Close(); err != nil {
        return fmt.Errorf("failed to close connection: %w", err)
    }
    return nil
}

func (r *RabbitMQ) GetRabbit() *amqp.Channel {
	return r.channel
}