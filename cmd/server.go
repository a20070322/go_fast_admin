package cmd

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/bootstrap"
	"github.com/a20070322/go_fast_admin/router"
	"github.com/labstack/gommon/color"
	"github.com/urfave/cli/v2"
)
var banner = `
 _______    ___           _______.___________.        ___       _______  .___  ___.  __  .__   __. 
|   ____|  /   \         /       |           |       /   \     |       \ |   \/   | |  | |  \ |  | 
|  |__    /  ^  \       |   (----.---|  |----.      /  ^  \    |  .--.  ||  \  /  | |  | |   \|  | 
|   __|  /  /_\  \       \   \       |  |          /  /_\  \   |  |  |  ||  |\/|  | |  | |  . .  |
|  |    /  _____  \  .----)   |      |  |         /  _____  \  |  '--'  ||  |  |  | |  | |  |\   |
|__|   /__/     \__\ |_______/       |__|        /__/     \__\ |_______/ |__|  |__| |__| |__| \__|
`
var Server = &cli.Command{
	Name:  "start",
	Usage: "start api server",
	Action: func(ctx *cli.Context) error {
		fmt.Println(color.Green(banner))
		bootstrap.Init()
		router.StartServer()
		return nil
	},
}