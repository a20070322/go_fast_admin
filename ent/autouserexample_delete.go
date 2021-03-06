// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/go_fast_admin/ent/autouserexample"
	"github.com/a20070322/go_fast_admin/ent/predicate"
)

// AutoUserExampleDelete is the builder for deleting a AutoUserExample entity.
type AutoUserExampleDelete struct {
	config
	hooks    []Hook
	mutation *AutoUserExampleMutation
}

// Where adds a new predicate to the AutoUserExampleDelete builder.
func (aued *AutoUserExampleDelete) Where(ps ...predicate.AutoUserExample) *AutoUserExampleDelete {
	aued.mutation.predicates = append(aued.mutation.predicates, ps...)
	return aued
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aued *AutoUserExampleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aued.hooks) == 0 {
		affected, err = aued.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AutoUserExampleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aued.mutation = mutation
			affected, err = aued.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aued.hooks) - 1; i >= 0; i-- {
			mut = aued.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aued.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aued *AutoUserExampleDelete) ExecX(ctx context.Context) int {
	n, err := aued.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aued *AutoUserExampleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: autouserexample.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: autouserexample.FieldID,
			},
		},
	}
	if ps := aued.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aued.driver, _spec)
}

// AutoUserExampleDeleteOne is the builder for deleting a single AutoUserExample entity.
type AutoUserExampleDeleteOne struct {
	aued *AutoUserExampleDelete
}

// Exec executes the deletion query.
func (auedo *AutoUserExampleDeleteOne) Exec(ctx context.Context) error {
	n, err := auedo.aued.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{autouserexample.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (auedo *AutoUserExampleDeleteOne) ExecX(ctx context.Context) {
	auedo.aued.ExecX(ctx)
}
