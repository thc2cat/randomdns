// randomdns
//
// generate random dns queries
// usefull when you want to test your DNS queries
// (unbound as transparent DNS resolver)

package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {

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
