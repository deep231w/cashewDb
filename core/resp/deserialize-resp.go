package resp

import (
	"bytes"
	"fmt"
	// "strconv"
)

// SIMPLESTRING byte = '+'
// ERRORS       byte = '-'
// INTEGERS     byte = ':'
// BULKSTRINGS  byte = '$'
// ARRAYS       byte = '*'
// NULL         byte = '_'

func DeserializeData(data []byte){

	switch data[0] {
		case SIMPLESTRING:
		case ERRORS:
		case INTEGERS:
		case BULKSTRINGS:
		case ARRAYS:
			array(data)
		case NULL:
	}
}

func bulkString(data []byte){
	fmt.Println("data in bulk string : ", string(data))

}
func array(data []byte){
	newlineidx:= bytes.Index(data, []byte("\r\n"))
	if newlineidx == -1{
		fmt.Println("MIssing in clrf \\r\\n ")
	}
	fmt.Println("len of arr : \n", newlineidx)
	fmt.Println("data in array deserializer :=" ,string(data))
}