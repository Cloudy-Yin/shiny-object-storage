package main

import (
	"log"
	"mq_es_cache/go-object-storage/lib/es"
	"mq_es_cache/go-object-storage/lib/utils"
	"os"
	"path/filepath"
	"strings"

	"mq_es_cache/go-object-storage/apiServer/objects"
)

func main() {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")

	for i := range files {
		hash := strings.Split(filepath.Base(files[i]), ".")[0]
		verify(hash)
	}
}

func verify(hash string) {
	log.Println("verify", hash)
	size, e := es.SearchHashSize(hash)
	if e != nil {
		log.Println(e)
		return
	}
	stream, e := objects.GetStream(hash, size)
	if e != nil {
		log.Println(e)
		return
	}
	d := utils.CalculateHash(stream)
	if d != hash {
		log.Printf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Close()
}
