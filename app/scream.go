package main

import (
	"flag"
	"fmt"
	"github.com/vodolaz095/scream"
	"os"
)

func main() {
	defer func() {
		msg := recover()
		fmt.Printf("Error: %v\n", msg)
		os.Exit(1)
	}()
	flag.StringVar(&scream.Cfg.Key, "key", "lalala", "secret key to make spammers sad")
	flag.StringVar(&scream.Cfg.Address, "listen", "0.0.0.0:8082", "address to listen for notifications")
	flag.Parse()

	err := scream.SanityCheck()
	if err != nil {
		panic(err)
	}
	err = scream.StartServer()
	if err != nil {
		panic(err)
	}
}
