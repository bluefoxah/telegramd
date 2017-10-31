package main

import (
	"github.com/youtube/vitess/go/vt/sqlparser"
	"fmt"
)

const (
	S = "SELECT * FROM auths WHERE id=1"
)

func main() {
	parsedSql, err := sqlparser.Parse(S)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n",parsedSql)
	fmt.Printf("%v\n",parsedSql)
	fmt.Printf("%v\n",sqlparser.String(parsedSql))
}
