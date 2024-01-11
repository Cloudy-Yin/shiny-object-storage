package main

import (
	"log"
	"mq_es_cache/go-object-storage/dataServer/heartbeat"
	"mq_es_cache/go-object-storage/dataServer/locate"
	"mq_es_cache/go-object-storage/dataServer/objects"
	"net/http"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:12356", nil))
}
