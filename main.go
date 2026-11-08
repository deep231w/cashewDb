package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleconn(conn net.Conn) {
	defer conn.Close()

	fmt.Println("connection recieved:", conn)

	reader := bufio.NewReader(conn)
	mesage, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Println("message is :=", mesage)
	
	response := fmt.Sprintln("ACK: yo i got yo message!")
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Println("Error during write !!")
	}

}
func main() {
	listner, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Println("Error accepting connection", err)
	}

	defer listner.Close()
	fmt.Println("TCP server runnign on port : 8000")
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
		}

		go handleconn(conn)
	}
}
