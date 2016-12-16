package main

import (
	// "fmt"
	"github.com/dantifelicita/training_project/talk_training"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	log.Printf("App starting ...")
	router := httprouter.New()

	router.GET("/v1/talks", talk_training.ReadTalks)
	router.POST("/v1/talks", talk_training.WriteTalks)

	log.Printf("App listen on 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
