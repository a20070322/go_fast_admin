// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/go_fast_admin/ent/admindict"
	"github.com/a20070322/go_fast_admin/ent/admindictkey"
	"github.com/a20070322/go_fast_admin/ent/predicate"
)

// AdminDictKeyUpdate is the builder for updating AdminDictKey entities.
type AdminDictKeyUpdate struct {
	config
	hooks    []Hook
	mutation *AdminDictKeyMutation
}

// Where adds a new predicate for the AdminDictKeyUpdate builder.
func (adku *AdminDictKeyUpdate) Where(ps ...predicate.AdminDictKey) *AdminDictKeyUpdate {
	adku.mutation.predicates = append(adku.mutation.predicates, ps...)
	return adku
}

// SetUpdatedAt sets the "updated_at" field.
func (adku *AdminDictKeyUpdate) SetUpdatedAt(t time.Time) *AdminDictKeyUpdate {
	adku.mutation.SetUpdatedAt(t)
	return adku
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (adku *AdminDictKeyUpdate) ClearUpdatedAt() *AdminDictKeyUpdate {
	adku.mutation.ClearUpdatedAt()
	return adku
}

// SetDeletedAt sets the "deleted_at" field.
func (adku *AdminDictKeyUpdate) SetDeletedAt(t time.Time) *AdminDictKeyUpdate {
	adku.mutation.SetDeletedAt(t)
	return adku
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (adku *AdminDictKeyUpdate) SetNillableDeletedAt(t *time.Time) *AdminDictKeyUpdate {
	if t != nil {
		adku.SetDeletedAt(*t)
	}
	return adku
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (adku *AdminDictKeyUpdate) ClearDeletedAt() *AdminDictKeyUpdate {
	adku.mutation.ClearDeletedAt()
	return adku
}

// SetDictLabel sets the "dict_label" field.
func (adku *AdminDictKeyUpdate) SetDictLabel(s string) *AdminDictKeyUpdate {
	adku.mutation.SetDictLabel(s)
	return adku
}

// SetDictCode sets the "dict_code" field.
func (adku *AdminDictKeyUpdate) SetDictCode(s string) *AdminDictKeyUpdate {
	adku.mutation.SetDictCode(s)
	return adku
}

// SetSort sets the "sort" field.
func (adku *AdminDictKeyUpdate) SetSort(i int) *AdminDictKeyUpdate {
	adku.mutation.ResetSort()
	adku.mutation.SetSort(i)
	return adku
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (adku *AdminDictKeyUpdate) SetNillableSort(i *int) *AdminDictKeyUpdate {
	if i != nil {
		adku.SetSort(*i)
	}
	return adku
}

// AddSort adds i to the "sort" field.
func (adku *AdminDictKeyUpdate) AddSort(i int) *AdminDictKeyUpdate {
	adku.mutation.AddSort(i)
	return adku
}

// ClearSort clears the value of the "sort" field.
func (adku *AdminDictKeyUpdate) ClearSort() *AdminDictKeyUpdate {
	adku.mutation.ClearSort()
	return adku
}

// SetRemarks sets the "remarks" field.
func (adku *AdminDictKeyUpdate) SetRemarks(s string) *AdminDictKeyUpdate {
	adku.mutation.SetRemarks(s)
	return adku
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (adku *AdminDictKeyUpdate) SetNillableRemarks(s *string) *AdminDictKeyUpdate {
	if s != nil {
		adku.SetRemarks(*s)
	}
	return adku
}

// ClearRemarks clears the value of the "remarks" field.
func (adku *AdminDictKeyUpdate) ClearRemarks() *AdminDictKeyUpdate {
	adku.mutation.ClearRemarks()
	return adku
}

// SetIsEnable sets the "is_enable" field.
func (adku *AdminDictKeyUpdate) SetIsEnable(b bool) *AdminDictKeyUpdate {
	adku.mutation.SetIsEnable(b)
	return adku
}

// SetPID sets the "P" edge to the AdminDict entity by ID.
func (adku *AdminDictKeyUpdate) SetPID(id int) *AdminDictKeyUpdate {
	adku.mutation.SetPID(id)
	return adku
}

// SetNillablePID sets the "P" edge to the AdminDict entity by ID if the given value is not nil.
func (adku *AdminDictKeyUpdate) SetNillablePID(id *int) *AdminDictKeyUpdate {
	if id != nil {
		adku = adku.SetPID(*id)
	}
	return adku
}

// SetP sets the "P" edge to the AdminDict entity.
func (adku *AdminDictKeyUpdate) SetP(a *AdminDict) *AdminDictKeyUpdate {
	return adku.SetPID(a.ID)
}

// Mutation returns the AdminDictKeyMutation object of the builder.
func (adku *AdminDictKeyUpdate) Mutation() *AdminDictKeyMutation {
	return adku.mutation
}

// ClearP clears the "P" edge to the AdminDict entity.
func (adku *AdminDictKeyUpdate) ClearP() *AdminDictKeyUpdate {
	adku.mutation.ClearP()
	return adku
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (adku *AdminDictKeyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	adku.defaults()
	if len(adku.hooks) == 0 {
		affected, err = adku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminDictKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			adku.mutation = mutation
			affected, err = adku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(adku.hooks) - 1; i >= 0; i-- {
			mut = adku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, adku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (adku *AdminDictKeyUpdate) SaveX(ctx context.Context) int {
	affected, err := adku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (adku *AdminDictKeyUpdate) Exec(ctx context.Context) error {
	_, err := adku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adku *AdminDictKeyUpdate) ExecX(ctx context.Context) {
	if err := adku.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (adku *AdminDictKeyUpdate) defaults() {
	if _, ok := adku.mutation.UpdatedAt(); !ok && !adku.mutation.UpdatedAtCleared() {
		v := admindictkey.UpdateDefaultUpdatedAt()
		adku.mutation.SetUpdatedAt(v)
	}
}

func (adku *AdminDictKeyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admindictkey.Table,
			Columns: admindictkey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admindictkey.FieldID,
			},
		},
	}
	if ps := adku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := adku.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admindictkey.FieldUpdatedAt,
		})
	}
	if adku.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: admindictkey.FieldUpdatedAt,
		})
	}
	if value, ok := adku.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admindictkey.FieldDeletedAt,
		})
	}
	if adku.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: admindictkey.FieldDeletedAt,
		})
	}
	if value, ok := adku.mutation.DictLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindictkey.FieldDictLabel,
		})
	}
	if value, ok := adku.mutation.DictCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindictkey.FieldDictCode,
		})
	}
	if value, ok := adku.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admindictkey.FieldSort,
		})
	}
	if value, ok := adku.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admindictkey.FieldSort,
		})
	}
	if adku.mutation.SortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: admindictkey.FieldSort,
		})
	}
	if value, ok := adku.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindictkey.FieldRemarks,
		})
	}
	if adku.mutation.RemarksCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: admindictkey.FieldRemarks,
		})
	}
	if value, ok := adku.mutation.IsEnable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: admindictkey.FieldIsEnable,
		})
	}
	if adku.mutation.PCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admindictkey.PTable,
			Columns: []string{admindictkey.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admindict.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := adku.mutation.PIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admindictkey.PTable,
			Columns: []string{admindictkey.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admindict.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, adku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admindictkey.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AdminDictKeyUpdateOne is the builder for updating a single AdminDictKey entity.
type AdminDictKeyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminDictKeyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (adkuo *AdminDictKeyUpdateOne) SetUpdatedAt(t time.Time) *AdminDictKeyUpdateOne {
	adkuo.mutation.SetUpdatedAt(t)
	return adkuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (adkuo *AdminDictKeyUpdateOne) ClearUpdatedAt() *AdminDictKeyUpdateOne {
	adkuo.mutation.ClearUpdatedAt()
	return adkuo
}

// SetDeletedAt sets the "deleted_at" field.
func (adkuo *AdminDictKeyUpdateOne) SetDeletedAt(t time.Time) *AdminDictKeyUpdateOne {
	adkuo.mutation.SetDeletedAt(t)
	return adkuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (adkuo *AdminDictKeyUpdateOne) SetNillableDeletedAt(t *time.Time) *AdminDictKeyUpdateOne {
	if t != nil {
		adkuo.SetDeletedAt(*t)
	}
	return adkuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (adkuo *AdminDictKeyUpdateOne) ClearDeletedAt() *AdminDictKeyUpdateOne {
	adkuo.mutation.ClearDeletedAt()
	return adkuo
}

// SetDictLabel sets the "dict_label" field.
func (adkuo *AdminDictKeyUpdateOne) SetDictLabel(s string) *AdminDictKeyUpdateOne {
	adkuo.mutation.SetDictLabel(s)
	return adkuo
}

// SetDictCode sets the "dict_code" field.
func (adkuo *AdminDictKeyUpdateOne) SetDictCode(s string) *AdminDictKeyUpdateOne {
	adkuo.mutation.SetDictCode(s)
	return adkuo
}

// SetSort sets the "sort" field.
func (adkuo *AdminDictKeyUpdateOne) SetSort(i int) *AdminDictKeyUpdateOne {
	adkuo.mutation.ResetSort()
	adkuo.mutation.SetSort(i)
	return adkuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (adkuo *AdminDictKeyUpdateOne) SetNillableSort(i *int) *AdminDictKeyUpdateOne {
	if i != nil {
		adkuo.SetSort(*i)
	}
	return adkuo
}

// AddSort adds i to the "sort" field.
func (adkuo *AdminDictKeyUpdateOne) AddSort(i int) *AdminDictKeyUpdateOne {
	adkuo.mutation.AddSort(i)
	return adkuo
}

// ClearSort clears the value of the "sort" field.
func (adkuo *AdminDictKeyUpdateOne) ClearSort() *AdminDictKeyUpdateOne {
	adkuo.mutation.ClearSort()
	return adkuo
}

// SetRemarks sets the "remarks" field.
func (adkuo *AdminDictKeyUpdateOne) SetRemarks(s string) *AdminDictKeyUpdateOne {
	adkuo.mutation.SetRemarks(s)
	return adkuo
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (adkuo *AdminDictKeyUpdateOne) SetNillableRemarks(s *string) *AdminDictKeyUpdateOne {
	if s != nil {
		adkuo.SetRemarks(*s)
	}
	return adkuo
}

// ClearRemarks clears the value of the "remarks" field.
func (adkuo *AdminDictKeyUpdateOne) ClearRemarks() *AdminDictKeyUpdateOne {
	adkuo.mutation.ClearRemarks()
	return adkuo
}

// SetIsEnable sets the "is_enable" field.
func (adkuo *AdminDictKeyUpdateOne) SetIsEnable(b bool) *AdminDictKeyUpdateOne {
	adkuo.mutation.SetIsEnable(b)
	return adkuo
}

// SetPID sets the "P" edge to the AdminDict entity by ID.
func (adkuo *AdminDictKeyUpdateOne) SetPID(id int) *AdminDictKeyUpdateOne {
	adkuo.mutation.SetPID(id)
	return adkuo
}

// SetNillablePID sets the "P" edge to the AdminDict entity by ID if the given value is not nil.
func (adkuo *AdminDictKeyUpdateOne) SetNillablePID(id *int) *AdminDictKeyUpdateOne {
	if id != nil {
		adkuo = adkuo.SetPID(*id)
	}
	return adkuo
}

// SetP sets the "P" edge to the AdminDict entity.
func (adkuo *AdminDictKeyUpdateOne) SetP(a *AdminDict) *AdminDictKeyUpdateOne {
	return adkuo.SetPID(a.ID)
}

// Mutation returns the AdminDictKeyMutation object of the builder.
func (adkuo *AdminDictKeyUpdateOne) Mutation() *AdminDictKeyMutation {
	return adkuo.mutation
}

// ClearP clears the "P" edge to the AdminDict entity.
func (adkuo *AdminDictKeyUpdateOne) ClearP() *AdminDictKeyUpdateOne {
	adkuo.mutation.ClearP()
	return adkuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (adkuo *AdminDictKeyUpdateOne) Select(field string, fields ...string) *AdminDictKeyUpdateOne {
	adkuo.fields = append([]string{field}, fields...)
	return adkuo
}

// Save executes the query and returns the updated AdminDictKey entity.
func (adkuo *AdminDictKeyUpdateOne) Save(ctx context.Context) (*AdminDictKey, error) {
	var (
		err  error
		node *AdminDictKey
	)
	adkuo.defaults()
	if len(adkuo.hooks) == 0 {
		node, err = adkuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminDictKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			adkuo.mutation = mutation
			node, err = adkuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(adkuo.hooks) - 1; i >= 0; i-- {
			mut = adkuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, adkuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (adkuo *AdminDictKeyUpdateOne) SaveX(ctx context.Context) *AdminDictKey {
	node, err := adkuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (adkuo *AdminDictKeyUpdateOne) Exec(ctx context.Context) error {
	_, err := adkuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adkuo *AdminDictKeyUpdateOne) ExecX(ctx context.Context) {
	if err := adkuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (adkuo *AdminDictKeyUpdateOne) defaults() {
	if _, ok := adkuo.mutation.UpdatedAt(); !ok && !adkuo.mutation.UpdatedAtCleared() {
		v := admindictkey.UpdateDefaultUpdatedAt()
		adkuo.mutation.SetUpdatedAt(v)
	}
}

func (adkuo *AdminDictKeyUpdateOne) sqlSave(ctx context.Context) (_node *AdminDictKey, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admindictkey.Table,
			Columns: admindictkey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admindictkey.FieldID,
			},
		},
	}
	id, ok := adkuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AdminDictKey.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := adkuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, admindictkey.FieldID)
		for _, f := range fields {
			if !admindictkey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != admindictkey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := adkuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := adkuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admindictkey.FieldUpdatedAt,
		})
	}
	if adkuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: admindictkey.FieldUpdatedAt,
		})
	}
	if value, ok := adkuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admindictkey.FieldDeletedAt,
		})
	}
	if adkuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: admindictkey.FieldDeletedAt,
		})
	}
	if value, ok := adkuo.mutation.DictLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindictkey.FieldDictLabel,
		})
	}
	if value, ok := adkuo.mutation.DictCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindictkey.FieldDictCode,
		})
	}
	if value, ok := adkuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admindictkey.FieldSort,
		})
	}
	if value, ok := adkuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admindictkey.FieldSort,
		})
	}
	if adkuo.mutation.SortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: admindictkey.FieldSort,
		})
	}
	if value, ok := adkuo.mutation.Remarks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admindictkey.FieldRemarks,
		})
	}
	if adkuo.mutation.RemarksCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: admindictkey.FieldRemarks,
		})
	}
	if value, ok := adkuo.mutation.IsEnable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: admindictkey.FieldIsEnable,
		})
	}
	if adkuo.mutation.PCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admindictkey.PTable,
			Columns: []string{admindictkey.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admindict.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := adkuo.mutation.PIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admindictkey.PTable,
			Columns: []string{admindictkey.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admindict.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdminDictKey{config: adkuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, adkuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admindictkey.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}