package main

import (
	"bufio"
	"log"
	"net"
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
		var command CommandStrategy
		switch cmd {
		case CREATE:
			command = NewCreateCommand(splittedMessage)
		case REMOVE:
			command = NewRemoveCommand(splittedMessage)
		case RENAME:
			command = NewRenameCommand(splittedMessage)
		case READDIR:
			command = NewReadDirCommand(splittedMessage)
		case EXIT:
			conn.Write([]byte("CLOSED\n"))
			conn.Close()
			return
		default:
			conn.Write([]byte("ERR\n"))
			continue
		}

		response, _ := command.Execute()
		conn.Write([]byte(response))
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
