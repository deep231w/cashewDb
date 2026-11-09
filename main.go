package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	cmdexec "github.com/deep231w/cashiedb/cmdExec"
	commandparser "github.com/deep231w/cashiedb/commandParser"
)

func handleconn(conn net.Conn) {
	defer conn.Close()

	fmt.Println("connection recieved:", &conn)

	reader := bufio.NewReader(conn)
	mesage, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error: %v", err)
	}

	cmdProtocol:= strings.TrimSpace(mesage)
	fmt.Println("message is :=", cmdProtocol)

	parsedArg,err := commandparser.CommandParser(&cmdProtocol)
	if err !=  nil{
		// conn.Write([]byte(err.Error()+"\n"))
		fmt.Fprintf(conn, "Error: %v\n", err.Error())
		return
	}

	res, err:=cmdexec.ExecCommand(parsedArg)
	if err != nil{
		fmt.Fprintf(conn, "Error: %v\n", err.Error())
		return
	}
	fmt.Println("result:", *res)
	response := fmt.Sprintf("ACK: yo i got yo message!,  here is your result: %v\n", *res)

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
