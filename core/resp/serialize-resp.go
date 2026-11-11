package resp

const (
	SIMPLESTRING byte = '+'
	ERRORS       byte = '-'
	INTEGERS     byte = ':'
	BULKSTRINGS  byte = '$'
	ARRAYS       byte = '*'
	NULL         byte = '_'
)
