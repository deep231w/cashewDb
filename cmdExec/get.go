package cmdexec

import "github.com/deep231w/cashiedb/db"

func Get(key *string) *string{
	res,ok:= db.Store[*key]
	if !ok {
		return nil
	}
	
	return &res
}