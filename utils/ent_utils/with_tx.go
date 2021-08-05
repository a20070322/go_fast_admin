package ent_utils

import (
	"context"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/pkg/errors"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rErr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

