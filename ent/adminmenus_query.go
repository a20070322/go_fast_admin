// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/go_fast_admin/ent/adminmenus"
	"github.com/a20070322/go_fast_admin/ent/adminrole"
	"github.com/a20070322/go_fast_admin/ent/predicate"
)

// AdminMenusQuery is the builder for querying AdminMenus entities.
type AdminMenusQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AdminMenus
	// eager-loading edges.
	withRole *AdminRoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminMenusQuery builder.
func (amq *AdminMenusQuery) Where(ps ...predicate.AdminMenus) *AdminMenusQuery {
	amq.predicates = append(amq.predicates, ps...)
	return amq
}

// Limit adds a limit step to the query.
func (amq *AdminMenusQuery) Limit(limit int) *AdminMenusQuery {
	amq.limit = &limit
	return amq
}

// Offset adds an offset step to the query.
func (amq *AdminMenusQuery) Offset(offset int) *AdminMenusQuery {
	amq.offset = &offset
	return amq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (amq *AdminMenusQuery) Unique(unique bool) *AdminMenusQuery {
	amq.unique = &unique
	return amq
}

// Order adds an order step to the query.
func (amq *AdminMenusQuery) Order(o ...OrderFunc) *AdminMenusQuery {
	amq.order = append(amq.order, o...)
	return amq
}

// QueryRole chains the current query on the "role" edge.
func (amq *AdminMenusQuery) QueryRole() *AdminRoleQuery {
	query := &AdminRoleQuery{config: amq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adminmenus.Table, adminmenus.FieldID, selector),
			sqlgraph.To(adminrole.Table, adminrole.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, adminmenus.RoleTable, adminmenus.RolePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(amq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AdminMenus entity from the query.
// Returns a *NotFoundError when no AdminMenus was found.
func (amq *AdminMenusQuery) First(ctx context.Context) (*AdminMenus, error) {
	nodes, err := amq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminmenus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (amq *AdminMenusQuery) FirstX(ctx context.Context) *AdminMenus {
	node, err := amq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminMenus ID from the query.
// Returns a *NotFoundError when no AdminMenus ID was found.
func (amq *AdminMenusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = amq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminmenus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (amq *AdminMenusQuery) FirstIDX(ctx context.Context) int {
	id, err := amq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminMenus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AdminMenus entity is not found.
// Returns a *NotFoundError when no AdminMenus entities are found.
func (amq *AdminMenusQuery) Only(ctx context.Context) (*AdminMenus, error) {
	nodes, err := amq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminmenus.Label}
	default:
		return nil, &NotSingularError{adminmenus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (amq *AdminMenusQuery) OnlyX(ctx context.Context) *AdminMenus {
	node, err := amq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminMenus ID in the query.
// Returns a *NotSingularError when exactly one AdminMenus ID is not found.
// Returns a *NotFoundError when no entities are found.
func (amq *AdminMenusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = amq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = &NotSingularError{adminmenus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (amq *AdminMenusQuery) OnlyIDX(ctx context.Context) int {
	id, err := amq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminMenusSlice.
func (amq *AdminMenusQuery) All(ctx context.Context) ([]*AdminMenus, error) {
	if err := amq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return amq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (amq *AdminMenusQuery) AllX(ctx context.Context) []*AdminMenus {
	nodes, err := amq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminMenus IDs.
func (amq *AdminMenusQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := amq.Select(adminmenus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (amq *AdminMenusQuery) IDsX(ctx context.Context) []int {
	ids, err := amq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (amq *AdminMenusQuery) Count(ctx context.Context) (int, error) {
	if err := amq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return amq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (amq *AdminMenusQuery) CountX(ctx context.Context) int {
	count, err := amq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (amq *AdminMenusQuery) Exist(ctx context.Context) (bool, error) {
	if err := amq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return amq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (amq *AdminMenusQuery) ExistX(ctx context.Context) bool {
	exist, err := amq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminMenusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (amq *AdminMenusQuery) Clone() *AdminMenusQuery {
	if amq == nil {
		return nil
	}
	return &AdminMenusQuery{
		config:     amq.config,
		limit:      amq.limit,
		offset:     amq.offset,
		order:      append([]OrderFunc{}, amq.order...),
		predicates: append([]predicate.AdminMenus{}, amq.predicates...),
		withRole:   amq.withRole.Clone(),
		// clone intermediate query.
		sql:  amq.sql.Clone(),
		path: amq.path,
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (amq *AdminMenusQuery) WithRole(opts ...func(*AdminRoleQuery)) *AdminMenusQuery {
	query := &AdminRoleQuery{config: amq.config}
	for _, opt := range opts {
		opt(query)
	}
	amq.withRole = query
	return amq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminMenus.Query().
//		GroupBy(adminmenus.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (amq *AdminMenusQuery) GroupBy(field string, fields ...string) *AdminMenusGroupBy {
	group := &AdminMenusGroupBy{config: amq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := amq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return amq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AdminMenus.Query().
//		Select(adminmenus.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (amq *AdminMenusQuery) Select(field string, fields ...string) *AdminMenusSelect {
	amq.fields = append([]string{field}, fields...)
	return &AdminMenusSelect{AdminMenusQuery: amq}
}

func (amq *AdminMenusQuery) prepareQuery(ctx context.Context) error {
	for _, f := range amq.fields {
		if !adminmenus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if amq.path != nil {
		prev, err := amq.path(ctx)
		if err != nil {
			return err
		}
		amq.sql = prev
	}
	return nil
}

func (amq *AdminMenusQuery) sqlAll(ctx context.Context) ([]*AdminMenus, error) {
	var (
		nodes       = []*AdminMenus{}
		_spec       = amq.querySpec()
		loadedTypes = [1]bool{
			amq.withRole != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AdminMenus{config: amq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, amq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := amq.withRole; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*AdminMenus, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Role = []*AdminRole{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*AdminMenus)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   adminmenus.RoleTable,
				Columns: adminmenus.RolePrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(adminmenus.RolePrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, amq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "role": %w`, err)
		}
		query.Where(adminrole.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "role" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Role = append(nodes[i].Edges.Role, n)
			}
		}
	}

	return nodes, nil
}

func (amq *AdminMenusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := amq.querySpec()
	return sqlgraph.CountNodes(ctx, amq.driver, _spec)
}

func (amq *AdminMenusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := amq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (amq *AdminMenusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adminmenus.Table,
			Columns: adminmenus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adminmenus.FieldID,
			},
		},
		From:   amq.sql,
		Unique: true,
	}
	if unique := amq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := amq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminmenus.FieldID)
		for i := range fields {
			if fields[i] != adminmenus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := amq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := amq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := amq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := amq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (amq *AdminMenusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(amq.driver.Dialect())
	t1 := builder.Table(adminmenus.Table)
	selector := builder.Select(t1.Columns(adminmenus.Columns...)...).From(t1)
	if amq.sql != nil {
		selector = amq.sql
		selector.Select(selector.Columns(adminmenus.Columns...)...)
	}
	for _, p := range amq.predicates {
		p(selector)
	}
	for _, p := range amq.order {
		p(selector)
	}
	if offset := amq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := amq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdminMenusGroupBy is the group-by builder for AdminMenus entities.
type AdminMenusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (amgb *AdminMenusGroupBy) Aggregate(fns ...AggregateFunc) *AdminMenusGroupBy {
	amgb.fns = append(amgb.fns, fns...)
	return amgb
}

// Scan applies the group-by query and scans the result into the given value.
func (amgb *AdminMenusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := amgb.path(ctx)
	if err != nil {
		return err
	}
	amgb.sql = query
	return amgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := amgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(amgb.fields) > 1 {
		return nil, errors.New("ent: AdminMenusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := amgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) StringsX(ctx context.Context) []string {
	v, err := amgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) StringX(ctx context.Context) string {
	v, err := amgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(amgb.fields) > 1 {
		return nil, errors.New("ent: AdminMenusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := amgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) IntsX(ctx context.Context) []int {
	v, err := amgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) IntX(ctx context.Context) int {
	v, err := amgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(amgb.fields) > 1 {
		return nil, errors.New("ent: AdminMenusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := amgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := amgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := amgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(amgb.fields) > 1 {
		return nil, errors.New("ent: AdminMenusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := amgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := amgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amgb *AdminMenusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amgb *AdminMenusGroupBy) BoolX(ctx context.Context) bool {
	v, err := amgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amgb *AdminMenusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range amgb.fields {
		if !adminmenus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := amgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := amgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (amgb *AdminMenusGroupBy) sqlQuery() *sql.Selector {
	selector := amgb.sql
	columns := make([]string, 0, len(amgb.fields)+len(amgb.fns))
	columns = append(columns, amgb.fields...)
	for _, fn := range amgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(amgb.fields...)
}

// AdminMenusSelect is the builder for selecting fields of AdminMenus entities.
type AdminMenusSelect struct {
	*AdminMenusQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ams *AdminMenusSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ams.prepareQuery(ctx); err != nil {
		return err
	}
	ams.sql = ams.AdminMenusQuery.sqlQuery(ctx)
	return ams.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ams *AdminMenusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ams.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ams.fields) > 1 {
		return nil, errors.New("ent: AdminMenusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ams.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ams *AdminMenusSelect) StringsX(ctx context.Context) []string {
	v, err := ams.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ams.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ams *AdminMenusSelect) StringX(ctx context.Context) string {
	v, err := ams.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ams.fields) > 1 {
		return nil, errors.New("ent: AdminMenusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ams.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ams *AdminMenusSelect) IntsX(ctx context.Context) []int {
	v, err := ams.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ams.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ams *AdminMenusSelect) IntX(ctx context.Context) int {
	v, err := ams.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ams.fields) > 1 {
		return nil, errors.New("ent: AdminMenusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ams.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ams *AdminMenusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ams.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ams.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ams *AdminMenusSelect) Float64X(ctx context.Context) float64 {
	v, err := ams.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ams.fields) > 1 {
		return nil, errors.New("ent: AdminMenusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ams.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ams *AdminMenusSelect) BoolsX(ctx context.Context) []bool {
	v, err := ams.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ams *AdminMenusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ams.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adminmenus.Label}
	default:
		err = fmt.Errorf("ent: AdminMenusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ams *AdminMenusSelect) BoolX(ctx context.Context) bool {
	v, err := ams.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ams *AdminMenusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ams.sqlQuery().Query()
	if err := ams.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ams *AdminMenusSelect) sqlQuery() sql.Querier {
	selector := ams.sql
	selector.Select(selector.Columns(ams.fields...)...)
	return selector
}
