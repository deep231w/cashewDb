package cmdexec

import (
	"errors"
	"strings"

	"github.com/deep231w/cashiedb/dto"
)

func ExecCommand(args *dto.Command) (*string, error){
	switch strings.ToUpper(args.Name) {
		case "SET" :
			res:=Set(&args.Args[0], &args.Args[1])
			return &res, nil
		case "GET" :
			res, err:=Get(&args.Args[0])
			if err != nil{
				return nil, errors.New("Not Found !")
			}
			return &res, nil
		case "DEL" :
			res, err:=Del(&args.Args[0])
			if err != nil{
				return nil, err
			}
			return &res, nil
		default :
			return nil, errors.New("Commad not matched !")
	}
}