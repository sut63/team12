// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClubappStatusQuery is the builder for querying ClubappStatus entities.
type ClubappStatusQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.ClubappStatus
	// eager-loading edges.
	withClubapplication *ClubapplicationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (csq *ClubappStatusQuery) Where(ps ...predicate.ClubappStatus) *ClubappStatusQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit adds a limit step to the query.
func (csq *ClubappStatusQuery) Limit(limit int) *ClubappStatusQuery {
	csq.limit = &limit
	return csq
}

// Offset adds an offset step to the query.
func (csq *ClubappStatusQuery) Offset(offset int) *ClubappStatusQuery {
	csq.offset = &offset
	return csq
}

// Order adds an order step to the query.
func (csq *ClubappStatusQuery) Order(o ...OrderFunc) *ClubappStatusQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// QueryClubapplication chains the current query on the clubapplication edge.
func (csq *ClubappStatusQuery) QueryClubapplication() *ClubapplicationQuery {
	query := &ClubapplicationQuery{config: csq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clubappstatus.Table, clubappstatus.FieldID, csq.sqlQuery()),
			sqlgraph.To(clubapplication.Table, clubapplication.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clubappstatus.ClubapplicationTable, clubappstatus.ClubapplicationColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClubappStatus entity in the query. Returns *NotFoundError when no clubappstatus was found.
func (csq *ClubappStatusQuery) First(ctx context.Context) (*ClubappStatus, error) {
	csSlice, err := csq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(csSlice) == 0 {
		return nil, &NotFoundError{clubappstatus.Label}
	}
	return csSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *ClubappStatusQuery) FirstX(ctx context.Context) *ClubappStatus {
	cs, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return cs
}

// FirstID returns the first ClubappStatus id in the query. Returns *NotFoundError when no id was found.
func (csq *ClubappStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{clubappstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (csq *ClubappStatusQuery) FirstXID(ctx context.Context) int {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ClubappStatus entity in the query, returns an error if not exactly one entity was returned.
func (csq *ClubappStatusQuery) Only(ctx context.Context) (*ClubappStatus, error) {
	csSlice, err := csq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(csSlice) {
	case 1:
		return csSlice[0], nil
	case 0:
		return nil, &NotFoundError{clubappstatus.Label}
	default:
		return nil, &NotSingularError{clubappstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *ClubappStatusQuery) OnlyX(ctx context.Context) *ClubappStatus {
	cs, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// OnlyID returns the only ClubappStatus id in the query, returns an error if not exactly one id was returned.
func (csq *ClubappStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = &NotSingularError{clubappstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *ClubappStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClubappStatusSlice.
func (csq *ClubappStatusQuery) All(ctx context.Context) ([]*ClubappStatus, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return csq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (csq *ClubappStatusQuery) AllX(ctx context.Context) []*ClubappStatus {
	csSlice, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return csSlice
}

// IDs executes the query and returns a list of ClubappStatus ids.
func (csq *ClubappStatusQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := csq.Select(clubappstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *ClubappStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *ClubappStatusQuery) Count(ctx context.Context) (int, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return csq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (csq *ClubappStatusQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *ClubappStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return csq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *ClubappStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *ClubappStatusQuery) Clone() *ClubappStatusQuery {
	return &ClubappStatusQuery{
		config:     csq.config,
		limit:      csq.limit,
		offset:     csq.offset,
		order:      append([]OrderFunc{}, csq.order...),
		unique:     append([]string{}, csq.unique...),
		predicates: append([]predicate.ClubappStatus{}, csq.predicates...),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

//  WithClubapplication tells the query-builder to eager-loads the nodes that are connected to
// the "clubapplication" edge. The optional arguments used to configure the query builder of the edge.
func (csq *ClubappStatusQuery) WithClubapplication(opts ...func(*ClubapplicationQuery)) *ClubappStatusQuery {
	query := &ClubapplicationQuery{config: csq.config}
	for _, opt := range opts {
		opt(query)
	}
	csq.withClubapplication = query
	return csq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ApplyStatus string `json:"apply_status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClubappStatus.Query().
//		GroupBy(clubappstatus.FieldApplyStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (csq *ClubappStatusQuery) GroupBy(field string, fields ...string) *ClubappStatusGroupBy {
	group := &ClubappStatusGroupBy{config: csq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return csq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		ApplyStatus string `json:"apply_status,omitempty"`
//	}
//
//	client.ClubappStatus.Query().
//		Select(clubappstatus.FieldApplyStatus).
//		Scan(ctx, &v)
//
func (csq *ClubappStatusQuery) Select(field string, fields ...string) *ClubappStatusSelect {
	selector := &ClubappStatusSelect{config: csq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return csq.sqlQuery(), nil
	}
	return selector
}

func (csq *ClubappStatusQuery) prepareQuery(ctx context.Context) error {
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *ClubappStatusQuery) sqlAll(ctx context.Context) ([]*ClubappStatus, error) {
	var (
		nodes       = []*ClubappStatus{}
		_spec       = csq.querySpec()
		loadedTypes = [1]bool{
			csq.withClubapplication != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &ClubappStatus{config: csq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := csq.withClubapplication; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ClubappStatus)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Clubapplication(func(s *sql.Selector) {
			s.Where(sql.InValues(clubappstatus.ClubapplicationColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.clubappstatus_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "clubappstatus_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "clubappstatus_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Clubapplication = append(node.Edges.Clubapplication, n)
		}
	}

	return nodes, nil
}

func (csq *ClubappStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *ClubappStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := csq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (csq *ClubappStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clubappstatus.Table,
			Columns: clubappstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clubappstatus.FieldID,
			},
		},
		From:   csq.sql,
		Unique: true,
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *ClubappStatusQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(clubappstatus.Table)
	selector := builder.Select(t1.Columns(clubappstatus.Columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(clubappstatus.Columns...)...)
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClubappStatusGroupBy is the builder for group-by ClubappStatus entities.
type ClubappStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *ClubappStatusGroupBy) Aggregate(fns ...AggregateFunc) *ClubappStatusGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the group-by query and scan the result into the given value.
func (csgb *ClubappStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := csgb.path(ctx)
	if err != nil {
		return err
	}
	csgb.sql = query
	return csgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := csgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := csgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = csgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) StringX(ctx context.Context) string {
	v, err := csgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := csgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = csgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) IntX(ctx context.Context) int {
	v, err := csgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := csgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = csgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := csgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := csgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (csgb *ClubappStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = csgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (csgb *ClubappStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := csgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (csgb *ClubappStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := csgb.sqlQuery().Query()
	if err := csgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (csgb *ClubappStatusGroupBy) sqlQuery() *sql.Selector {
	selector := csgb.sql
	columns := make([]string, 0, len(csgb.fields)+len(csgb.fns))
	columns = append(columns, csgb.fields...)
	for _, fn := range csgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(csgb.fields...)
}

// ClubappStatusSelect is the builder for select fields of ClubappStatus entities.
type ClubappStatusSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (css *ClubappStatusSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := css.path(ctx)
	if err != nil {
		return err
	}
	css.sql = query
	return css.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (css *ClubappStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := css.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (css *ClubappStatusSelect) StringsX(ctx context.Context) []string {
	v, err := css.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = css.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (css *ClubappStatusSelect) StringX(ctx context.Context) string {
	v, err := css.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (css *ClubappStatusSelect) IntsX(ctx context.Context) []int {
	v, err := css.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = css.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (css *ClubappStatusSelect) IntX(ctx context.Context) int {
	v, err := css.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (css *ClubappStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := css.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = css.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (css *ClubappStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := css.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: ClubappStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (css *ClubappStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := css.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (css *ClubappStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = css.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{clubappstatus.Label}
	default:
		err = fmt.Errorf("ent: ClubappStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (css *ClubappStatusSelect) BoolX(ctx context.Context) bool {
	v, err := css.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (css *ClubappStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := css.sqlQuery().Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (css *ClubappStatusSelect) sqlQuery() sql.Querier {
	selector := css.sql
	selector.Select(selector.Columns(css.fields...)...)
	return selector
}
