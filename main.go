package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fvbock/endless"
)

//SayHello is default response function
func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {

	http.HandleFunc("/", SayHello)
	err := endless.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8888 stopped")

	os.Exit(0)
}
