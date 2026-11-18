package resp

import (
	"reflect"
	"testing"
)

func TestDeserializer(t *testing.T){
	inputs:= []struct{
		Data string
		Type byte
	}{
		{Data: "*1\r\n$4\r\nping\r\n", Type: ARRAYS},
		{Data: "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n", Type: ARRAYS},
		{Data: "*2\r\n$3\r\nget\r\n$3\r\nkey\r\n", Type: ARRAYS},
	}

	expectedResult:= []struct{
		Data interface{}
		Type byte
	}{
		{Data: []*RespDataType{{Value: "ping", Type: BULKSTRINGS}}, Type: ARRAYS},
		{Data: []*RespDataType{{Value: "echo", Type: BULKSTRINGS}, {Value: "hello world", Type: BULKSTRINGS}}, Type: ARRAYS},
		{Data: []*RespDataType{{Value: "get", Type: BULKSTRINGS}, {Value: "key", Type: BULKSTRINGS}}, Type: ARRAYS},
	}
	testDeserializer:= NewRespDeserializer()

	deserializeFuncs := map[byte]func([]byte, int) (*RespDataType, int, error){
		SIMPLESTRING: testDeserializer.SimpleString,
		ERRORS:       testDeserializer.SimpleErrors,
		INTEGERS:     testDeserializer.Integers,
		BULKSTRINGS:  testDeserializer.BulkString,
		ARRAYS:       testDeserializer.Arrays,
	}

	for i,v:= range inputs {
		t.Log(i)

		var res *RespDataType
		var err error
		switch v.Type{
		case SIMPLESTRING:
			res,_,err=deserializeFuncs[SIMPLESTRING]([]byte(v.Data),0)
		case BULKSTRINGS:
			res, _, err= deserializeFuncs[BULKSTRINGS]([]byte(v.Data),0)
		case ERRORS:
			res, _, err= deserializeFuncs[ERRORS]([]byte(v.Data),0)
		case ARRAYS:
			res,_,err= deserializeFuncs[ARRAYS]([]byte(v.Data),0)
		}

		if res == nil {
			res= &RespDataType{Value: nil, Type: v.Type}
		}

		if err != nil {
			t.Logf("Error :  deserialising index %v, %v : %v", i, v.Data, err)
		}
		if !reflect.DeepEqual(expectedResult[i].Data, res.Value) {
			t.Fatalf("For input %v, expected %v, got %v", v.Data, expectedResult[i].Data, res.Value)
		}

		if res.Type != expectedResult[i].Type {
			t.Fatalf("For input %v, expected type %v, got %v", v.Data, expectedResult[i].Type, res.Type)
		}
	}
}