// randomdns
//
// generate random dns queries
// usefull when you want to test your DNS queries
// (unbound as transparent DNS resolver in a FreeBSD NAT gateway)
//

package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Println("\nBye")
			os.Exit(0)
		}
	}()
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		ipstring := ""
		for i := 1; i < 4; i++ {
			ipstring += strconv.Itoa(rand.Intn(235) + 10)
			ipstring += "."
		}
		ipstring += strconv.Itoa(rand.Intn(9) + 1)

		fmt.Print(ipstring, " : ")
		hostname, _ := net.LookupAddr(ipstring)
		fmt.Print(hostname, "\n")
		time.Sleep(50 * time.Millisecond)
	}
}
