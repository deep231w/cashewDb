package commandparser

import (
	"errors"
	"strings"

	"github.com/deep231w/cashiedb/dto"
)

//1st arg : operation command (SET/GET/DEL)
//2nd arg : Key (ex: name, age, rollNo, etc...)
//3rd arg : value
/*
-so the example command will be
	SET name deepak
	GET name
	DEL name
*/


func CommandParser(cmd *string) (*dto.Command,error){
	cmdArr:= strings.Split(*cmd, " ")

	if len(cmdArr) > 3 {
		return &dto.Command{},errors.New("Too many arguments !!")
	}
	cmdArg:= cmdArr[0]

	var restArgs []string
	for i:=1; i<len(cmdArr); i++{
		restArgs=append(restArgs,cmdArr[i])
	}
	
	structedArg:= &dto.Command{
		Name: cmdArg,
		Args: restArgs,
	}

	return structedArg, nil
}