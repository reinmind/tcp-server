package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	// define run arguments
	// or
	// var p int
	// flag.IntVar(&p,"p",8099,"server listen port")
	port := flag.String("p", "8099", "server listen port")
	flag.Parse()

	l, err := net.Listen("tcp4", ":"+*port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err2 := l.Accept()
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		go handleConnection(c)
	}
}
