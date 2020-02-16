package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os/exec"
	"time"

	"github.com/WesleyMonten/Go-Exercises/Section11/Lab01/cmd"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
)

func main() {
	ln, err := net.Listen("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Infof("Command server listening %v", localhost)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	decoder := json.NewDecoder(conn)
	cmd := new(cmd.Command)
	err := decoder.Decode(cmd)

	if err != nil {
		io.WriteString(conn, fmt.Sprintf("Decode error: %v", err))
		return
	}

	conn.SetDeadline(time.Now().Add(1 * time.Minute))

	cmd2 := exec.Command(cmd.Name, cmd.Args...)
	cmd2.Stdin = conn
	cmd2.Stdout = conn
	cmd2.Stderr = conn

	err = cmd2.Run()
	if err != nil {
		log.Error("Error running command")
		return
	}
}
