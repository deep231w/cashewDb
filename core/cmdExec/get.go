package cmdexec

import (
	"errors"
	"fmt"

	"github.com/deep231w/cashiedb/db"
)

func Get(key *string) (string, error){
	res,ok:= db.Store[*key]
	fmt.Println("all db : ", db.Store)
	fmt.Println("key , found: ", *key,res)
	if !ok {
		return "", errors.New("Not found")
	}

	return res, nil
}