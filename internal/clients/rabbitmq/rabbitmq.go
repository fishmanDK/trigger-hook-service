package rabbitmq

import (
	"encoding/json"
	"github.com/wagslane/go-rabbitmq"
	"log"
)

type RabbitMQPublisher struct {
	publisher *rabbitmq.Publisher
}

func NewRabbitMQClient() (*RabbitMQPublisher, error) {
	const op = "rabbitmq.NewNUTS"

	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost:5672/",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	//defer publisher.Close()

	return &RabbitMQPublisher{
		publisher: publisher,
	}, nil
}

type Message struct {
	BannerID  int64 `json:"bannerID,omitempty"`
	TagID     int64 `json:"tagID,omitempty"`
	FeatureID int64 `json:"featureID,omitempty"`
}

func (rq *RabbitMQPublisher) PublishMessage(message Message) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = rq.publisher.Publish(
		messageBytes,
		[]string{"my_routing_key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("events"),
	)
	if err != nil {
		log.Println(err)
	}
	return nil
}
