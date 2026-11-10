package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	cmdexec "github.com/deep231w/cashiedb/core/cmdExec"
	commandparser "github.com/deep231w/cashiedb/core/commandParser"
)

func handleconn(conn net.Conn ) {
	defer conn.Close()
	
	log.Println("connection recieved:", &conn)
	reader := bufio.NewReader(conn)
	// m,_:=reader.ReadString('\n')
	// fmt.Println("message recieved",m)
	for{
		mesage, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error in nil : %v", err)
			if err ==io.EOF{
				log.Println("client disconnected !!")
			}else{
				log.Println("Error during read : ", err)
			}
			// return 
		}

		cmdProtocol:= strings.TrimSpace(mesage)
		fmt.Println("message is :=", cmdProtocol)

		parsedArg,err := commandparser.CommandParser(&cmdProtocol)
		if err !=  nil{
			// conn.Write([]byte(err.Error()+"\n"))
			fmt.Fprintf(conn, "Error during parsing cmd : %v\n", err.Error())
			continue
		}

		res, err:=cmdexec.ExecCommand(parsedArg)
		if err != nil{
			fmt.Fprintf(conn, "Error during cmd : %v\n", err.Error())
			continue
		}
		fmt.Println("result:", *res)
		response := fmt.Sprintf("ACK: yo i got yo message!,  here is your result: %v\n", *res)

		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Println("Error during write !!")
			continue
		}
	}
	

}


func main() {
	listner, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Println("Error accepting connection", err)
	}

	connClients:= int64(0)

	defer listner.Close()
	fmt.Println("TCP server runnign on port : 8000")
	for {
		conn, err := listner.Accept()
		connClients+=1
		log.Println("Total connections : ", connClients)
		if err != nil {
			log.Println("Error accepting connection:", err)
		}
		go handleconn(conn)
	}
}
