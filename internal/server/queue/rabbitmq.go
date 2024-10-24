package queue

import (
    "encoding/json"
    amqp "github.com/rabbitmq/amqp091-go"
    "github.com/device-sec/internal/models"
)

type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    _, err = ch.QueueDeclare(
        "commands",
        true,
        false,
        false,
        false,
        nil,
    )

    return &RabbitMQ{
        conn:    conn,
        channel: ch,
    }, nil
}

func (r *RabbitMQ) PublishCommand(cmd *models.Command) error {
    body, err := json.Marshal(cmd)
    if err != nil {
        return err
    }

    return r.channel.Publish(
        "",
        "commands",
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:       body,
        },
    )
}

func (r *RabbitMQ) Close() error {
    if err := r.channel.Close(); err != nil {
        return err
    }
    return r.conn.Close()
}