package main

import (
	"bufio"
	"github.com/golang/go/src/pkg/fmt"
	"net"
	"time"
)

func main()  {
		listener,err := net.Listen("tcp",":9000")
		if err != nil {
			panic(err)
		}

		for {
			conn,err := listener.Accept()
			if err != nil {
				panic(err)
			}

			go handle(conn)
		}
}

func handle(conn net.Conn)  {

	err := conn.SetDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		fmt.Printf("Connection time out")
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("Code got here")

}
