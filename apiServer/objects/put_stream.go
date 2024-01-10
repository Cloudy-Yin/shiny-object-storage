package objects

import (
	"fmt"
	"mq_es_cache/go-object-storage/apiServer/heartbeat"
	"mq_es_cache/go-object-storage/lib/rs"
)

func putStream(hash string, size int64) (*rs.RSPutStream, error) {
	servers := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(servers) != rs.ALL_SHARDS {
		return nil, fmt.Errorf("cannot find enough dataServer")
	}

	return rs.NewRSPutStream(servers, hash, size)
}
