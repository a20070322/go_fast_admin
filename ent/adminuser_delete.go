// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/go_fast_admin/ent/adminuser"
	"github.com/a20070322/go_fast_admin/ent/predicate"
)

// AdminUserDelete is the builder for deleting a AdminUser entity.
type AdminUserDelete struct {
	config
	hooks    []Hook
	mutation *AdminUserMutation
}

// Where adds a new predicate to the AdminUserDelete builder.
func (aud *AdminUserDelete) Where(ps ...predicate.AdminUser) *AdminUserDelete {
	aud.mutation.predicates = append(aud.mutation.predicates, ps...)
	return aud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aud *AdminUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aud.hooks) == 0 {
		affected, err = aud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aud.mutation = mutation
			affected, err = aud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aud.hooks) - 1; i >= 0; i-- {
			mut = aud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aud *AdminUserDelete) ExecX(ctx context.Context) int {
	n, err := aud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aud *AdminUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: adminuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: adminuser.FieldID,
			},
		},
	}
	if ps := aud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aud.driver, _spec)
}

// AdminUserDeleteOne is the builder for deleting a single AdminUser entity.
type AdminUserDeleteOne struct {
	aud *AdminUserDelete
}

// Exec executes the deletion query.
func (audo *AdminUserDeleteOne) Exec(ctx context.Context) error {
	n, err := audo.aud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adminuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (audo *AdminUserDeleteOne) ExecX(ctx context.Context) {
	audo.aud.ExecX(ctx)
}
