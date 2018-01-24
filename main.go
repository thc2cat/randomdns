// randomdns
//
// generate random dns queries
// usefull when you want to test your DNS queries
// without cache

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
		a := strconv.Itoa(rand.Intn(255))
		b := strconv.Itoa(rand.Intn(255))
		c := strconv.Itoa(rand.Intn(255))
		d := strconv.Itoa(rand.Intn(255))
		ipstring := a + "." + b + "." + c + "." + d
		fmt.Print(ipstring, " : ")
		hostname, _ := net.LookupAddr(ipstring)
		fmt.Print(hostname, "\n")
		time.Sleep(50 * time.Millisecond)
	}
}
