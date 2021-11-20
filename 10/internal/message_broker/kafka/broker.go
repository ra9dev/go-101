package kafka

import (
	"context"
	lru "github.com/hashicorp/golang-lru"
	"lecture-10/internal/message_broker"
)

type Broker struct {
	brokers  []string // адреса кластера кафки, к которым мы подклюаемся
	clientID string   // айди каждой реплики сервиса, чтобы она могла читать сообщения со всеми вместе параллельно

	cacheBroker message_broker.CacheBroker // субброкер кэша
	cache       *lru.TwoQueueCache         // ссылка на кэш приложения
}

func NewBroker(brokers []string, cache *lru.TwoQueueCache, clientID string) message_broker.MessageBroker {
	return &Broker{brokers: brokers, cache: cache, clientID: clientID}
}

func (b *Broker) Connect(ctx context.Context) error {
	brokers := []message_broker.BrokerWithClient{b.Cache()}

	for _, broker := range brokers {
		if err := broker.Connect(ctx, b.brokers); err != nil {
			return err
		}
	}

	return nil
}

func (b *Broker) Close() error {
	brokers := []message_broker.BrokerWithClient{b.Cache()}
	for _, broker := range brokers {
		if err := broker.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (b *Broker) Cache() message_broker.CacheBroker {
	if b.cacheBroker == nil {
		b.cacheBroker = NewCacheBroker(b.cache, b.clientID)
	}

	return b.cacheBroker
}
