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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/attachedfile"
	"entgo.io/ent/entc/integration/edgeschema/ent/file"
	"entgo.io/ent/entc/integration/edgeschema/ent/process"
	"entgo.io/ent/schema/field"
)

// AttachedFileCreate is the builder for creating a AttachedFile entity.
type AttachedFileCreate struct {
	config
	mutation *AttachedFileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAttachTime sets the "attach_time" field.
func (_c *AttachedFileCreate) SetAttachTime(v time.Time) *AttachedFileCreate {
	_c.mutation.SetAttachTime(v)
	return _c
}

// SetNillableAttachTime sets the "attach_time" field if the given value is not nil.
func (_c *AttachedFileCreate) SetNillableAttachTime(v *time.Time) *AttachedFileCreate {
	if v != nil {
		_c.SetAttachTime(*v)
	}
	return _c
}

// SetFID sets the "f_id" field.
func (_c *AttachedFileCreate) SetFID(v int) *AttachedFileCreate {
	_c.mutation.SetFID(v)
	return _c
}

// SetProcID sets the "proc_id" field.
func (_c *AttachedFileCreate) SetProcID(v int) *AttachedFileCreate {
	_c.mutation.SetProcID(v)
	return _c
}

// SetFiID sets the "fi" edge to the File entity by ID.
func (_c *AttachedFileCreate) SetFiID(id int) *AttachedFileCreate {
	_c.mutation.SetFiID(id)
	return _c
}

// SetFi sets the "fi" edge to the File entity.
func (_c *AttachedFileCreate) SetFi(v *File) *AttachedFileCreate {
	return _c.SetFiID(v.ID)
}

// SetProc sets the "proc" edge to the Process entity.
func (_c *AttachedFileCreate) SetProc(v *Process) *AttachedFileCreate {
	return _c.SetProcID(v.ID)
}

// Mutation returns the AttachedFileMutation object of the builder.
func (_c *AttachedFileCreate) Mutation() *AttachedFileMutation {
	return _c.mutation
}

// Save creates the AttachedFile in the database.
func (_c *AttachedFileCreate) Save(ctx context.Context) (*AttachedFile, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *AttachedFileCreate) SaveX(ctx context.Context) *AttachedFile {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *AttachedFileCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *AttachedFileCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *AttachedFileCreate) defaults() {
	if _, ok := _c.mutation.AttachTime(); !ok {
		v := attachedfile.DefaultAttachTime()
		_c.mutation.SetAttachTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *AttachedFileCreate) check() error {
	if _, ok := _c.mutation.AttachTime(); !ok {
		return &ValidationError{Name: "attach_time", err: errors.New(`ent: missing required field "AttachedFile.attach_time"`)}
	}
	if _, ok := _c.mutation.FID(); !ok {
		return &ValidationError{Name: "f_id", err: errors.New(`ent: missing required field "AttachedFile.f_id"`)}
	}
	if _, ok := _c.mutation.ProcID(); !ok {
		return &ValidationError{Name: "proc_id", err: errors.New(`ent: missing required field "AttachedFile.proc_id"`)}
	}
	if len(_c.mutation.FiIDs()) == 0 {
		return &ValidationError{Name: "fi", err: errors.New(`ent: missing required edge "AttachedFile.fi"`)}
	}
	if len(_c.mutation.ProcIDs()) == 0 {
		return &ValidationError{Name: "proc", err: errors.New(`ent: missing required edge "AttachedFile.proc"`)}
	}
	return nil
}

func (_c *AttachedFileCreate) sqlSave(ctx context.Context) (*AttachedFile, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *AttachedFileCreate) createSpec() (*AttachedFile, *sqlgraph.CreateSpec) {
	var (
		_node = &AttachedFile{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(attachedfile.Table, sqlgraph.NewFieldSpec(attachedfile.FieldID, field.TypeInt))
	)
	_spec.OnConflict = _c.conflict
	if value, ok := _c.mutation.AttachTime(); ok {
		_spec.SetField(attachedfile.FieldAttachTime, field.TypeTime, value)
		_node.AttachTime = value
	}
	if nodes := _c.mutation.FiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attachedfile.FiTable,
			Columns: []string{attachedfile.FiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.ProcIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attachedfile.ProcTable,
			Columns: []string{attachedfile.ProcColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProcID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AttachedFile.Create().
//		SetAttachTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttachedFileUpsert) {
//			SetAttachTime(v+v).
//		}).
//		Exec(ctx)
func (_c *AttachedFileCreate) OnConflict(opts ...sql.ConflictOption) *AttachedFileUpsertOne {
	_c.conflict = opts
	return &AttachedFileUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *AttachedFileCreate) OnConflictColumns(columns ...string) *AttachedFileUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &AttachedFileUpsertOne{
		create: _c,
	}
}

type (
	// AttachedFileUpsertOne is the builder for "upsert"-ing
	//  one AttachedFile node.
	AttachedFileUpsertOne struct {
		create *AttachedFileCreate
	}

	// AttachedFileUpsert is the "OnConflict" setter.
	AttachedFileUpsert struct {
		*sql.UpdateSet
	}
)

// SetAttachTime sets the "attach_time" field.
func (u *AttachedFileUpsert) SetAttachTime(v time.Time) *AttachedFileUpsert {
	u.Set(attachedfile.FieldAttachTime, v)
	return u
}

// UpdateAttachTime sets the "attach_time" field to the value that was provided on create.
func (u *AttachedFileUpsert) UpdateAttachTime() *AttachedFileUpsert {
	u.SetExcluded(attachedfile.FieldAttachTime)
	return u
}

// SetFID sets the "f_id" field.
func (u *AttachedFileUpsert) SetFID(v int) *AttachedFileUpsert {
	u.Set(attachedfile.FieldFID, v)
	return u
}

// UpdateFID sets the "f_id" field to the value that was provided on create.
func (u *AttachedFileUpsert) UpdateFID() *AttachedFileUpsert {
	u.SetExcluded(attachedfile.FieldFID)
	return u
}

// SetProcID sets the "proc_id" field.
func (u *AttachedFileUpsert) SetProcID(v int) *AttachedFileUpsert {
	u.Set(attachedfile.FieldProcID, v)
	return u
}

// UpdateProcID sets the "proc_id" field to the value that was provided on create.
func (u *AttachedFileUpsert) UpdateProcID() *AttachedFileUpsert {
	u.SetExcluded(attachedfile.FieldProcID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AttachedFileUpsertOne) UpdateNewValues() *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AttachedFileUpsertOne) Ignore() *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttachedFileUpsertOne) DoNothing() *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttachedFileCreate.OnConflict
// documentation for more info.
func (u *AttachedFileUpsertOne) Update(set func(*AttachedFileUpsert)) *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttachedFileUpsert{UpdateSet: update})
	}))
	return u
}

// SetAttachTime sets the "attach_time" field.
func (u *AttachedFileUpsertOne) SetAttachTime(v time.Time) *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetAttachTime(v)
	})
}

// UpdateAttachTime sets the "attach_time" field to the value that was provided on create.
func (u *AttachedFileUpsertOne) UpdateAttachTime() *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateAttachTime()
	})
}

// SetFID sets the "f_id" field.
func (u *AttachedFileUpsertOne) SetFID(v int) *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetFID(v)
	})
}

// UpdateFID sets the "f_id" field to the value that was provided on create.
func (u *AttachedFileUpsertOne) UpdateFID() *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateFID()
	})
}

// SetProcID sets the "proc_id" field.
func (u *AttachedFileUpsertOne) SetProcID(v int) *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetProcID(v)
	})
}

// UpdateProcID sets the "proc_id" field to the value that was provided on create.
func (u *AttachedFileUpsertOne) UpdateProcID() *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateProcID()
	})
}

// Exec executes the query.
func (u *AttachedFileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttachedFileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttachedFileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AttachedFileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AttachedFileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AttachedFileCreateBulk is the builder for creating many AttachedFile entities in bulk.
type AttachedFileCreateBulk struct {
	config
	err      error
	builders []*AttachedFileCreate
	conflict []sql.ConflictOption
}

// Save creates the AttachedFile entities in the database.
func (_c *AttachedFileCreateBulk) Save(ctx context.Context) ([]*AttachedFile, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*AttachedFile, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachedFileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = _c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *AttachedFileCreateBulk) SaveX(ctx context.Context) []*AttachedFile {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *AttachedFileCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *AttachedFileCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AttachedFile.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttachedFileUpsert) {
//			SetAttachTime(v+v).
//		}).
//		Exec(ctx)
func (_c *AttachedFileCreateBulk) OnConflict(opts ...sql.ConflictOption) *AttachedFileUpsertBulk {
	_c.conflict = opts
	return &AttachedFileUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *AttachedFileCreateBulk) OnConflictColumns(columns ...string) *AttachedFileUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &AttachedFileUpsertBulk{
		create: _c,
	}
}

// AttachedFileUpsertBulk is the builder for "upsert"-ing
// a bulk of AttachedFile nodes.
type AttachedFileUpsertBulk struct {
	create *AttachedFileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AttachedFileUpsertBulk) UpdateNewValues() *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AttachedFileUpsertBulk) Ignore() *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttachedFileUpsertBulk) DoNothing() *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttachedFileCreateBulk.OnConflict
// documentation for more info.
func (u *AttachedFileUpsertBulk) Update(set func(*AttachedFileUpsert)) *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttachedFileUpsert{UpdateSet: update})
	}))
	return u
}

// SetAttachTime sets the "attach_time" field.
func (u *AttachedFileUpsertBulk) SetAttachTime(v time.Time) *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetAttachTime(v)
	})
}

// UpdateAttachTime sets the "attach_time" field to the value that was provided on create.
func (u *AttachedFileUpsertBulk) UpdateAttachTime() *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateAttachTime()
	})
}

// SetFID sets the "f_id" field.
func (u *AttachedFileUpsertBulk) SetFID(v int) *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetFID(v)
	})
}

// UpdateFID sets the "f_id" field to the value that was provided on create.
func (u *AttachedFileUpsertBulk) UpdateFID() *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateFID()
	})
}

// SetProcID sets the "proc_id" field.
func (u *AttachedFileUpsertBulk) SetProcID(v int) *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetProcID(v)
	})
}

// UpdateProcID sets the "proc_id" field to the value that was provided on create.
func (u *AttachedFileUpsertBulk) UpdateProcID() *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateProcID()
	})
}

// Exec executes the query.
func (u *AttachedFileUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AttachedFileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttachedFileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttachedFileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
