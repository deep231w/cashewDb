package resp

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

)

/*
DATA TYPES STARTS WITH:-

	SIMPLESTRING byte = '+'
	ERRORS       byte = '-'
	INTEGERS     byte = ':'
	BULKSTRINGS  byte = '$'
	ARRAYS       byte = '*'
	NULL         byte = '_'

*/

/*
	EXAMPLE RESPED DATA :=
		deserialized data:  *2\r\n$3\r\nlsl\r\n$2\r\nls\r\n
		after serialize: ["lsl","ls"]
*/

type RespDeserializer struct{}
func (rd RespDeserializer) DeserializeData(data []byte) ( *RespDataType,error) {
	var result *RespDataType 
	var err error
	switch data[0] {
		case SIMPLESTRING:
		case ERRORS:
		case INTEGERS:
		case BULKSTRINGS:
			result,_, err=rd.BulkString(data, 0)
		case ARRAYS:
			result, _, err= rd.Arrays(data, 0)
		case NULL:
	}
	return result, err
}

func (rd RespDeserializer) SimpleString(data []byte, start int)(*RespDataType,int, error){
	return nil, -1, nil
}

func (rd RespDeserializer) Errors(data []byte, start int)(*RespDataType,int, error){
	return nil, -1, nil
}

func (rd RespDeserializer) Integers(data []byte, start int)(*RespDataType, int ,error){
	return nil, -1, nil
}

//$<length>\r\n<element>\r\n
func (rd RespDeserializer) BulkString(data []byte, start int) (*RespDataType,int, error){
	fmt.Println("inside string deser-----------------")
	fmt.Printf("data in bulk string : %q\n", data)
	if data[0] != BULKSTRINGS {
		fmt.Println("Error 1 in string type check $ missing!!")
	}

	crlfIdx:= bytes.Index(data, []byte("\r\n"))
	
	lengthOfelement, err:= strconv.Atoi(string(data[1:crlfIdx]))
	
	if err != nil{
		log.Println("error during convert ,  in bulkstring !!")
	}
	pos:= crlfIdx+2
	fmt.Printf("length of str in bulkstrng: %d and pos is : %d & newlineidx is : %d\n", lengthOfelement, pos, crlfIdx)
	tempData:= data[pos:pos+lengthOfelement]

	fmt.Printf("temp data in bulkstring : %q\n", tempData)

	return nil, -1,nil
}

//*<length-of-array(int)>\r\nelement1...elementN\r\n
func (rd RespDeserializer) Arrays(data []byte, start int)(*RespDataType, int, error){
	fmt.Println("inside array deser----------------")
	crlfIdx:= bytes.Index(data, []byte("\r\n"))
	fmt.Println("crlfidx: ", crlfIdx)
	if crlfIdx == -1{
		fmt.Println("MIssing in clrf \\r\\n ")
	}
	lengthOfarr, err:= strconv.Atoi(string(data[1:crlfIdx]))
	if err!= nil {
		log.Println("Error during convert : ", err)
	}
	fmt.Printf("len of arr : %d\n", lengthOfarr)

	result:= make([]string,0, lengthOfarr)
	
	deserialesedFuncs := map[byte]func([]byte, int)(*RespDataType, int, error){
		SIMPLESTRING:	rd.SimpleString,
		ERRORS   	: 	rd.Errors,
		INTEGERS    : 	rd.BulkString,
		BULKSTRINGS :	rd.BulkString,
		ARRAYS      :	rd.Arrays,
	}

	pos:= crlfIdx+2
	arrIdx:=0
	for arrIdx < lengthOfarr {
		deserialiseFunc, ok:= deserialesedFuncs[data[pos]]
		if !ok{
			return nil, -1, fmt.Errorf("Not found data type %c \n", data[pos])
		}

		fmt.Printf("old data : = %q , and pos is : %d\n", data, pos )
		newData:= data[pos:]
		fmt.Printf("new data extrcted: = %q\n", newData)

		res, next, err:=deserialiseFunc(newData, pos)
		if err != nil {
			return nil, -1, err
		}
		fmt.Println("res is : ", res)

		pos+=next
		fmt.Println("after it pos is : ---", pos)
		arrIdx++
	}

	fmt.Println("final result: ", result)
	fmt.Println("data in array deserializer :=" ,string(data))
	return nil, -1, nil
}