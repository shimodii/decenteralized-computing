package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
	"os"
)

func encryptMessage(message string) string {
	rand.Seed(time.Now().UnixNano())
	encrypted := ""

	for _, char := range message {
		encrypted += string(char)
		if rand.Intn(2) == 0 {
			encrypted += "#"
		}
	}
	return encrypted
}

func main() {
	buffer := make([]string, 2)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Server is running on port 8080")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	encryptedMessage := encryptMessage(message)
	fmt.Println("Encrypted message:", encryptedMessage)

	mid := len(encryptedMessage) / 2
	part1 := encryptedMessage[:mid]
	part2 := encryptedMessage[mid:]

	for i := 0; i < 2; i++ {
		client, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer client.Close()

		go func(client net.Conn, part string, index int) {
			fmt.Fprintf(client, part+"\n")

			response, _ := bufio.NewReader(client).ReadString('\n')
			buffer[index] = strings.TrimSpace(response)
		}(client, []string{part1, part2}[i], i)
	}

	time.Sleep(5 * time.Second)
	finalMessage := buffer[0] + buffer[1]
	fmt.Println("Final Decrypted Message:", finalMessage)
}
