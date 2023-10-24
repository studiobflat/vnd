package redis

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

type Subscriber interface {
	Close() error
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
	GroupID() string
	Topic() string
}

type streamSubscriber struct {
	*redisstream.Subscriber
	topic   string
	groupID string
}

func NewSubscriber(config *SubConfig, client redis.UniversalClient, topic string) (Subscriber, error) {
	subscriber, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client:        client,
			Unmarshaller:  redisstream.DefaultMarshallerUnmarshaller{},
			ConsumerGroup: config.ConsumerGroup,
		},
		watermill.NewStdLogger(config.LoggerDebug, config.LoggerTrace),
	)

	if err != nil {
		return nil, err
	}

	return &streamSubscriber{
		topic:      topic,
		Subscriber: subscriber,
	}, nil
}

func (r *streamSubscriber) GroupID() string {
	return r.groupID
}

func (r *streamSubscriber) Topic() string {
	return r.topic
}
