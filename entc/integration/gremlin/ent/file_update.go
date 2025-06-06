// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/file"
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (_u *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetSetID sets the "set_id" field.
func (_u *FileUpdate) SetSetID(v int) *FileUpdate {
	_u.mutation.ResetSetID()
	_u.mutation.SetSetID(v)
	return _u
}

// SetNillableSetID sets the "set_id" field if the given value is not nil.
func (_u *FileUpdate) SetNillableSetID(v *int) *FileUpdate {
	if v != nil {
		_u.SetSetID(*v)
	}
	return _u
}

// AddSetID adds value to the "set_id" field.
func (_u *FileUpdate) AddSetID(v int) *FileUpdate {
	_u.mutation.AddSetID(v)
	return _u
}

// ClearSetID clears the value of the "set_id" field.
func (_u *FileUpdate) ClearSetID() *FileUpdate {
	_u.mutation.ClearSetID()
	return _u
}

// SetSize sets the "size" field.
func (_u *FileUpdate) SetSize(v int) *FileUpdate {
	_u.mutation.ResetSize()
	_u.mutation.SetSize(v)
	return _u
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (_u *FileUpdate) SetNillableSize(v *int) *FileUpdate {
	if v != nil {
		_u.SetSize(*v)
	}
	return _u
}

// AddSize adds value to the "size" field.
func (_u *FileUpdate) AddSize(v int) *FileUpdate {
	_u.mutation.AddSize(v)
	return _u
}

// SetName sets the "name" field.
func (_u *FileUpdate) SetName(v string) *FileUpdate {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *FileUpdate) SetNillableName(v *string) *FileUpdate {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetUser sets the "user" field.
func (_u *FileUpdate) SetUser(v string) *FileUpdate {
	_u.mutation.SetUser(v)
	return _u
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (_u *FileUpdate) SetNillableUser(v *string) *FileUpdate {
	if v != nil {
		_u.SetUser(*v)
	}
	return _u
}

// ClearUser clears the value of the "user" field.
func (_u *FileUpdate) ClearUser() *FileUpdate {
	_u.mutation.ClearUser()
	return _u
}

// SetGroup sets the "group" field.
func (_u *FileUpdate) SetGroup(v string) *FileUpdate {
	_u.mutation.SetGroup(v)
	return _u
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (_u *FileUpdate) SetNillableGroup(v *string) *FileUpdate {
	if v != nil {
		_u.SetGroup(*v)
	}
	return _u
}

// ClearGroup clears the value of the "group" field.
func (_u *FileUpdate) ClearGroup() *FileUpdate {
	_u.mutation.ClearGroup()
	return _u
}

// SetOp sets the "op" field.
func (_u *FileUpdate) SetOp(v bool) *FileUpdate {
	_u.mutation.SetOpField(v)
	return _u
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (_u *FileUpdate) SetNillableOp(v *bool) *FileUpdate {
	if v != nil {
		_u.SetOp(*v)
	}
	return _u
}

// ClearOp clears the value of the "op" field.
func (_u *FileUpdate) ClearOp() *FileUpdate {
	_u.mutation.ClearOp()
	return _u
}

// SetFieldID sets the "field_id" field.
func (_u *FileUpdate) SetFieldID(v int) *FileUpdate {
	_u.mutation.ResetFieldID()
	_u.mutation.SetFieldID(v)
	return _u
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (_u *FileUpdate) SetNillableFieldID(v *int) *FileUpdate {
	if v != nil {
		_u.SetFieldID(*v)
	}
	return _u
}

// AddFieldID adds value to the "field_id" field.
func (_u *FileUpdate) AddFieldID(v int) *FileUpdate {
	_u.mutation.AddFieldID(v)
	return _u
}

// ClearFieldID clears the value of the "field_id" field.
func (_u *FileUpdate) ClearFieldID() *FileUpdate {
	_u.mutation.ClearFieldID()
	return _u
}

// SetCreateTime sets the "create_time" field.
func (_u *FileUpdate) SetCreateTime(v time.Time) *FileUpdate {
	_u.mutation.SetCreateTime(v)
	return _u
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (_u *FileUpdate) SetNillableCreateTime(v *time.Time) *FileUpdate {
	if v != nil {
		_u.SetCreateTime(*v)
	}
	return _u
}

// ClearCreateTime clears the value of the "create_time" field.
func (_u *FileUpdate) ClearCreateTime() *FileUpdate {
	_u.mutation.ClearCreateTime()
	return _u
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_u *FileUpdate) SetOwnerID(id string) *FileUpdate {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_u *FileUpdate) SetNillableOwnerID(id *string) *FileUpdate {
	if id != nil {
		_u = _u.SetOwnerID(*id)
	}
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *FileUpdate) SetOwner(v *User) *FileUpdate {
	return _u.SetOwnerID(v.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (_u *FileUpdate) SetTypeID(id string) *FileUpdate {
	_u.mutation.SetTypeID(id)
	return _u
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (_u *FileUpdate) SetNillableTypeID(id *string) *FileUpdate {
	if id != nil {
		_u = _u.SetTypeID(*id)
	}
	return _u
}

// SetType sets the "type" edge to the FileType entity.
func (_u *FileUpdate) SetType(v *FileType) *FileUpdate {
	return _u.SetTypeID(v.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (_u *FileUpdate) AddFieldIDs(ids ...string) *FileUpdate {
	_u.mutation.AddFieldIDs(ids...)
	return _u
}

// AddField adds the "field" edges to the FieldType entity.
func (_u *FileUpdate) AddField(v ...*FieldType) *FileUpdate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (_u *FileUpdate) Mutation() *FileMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *FileUpdate) ClearOwner() *FileUpdate {
	_u.mutation.ClearOwner()
	return _u
}

// ClearType clears the "type" edge to the FileType entity.
func (_u *FileUpdate) ClearType() *FileUpdate {
	_u.mutation.ClearType()
	return _u
}

// ClearFieldEdge clears all "field" edges to the FieldType entity.
func (_u *FileUpdate) ClearFieldEdge() *FileUpdate {
	_u.mutation.ClearFieldEdge()
	return _u
}

// RemoveFieldIDs removes the "field" edge to FieldType entities by IDs.
func (_u *FileUpdate) RemoveFieldIDs(ids ...string) *FileUpdate {
	_u.mutation.RemoveFieldIDs(ids...)
	return _u
}

// RemoveField removes "field" edges to FieldType entities.
func (_u *FileUpdate) RemoveField(v ...*FieldType) *FileUpdate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveFieldIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *FileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.gremlinSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *FileUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *FileUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *FileUpdate) check() error {
	if v, ok := _u.mutation.SetID(); ok {
		if err := file.SetIDValidator(v); err != nil {
			return &ValidationError{Name: "set_id", err: fmt.Errorf(`ent: validator failed for field "File.set_id": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "File.size": %w`, err)}
		}
	}
	return nil
}

func (_u *FileUpdate) gremlinSave(ctx context.Context) (int, error) {
	if err := _u.check(); err != nil {
		return 0, err
	}
	res := &gremlin.Response{}
	query, bindings := _u.gremlin().Query()
	if err := _u.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	_u.mutation.done = true
	return res.ReadInt()
}

func (_u *FileUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(file.Label)
	for _, p := range _u.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := _u.mutation.SetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, value)
	}
	if value, ok := _u.mutation.AddedSetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, __.Union(__.Values(file.FieldSetID), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := _u.mutation.AddedSize(); ok {
		v.Property(dsl.Single, file.FieldSize, __.Union(__.Values(file.FieldSize), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := _u.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := _u.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := _u.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	if value, ok := _u.mutation.FieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, value)
	}
	if value, ok := _u.mutation.AddedFieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, __.Union(__.Values(file.FieldFieldID), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.CreateTime(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldCreateTime, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldCreateTime, value)),
		})
		v.Property(dsl.Single, file.FieldCreateTime, value)
	}
	var properties []any
	if _u.mutation.SetIDCleared() {
		properties = append(properties, file.FieldSetID)
	}
	if _u.mutation.UserCleared() {
		properties = append(properties, file.FieldUser)
	}
	if _u.mutation.GroupCleared() {
		properties = append(properties, file.FieldGroup)
	}
	if _u.mutation.OpCleared() {
		properties = append(properties, file.FieldOp)
	}
	if _u.mutation.FieldIDCleared() {
		properties = append(properties, file.FieldFieldID)
	}
	if _u.mutation.CreateTimeCleared() {
		properties = append(properties, file.FieldCreateTime)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if _u.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	if _u.mutation.TypeCleared() {
		tr := rv.Clone().InE(filetype.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.TypeIDs() {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range _u.mutation.RemovedFieldIDs() {
		tr := rv.Clone().OutE(file.FieldLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.FieldIDs() {
		v.AddE(file.FieldLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(file.FieldLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(file.Label, file.FieldLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetSetID sets the "set_id" field.
func (_u *FileUpdateOne) SetSetID(v int) *FileUpdateOne {
	_u.mutation.ResetSetID()
	_u.mutation.SetSetID(v)
	return _u
}

// SetNillableSetID sets the "set_id" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableSetID(v *int) *FileUpdateOne {
	if v != nil {
		_u.SetSetID(*v)
	}
	return _u
}

// AddSetID adds value to the "set_id" field.
func (_u *FileUpdateOne) AddSetID(v int) *FileUpdateOne {
	_u.mutation.AddSetID(v)
	return _u
}

// ClearSetID clears the value of the "set_id" field.
func (_u *FileUpdateOne) ClearSetID() *FileUpdateOne {
	_u.mutation.ClearSetID()
	return _u
}

// SetSize sets the "size" field.
func (_u *FileUpdateOne) SetSize(v int) *FileUpdateOne {
	_u.mutation.ResetSize()
	_u.mutation.SetSize(v)
	return _u
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableSize(v *int) *FileUpdateOne {
	if v != nil {
		_u.SetSize(*v)
	}
	return _u
}

// AddSize adds value to the "size" field.
func (_u *FileUpdateOne) AddSize(v int) *FileUpdateOne {
	_u.mutation.AddSize(v)
	return _u
}

// SetName sets the "name" field.
func (_u *FileUpdateOne) SetName(v string) *FileUpdateOne {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableName(v *string) *FileUpdateOne {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetUser sets the "user" field.
func (_u *FileUpdateOne) SetUser(v string) *FileUpdateOne {
	_u.mutation.SetUser(v)
	return _u
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableUser(v *string) *FileUpdateOne {
	if v != nil {
		_u.SetUser(*v)
	}
	return _u
}

// ClearUser clears the value of the "user" field.
func (_u *FileUpdateOne) ClearUser() *FileUpdateOne {
	_u.mutation.ClearUser()
	return _u
}

// SetGroup sets the "group" field.
func (_u *FileUpdateOne) SetGroup(v string) *FileUpdateOne {
	_u.mutation.SetGroup(v)
	return _u
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableGroup(v *string) *FileUpdateOne {
	if v != nil {
		_u.SetGroup(*v)
	}
	return _u
}

// ClearGroup clears the value of the "group" field.
func (_u *FileUpdateOne) ClearGroup() *FileUpdateOne {
	_u.mutation.ClearGroup()
	return _u
}

// SetOp sets the "op" field.
func (_u *FileUpdateOne) SetOp(v bool) *FileUpdateOne {
	_u.mutation.SetOpField(v)
	return _u
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableOp(v *bool) *FileUpdateOne {
	if v != nil {
		_u.SetOp(*v)
	}
	return _u
}

// ClearOp clears the value of the "op" field.
func (_u *FileUpdateOne) ClearOp() *FileUpdateOne {
	_u.mutation.ClearOp()
	return _u
}

// SetFieldID sets the "field_id" field.
func (_u *FileUpdateOne) SetFieldID(v int) *FileUpdateOne {
	_u.mutation.ResetFieldID()
	_u.mutation.SetFieldID(v)
	return _u
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableFieldID(v *int) *FileUpdateOne {
	if v != nil {
		_u.SetFieldID(*v)
	}
	return _u
}

// AddFieldID adds value to the "field_id" field.
func (_u *FileUpdateOne) AddFieldID(v int) *FileUpdateOne {
	_u.mutation.AddFieldID(v)
	return _u
}

// ClearFieldID clears the value of the "field_id" field.
func (_u *FileUpdateOne) ClearFieldID() *FileUpdateOne {
	_u.mutation.ClearFieldID()
	return _u
}

// SetCreateTime sets the "create_time" field.
func (_u *FileUpdateOne) SetCreateTime(v time.Time) *FileUpdateOne {
	_u.mutation.SetCreateTime(v)
	return _u
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (_u *FileUpdateOne) SetNillableCreateTime(v *time.Time) *FileUpdateOne {
	if v != nil {
		_u.SetCreateTime(*v)
	}
	return _u
}

// ClearCreateTime clears the value of the "create_time" field.
func (_u *FileUpdateOne) ClearCreateTime() *FileUpdateOne {
	_u.mutation.ClearCreateTime()
	return _u
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_u *FileUpdateOne) SetOwnerID(id string) *FileUpdateOne {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_u *FileUpdateOne) SetNillableOwnerID(id *string) *FileUpdateOne {
	if id != nil {
		_u = _u.SetOwnerID(*id)
	}
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *FileUpdateOne) SetOwner(v *User) *FileUpdateOne {
	return _u.SetOwnerID(v.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (_u *FileUpdateOne) SetTypeID(id string) *FileUpdateOne {
	_u.mutation.SetTypeID(id)
	return _u
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (_u *FileUpdateOne) SetNillableTypeID(id *string) *FileUpdateOne {
	if id != nil {
		_u = _u.SetTypeID(*id)
	}
	return _u
}

// SetType sets the "type" edge to the FileType entity.
func (_u *FileUpdateOne) SetType(v *FileType) *FileUpdateOne {
	return _u.SetTypeID(v.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (_u *FileUpdateOne) AddFieldIDs(ids ...string) *FileUpdateOne {
	_u.mutation.AddFieldIDs(ids...)
	return _u
}

// AddField adds the "field" edges to the FieldType entity.
func (_u *FileUpdateOne) AddField(v ...*FieldType) *FileUpdateOne {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (_u *FileUpdateOne) Mutation() *FileMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *FileUpdateOne) ClearOwner() *FileUpdateOne {
	_u.mutation.ClearOwner()
	return _u
}

// ClearType clears the "type" edge to the FileType entity.
func (_u *FileUpdateOne) ClearType() *FileUpdateOne {
	_u.mutation.ClearType()
	return _u
}

// ClearFieldEdge clears all "field" edges to the FieldType entity.
func (_u *FileUpdateOne) ClearFieldEdge() *FileUpdateOne {
	_u.mutation.ClearFieldEdge()
	return _u
}

// RemoveFieldIDs removes the "field" edge to FieldType entities by IDs.
func (_u *FileUpdateOne) RemoveFieldIDs(ids ...string) *FileUpdateOne {
	_u.mutation.RemoveFieldIDs(ids...)
	return _u
}

// RemoveField removes "field" edges to FieldType entities.
func (_u *FileUpdateOne) RemoveField(v ...*FieldType) *FileUpdateOne {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveFieldIDs(ids...)
}

// Where appends a list predicates to the FileUpdate builder.
func (_u *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated File entity.
func (_u *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, _u.gremlinSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *FileUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *FileUpdateOne) check() error {
	if v, ok := _u.mutation.SetID(); ok {
		if err := file.SetIDValidator(v); err != nil {
			return &ValidationError{Name: "set_id", err: fmt.Errorf(`ent: validator failed for field "File.set_id": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "File.size": %w`, err)}
		}
	}
	return nil
}

func (_u *FileUpdateOne) gremlinSave(ctx context.Context) (*File, error) {
	if err := _u.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	query, bindings := _u.gremlin(id).Query()
	if err := _u.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	_u.mutation.done = true
	_m := &File{config: _u.config}
	if err := _m.FromResponse(res); err != nil {
		return nil, err
	}
	return _m, nil
}

func (_u *FileUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := _u.mutation.SetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, value)
	}
	if value, ok := _u.mutation.AddedSetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, __.Union(__.Values(file.FieldSetID), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := _u.mutation.AddedSize(); ok {
		v.Property(dsl.Single, file.FieldSize, __.Union(__.Values(file.FieldSize), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := _u.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := _u.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := _u.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	if value, ok := _u.mutation.FieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, value)
	}
	if value, ok := _u.mutation.AddedFieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, __.Union(__.Values(file.FieldFieldID), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.CreateTime(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldCreateTime, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldCreateTime, value)),
		})
		v.Property(dsl.Single, file.FieldCreateTime, value)
	}
	var properties []any
	if _u.mutation.SetIDCleared() {
		properties = append(properties, file.FieldSetID)
	}
	if _u.mutation.UserCleared() {
		properties = append(properties, file.FieldUser)
	}
	if _u.mutation.GroupCleared() {
		properties = append(properties, file.FieldGroup)
	}
	if _u.mutation.OpCleared() {
		properties = append(properties, file.FieldOp)
	}
	if _u.mutation.FieldIDCleared() {
		properties = append(properties, file.FieldFieldID)
	}
	if _u.mutation.CreateTimeCleared() {
		properties = append(properties, file.FieldCreateTime)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if _u.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	if _u.mutation.TypeCleared() {
		tr := rv.Clone().InE(filetype.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.TypeIDs() {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range _u.mutation.RemovedFieldIDs() {
		tr := rv.Clone().OutE(file.FieldLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.FieldIDs() {
		v.AddE(file.FieldLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(file.FieldLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(file.Label, file.FieldLabel, id)),
		})
	}
	if len(_u.fields) > 0 {
		fields := make([]any, 0, len(_u.fields)+1)
		fields = append(fields, true)
		for _, f := range _u.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
