package redis

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/thienhaole92/vnd/logger"
)

type Publisher interface {
	Publish(topic string, msgs ...*message.Message) error
	Close() error
	Topic() string
}

type streamPublisher struct {
	*redisstream.Publisher
	topic string
}

func NewPublisher(config *PubConfig, topic string) (Publisher, error) {
	log := logger.GetLogger("NewPublisher")
	defer log.Sync()

	c, err := NewConfig()
	if err != nil {
		return nil, err
	}

	log.Infow("loaded redis publisher config")

	r, err := NewRedis(c)
	if err != nil {
		return nil, err
	}
	log.Infow("redis publisher connected")

	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client:     r,
			Marshaller: redisstream.DefaultMarshallerUnmarshaller{},
		},
		watermill.NewStdLogger(config.LoggerDebug, config.LoggerTrace),
	)

	if err != nil {
		return nil, err
	}

	return &streamPublisher{
		Publisher: publisher,
		topic:     topic,
	}, nil
}

func (r *streamPublisher) Topic() string {
	return r.topic
}
