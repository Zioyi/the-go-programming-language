package main

import (
	"database/sql"
	"fmt"

	perrors "github.com/pkg/errors"
)

func foo() error {
	return perrors.WithStack(sql.ErrNoRows)
}

func bar() error {
	return perrors.WithMessage(foo(), "bar failed")
}

func main1() {
	err := bar()
	if perrors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}

}
