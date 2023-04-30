package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    scanner.Scan()
    fmt.Println(scanner.Text())

    scanner = bufio.NewScanner(os.Stdin)
    fmt.Print("Enter your name: ")
    scanner.Scan()
    name := scanner.Text()

    go func() {
        for {
            scanner := bufio.NewScanner(conn)
            if scanner.Scan() {
                fmt.Println(scanner.Text())
            } else {
                break
            }
        }
    }()

    for {
        scanner = bufio.NewScanner(os.Stdin)
        fmt.Print("> ")
        if scanner.Scan() {
            message := scanner.Text()
            fmt.Fprintf(conn, "%s: %s\n", name, message)
        } else {
            break
        }
    }

	
}
