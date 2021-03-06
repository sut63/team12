// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/OMENX/app/ent/predicate"
	"github.com/OMENX/app/ent/purpose"
	"github.com/OMENX/app/ent/room"
	"github.com/OMENX/app/ent/roomuse"
	"github.com/OMENX/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomuseQuery is the builder for querying Roomuse entities.
type RoomuseQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Roomuse
	// eager-loading edges.
	withRooms    *RoomQuery
	withPurposes *PurposeQuery
	withUsers    *UserQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rq *RoomuseQuery) Where(ps ...predicate.Roomuse) *RoomuseQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RoomuseQuery) Limit(limit int) *RoomuseQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RoomuseQuery) Offset(offset int) *RoomuseQuery {
	rq.offset = &offset
	return rq
}

// Order adds an order step to the query.
func (rq *RoomuseQuery) Order(o ...OrderFunc) *RoomuseQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryRooms chains the current query on the rooms edge.
func (rq *RoomuseQuery) QueryRooms() *RoomQuery {
	query := &RoomQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomuse.Table, roomuse.FieldID, rq.sqlQuery()),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomuse.RoomsTable, roomuse.RoomsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPurposes chains the current query on the purposes edge.
func (rq *RoomuseQuery) QueryPurposes() *PurposeQuery {
	query := &PurposeQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomuse.Table, roomuse.FieldID, rq.sqlQuery()),
			sqlgraph.To(purpose.Table, purpose.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomuse.PurposesTable, roomuse.PurposesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the users edge.
func (rq *RoomuseQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomuse.Table, roomuse.FieldID, rq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomuse.UsersTable, roomuse.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Roomuse entity in the query. Returns *NotFoundError when no roomuse was found.
func (rq *RoomuseQuery) First(ctx context.Context) (*Roomuse, error) {
	rs, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rs) == 0 {
		return nil, &NotFoundError{roomuse.Label}
	}
	return rs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoomuseQuery) FirstX(ctx context.Context) *Roomuse {
	r, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return r
}

// FirstID returns the first Roomuse id in the query. Returns *NotFoundError when no id was found.
func (rq *RoomuseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{roomuse.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rq *RoomuseQuery) FirstXID(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Roomuse entity in the query, returns an error if not exactly one entity was returned.
func (rq *RoomuseQuery) Only(ctx context.Context) (*Roomuse, error) {
	rs, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rs) {
	case 1:
		return rs[0], nil
	case 0:
		return nil, &NotFoundError{roomuse.Label}
	default:
		return nil, &NotSingularError{roomuse.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoomuseQuery) OnlyX(ctx context.Context) *Roomuse {
	r, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// OnlyID returns the only Roomuse id in the query, returns an error if not exactly one id was returned.
func (rq *RoomuseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = &NotSingularError{roomuse.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoomuseQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Roomuses.
func (rq *RoomuseQuery) All(ctx context.Context) ([]*Roomuse, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoomuseQuery) AllX(ctx context.Context) []*Roomuse {
	rs, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// IDs executes the query and returns a list of Roomuse ids.
func (rq *RoomuseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(roomuse.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoomuseQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoomuseQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoomuseQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoomuseQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoomuseQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoomuseQuery) Clone() *RoomuseQuery {
	return &RoomuseQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		unique:     append([]string{}, rq.unique...),
		predicates: append([]predicate.Roomuse{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

//  WithRooms tells the query-builder to eager-loads the nodes that are connected to
// the "rooms" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomuseQuery) WithRooms(opts ...func(*RoomQuery)) *RoomuseQuery {
	query := &RoomQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withRooms = query
	return rq
}

//  WithPurposes tells the query-builder to eager-loads the nodes that are connected to
// the "purposes" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomuseQuery) WithPurposes(opts ...func(*PurposeQuery)) *RoomuseQuery {
	query := &PurposeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withPurposes = query
	return rq
}

//  WithUsers tells the query-builder to eager-loads the nodes that are connected to
// the "users" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomuseQuery) WithUsers(opts ...func(*UserQuery)) *RoomuseQuery {
	query := &UserQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withUsers = query
	return rq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		People int `json:"people,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Roomuse.Query().
//		GroupBy(roomuse.FieldPeople).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RoomuseQuery) GroupBy(field string, fields ...string) *RoomuseGroupBy {
	group := &RoomuseGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		People int `json:"people,omitempty"`
//	}
//
//	client.Roomuse.Query().
//		Select(roomuse.FieldPeople).
//		Scan(ctx, &v)
//
func (rq *RoomuseQuery) Select(field string, fields ...string) *RoomuseSelect {
	selector := &RoomuseSelect{config: rq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return selector
}

func (rq *RoomuseQuery) prepareQuery(ctx context.Context) error {
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RoomuseQuery) sqlAll(ctx context.Context) ([]*Roomuse, error) {
	var (
		nodes       = []*Roomuse{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withRooms != nil,
			rq.withPurposes != nil,
			rq.withUsers != nil,
		}
	)
	if rq.withRooms != nil || rq.withPurposes != nil || rq.withUsers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, roomuse.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Roomuse{config: rq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withRooms; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomuse)
		for i := range nodes {
			if fk := nodes[i].room_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(room.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Rooms = n
			}
		}
	}

	if query := rq.withPurposes; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomuse)
		for i := range nodes {
			if fk := nodes[i].purpose_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(purpose.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "purpose_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Purposes = n
			}
		}
	}

	if query := rq.withUsers; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomuse)
		for i := range nodes {
			if fk := nodes[i].UserID; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "UserID" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Users = n
			}
		}
	}

	return nodes, nil
}

func (rq *RoomuseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoomuseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rq *RoomuseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomuse.Table,
			Columns: roomuse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomuse.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RoomuseQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(roomuse.Table)
	selector := builder.Select(t1.Columns(roomuse.Columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(roomuse.Columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoomuseGroupBy is the builder for group-by Roomuse entities.
type RoomuseGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoomuseGroupBy) Aggregate(fns ...AggregateFunc) *RoomuseGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rgb *RoomuseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *RoomuseGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomuseGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *RoomuseGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *RoomuseGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomuseGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *RoomuseGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *RoomuseGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomuseGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *RoomuseGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *RoomuseGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomuseGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *RoomuseGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomuseGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *RoomuseGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *RoomuseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rgb.sqlQuery().Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RoomuseGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql
	columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
	columns = append(columns, rgb.fields...)
	for _, fn := range rgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rgb.fields...)
}

// RoomuseSelect is the builder for select fields of Roomuse entities.
type RoomuseSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rs *RoomuseSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rs.path(ctx)
	if err != nil {
		return err
	}
	rs.sql = query
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *RoomuseSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomuseSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *RoomuseSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *RoomuseSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomuseSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *RoomuseSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *RoomuseSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomuseSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *RoomuseSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *RoomuseSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomuseSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *RoomuseSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rs *RoomuseSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomuse.Label}
	default:
		err = fmt.Errorf("ent: RoomuseSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *RoomuseSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *RoomuseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sqlQuery().Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rs *RoomuseSelect) sqlQuery() sql.Querier {
	selector := rs.sql
	selector.Select(selector.Columns(rs.fields...)...)
	return selector
}
