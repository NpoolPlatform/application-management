package rollback

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/application-management/pkg/db/ent"
)

func WithTX(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			panic(v)
		}
	}()
	resp, err := fn(tx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}
	return resp, nil
}
