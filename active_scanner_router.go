package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"net/http"
)

var ActiveScanner = cli.Command{
	Name:  "active_scan",
	Usage: "active scan online",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		router := StartRouter()
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			fmt.Errorf("err:%v", err)
		}
		return nil
	},
}

func StartRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/activeScanner", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello, world"))
	}).Methods("GET")

	return router
}
