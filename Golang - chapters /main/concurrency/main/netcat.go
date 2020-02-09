package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	//conn, err := net.Dial("tcp", "localhost:8000")
	conn, err := net.Dial("tcp", "localhost:" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	go mustCopy(os.Stdout, conn)

}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

