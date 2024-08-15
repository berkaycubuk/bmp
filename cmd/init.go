package cmd

import (
	"log"
	"net/http"

	"github.com/berkaycubuk/bmp/database"
)

func Execute() error {
	db, err := database.NewSQLiteDB("database/database.sqlite")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	_, err = initRouter()
	if err != nil {
		return err
	}

	return nil
}

func initRouter() (*http.ServeMux, error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web/static")))

	log.Printf("Server started, and listenin on port: 8080")
	err := http.ListenAndServe(":8080", mux)
	return mux, err
}
