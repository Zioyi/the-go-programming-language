package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func foo2() error {
	return fmt.Errorf("foo failed: %w", sql.ErrNoRows)
}

func bar2() error {
	if err := foo2(); err != nil {
		return fmt.Errorf("bar failed: %w", foo())
	}
	return nil
}

func main2() {
	err := bar2()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("data not found, %+v\n", err)
		return
	}
}
