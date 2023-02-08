package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func main() {
	if errors.Is(unwrapErrByV(), sql.ErrNoRows) {
		fmt.Println("unwrapErrByV is true")
	}

	if errors.Is(unwrapErrByW(), sql.ErrNoRows) {
		fmt.Println("unwrapErrByW is true")
	}

	if errors.As(unwrapErrByV(), &sql.ErrNoRows) {
		fmt.Println("unwrapErrByV as true")
	}

	if errors.As(unwrapErrByW(), &sql.ErrNoRows) {
		fmt.Println("unwrapErrByV as true")
	}

	if errors.Is(fmt.Errorf("double unwrap %w", unwrapErrByW()), sql.ErrNoRows) {
		fmt.Println("double unwrapErrByW is true")
	}

}

func getSqlErrors() error {
	return sql.ErrNoRows
}

func unwrapErrByV() error {
	return fmt.Errorf("unwrapErr: %v", getSqlErrors())
}

func unwrapErrByW() error {
	return fmt.Errorf("unwrapErr: %w", getSqlErrors())
}
