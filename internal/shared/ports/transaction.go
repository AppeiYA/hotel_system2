package ports

import "context"

// TransactionManagerInt is the port the application/service layer
// depends on. It has no knowledge of sqlx, postgres, or any other
// infra detail — application code imports this, never the adapter.
type TransactionManagerInt interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}