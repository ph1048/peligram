package main

import (
	"fmt"
	"pgn/pgconf"
	"pgn/tgapi"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Notify struct {
		Msg string `required help:"Notification text to send"`
	} `cmd help:"Send a single notification to default contact"`
}

func main() {
	conf := pgconf.GetConf()
	ctx := kong.Parse(&CLI)
	if ctx.Command() == "notify" {
		t := tgapi.New(conf.Token, conf.TargetUid)
		t.SendMessage(CLI.Notify.Msg)
		fmt.Println("Done!")
	}
}
