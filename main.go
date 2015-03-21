package main

import ("net"
        "fmt"
)

func main() {
   conn, err := net.Dial("tcp", "localhost:6667")
   if err != nil {
      fmt.Printf("Couldn't connect to host")
   } else {
      fmt.Fprintf(conn, "USER alpha-centauri * 0 :alpha-centauri\n")
      fmt.Fprintf(conn, "NICK alpha-centauri")
      fmt.Fprintf(conn, "JOIN #jasper")
   }
}
