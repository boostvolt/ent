// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"github.com/google/uuid"
)

// Blob is the model entity for the Blob schema.
type Blob struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Count holds the value of the "count" field.
	Count int `json:"count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlobQuery when eager-loading is set.
	Edges        BlobEdges `json:"edges"`
	blob_parent  *uuid.UUID
	selectValues sql.SelectValues
}

// BlobEdges holds the relations/edges for other nodes in the graph.
type BlobEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Blob `json:"parent,omitempty"`
	// Links holds the value of the links edge.
	Links []*Blob `json:"links,omitempty"`
	// BlobLinks holds the value of the blob_links edge.
	BlobLinks []*BlobLink `json:"blob_links,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlobEdges) ParentOrErr() (*Blob, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: blob.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// LinksOrErr returns the Links value or an error if the edge
// was not loaded in eager-loading.
func (e BlobEdges) LinksOrErr() ([]*Blob, error) {
	if e.loadedTypes[1] {
		return e.Links, nil
	}
	return nil, &NotLoadedError{edge: "links"}
}

// BlobLinksOrErr returns the BlobLinks value or an error if the edge
// was not loaded in eager-loading.
func (e BlobEdges) BlobLinksOrErr() ([]*BlobLink, error) {
	if e.loadedTypes[2] {
		return e.BlobLinks, nil
	}
	return nil, &NotLoadedError{edge: "blob_links"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blob) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blob.FieldCount:
			values[i] = new(sql.NullInt64)
		case blob.FieldID, blob.FieldUUID:
			values[i] = new(uuid.UUID)
		case blob.ForeignKeys[0]: // blob_parent
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blob fields.
func (_m *Blob) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blob.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_m.ID = *value
			}
		case blob.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				_m.UUID = *value
			}
		case blob.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				_m.Count = int(value.Int64)
			}
		case blob.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field blob_parent", values[i])
			} else if value.Valid {
				_m.blob_parent = new(uuid.UUID)
				*_m.blob_parent = *value.S.(*uuid.UUID)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Blob.
// This includes values selected through modifiers, order, etc.
func (_m *Blob) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Blob entity.
func (_m *Blob) QueryParent() *BlobQuery {
	return NewBlobClient(_m.config).QueryParent(_m)
}

// QueryLinks queries the "links" edge of the Blob entity.
func (_m *Blob) QueryLinks() *BlobQuery {
	return NewBlobClient(_m.config).QueryLinks(_m)
}

// QueryBlobLinks queries the "blob_links" edge of the Blob entity.
func (_m *Blob) QueryBlobLinks() *BlobLinkQuery {
	return NewBlobClient(_m.config).QueryBlobLinks(_m)
}

// Update returns a builder for updating this Blob.
// Note that you need to call Blob.Unwrap() before calling this method if this Blob
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Blob) Update() *BlobUpdateOne {
	return NewBlobClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Blob entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Blob) Unwrap() *Blob {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blob is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Blob) String() string {
	var builder strings.Builder
	builder.WriteString("Blob(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", _m.UUID))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", _m.Count))
	builder.WriteByte(')')
	return builder.String()
}

// Blobs is a parsable slice of Blob.
type Blobs []*Blob
