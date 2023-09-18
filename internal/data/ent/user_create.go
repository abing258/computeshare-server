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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCountryCallCoding sets the "country_call_coding" field.
func (uc *UserCreate) SetCountryCallCoding(s string) *UserCreate {
	uc.mutation.SetCountryCallCoding(s)
	return uc
}

// SetTelephoneNumber sets the "telephone_number" field.
func (uc *UserCreate) SetTelephoneNumber(s string) *UserCreate {
	uc.mutation.SetTelephoneNumber(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetCreateDate sets the "create_date" field.
func (uc *UserCreate) SetCreateDate(t time.Time) *UserCreate {
	uc.mutation.SetCreateDate(t)
	return uc
}

// SetNillableCreateDate sets the "create_date" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateDate(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateDate(*t)
	}
	return uc
}

// SetLastLoginDate sets the "last_login_date" field.
func (uc *UserCreate) SetLastLoginDate(t time.Time) *UserCreate {
	uc.mutation.SetLastLoginDate(t)
	return uc
}

// SetNillableLastLoginDate sets the "last_login_date" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastLoginDate(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastLoginDate(*t)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetIcon sets the "icon" field.
func (uc *UserCreate) SetIcon(s string) *UserCreate {
	uc.mutation.SetIcon(s)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreateDate(); !ok {
		v := user.DefaultCreateDate()
		uc.mutation.SetCreateDate(v)
	}
	if _, ok := uc.mutation.LastLoginDate(); !ok {
		v := user.DefaultLastLoginDate()
		uc.mutation.SetLastLoginDate(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CountryCallCoding(); !ok {
		return &ValidationError{Name: "country_call_coding", err: errors.New(`ent: missing required field "User.country_call_coding"`)}
	}
	if v, ok := uc.mutation.CountryCallCoding(); ok {
		if err := user.CountryCallCodingValidator(v); err != nil {
			return &ValidationError{Name: "country_call_coding", err: fmt.Errorf(`ent: validator failed for field "User.country_call_coding": %w`, err)}
		}
	}
	if _, ok := uc.mutation.TelephoneNumber(); !ok {
		return &ValidationError{Name: "telephone_number", err: errors.New(`ent: missing required field "User.telephone_number"`)}
	}
	if v, ok := uc.mutation.TelephoneNumber(); ok {
		if err := user.TelephoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "telephone_number", err: fmt.Errorf(`ent: validator failed for field "User.telephone_number": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreateDate(); !ok {
		return &ValidationError{Name: "create_date", err: errors.New(`ent: missing required field "User.create_date"`)}
	}
	if _, ok := uc.mutation.LastLoginDate(); !ok {
		return &ValidationError{Name: "last_login_date", err: errors.New(`ent: missing required field "User.last_login_date"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "User.icon"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
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
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.CountryCallCoding(); ok {
		_spec.SetField(user.FieldCountryCallCoding, field.TypeString, value)
		_node.CountryCallCoding = value
	}
	if value, ok := uc.mutation.TelephoneNumber(); ok {
		_spec.SetField(user.FieldTelephoneNumber, field.TypeString, value)
		_node.TelephoneNumber = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.CreateDate(); ok {
		_spec.SetField(user.FieldCreateDate, field.TypeTime, value)
		_node.CreateDate = value
	}
	if value, ok := uc.mutation.LastLoginDate(); ok {
		_spec.SetField(user.FieldLastLoginDate, field.TypeTime, value)
		_node.LastLoginDate = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Icon(); ok {
		_spec.SetField(user.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
