package postgres

import "time"

// Option - function that defines the option of the postgres connection
type Option func(*Postgres)

// MaxPoolSize - sets the maximum size of the pool
func MaxPoolSize(size int) Option {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

// ConnAttempts - defines the number of times to connect to db
func ConnAttempts(attempts int) Option {
	return func(c *Postgres) {
		c.connAttempts = attempts
	}
}

// ConnTimeout - defines the timeout of the connection
func ConnTimeout(timeout time.Duration) Option {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}
