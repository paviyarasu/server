package main

import (
	"bufio"
	"github.com/golang/go/src/pkg/fmt"
	"net"
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

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("Code got here")

}
