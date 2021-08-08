package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	ADDRESS = "127.0.0.1:5000"
)

const (
	OK     = "OK"
	ERR    = "ERR"
	CLOSED = "CLOSED"
)

func main() {
	conn, err := net.Dial("tcp", ADDRESS)
	if err != nil {
		log.Fatalf("Error connecting to server %v\n", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			message := scanner.Text()
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error scanning %v\n", err)
			}

			log.Printf("%v\n", message)
			_, err := conn.Write([]byte(message + "\n")) // @@@
			if err != nil {
				log.Fatalf("Error writing...%v\n", err) // @@@
			}

			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading response %v\n", err)
			}
			log.Printf("Response %v\n", response)

			if response == CLOSED+"\n" {
				break
			}
		}
	}

	conn.Close()
}
