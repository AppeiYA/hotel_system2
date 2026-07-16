package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"hotel_system2/internal/shared/config"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

// DB wraps *sqlx.DB so we have a concrete, swappable connection type
// throughout the adapter layer.
type DB struct {
	*sqlx.DB
}

func ConnectDB(cfg *config.Config) (*DB, error) {
	sqlxDB, err := sqlx.Connect("postgres", cfg.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	if err := sqlxDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	sqlxDB.SetMaxOpenConns(cfg.DBMAXOpenConns)
	sqlxDB.SetMaxIdleConns(cfg.DBMAXIdleConns)
	sqlxDB.SetConnMaxLifetime(time.Duration(cfg.DBConnMAXLife) * time.Second)

	log.Println("Database Connected")

	return &DB{sqlxDB}, nil
}

// GenericJSON adapts any T to a Postgres jsonb column via Scan/Value.
type GenericJSON[T any] struct {
	Data T
}

func (j *GenericJSON[T]) Scan(value any) error {
	if value == nil {
		j.Data = *new(T)
		return nil
	}

	var data []byte
	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("cannot scan %T into JSON", value)
	}

	return json.Unmarshal(data, &j.Data)
}

func (j GenericJSON[T]) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}

// compile-time check: *DB must satisfy Executor (via its embedded *sqlx.DB)
var _ Executor = (*DB)(nil)