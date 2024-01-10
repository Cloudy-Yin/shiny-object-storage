package main

import (
	"log"
	"mq_es_cache/go-object-storage/dataServer/heartbeat"
	"mq_es_cache/go-object-storage/dataServer/locate"
	"mq_es_cache/go-object-storage/dataServer/objects"
	"mq_es_cache/go-object-storage/dataServer/temp"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
