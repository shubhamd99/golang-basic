package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", GetRouter)
	// {"level":"info","ts":1622283243.512713,"caller":"go-modules_example/get.go:11","msg":"success"}
	log.Fatal(http.ListenAndServe(":9090", router))
}
