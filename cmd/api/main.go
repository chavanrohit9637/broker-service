package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

func main() {
	fmt.Printf("server starting on port %v", port)

	config := Config{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: config.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
