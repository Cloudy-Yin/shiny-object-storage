package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.EscapedPath(), strings.Split(r.URL.EscapedPath(), "/")[2])
	f, e := os.Create("/Users/yinhuile/data" + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
