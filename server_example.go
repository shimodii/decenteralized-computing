package main

import (
    "fmt"
    "net"
    "os"
    "log"
)

func main() {
    // Create or open a file to save IP addresses
    file, err := os.OpenFile("client_ips.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Listen on a port
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer ln.Close()

    fmt.Println("Server is listening on port 8080")

    for {
        // Accept connections
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err)
            return
        }
        
        // Save the client's IP address
        clientIP := conn.RemoteAddr().String()
        if _, err := file.WriteString(clientIP + "\n"); err != nil {
            log.Println("Error saving IP address:", err)
        }

        // Handle connections in a new goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    // Close the connection when you're done
    defer conn.Close()

    // Read data from the connection
    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading:", err)
            return
        }
        // Print the received data
        fmt.Println("Received:", string(buf[:n]))

        // Send a response
        _, err = conn.Write([]byte("Message received"))
        if err != nil {
            fmt.Println("Error writing:", err)
            return
        }
    }
}
