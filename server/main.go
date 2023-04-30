package main

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
    conn     net.Conn
    username string
}

var clients []Client

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        client := Client{conn: conn}

        clients = append(clients, client)

        go handleClient(client)
    }
}

func handleClient(client Client) {
	defer client.conn.Close()

	client.conn.Write([]byte("Welcome to the chat. Hacking the terminal? What's your name?"))

	scanner := bufio.NewScanner(client.conn)

	scanner.Scan()

	for {
		scanner := bufio.NewScanner(client.conn)
		if !scanner.Scan() {
			// Client disconnected
			break
		}
		message := scanner.Text()
		client.username = scanner.Text()
		log.Printf("%s: %s\n", client.username, message)
	}

	

	for {
		msg, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil {
			log.Printf("Error reading message from client %s: %v\n", client.username, err)
			return
		}

		log.Printf("%s: %s", client.username, msg)

		for _, c := range clients {
			if c.conn == client.conn {
				continue
			}
			_, err := c.conn.Write([]byte(client.username + ": " + msg))
			if err != nil {
				log.Printf("Error sending message to client %s: %v\n", c.username, err)
			}
		}
	}
}

func broadcast(message string) {
    for _, client := range clients {
        if client.conn != nil {
            client.conn.Write([]byte(message + "\n"))
        }
    }
}
