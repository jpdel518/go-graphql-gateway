// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/jpdel518/go-graphql-gateway/group/ent/group"
	"github.com/jpdel518/go-graphql-gateway/group/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGroup = "Group"
)

// GroupMutation represents an operation that mutates the Group nodes in the graph.
type GroupMutation struct {
	config
	op            Op
	typ           string
	id            *int
	created_at    *time.Time
	updated_at    *time.Time
	name          *string
	description   *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Group, error)
	predicates    []predicate.Group
}

var _ ent.Mutation = (*GroupMutation)(nil)

// groupOption allows management of the mutation configuration using functional options.
type groupOption func(*GroupMutation)

// newGroupMutation creates new mutation for the Group entity.
func newGroupMutation(c config, op Op, opts ...groupOption) *GroupMutation {
	m := &GroupMutation{
		config:        c,
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGroupID sets the ID field of the mutation.
func withGroupID(id int) groupOption {
	return func(m *GroupMutation) {
		var (
			err   error
			once  sync.Once
			value *Group
		)
		m.oldValue = func(ctx context.Context) (*Group, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Group.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGroup sets the old Group of the mutation.
func withGroup(node *Group) groupOption {
	return func(m *GroupMutation) {
		m.oldValue = func(context.Context) (*Group, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GroupMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *GroupMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Group.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *GroupMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *GroupMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *GroupMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *GroupMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *GroupMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *GroupMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *GroupMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *GroupMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *GroupMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *GroupMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *GroupMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *GroupMutation) ResetDescription() {
	m.description = nil
}

// Where appends a list predicates to the GroupMutation builder.
func (m *GroupMutation) Where(ps ...predicate.Group) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the GroupMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *GroupMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Group, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *GroupMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *GroupMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, group.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, group.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, group.FieldName)
	}
	if m.description != nil {
		fields = append(fields, group.FieldDescription)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case group.FieldCreatedAt:
		return m.CreatedAt()
	case group.FieldUpdatedAt:
		return m.UpdatedAt()
	case group.FieldName:
		return m.Name()
	case group.FieldDescription:
		return m.Description()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GroupMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case group.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case group.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case group.FieldName:
		return m.OldName(ctx)
	case group.FieldDescription:
		return m.OldDescription(ctx)
	}
	return nil, fmt.Errorf("unknown Group field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case group.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case group.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case group.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case group.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GroupMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GroupMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	case group.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case group.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case group.FieldName:
		m.ResetName()
		return nil
	case group.FieldDescription:
		m.ResetDescription()
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Group edge %s", name)
}
