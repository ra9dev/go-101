package message_broker

type CacheBroker interface {
	BrokerWithClient
	Producer() CacheProducer
}

type CacheProducer interface {
	Remove(key interface{}) error
	Purge() error
}
