package main

import (
	"log"
	"net/http"

	"github.com/dempsey-ycr/cock/library/util"
)

type MHander struct {
}

func (*MHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusConflict)
	_, err := w.Write([]byte("hello world..."))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("-------------**-----------")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello ....."))
	_ = util.Stob("hello")
	log.Println("-----------h-------")
}

func main() {
	hander := new(MHander)
	http.Handle("/xu", hander)
	http.HandleFunc("/test", Hello)

	_ = http.ListenAndServe(":8080", nil)
}