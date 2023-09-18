// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeinstance"
)

// ComputeInstanceCreate is the builder for creating a ComputeInstance entity.
type ComputeInstanceCreate struct {
	config
	mutation *ComputeInstanceMutation
	hooks    []Hook
}

// SetOwner sets the "owner" field.
func (cic *ComputeInstanceCreate) SetOwner(s string) *ComputeInstanceCreate {
	cic.mutation.SetOwner(s)
	return cic
}

// SetName sets the "name" field.
func (cic *ComputeInstanceCreate) SetName(s string) *ComputeInstanceCreate {
	cic.mutation.SetName(s)
	return cic
}

// SetCore sets the "core" field.
func (cic *ComputeInstanceCreate) SetCore(s string) *ComputeInstanceCreate {
	cic.mutation.SetCore(s)
	return cic
}

// SetMemory sets the "memory" field.
func (cic *ComputeInstanceCreate) SetMemory(s string) *ComputeInstanceCreate {
	cic.mutation.SetMemory(s)
	return cic
}

// SetImage sets the "image" field.
func (cic *ComputeInstanceCreate) SetImage(s string) *ComputeInstanceCreate {
	cic.mutation.SetImage(s)
	return cic
}

// SetExpirationTime sets the "expiration_time" field.
func (cic *ComputeInstanceCreate) SetExpirationTime(t time.Time) *ComputeInstanceCreate {
	cic.mutation.SetExpirationTime(t)
	return cic
}

// SetStatus sets the "status" field.
func (cic *ComputeInstanceCreate) SetStatus(i int8) *ComputeInstanceCreate {
	cic.mutation.SetStatus(i)
	return cic
}

// SetContainerID sets the "container_id" field.
func (cic *ComputeInstanceCreate) SetContainerID(s string) *ComputeInstanceCreate {
	cic.mutation.SetContainerID(s)
	return cic
}

// SetNillableContainerID sets the "container_id" field if the given value is not nil.
func (cic *ComputeInstanceCreate) SetNillableContainerID(s *string) *ComputeInstanceCreate {
	if s != nil {
		cic.SetContainerID(*s)
	}
	return cic
}

// SetPeerID sets the "peer_id" field.
func (cic *ComputeInstanceCreate) SetPeerID(s string) *ComputeInstanceCreate {
	cic.mutation.SetPeerID(s)
	return cic
}

// SetNillablePeerID sets the "peer_id" field if the given value is not nil.
func (cic *ComputeInstanceCreate) SetNillablePeerID(s *string) *ComputeInstanceCreate {
	if s != nil {
		cic.SetPeerID(*s)
	}
	return cic
}

// SetID sets the "id" field.
func (cic *ComputeInstanceCreate) SetID(u uuid.UUID) *ComputeInstanceCreate {
	cic.mutation.SetID(u)
	return cic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cic *ComputeInstanceCreate) SetNillableID(u *uuid.UUID) *ComputeInstanceCreate {
	if u != nil {
		cic.SetID(*u)
	}
	return cic
}

// Mutation returns the ComputeInstanceMutation object of the builder.
func (cic *ComputeInstanceCreate) Mutation() *ComputeInstanceMutation {
	return cic.mutation
}

// Save creates the ComputeInstance in the database.
func (cic *ComputeInstanceCreate) Save(ctx context.Context) (*ComputeInstance, error) {
	cic.defaults()
	return withHooks(ctx, cic.sqlSave, cic.mutation, cic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cic *ComputeInstanceCreate) SaveX(ctx context.Context) *ComputeInstance {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cic *ComputeInstanceCreate) Exec(ctx context.Context) error {
	_, err := cic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cic *ComputeInstanceCreate) ExecX(ctx context.Context) {
	if err := cic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cic *ComputeInstanceCreate) defaults() {
	if _, ok := cic.mutation.ID(); !ok {
		v := computeinstance.DefaultID()
		cic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cic *ComputeInstanceCreate) check() error {
	if _, ok := cic.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "ComputeInstance.owner"`)}
	}
	if v, ok := cic.mutation.Owner(); ok {
		if err := computeinstance.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.owner": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ComputeInstance.name"`)}
	}
	if v, ok := cic.mutation.Name(); ok {
		if err := computeinstance.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.name": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Core(); !ok {
		return &ValidationError{Name: "core", err: errors.New(`ent: missing required field "ComputeInstance.core"`)}
	}
	if v, ok := cic.mutation.Core(); ok {
		if err := computeinstance.CoreValidator(v); err != nil {
			return &ValidationError{Name: "core", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.core": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Memory(); !ok {
		return &ValidationError{Name: "memory", err: errors.New(`ent: missing required field "ComputeInstance.memory"`)}
	}
	if v, ok := cic.mutation.Memory(); ok {
		if err := computeinstance.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.memory": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "ComputeInstance.image"`)}
	}
	if v, ok := cic.mutation.Image(); ok {
		if err := computeinstance.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "ComputeInstance.image": %w`, err)}
		}
	}
	if _, ok := cic.mutation.ExpirationTime(); !ok {
		return &ValidationError{Name: "expiration_time", err: errors.New(`ent: missing required field "ComputeInstance.expiration_time"`)}
	}
	if _, ok := cic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ComputeInstance.status"`)}
	}
	return nil
}

func (cic *ComputeInstanceCreate) sqlSave(ctx context.Context) (*ComputeInstance, error) {
	if err := cic.check(); err != nil {
		return nil, err
	}
	_node, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cic.mutation.id = &_node.ID
	cic.mutation.done = true
	return _node, nil
}

func (cic *ComputeInstanceCreate) createSpec() (*ComputeInstance, *sqlgraph.CreateSpec) {
	var (
		_node = &ComputeInstance{config: cic.config}
		_spec = sqlgraph.NewCreateSpec(computeinstance.Table, sqlgraph.NewFieldSpec(computeinstance.FieldID, field.TypeUUID))
	)
	if id, ok := cic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cic.mutation.Owner(); ok {
		_spec.SetField(computeinstance.FieldOwner, field.TypeString, value)
		_node.Owner = value
	}
	if value, ok := cic.mutation.Name(); ok {
		_spec.SetField(computeinstance.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cic.mutation.Core(); ok {
		_spec.SetField(computeinstance.FieldCore, field.TypeString, value)
		_node.Core = value
	}
	if value, ok := cic.mutation.Memory(); ok {
		_spec.SetField(computeinstance.FieldMemory, field.TypeString, value)
		_node.Memory = value
	}
	if value, ok := cic.mutation.Image(); ok {
		_spec.SetField(computeinstance.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := cic.mutation.ExpirationTime(); ok {
		_spec.SetField(computeinstance.FieldExpirationTime, field.TypeTime, value)
		_node.ExpirationTime = value
	}
	if value, ok := cic.mutation.Status(); ok {
		_spec.SetField(computeinstance.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := cic.mutation.ContainerID(); ok {
		_spec.SetField(computeinstance.FieldContainerID, field.TypeString, value)
		_node.ContainerID = value
	}
	if value, ok := cic.mutation.PeerID(); ok {
		_spec.SetField(computeinstance.FieldPeerID, field.TypeString, value)
		_node.PeerID = value
	}
	return _node, _spec
}

// ComputeInstanceCreateBulk is the builder for creating many ComputeInstance entities in bulk.
type ComputeInstanceCreateBulk struct {
	config
	builders []*ComputeInstanceCreate
}

// Save creates the ComputeInstance entities in the database.
func (cicb *ComputeInstanceCreateBulk) Save(ctx context.Context) ([]*ComputeInstance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cicb.builders))
	nodes := make([]*ComputeInstance, len(cicb.builders))
	mutators := make([]Mutator, len(cicb.builders))
	for i := range cicb.builders {
		func(i int, root context.Context) {
			builder := cicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ComputeInstanceMutation)
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
					_, err = mutators[i+1].Mutate(root, cicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, cicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cicb *ComputeInstanceCreateBulk) SaveX(ctx context.Context) []*ComputeInstance {
	v, err := cicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cicb *ComputeInstanceCreateBulk) Exec(ctx context.Context) error {
	_, err := cicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicb *ComputeInstanceCreateBulk) ExecX(ctx context.Context) {
	if err := cicb.Exec(ctx); err != nil {
		panic(err)
	}
}
