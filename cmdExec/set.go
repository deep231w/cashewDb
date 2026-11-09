package cmdexec

import (
	"fmt"

	"github.com/deep231w/cashiedb/db"
)

func Set(key *string, val *string) string{
	db.Store[*key]=*val
	fmt.Printf("updated db : %v",db.Store)

	return "Value set successfully !!, "+ *key +":"+ *val
}