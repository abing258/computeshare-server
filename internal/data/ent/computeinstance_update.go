// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeinstance"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ComputeInstanceUpdate is the builder for updating ComputeInstance entities.
type ComputeInstanceUpdate struct {
	config
	hooks    []Hook
	mutation *ComputeInstanceMutation
}

// Where appends a list predicates to the ComputeInstanceUpdate builder.
func (ciu *ComputeInstanceUpdate) Where(ps ...predicate.ComputeInstance) *ComputeInstanceUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetOwner sets the "owner" field.
func (ciu *ComputeInstanceUpdate) SetOwner(s string) *ComputeInstanceUpdate {
	ciu.mutation.SetOwner(s)
	return ciu
}

// SetName sets the "name" field.
func (ciu *ComputeInstanceUpdate) SetName(s string) *ComputeInstanceUpdate {
	ciu.mutation.SetName(s)
	return ciu
}

// SetCore sets the "core" field.
func (ciu *ComputeInstanceUpdate) SetCore(s string) *ComputeInstanceUpdate {
	ciu.mutation.SetCore(s)
	return ciu
}

// SetMemory sets the "memory" field.
func (ciu *ComputeInstanceUpdate) SetMemory(s string) *ComputeInstanceUpdate {
	ciu.mutation.SetMemory(s)
	return ciu
}

// SetImage sets the "image" field.
func (ciu *ComputeInstanceUpdate) SetImage(s string) *ComputeInstanceUpdate {
	ciu.mutation.SetImage(s)
	return ciu
}

// SetExpirationTime sets the "expiration_time" field.
func (ciu *ComputeInstanceUpdate) SetExpirationTime(t time.Time) *ComputeInstanceUpdate {
	ciu.mutation.SetExpirationTime(t)
	return ciu
}

// SetStatus sets the "status" field.
func (ciu *ComputeInstanceUpdate) SetStatus(i int8) *ComputeInstanceUpdate {
	ciu.mutation.ResetStatus()
	ciu.mutation.SetStatus(i)
	return ciu
}

// AddStatus adds i to the "status" field.
func (ciu *ComputeInstanceUpdate) AddStatus(i int8) *ComputeInstanceUpdate {
	ciu.mutation.AddStatus(i)
	return ciu
}

// SetContainerID sets the "container_id" field.
func (ciu *ComputeInstanceUpdate) SetContainerID(s string) *ComputeInstanceUpdate {
	ciu.mutation.SetContainerID(s)
	return ciu
}

// SetNillableContainerID sets the "container_id" field if the given value is not nil.
func (ciu *ComputeInstanceUpdate) SetNillableContainerID(s *string) *ComputeInstanceUpdate {
	if s != nil {
		ciu.SetContainerID(*s)
	}
	return ciu
}

// ClearContainerID clears the value of the "container_id" field.
func (ciu *ComputeInstanceUpdate) ClearContainerID() *ComputeInstanceUpdate {
	ciu.mutation.ClearContainerID()
	return ciu
}

// SetPeerID sets the "peer_id" field.
func (ciu *ComputeInstanceUpdate) SetPeerID(s string) *ComputeInstanceUpdate {
	ciu.mutation.SetPeerID(s)
	return ciu
}

// SetNillablePeerID sets the "peer_id" field if the given value is not nil.
func (ciu *ComputeInstanceUpdate) SetNillablePeerID(s *string) *ComputeInstanceUpdate {
	if s != nil {
		ciu.SetPeerID(*s)
	}
	return ciu
}

// ClearPeerID clears the value of the "peer_id" field.
func (ciu *ComputeInstanceUpdate) ClearPeerID() *ComputeInstanceUpdate {
	ciu.mutation.ClearPeerID()
	return ciu
}

// Mutation returns the ComputeInstanceMutation object of the builder.
func (ciu *ComputeInstanceUpdate) Mutation() *ComputeInstanceMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *ComputeInstanceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ciu.sqlSave, ciu.mutation, ciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *ComputeInstanceUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *ComputeInstanceUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *ComputeInstanceUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *ComputeInstanceUpdate) check() error {
	if v, ok := ciu.mutation.Owner(); ok {
		if err := computeinstance.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.owner": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Name(); ok {
		if err := computeinstance.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.name": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Core(); ok {
		if err := computeinstance.CoreValidator(v); err != nil {
			return &ValidationError{Name: "core", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.core": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Memory(); ok {
		if err := computeinstance.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.memory": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Image(); ok {
		if err := computeinstance.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.image": %w`, err)}
		}
	}
	return nil
}

func (ciu *ComputeInstanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ciu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(computeinstance.Table, computeinstance.Columns, sqlgraph.NewFieldSpec(computeinstance.FieldID, field.TypeUUID))
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.Owner(); ok {
		_spec.SetField(computeinstance.FieldOwner, field.TypeString, value)
	}
	if value, ok := ciu.mutation.Name(); ok {
		_spec.SetField(computeinstance.FieldName, field.TypeString, value)
	}
	if value, ok := ciu.mutation.Core(); ok {
		_spec.SetField(computeinstance.FieldCore, field.TypeString, value)
	}
	if value, ok := ciu.mutation.Memory(); ok {
		_spec.SetField(computeinstance.FieldMemory, field.TypeString, value)
	}
	if value, ok := ciu.mutation.Image(); ok {
		_spec.SetField(computeinstance.FieldImage, field.TypeString, value)
	}
	if value, ok := ciu.mutation.ExpirationTime(); ok {
		_spec.SetField(computeinstance.FieldExpirationTime, field.TypeTime, value)
	}
	if value, ok := ciu.mutation.Status(); ok {
		_spec.SetField(computeinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ciu.mutation.AddedStatus(); ok {
		_spec.AddField(computeinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ciu.mutation.ContainerID(); ok {
		_spec.SetField(computeinstance.FieldContainerID, field.TypeString, value)
	}
	if ciu.mutation.ContainerIDCleared() {
		_spec.ClearField(computeinstance.FieldContainerID, field.TypeString)
	}
	if value, ok := ciu.mutation.PeerID(); ok {
		_spec.SetField(computeinstance.FieldPeerID, field.TypeString, value)
	}
	if ciu.mutation.PeerIDCleared() {
		_spec.ClearField(computeinstance.FieldPeerID, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computeinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ciu.mutation.done = true
	return n, nil
}

// ComputeInstanceUpdateOne is the builder for updating a single ComputeInstance entity.
type ComputeInstanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ComputeInstanceMutation
}

// SetOwner sets the "owner" field.
func (ciuo *ComputeInstanceUpdateOne) SetOwner(s string) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetOwner(s)
	return ciuo
}

// SetName sets the "name" field.
func (ciuo *ComputeInstanceUpdateOne) SetName(s string) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetName(s)
	return ciuo
}

// SetCore sets the "core" field.
func (ciuo *ComputeInstanceUpdateOne) SetCore(s string) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetCore(s)
	return ciuo
}

// SetMemory sets the "memory" field.
func (ciuo *ComputeInstanceUpdateOne) SetMemory(s string) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetMemory(s)
	return ciuo
}

// SetImage sets the "image" field.
func (ciuo *ComputeInstanceUpdateOne) SetImage(s string) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetImage(s)
	return ciuo
}

// SetExpirationTime sets the "expiration_time" field.
func (ciuo *ComputeInstanceUpdateOne) SetExpirationTime(t time.Time) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetExpirationTime(t)
	return ciuo
}

// SetStatus sets the "status" field.
func (ciuo *ComputeInstanceUpdateOne) SetStatus(i int8) *ComputeInstanceUpdateOne {
	ciuo.mutation.ResetStatus()
	ciuo.mutation.SetStatus(i)
	return ciuo
}

// AddStatus adds i to the "status" field.
func (ciuo *ComputeInstanceUpdateOne) AddStatus(i int8) *ComputeInstanceUpdateOne {
	ciuo.mutation.AddStatus(i)
	return ciuo
}

// SetContainerID sets the "container_id" field.
func (ciuo *ComputeInstanceUpdateOne) SetContainerID(s string) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetContainerID(s)
	return ciuo
}

// SetNillableContainerID sets the "container_id" field if the given value is not nil.
func (ciuo *ComputeInstanceUpdateOne) SetNillableContainerID(s *string) *ComputeInstanceUpdateOne {
	if s != nil {
		ciuo.SetContainerID(*s)
	}
	return ciuo
}

// ClearContainerID clears the value of the "container_id" field.
func (ciuo *ComputeInstanceUpdateOne) ClearContainerID() *ComputeInstanceUpdateOne {
	ciuo.mutation.ClearContainerID()
	return ciuo
}

// SetPeerID sets the "peer_id" field.
func (ciuo *ComputeInstanceUpdateOne) SetPeerID(s string) *ComputeInstanceUpdateOne {
	ciuo.mutation.SetPeerID(s)
	return ciuo
}

// SetNillablePeerID sets the "peer_id" field if the given value is not nil.
func (ciuo *ComputeInstanceUpdateOne) SetNillablePeerID(s *string) *ComputeInstanceUpdateOne {
	if s != nil {
		ciuo.SetPeerID(*s)
	}
	return ciuo
}

// ClearPeerID clears the value of the "peer_id" field.
func (ciuo *ComputeInstanceUpdateOne) ClearPeerID() *ComputeInstanceUpdateOne {
	ciuo.mutation.ClearPeerID()
	return ciuo
}

// Mutation returns the ComputeInstanceMutation object of the builder.
func (ciuo *ComputeInstanceUpdateOne) Mutation() *ComputeInstanceMutation {
	return ciuo.mutation
}

// Where appends a list predicates to the ComputeInstanceUpdate builder.
func (ciuo *ComputeInstanceUpdateOne) Where(ps ...predicate.ComputeInstance) *ComputeInstanceUpdateOne {
	ciuo.mutation.Where(ps...)
	return ciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *ComputeInstanceUpdateOne) Select(field string, fields ...string) *ComputeInstanceUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated ComputeInstance entity.
func (ciuo *ComputeInstanceUpdateOne) Save(ctx context.Context) (*ComputeInstance, error) {
	return withHooks(ctx, ciuo.sqlSave, ciuo.mutation, ciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *ComputeInstanceUpdateOne) SaveX(ctx context.Context) *ComputeInstance {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *ComputeInstanceUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *ComputeInstanceUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *ComputeInstanceUpdateOne) check() error {
	if v, ok := ciuo.mutation.Owner(); ok {
		if err := computeinstance.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.owner": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Name(); ok {
		if err := computeinstance.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.name": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Core(); ok {
		if err := computeinstance.CoreValidator(v); err != nil {
			return &ValidationError{Name: "core", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.core": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Memory(); ok {
		if err := computeinstance.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.memory": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Image(); ok {
		if err := computeinstance.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.image": %w`, err)}
		}
	}
	return nil
}

func (ciuo *ComputeInstanceUpdateOne) sqlSave(ctx context.Context) (_node *ComputeInstance, err error) {
	if err := ciuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(computeinstance.Table, computeinstance.Columns, sqlgraph.NewFieldSpec(computeinstance.FieldID, field.TypeUUID))
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ComputeInstance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, computeinstance.FieldID)
		for _, f := range fields {
			if !computeinstance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != computeinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.Owner(); ok {
		_spec.SetField(computeinstance.FieldOwner, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.Name(); ok {
		_spec.SetField(computeinstance.FieldName, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.Core(); ok {
		_spec.SetField(computeinstance.FieldCore, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.Memory(); ok {
		_spec.SetField(computeinstance.FieldMemory, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.Image(); ok {
		_spec.SetField(computeinstance.FieldImage, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.ExpirationTime(); ok {
		_spec.SetField(computeinstance.FieldExpirationTime, field.TypeTime, value)
	}
	if value, ok := ciuo.mutation.Status(); ok {
		_spec.SetField(computeinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ciuo.mutation.AddedStatus(); ok {
		_spec.AddField(computeinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ciuo.mutation.ContainerID(); ok {
		_spec.SetField(computeinstance.FieldContainerID, field.TypeString, value)
	}
	if ciuo.mutation.ContainerIDCleared() {
		_spec.ClearField(computeinstance.FieldContainerID, field.TypeString)
	}
	if value, ok := ciuo.mutation.PeerID(); ok {
		_spec.SetField(computeinstance.FieldPeerID, field.TypeString, value)
	}
	if ciuo.mutation.PeerIDCleared() {
		_spec.ClearField(computeinstance.FieldPeerID, field.TypeString)
	}
	_node = &ComputeInstance{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computeinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ciuo.mutation.done = true
	return _node, nil
}
