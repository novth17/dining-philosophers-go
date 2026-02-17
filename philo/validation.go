package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
    ErrNegativeValue = errors.New("value must be strictly positive")
    ErrInvalidArgCount = errors.New("invalid number of arguments")
)

func validateArgs(args []string) error {

    if len(args) != 4 && len(args) != 5 {
        return fmt.Errorf("%w: expected 4 or 5, got %d", ErrInvalidArgCount, len(args))
    }

    for i := 0; i < len(args); i++ {
        _, err := parseValidNumber(args[i])
        if err != nil {
            return fmt.Errorf("argument %d: '%s': %w", i + 1, args[i], err)     }
    }
    return nil
}

func fatal(err error) {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
}

func parseValidNumber(str string) (int, error) {
    number, err := strconv.Atoi(str)
    if err != nil {
        if errors.Is(err, strconv.ErrSyntax) {
            return 0, fmt.Errorf("contains non-numeric characters")
        }
        if errors.Is(err, strconv.ErrRange) {
            return 0, fmt.Errorf("is too large for an integer")
        }
        return 0, err
    }
    if number <= 0 {
        return 0, ErrNegativeValue
    }
    return number, nil
}