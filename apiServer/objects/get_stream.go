package objects

import (
	"fmt"
	"io"
	"mq_es_cache/go-object-storage/apiServer/locate"
	"mq_es_cache/go-object-storage/lib/objectstream"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
