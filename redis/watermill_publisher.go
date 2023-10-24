package redis

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/thienhaole92/vnd/publisher"
)

type WatermillPublisher struct {
	pub Publisher
}

func NewWatermillPublisher(pub Publisher) (publisher.Publisher, error) {
	return &WatermillPublisher{
		pub: pub,
	}, nil
}

func (w *WatermillPublisher) Close() error {
	return w.pub.Close()
}

func (w *WatermillPublisher) PublishMessage(messages ...string) error {
	msgs := make([]*message.Message, 0)

	for _, v := range messages {
		msgs = append(msgs, message.NewMessage(watermill.NewUUID(), []byte(v)))
	}

	return w.pub.Publish(
		w.pub.Topic(),
		msgs...,
	)
}

func (w *WatermillPublisher) Topic() string {
	return w.pub.Topic()
}
