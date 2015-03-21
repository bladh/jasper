package main

import (
	"fmt"
	"net"
   "bufio"
)

func main() {
	conn, err := net.Dial("tcp", "irc.underworld.no:6667")
	if err != nil {
		fmt.Printf("Couldn't connect to host")
	} else {
		fmt.Fprintf(conn, "USER alpha-centauri * 0 :alpha-centauri\n")
		fmt.Fprintf(conn, "NICK alpha-centauri\n")
		fmt.Fprintf(conn, "JOIN #jasper\n")
	}

   reader := bufio.NewReader(conn)	
	for {
		line, _, err := reader.ReadLine()
		if err == nil {
         fmt.Printf("%s\n", line)
      }
	}
}
