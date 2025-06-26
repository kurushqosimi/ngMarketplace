package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = 1 * time.Second
)

// PostgresErr is a wrapper around pgconn.PgError for custom error formatting.
type PostgresErr struct {
	*pgconn.PgError
}

// IsPgErr checks whether an error is pgx error
func IsPgErr(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr)
}

// Conv2CustomErr coverts pgx error to out custom error
func Conv2CustomErr(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return &PostgresErr{PgError: pgErr}
	}
	return fmt.Errorf("failed to conv: %w", err) // never likely to happen
}

// Error - custom error function
func (p *PostgresErr) Error() string {
	if p.PgError == nil {
		return "unknown postgres error"
	}
	msg := p.Severity + ": " + p.Message + " (SQLSTATE " + p.Code + ")"
	if p.Detail != "" {
		msg += ", Detail: " + p.Detail
	}
	if p.Hint != "" {
		msg += ", Hint: " + p.Hint
	}
	return msg
}

var (
	ErrNoRows = pgx.ErrNoRows
)

// Postgres - wrapper to work with the db
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

// New initializes a new Postgres instance with connection pooling.
// It tries to connect up to connAttempts times with a timeout between attempts.
// Returns an error if the connection cannot be established.
func New(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to coonect, attmepts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

// Close - closes connection to the database
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
