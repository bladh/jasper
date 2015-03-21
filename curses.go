package main

import "github.com/tncardoso/gocurses"

func main() {
	gocurses.Initscr()
	defer gocurses.End()
	gocurses.Cbreak()
	gocurses.Noecho()
	gocurses.Stdscr.Keypad(true)

	gocurses.Attron(gocurses.A_BOLD)
	gocurses.Addstr("Hello World")
	gocurses.Refresh()

	gocurses.Stdscr.Getch()
}
