package main

import (
	"code.google.com/p/goncurses"
	"log"
	"time"
)

func main() {
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer goncurses.End()
	stdscr.Print("Hello goncurses")
	stdscr.Refresh()
	time.Sleep(2 * time.Second)
}
