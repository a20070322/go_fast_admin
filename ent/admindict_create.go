// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/go_fast_admin/ent/admindict"
	"github.com/a20070322/go_fast_admin/ent/admindictkey"
)

// AdminDictCreate is the builder for creating a AdminDict entity.
type AdminDictCreate struct {
	config
	mutation *AdminDictMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (adc *AdminDictCreate) SetCreatedAt(t time.Time) *AdminDictCreate {
	adc.mutation.SetCreatedAt(t)
	return adc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (adc *AdminDictCreate) SetNillableCreatedAt(t *time.Time) *AdminDictCreate {
	if t != nil {
		adc.SetCreatedAt(*t)
	}
	return adc
}

// SetUpdatedAt sets the "updated_at" field.
func (adc *AdminDictCreate) SetUpdatedAt(t time.Time) *AdminDictCreate {
	adc.mutation.SetUpdatedAt(t)
	return adc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (adc *AdminDictCreate) SetNillableUpdatedAt(t *time.Time) *AdminDictCreate {
	if t != nil {
		adc.SetUpdatedAt(*t)
	}
	return adc
}

// SetDeletedAt sets the "deleted_at" field.
func (adc *AdminDictCreate) SetDeletedAt(t time.Time) *AdminDictCreate {
	adc.mutation.SetDeletedAt(t)
	return adc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (adc *AdminDictCreate) SetNillableDeletedAt(t *time.Time) *AdminDictCreate {
	if t != nil {
		adc.SetDeletedAt(*t)
	}
	return adc
}

// SetDictType sets the "dict_type" field.
func (adc *AdminDictCreate) SetDictType(s string) *AdminDictCreate {
	adc.mutation.SetDictType(s)
	return adc
}

// SetDictName sets the "dict_name" field.
func (adc *AdminDictCreate) SetDictName(s string) *AdminDictCreate {
	adc.mutation.SetDictName(s)
	return adc
}

// SetRemarks sets the "remarks" field.
func (adc *AdminDictCreate) SetRemarks(s string) *AdminDictCreate {
	adc.mutation.SetRemarks(s)
	return adc
}

// SetIsEnable sets the "is_enable" field.
func (adc *AdminDictCreate) SetIsEnable(b bool) *AdminDictCreate {
	adc.mutation.SetIsEnable(b)
	return adc
}

// AddKeyIDs adds the "key" edge to the AdminDictKey entity by IDs.
func (adc *AdminDictCreate) AddKeyIDs(ids ...int) *AdminDictCreate {
	adc.mutation.AddKeyIDs(ids...)
	return adc
}

// AddKey adds the "key" edges to the AdminDictKey entity.
func (adc *AdminDictCreate) AddKey(a ...*AdminDictKey) *AdminDictCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return adc.AddKeyIDs(ids...)
}

// Mutation returns the AdminDictMutation object of the builder.
func (adc *AdminDictCreate) Mutation() *AdminDictMutation {
	return adc.mutation
}

// Save creates the AdminDict in the database.
func (adc *AdminDictCreate) Save(ctx context.Context) (*AdminDict, error) {
	var (
		err  error
		node *AdminDict
	)
	adc.defaults()
	if len(adc.hooks) == 0 {
		if err = adc.check(); err != nil {
			return nil, err
		}
		node, err = adc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminDictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = adc.check(); err != nil {
				return nil, err
			}
			adc.mutation = mutation
			node, err = adc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(adc.hooks) - 1; i >= 0; i-- {
			mut = adc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, adc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (adc *AdminDictCreate) SaveX(ctx context.Context) *AdminDict {
	v, err := adc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (adc *AdminDictCreate) defaults() {
	if _, ok := adc.mutation.CreatedAt(); !ok {
		v := admindict.DefaultCreatedAt()
		adc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (adc *AdminDictCreate) check() error {
	if _, ok := adc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := adc.mutation.DictType(); !ok {
		return &ValidationError{Name: "dict_type", err: errors.New("ent: missing required field \"dict_type\"")}
	}
	if _, ok := adc.mutation.DictName(); !ok {
		return &ValidationError{Name: "dict_name", err: errors.New("ent: missing required field \"dict_name\"")}
	}
	if _, ok := adc.mutation.Remarks(); !ok {
		return &ValidationError{Name: "remarks", err: errors.New("ent: missing required field \"remarks\"")}
	}
	if _, ok := adc.mutation.IsEnable(); !ok {
		return &ValidationError{Name: "is_enable", err: errors.New("ent: missing required field \"is_enable\"")}
	}
	return nil
}

func (adc *AdminDictCreate) sqlSave(ctx context.Context) (*AdminDict, error) {
	_node, _spec := adc.createSpec()
	if err := sqlgraph.CreateNode(ctx, adc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (adc *AdminDictCreate) createSpec() (*AdminDict, *sqlgraph.CreateSpec) {
	var (
		_node = &AdminDict{config: adc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: admindict.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admindict.FieldID,
			},
		}
	)
	if value, ok := adc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admindict.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := adc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admindict.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := adc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admindict.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := adc.mutation.DictType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindict.FieldDictType,
		})
		_node.DictType = value
	}
	if value, ok := adc.mutation.DictName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindict.FieldDictName,
		})
		_node.DictName = value
	}
	if value, ok := adc.mutation.Remarks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindict.FieldRemarks,
		})
		_node.Remarks = value
	}
	if value, ok := adc.mutation.IsEnable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: admindict.FieldIsEnable,
		})
		_node.IsEnable = value
	}
	if nodes := adc.mutation.KeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admindict.KeyTable,
			Columns: []string{admindict.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admindictkey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdminDictCreateBulk is the builder for creating many AdminDict entities in bulk.
type AdminDictCreateBulk struct {
	config
	builders []*AdminDictCreate
}

// Save creates the AdminDict entities in the database.
func (adcb *AdminDictCreateBulk) Save(ctx context.Context) ([]*AdminDict, error) {
	specs := make([]*sqlgraph.CreateSpec, len(adcb.builders))
	nodes := make([]*AdminDict, len(adcb.builders))
	mutators := make([]Mutator, len(adcb.builders))
	for i := range adcb.builders {
		func(i int, root context.Context) {
			builder := adcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminDictMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, adcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, adcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, adcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (adcb *AdminDictCreateBulk) SaveX(ctx context.Context) []*AdminDict {
	v, err := adcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
