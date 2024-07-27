import (
    "context"
    "encoding/json"

    amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
    bytes, err := JSON.Marshal(val)
    if err != nil {
        return err
    }   

    message = amqp.Publishing {
        ContentType: "application/json",
        Body: bytes,
    }

    ch.PublishWithContext(context.Background(), exchange, key, false, false, message)

    return nil
}
