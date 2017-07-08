package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	router.GET("/insertclaps/:input", InsertClapsGetRequest)
	router.POST("/insertclaps/", InsertClapsPostRequest)

	router.GET("/spam/:emoji/:num", SpamGetRequest)
	router.POST("/spam/", SpamPostRequest)

	port := os.Getenv("PORT")
	if port != "" {
		fmt.Println("Running in PRODUCTION at port :" + port + "......")
		log.Fatal(http.ListenAndServe(":"+port, router))
	} else {
		fmt.Println("Running locally at port localhost:8080......")
		log.Fatal(http.ListenAndServe(":8080", router))
	}

}
