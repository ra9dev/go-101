Брокеры сообщений:

- Kafka - deployments/kafka/docker-compose.yml
- Kafka Offset Explorer - https://kafkatool.com/download.html
- RabbitMQ - https://www.rabbitmq.com/download.html
- Nats
  - Installation - https://docs.nats.io/nats-server/installation
  - Releases - https://github.com/nats-io/nats-server/releases/tag/v2.6.5 (скачиваем архив для винды, разархивируем и запускаем exe-файлы)

Для RabbitMQ или Nats имплементируем интерфейсы из internal/message_broker