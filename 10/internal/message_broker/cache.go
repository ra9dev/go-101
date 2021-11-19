package message_broker

type CacheBroker interface {
	BrokerWithClient
	Remove(key interface{}) error
	Purge() error
}
