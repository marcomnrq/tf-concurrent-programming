package node

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

var serverHost string
var nodeHost string

func node() {
	// Nodo
	fmt.Print("[Nodo] Ingrese puerto TCP para este nodo: ")
	bIn := bufio.NewReader(os.Stdin)
	port, _ := bIn.ReadString('\n')
	port = strings.TrimSpace(port)
	nodeHost = fmt.Sprintf("localhost:%s", port)
	// Servidor
	fmt.Print("[Servidor] Ingrese puerto del servidor: ")
	server, _ := bIn.ReadString('\n')
	server = strings.TrimSpace(server)
	serverHost = fmt.Sprintf("localhost:%s", server)

	initNode()
}

func initNode() {
	ln, _ := net.Listen("tcp", nodeHost)
	fmt.Println("[Nodo] Escuchando en ", nodeHost)
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handleConnection(con)
	}
}

func handleConnection(conn net.Conn) {
	buf, read_err := ioutil.ReadAll(conn)
	if read_err != nil {
		fmt.Println("failed:", read_err)
		return
	}
	fmt.Println("Got: ", string(buf))

	_, write_err := conn.Write([]byte("Message received.\n"))
	if write_err != nil {
		fmt.Println("failed:", write_err)
		return
	}
	conn.Close()
}
