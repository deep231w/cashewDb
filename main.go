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
	"github.com/deep231w/cashiedb/core/resp"
)

func handleconn(conn net.Conn ) {
	defer conn.Close()
	
	log.Println("connection recieved:", &conn)
	reader := bufio.NewReader(conn)
	// m,_:=reader.ReadString('\n')
	// fmt.Println("message recieved",m)
	for {
		mesage, err := reader.ReadString('\n')
		fmt.Println("message before parsing : ", mesage)
		if err != nil {
			if err ==io.EOF{
				log.Println("client disconnected !!")
			}else{
				log.Println("Error during read : ", err)
			}
			return 
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
		// go handleconn(conn)
		go testHandler(conn)
	}
}


//test redis cli 

func testHandler(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	buff:= make([]byte, 1024)
	for{
		// message, err:= reader.ReadString('\n')
		n, err:=reader.Read(buff)

		// fmt.Println("message is ", message)
		
		if err!= nil{
			if err == io.EOF{
				return
			}
			fmt.Println("Read error")
		}

		serialisedData:= buff[:n]
		fmt.Printf("recieved serdata: %q\n", serialisedData)
		// resp.RespDeserializer(serialisedData)
		rd:= resp.RespDeserializer{}
		rd.DeserializeData(serialisedData)
		
		fmt.Println("data recieved is = ", string(serialisedData))

		fmt.Printf("data recieved is = %q\n", serialisedData)

		fmt.Fprintf(conn,"$%d\r\n%s\r\n", len(serialisedData),serialisedData)
	}
	
}