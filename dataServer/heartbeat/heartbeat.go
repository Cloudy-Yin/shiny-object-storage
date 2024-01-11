package heartbeat

import (
	"mq_es_cache/go-object-storage/lib/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", "127.0.0.1:12356")
		time.Sleep(5 * time.Second)
	}
}
