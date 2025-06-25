package postgres

import (
	"fmt"
)

// ErrCommit returns a formatted error indicating failure to commit a transaction.
func ErrCommit(op string, err error) error {
	return fmt.Errorf("%s: failed to commit Tx: %w", op, err)
}

// ErrRollback returns a formatted error indicating failure rollback a transaction.
func ErrRollback(op string, err error) error {
	return fmt.Errorf("%s: failed to rollback Tx: %w", op, err)
}

// ErrCreateTx returns a formatted error indicating failure to create a transaction.
func ErrCreateTx(op string, err error) error {
	return fmt.Errorf("%s: failed to create Tx: %w", op, err)
}

// ErrCreateQuery returns a formatted error indicating failure to create query.
func ErrCreateQuery(op string, err error) error {
	return fmt.Errorf("%s: failed to create SQL Query: %w", op, err)
}

// ErrScan returns a formatted error indicating failure to scan.
func ErrScan(op string, err error) error {
	return fmt.Errorf("%s: failed to scan: %w", op, err)
}

// ErrExec returns a formatted error indicating failure to execute the query.
func ErrExec(op string, err error) error {
	return fmt.Errorf("%s: failed to execute: %w", op, err)
}

// ErrDoQuery returns a formatted error indicating failure to do the query.
func ErrDoQuery(op string, err error) error {
	return fmt.Errorf("%s: failed to query: %w", op, err)
}

// ErrReadRows returns a formatted error indicating failure to read rows.
func ErrReadRows(op string, err error) error {
	return fmt.Errorf("%s: failed to scan rows: %w", op, err)
}
