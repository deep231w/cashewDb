package cmdexec

import (
	"errors"

	"github.com/deep231w/cashiedb/db"
)

func Del(key *string) (string, error){
	if _, ok:=db.Store[*key]; !ok{
		return "", errors.New("Key not present")
	}

	delete(db.Store, *key)

	return "Deleted successfully", nil
}