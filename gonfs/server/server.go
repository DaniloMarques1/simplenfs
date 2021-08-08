package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	IP   = "127.0.0.1"
	PORT = "5000"
)

const (
	READDIR = "readdir"
	CREATE  = "create"
	REMOVE  = "remove"
	RENAME  = "rename"
	EXIT    = "exit"
)

func main() {
	listener, err := net.Listen("tcp", IP+":"+PORT)
	if err != nil {
		log.Fatalf("Error starting server %v\n", err)
	}
	log.Printf("Starting server...\n")

	for {
		conn, err := listener.Accept()
		log.Printf("New connection received\n")
		if err != nil {
			log.Fatalf("Error opening connection %v\n", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		rawMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading message %v\n", err)
		}

		splittedMessage := readMessage(rawMessage)
		cmd := splittedMessage[0] // will be the nfs command
		switch cmd {
		case CREATE:
			if len(splittedMessage) >= 2 {
				arg := splittedMessage[1]
				if arg != "" {
					err := os.Mkdir(arg, os.ModePerm)
					if err != nil {
						log.Printf("Error creating the dir %v\n", err)
						conn.Write([]byte("ERR\n"))
					} else {
						conn.Write([]byte("OK\n"))
					}
				}
			}
		case REMOVE:
			log.Println("Removendo....")
			if len(splittedMessage) >= 2 {
				arg := splittedMessage[1]
				if arg != "" {
					err := os.Remove(arg)
					if err != nil {
						log.Printf("Error conn %v\n", err)
						conn.Write([]byte("ERROR\n"))
					} else {
						conn.Write([]byte("OK\n"))
					}
				}
			}
		case RENAME:
			if len(splittedMessage) >= 3 {
				arg1 := splittedMessage[1]
				arg2 := splittedMessage[2]
				if arg1 != "" && arg2 != "" {
					err := os.Rename(arg1, arg2)
					if err != nil {
						log.Printf("Error conn %v\n", err)
						conn.Write([]byte("ERROR\n"))
					} else {
						conn.Write([]byte("OK\n"))
					}
				}
			}
		case READDIR:
			arg := "."
			if len(splittedMessage) >= 2 {
				arg = splittedMessage[1]
			}

			entries, err := os.ReadDir(arg)
			if err != nil {
				log.Printf("Error reading the given directory%v\n")
				conn.Write([]byte("ERROR\n"))
			} else {
				var response string
				for idx, file := range entries {
					response += fmt.Sprintf("%v- %v;", idx+1, file.Name())
					log.Println(file.Name()) // this is what i should return
				}

				log.Println(response)
				conn.Write([]byte(response+"\n"))
			}

		case EXIT:
			conn.Write([]byte("CLOSED\n"))
			conn.Close()
			return
		}
	}

}

func readMessage(rawMessage string) []string {
	rawMessage = strings.Replace(rawMessage, "\n", "", 1)
	log.Printf("Raw message %v\n", rawMessage)

	splittedMessage := strings.Split(rawMessage, " ")

	return splittedMessage
}

// for READDIR, it will be the dir names separeted by \n
func sendMessage() {
}
