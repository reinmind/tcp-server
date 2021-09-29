package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

func handleConnection(conn net.Conn) {

	fmt.Printf("Servering %s\n", conn.RemoteAddr().String())
	for {
		clientData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		s := strings.TrimSpace(string(clientData))
		// accept
		if s == "quit" {
			u := "connection closed\n"
			conn.Write([]byte(string(u)))
			break
		} else {
			prefix := "/bin/"
			sa := strings.Split(s, " ")
			sa[0] = prefix + sa[0]
			app := sa[0]
			// execute remote command
			cmd := exec.Command(app, sa...)
			b, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(string(b))
			conn.Write(b)
			fmt.Printf("%s: %s\n", conn.RemoteAddr().String(), s)
		}
		// send a random integer to client
		result := strconv.Itoa(rand.Intn(100)) + "\n"
		conn.Write([]byte(string(result)))
	}
	conn.Close()
}
