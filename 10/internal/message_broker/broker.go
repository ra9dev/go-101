package message_broker

import "context"

type (
	// Services: orders, logistics
	// Шаги обработки заказа:
	//  1) Оформление заказа
	//  2) Orders обработал и сохранил заказ
	//  3) Orders отправляет на доставку заказы в большом количестве в MessageBroker
	//  4) Logistics с нужной ему скоростью считывает заказы с MessageBroker и оформляет доставку

	// Cache invalidating
	// 1) S1, S2, S3 - знают друг о друге и кидают запросы на инвалидацию (P2P, etcd)
	// 2) MessageBroker - S1 кидает сообщение, а S2 и S3 подписаны на очередь сообщений, в очередь шлются команды

	MessageBroker interface {
		Connect(ctx context.Context) error
		Close() error

		Cache() CacheBroker // Сюда отправляются команды, связанные непосредственно с кэшом
	}

	BrokerWithClient interface {
		Connect(ctx context.Context, brokers []string) error
		Close() error
	}
)
