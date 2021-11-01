package main

import (
	"fmt"
	fmux "github.com/hashicorp/yamux/src/multi"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "www.tqautotech.com:80")


	session, err := fmux.Client(conn, nil)



	fmt.Println("test %v %v", session, err)
}
