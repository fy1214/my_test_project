package router

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"net/http"
)

var ActiveScanner = cli.Command{
	Name:  "active_scan",
	Usage: "active scan online",
	Flags: []cli.Flag{
		//cli.StringFlag{
		//	Name:  "conf",
		//	Value: "",
		//	Usage: "conf path",
		//},
	},
	Action: func(c *cli.Context) error {
		confPath := c.String("conf")
		router := StartRouter(confPath)
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			fmt.Errorf("err:%v", err)
		}
		return nil
	},
}

func StartRouter(confPath string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello, world"))
	}).Methods("GET")

	return router
}
