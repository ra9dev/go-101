package message_broker

import "context"

type (
	MessageBroker interface {
		Connect(ctx context.Context) error
		Close() error

		Cache() CacheBroker
	}

	BrokerWithClient interface {
		Connect(ctx context.Context, brokers []string) error
		Close() error
	}
)
