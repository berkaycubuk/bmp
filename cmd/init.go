package cmd

import (
	"log"
	"net/http"
)

func Execute() error {
	http.Handle("/", http.FileServer(http.Dir("web/static")))

	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}
