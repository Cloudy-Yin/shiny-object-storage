package main

import (
	"log"
	"mq_es_cache/go-object-storage/apiServer/heartbeat"
	"mq_es_cache/go-object-storage/apiServer/locate"
	"mq_es_cache/go-object-storage/apiServer/objects"
	"mq_es_cache/go-object-storage/apiServer/temp"
	"mq_es_cache/go-object-storage/apiServer/versions"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
