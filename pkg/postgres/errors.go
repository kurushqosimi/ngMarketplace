package postgres

import (
	"fmt"
)

func ErrCommit(op string, err error) error {
	return fmt.Errorf("%s: failed to commit Tx: %w", op, err)
}

func ErrRollback(op string, err error) error {
	return fmt.Errorf("%s: failed to rollback Tx: %w", op, err)
}

func ErrCreateTx(op string, err error) error {
	return fmt.Errorf("%s: failed to create Tx: %w", op, err)
}

func ErrCreateQuery(op string, err error) error {
	return fmt.Errorf("%s: failed to create SQL Query: %w", op, err)
}

func ErrScan(op string, err error) error {
	return fmt.Errorf("%s: failed to scan: %w", op, err)
}

func ErrExec(op string, err error) error {
	return fmt.Errorf("%s: failed to execute: %w", op, err)
}

func ErrDoQuery(op string, err error) error {
	return fmt.Errorf("%s: failed to query: %w", op, err)
}

func ErrReadRows(op string, err error) error {
	return fmt.Errorf("%s: failed to scan rows: %w", op, err)
}
