package main

import (
	"github.com/codegangsta/cli"
	"os"
	"test_project/router"
)

const (
	AppName = "tos scanner kv tool"
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Version = "scanner kv v1 "
	app.Usage = "tos scanner kv tool"
	app.Writer = os.Stdout

	app.Commands = []cli.Command{
		router.ActiveScanner,
	}
	_ = app.Run(os.Args)
}
