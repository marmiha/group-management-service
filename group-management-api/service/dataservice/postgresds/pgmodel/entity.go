package pgmodel

import (
	"context"
	"time"
)

type EntityID int64
type Entity struct {
	CreatedAt time.Time `pg:"default:now()"`
	UpdatedAt time.Time `pg:"default:now()"`
}

// Update time before saving.
func (e Entity) BeforeUpdate(ctx context.Context) (context.Context, error) {
	e.UpdatedAt = time.Now()
	return ctx, nil
}
