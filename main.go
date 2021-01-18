package main

import (
	"fmt"
	"os"
	"pgn/pgconf"
	"pgn/tgapi"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Usage: pgn <text>")
		return
	}
	message := os.Args[1]
	conf := pgconf.GetConf()
	t := tgapi.New(conf.Token, conf.TargetUid)
	t.SendMessage(message)
}
