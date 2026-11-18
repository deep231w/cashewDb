package resp

import (
	"reflect"
	"testing"
)

func TestDeserializeData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *RespDataType
		wantErr  bool
	}{
		// Simple strings
		{
			name:  "simple string",
			input: "+OK\r\n",
			expected: &RespDataType{
				Value: []byte("OK"),
				Type:  SIMPLESTRING,
			},
		},
		{
			name:  "simple string with spaces",
			input: "+hello world\r\n",
			expected: &RespDataType{
				Value: []byte("hello world"),
				Type:  SIMPLESTRING,
			},
		},
		{
			name:  "empty simple string",
			input: "+\r\n",
			expected: &RespDataType{
				Value: []byte{},
				Type:  SIMPLESTRING,
			},
		},

		// Errors
		{
			name:  "simple error",
			input: "-ERR unknown command\r\n",
			expected: &RespDataType{
				Value: "ERR unknown command",
				Type:  ERRORS,
			},
		},
		{
			name:  "error without message",
			input: "-ERR\r\n",
			expected: &RespDataType{
				Value: "ERR",
				Type:  ERRORS,
			},
		},

		// Integers
		{
			name:  "positive integer",
			input: ":1000\r\n",
			expected: &RespDataType{
				Value: 1000,
				Type:  INTEGERS,
			},
		},
		{
			name:  "negative integer",
			input: ":-100\r\n",
			expected: &RespDataType{
				Value: -100,
				Type:  INTEGERS,
			},
		},
		{
			name:  "zero integer",
			input: ":0\r\n",
			expected: &RespDataType{
				Value: 0,
				Type:  INTEGERS,
			},
		},

		// Bulk strings
		{
			name:  "bulk string",
			input: "$4\r\nping\r\n",
			expected: &RespDataType{
				Value: []byte("ping"),
				Type:  BULKSTRINGS,
			},
		},
		{
			name:  "bulk string with spaces",
			input: "$11\r\nhello world\r\n",
			expected: &RespDataType{
				Value: []byte("hello world"),
				Type:  BULKSTRINGS,
			},
		},
		{
			name:  "empty bulk string",
			input: "$0\r\n\r\n",
			expected: &RespDataType{
				Value: []byte{},
				Type:  BULKSTRINGS,
			},
		},

		// Arrays
		{
			name:  "array with one element",
			input: "*1\r\n$4\r\nping\r\n",
			expected: &RespDataType{
				Value: []*RespDataType{
					{Value: []byte("ping"), Type: BULKSTRINGS},
				},
				Type: ARRAYS,
			},
		},
		{
			name:  "array with two elements",
			input: "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n",
			expected: &RespDataType{
				Value: []*RespDataType{
					{Value: []byte("echo"), Type: BULKSTRINGS},
					{Value: []byte("hello world"), Type: BULKSTRINGS},
				},
				Type: ARRAYS,
			},
		},
		{
			name:  "redis GET command",
			input: "*2\r\n$3\r\nget\r\n$3\r\nkey\r\n",
			expected: &RespDataType{
				Value: []*RespDataType{
					{Value: []byte("get"), Type: BULKSTRINGS},
					{Value: []byte("key"), Type: BULKSTRINGS},
				},
				Type: ARRAYS,
			},
		},
		{
			name:  "redis SET command",
			input: "*3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
			expected: &RespDataType{
				Value: []*RespDataType{
					{Value: []byte("set"), Type: BULKSTRINGS},
					{Value: []byte("key"), Type: BULKSTRINGS},
					{Value: []byte("value"), Type: BULKSTRINGS},
				},
				Type: ARRAYS,
			},
		},
		{
			name:  "empty array",
			input: "*0\r\n",
			expected: &RespDataType{
				Value: []*RespDataType{},
				Type:  ARRAYS,
			},
		},
		{
			name:  "nested array",
			input: "*2\r\n*1\r\n$4\r\nping\r\n*1\r\n$4\r\npong\r\n",
			expected: &RespDataType{
				Value: []*RespDataType{
					{
						Value: []*RespDataType{
							{Value: []byte("ping"), Type: BULKSTRINGS},
						},
						Type: ARRAYS,
					},
					{
						Value: []*RespDataType{
							{Value: []byte("pong"), Type: BULKSTRINGS},
						},
						Type: ARRAYS,
					},
				},
				Type: ARRAYS,
			},
		},

		// Null (RESP3)
		{
			name:  "null",
			input: "_\r\n",
			expected: &RespDataType{
				Value: nil,
				Type:  NULL,
			},
		},

		// Invalid input
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "invalid prefix",
			input:   "?invalid\r\n",
			wantErr: true,
		},
		{
			name:    "incomplete simple string",
			input:   "+OK",
			wantErr: true,
		},
		{
			name:    "invalid integer",
			input:   ":abc\r\n",
			wantErr: true,
		},
		{
			name:    "invalid bulk string length",
			input:   "$abc\r\nhello\r\n",
			wantErr: true,
		},
		{
			name:    "truncated bulk string",
			input:   "$10\r\nhello\r\n",
			wantErr: true,
		},
		{
			name:    "invalid array length",
			input:   "*abc\r\n",
			wantErr: true,
		},
		{
			name:    "truncated array",
			input:   "*2\r\n$3\r\nget\r\n",
			wantErr: true,
		},
	}

	rd := NewRespDeserializer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rd.DeserializeData([]byte(tt.input))

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf(
					"\nInput:    %q\nExpected: %#v\nGot:      %#v",
					tt.input,
					tt.expected,
					got,
				)
			}
		})
	}
}
