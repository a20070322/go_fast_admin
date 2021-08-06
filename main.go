package main

import (
	"github.com/a20070322/go_fast_admin/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main()  {
	app := &cli.App{
		Name:    "go_fast_admin",
		Usage:   "api server",
		Version: "1.0.0",
		Commands: []*cli.Command{
			cmd.Server,
			cmd.EntMigrate,
			cmd.AutoCurd,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}