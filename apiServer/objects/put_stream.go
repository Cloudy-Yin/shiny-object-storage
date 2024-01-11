package objects

import (
	"fmt"
	"mq_es_cache/go-object-storage/apiServer/heartbeat"
	"mq_es_cache/go-object-storage/lib/objectstream"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
