package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


//domain layer
var (
	ErrInvalidArgCount = errors.New("invalid argument count")
	ErrNonNumeric      = errors.New("non numeric input")
	ErrTooLarge        = errors.New("number too large")
	ErrNonPositive     = errors.New("value must be strictly positive")
)

func parseValidNumber(str string) (int, error) {
	number, err := strconv.Atoi(str)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return 0, ErrNonNumeric
		}
		if errors.Is(err, strconv.ErrRange) {
			return 0, ErrTooLarge
		}
		return 0, err
	}

	if number <= 0 {
		return 0, ErrNonPositive
	}

	return number, nil
}

// application layer - use casesF
// Orchestrates validation and wraps errors with context.
func validateArgs(args []string) error {
	if len(args) != 4 && len(args) != 5 {
		return fmt.Errorf("%w: expected 4 or 5, got %d",
			ErrInvalidArgCount,
			len(args),
		)
	}

	for i := 0; i < len(args); i++ {
		_, err := parseValidNumber(args[i])
		if err != nil {
			return fmt.Errorf("argument %d ('%s'): %w",
				i+1,
				args[i],
				err,
			)
		}
	}
	return nil
}

// ===== PRESENTATION LAYER =====
func fatal(err error) {
	fmt.Fprintln(os.Stderr, "Error:", mapErrorToMessage(err))
	os.Exit(1)
}

func mapErrorToMessage(err error) string {
	switch {
	case errors.Is(err, ErrInvalidArgCount):
		return "Wrong number of arguments."
	case errors.Is(err, ErrNonNumeric):
		return "Arguments must be numeric."
	case errors.Is(err, ErrTooLarge):
		return "One of the numbers is too large."
	case errors.Is(err, ErrNonPositive):
		return "All values must be strictly positive."
	default:
		return err.Error()
	}
}
