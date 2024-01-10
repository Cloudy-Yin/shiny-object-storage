package locate

import (
	"encoding/json"
	"mq_es_cache/go-object-storage/lib/rabbitmq"
	"mq_es_cache/go-object-storage/lib/rs"
	"mq_es_cache/go-object-storage/lib/types"
	"time"
)

func Locate(name string) (locateInfo map[int]string) {
	q := rabbitmq.New("amqp://admin:admin123@127.0.0.1:5672")
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	locateInfo = make(map[int]string)
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}
	return
}

func Exist(name string) bool {
	return len(Locate(name)) >= rs.DATA_SHARDS
}
