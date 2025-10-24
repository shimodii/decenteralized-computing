package main

import (
    "fmt"
    "net"
)

func main() {
    // Connect to the server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    // Read data from standard input and send to server
    for {
        fmt.Print("Enter message: ")
        var msg string
        _, err := fmt.Scanln(&msg)
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }

        _, err = conn.Write([]byte(msg))
        if err != nil {
            fmt.Println("Error writing:", err)
            return
        }

        // Read response from server
        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading:", err)
            return
        }
        fmt.Println("Server response:", string(buf[:n]))
    }
}

