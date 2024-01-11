package locate

import (
	"mq_es_cache/go-object-storage/lib/rabbitmq"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate("/Users/yinhuile/data" + "/objects/" + object) {
			q.Send(msg.ReplyTo, "127.0.0.0:12356")
		}
	}
}
