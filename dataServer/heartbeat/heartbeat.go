package heartbeat

import (
	"mq_es_cache/go-object-storage/lib/rabbitmq"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New("amqp://admin:admin123@127.0.0.1:5672")
	defer q.Close()
	for {
		q.Publish("apiServers", "127.0.0.1:12356")
		time.Sleep(5 * time.Second)
	}
}
