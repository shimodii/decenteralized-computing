package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func decryptMessage(message string) string {
	return strings.ReplaceAll(message, "#", "")
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message, _ := bufio.NewReader(conn).ReadString('\n')
	message = strings.TrimSpace(message)
	fmt.Println("Received Encrypted Part:", message)

	decrypted := decryptMessage(message)
	fmt.Println("Decrypted Part:", decrypted)

	fmt.Fprintf(conn, decrypted+"\n")
}
