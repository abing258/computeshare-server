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
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cyclerecharge"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// CycleRechargeUpdate is the builder for updating CycleRecharge entities.
type CycleRechargeUpdate struct {
	config
	hooks    []Hook
	mutation *CycleRechargeMutation
}

// Where appends a list predicates to the CycleRechargeUpdate builder.
func (cru *CycleRechargeUpdate) Where(ps ...predicate.CycleRecharge) *CycleRechargeUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetFkUserID sets the "fk_user_id" field.
func (cru *CycleRechargeUpdate) SetFkUserID(u uuid.UUID) *CycleRechargeUpdate {
	cru.mutation.SetFkUserID(u)
	return cru
}

// SetOutTradeNo sets the "out_trade_no" field.
func (cru *CycleRechargeUpdate) SetOutTradeNo(s string) *CycleRechargeUpdate {
	cru.mutation.SetOutTradeNo(s)
	return cru
}

// SetAlipayTradeNo sets the "alipay_trade_no" field.
func (cru *CycleRechargeUpdate) SetAlipayTradeNo(s string) *CycleRechargeUpdate {
	cru.mutation.SetAlipayTradeNo(s)
	return cru
}

// SetRechargeChannel sets the "recharge_channel" field.
func (cru *CycleRechargeUpdate) SetRechargeChannel(s string) *CycleRechargeUpdate {
	cru.mutation.SetRechargeChannel(s)
	return cru
}

// SetRedeemCode sets the "redeem_code" field.
func (cru *CycleRechargeUpdate) SetRedeemCode(s string) *CycleRechargeUpdate {
	cru.mutation.SetRedeemCode(s)
	return cru
}

// SetState sets the "state" field.
func (cru *CycleRechargeUpdate) SetState(s string) *CycleRechargeUpdate {
	cru.mutation.SetState(s)
	return cru
}

// SetPayAmount sets the "pay_amount" field.
func (cru *CycleRechargeUpdate) SetPayAmount(f float64) *CycleRechargeUpdate {
	cru.mutation.ResetPayAmount()
	cru.mutation.SetPayAmount(f)
	return cru
}

// AddPayAmount adds f to the "pay_amount" field.
func (cru *CycleRechargeUpdate) AddPayAmount(f float64) *CycleRechargeUpdate {
	cru.mutation.AddPayAmount(f)
	return cru
}

// SetTotalAmount sets the "total_amount" field.
func (cru *CycleRechargeUpdate) SetTotalAmount(f float64) *CycleRechargeUpdate {
	cru.mutation.ResetTotalAmount()
	cru.mutation.SetTotalAmount(f)
	return cru
}

// AddTotalAmount adds f to the "total_amount" field.
func (cru *CycleRechargeUpdate) AddTotalAmount(f float64) *CycleRechargeUpdate {
	cru.mutation.AddTotalAmount(f)
	return cru
}

// SetBuyCycle sets the "buy_cycle" field.
func (cru *CycleRechargeUpdate) SetBuyCycle(f float64) *CycleRechargeUpdate {
	cru.mutation.ResetBuyCycle()
	cru.mutation.SetBuyCycle(f)
	return cru
}

// AddBuyCycle adds f to the "buy_cycle" field.
func (cru *CycleRechargeUpdate) AddBuyCycle(f float64) *CycleRechargeUpdate {
	cru.mutation.AddBuyCycle(f)
	return cru
}

// SetCreateTime sets the "create_time" field.
func (cru *CycleRechargeUpdate) SetCreateTime(t time.Time) *CycleRechargeUpdate {
	cru.mutation.SetCreateTime(t)
	return cru
}

// SetUpdateTime sets the "update_time" field.
func (cru *CycleRechargeUpdate) SetUpdateTime(t time.Time) *CycleRechargeUpdate {
	cru.mutation.SetUpdateTime(t)
	return cru
}

// Mutation returns the CycleRechargeMutation object of the builder.
func (cru *CycleRechargeUpdate) Mutation() *CycleRechargeMutation {
	return cru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CycleRechargeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CycleRechargeUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CycleRechargeUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CycleRechargeUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cru *CycleRechargeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cyclerecharge.Table, cyclerecharge.Columns, sqlgraph.NewFieldSpec(cyclerecharge.FieldID, field.TypeUUID))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.FkUserID(); ok {
		_spec.SetField(cyclerecharge.FieldFkUserID, field.TypeUUID, value)
	}
	if value, ok := cru.mutation.OutTradeNo(); ok {
		_spec.SetField(cyclerecharge.FieldOutTradeNo, field.TypeString, value)
	}
	if value, ok := cru.mutation.AlipayTradeNo(); ok {
		_spec.SetField(cyclerecharge.FieldAlipayTradeNo, field.TypeString, value)
	}
	if value, ok := cru.mutation.RechargeChannel(); ok {
		_spec.SetField(cyclerecharge.FieldRechargeChannel, field.TypeString, value)
	}
	if value, ok := cru.mutation.RedeemCode(); ok {
		_spec.SetField(cyclerecharge.FieldRedeemCode, field.TypeString, value)
	}
	if value, ok := cru.mutation.State(); ok {
		_spec.SetField(cyclerecharge.FieldState, field.TypeString, value)
	}
	if value, ok := cru.mutation.PayAmount(); ok {
		_spec.SetField(cyclerecharge.FieldPayAmount, field.TypeFloat64, value)
	}
	if value, ok := cru.mutation.AddedPayAmount(); ok {
		_spec.AddField(cyclerecharge.FieldPayAmount, field.TypeFloat64, value)
	}
	if value, ok := cru.mutation.TotalAmount(); ok {
		_spec.SetField(cyclerecharge.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := cru.mutation.AddedTotalAmount(); ok {
		_spec.AddField(cyclerecharge.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := cru.mutation.BuyCycle(); ok {
		_spec.SetField(cyclerecharge.FieldBuyCycle, field.TypeFloat64, value)
	}
	if value, ok := cru.mutation.AddedBuyCycle(); ok {
		_spec.AddField(cyclerecharge.FieldBuyCycle, field.TypeFloat64, value)
	}
	if value, ok := cru.mutation.CreateTime(); ok {
		_spec.SetField(cyclerecharge.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := cru.mutation.UpdateTime(); ok {
		_spec.SetField(cyclerecharge.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cyclerecharge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// CycleRechargeUpdateOne is the builder for updating a single CycleRecharge entity.
type CycleRechargeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CycleRechargeMutation
}

// SetFkUserID sets the "fk_user_id" field.
func (cruo *CycleRechargeUpdateOne) SetFkUserID(u uuid.UUID) *CycleRechargeUpdateOne {
	cruo.mutation.SetFkUserID(u)
	return cruo
}

// SetOutTradeNo sets the "out_trade_no" field.
func (cruo *CycleRechargeUpdateOne) SetOutTradeNo(s string) *CycleRechargeUpdateOne {
	cruo.mutation.SetOutTradeNo(s)
	return cruo
}

// SetAlipayTradeNo sets the "alipay_trade_no" field.
func (cruo *CycleRechargeUpdateOne) SetAlipayTradeNo(s string) *CycleRechargeUpdateOne {
	cruo.mutation.SetAlipayTradeNo(s)
	return cruo
}

// SetRechargeChannel sets the "recharge_channel" field.
func (cruo *CycleRechargeUpdateOne) SetRechargeChannel(s string) *CycleRechargeUpdateOne {
	cruo.mutation.SetRechargeChannel(s)
	return cruo
}

// SetRedeemCode sets the "redeem_code" field.
func (cruo *CycleRechargeUpdateOne) SetRedeemCode(s string) *CycleRechargeUpdateOne {
	cruo.mutation.SetRedeemCode(s)
	return cruo
}

// SetState sets the "state" field.
func (cruo *CycleRechargeUpdateOne) SetState(s string) *CycleRechargeUpdateOne {
	cruo.mutation.SetState(s)
	return cruo
}

// SetPayAmount sets the "pay_amount" field.
func (cruo *CycleRechargeUpdateOne) SetPayAmount(f float64) *CycleRechargeUpdateOne {
	cruo.mutation.ResetPayAmount()
	cruo.mutation.SetPayAmount(f)
	return cruo
}

// AddPayAmount adds f to the "pay_amount" field.
func (cruo *CycleRechargeUpdateOne) AddPayAmount(f float64) *CycleRechargeUpdateOne {
	cruo.mutation.AddPayAmount(f)
	return cruo
}

// SetTotalAmount sets the "total_amount" field.
func (cruo *CycleRechargeUpdateOne) SetTotalAmount(f float64) *CycleRechargeUpdateOne {
	cruo.mutation.ResetTotalAmount()
	cruo.mutation.SetTotalAmount(f)
	return cruo
}

// AddTotalAmount adds f to the "total_amount" field.
func (cruo *CycleRechargeUpdateOne) AddTotalAmount(f float64) *CycleRechargeUpdateOne {
	cruo.mutation.AddTotalAmount(f)
	return cruo
}

// SetBuyCycle sets the "buy_cycle" field.
func (cruo *CycleRechargeUpdateOne) SetBuyCycle(f float64) *CycleRechargeUpdateOne {
	cruo.mutation.ResetBuyCycle()
	cruo.mutation.SetBuyCycle(f)
	return cruo
}

// AddBuyCycle adds f to the "buy_cycle" field.
func (cruo *CycleRechargeUpdateOne) AddBuyCycle(f float64) *CycleRechargeUpdateOne {
	cruo.mutation.AddBuyCycle(f)
	return cruo
}

// SetCreateTime sets the "create_time" field.
func (cruo *CycleRechargeUpdateOne) SetCreateTime(t time.Time) *CycleRechargeUpdateOne {
	cruo.mutation.SetCreateTime(t)
	return cruo
}

// SetUpdateTime sets the "update_time" field.
func (cruo *CycleRechargeUpdateOne) SetUpdateTime(t time.Time) *CycleRechargeUpdateOne {
	cruo.mutation.SetUpdateTime(t)
	return cruo
}

// Mutation returns the CycleRechargeMutation object of the builder.
func (cruo *CycleRechargeUpdateOne) Mutation() *CycleRechargeMutation {
	return cruo.mutation
}

// Where appends a list predicates to the CycleRechargeUpdate builder.
func (cruo *CycleRechargeUpdateOne) Where(ps ...predicate.CycleRecharge) *CycleRechargeUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CycleRechargeUpdateOne) Select(field string, fields ...string) *CycleRechargeUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CycleRecharge entity.
func (cruo *CycleRechargeUpdateOne) Save(ctx context.Context) (*CycleRecharge, error) {
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CycleRechargeUpdateOne) SaveX(ctx context.Context) *CycleRecharge {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CycleRechargeUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CycleRechargeUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cruo *CycleRechargeUpdateOne) sqlSave(ctx context.Context) (_node *CycleRecharge, err error) {
	_spec := sqlgraph.NewUpdateSpec(cyclerecharge.Table, cyclerecharge.Columns, sqlgraph.NewFieldSpec(cyclerecharge.FieldID, field.TypeUUID))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CycleRecharge.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cyclerecharge.FieldID)
		for _, f := range fields {
			if !cyclerecharge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cyclerecharge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.FkUserID(); ok {
		_spec.SetField(cyclerecharge.FieldFkUserID, field.TypeUUID, value)
	}
	if value, ok := cruo.mutation.OutTradeNo(); ok {
		_spec.SetField(cyclerecharge.FieldOutTradeNo, field.TypeString, value)
	}
	if value, ok := cruo.mutation.AlipayTradeNo(); ok {
		_spec.SetField(cyclerecharge.FieldAlipayTradeNo, field.TypeString, value)
	}
	if value, ok := cruo.mutation.RechargeChannel(); ok {
		_spec.SetField(cyclerecharge.FieldRechargeChannel, field.TypeString, value)
	}
	if value, ok := cruo.mutation.RedeemCode(); ok {
		_spec.SetField(cyclerecharge.FieldRedeemCode, field.TypeString, value)
	}
	if value, ok := cruo.mutation.State(); ok {
		_spec.SetField(cyclerecharge.FieldState, field.TypeString, value)
	}
	if value, ok := cruo.mutation.PayAmount(); ok {
		_spec.SetField(cyclerecharge.FieldPayAmount, field.TypeFloat64, value)
	}
	if value, ok := cruo.mutation.AddedPayAmount(); ok {
		_spec.AddField(cyclerecharge.FieldPayAmount, field.TypeFloat64, value)
	}
	if value, ok := cruo.mutation.TotalAmount(); ok {
		_spec.SetField(cyclerecharge.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := cruo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(cyclerecharge.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := cruo.mutation.BuyCycle(); ok {
		_spec.SetField(cyclerecharge.FieldBuyCycle, field.TypeFloat64, value)
	}
	if value, ok := cruo.mutation.AddedBuyCycle(); ok {
		_spec.AddField(cyclerecharge.FieldBuyCycle, field.TypeFloat64, value)
	}
	if value, ok := cruo.mutation.CreateTime(); ok {
		_spec.SetField(cyclerecharge.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := cruo.mutation.UpdateTime(); ok {
		_spec.SetField(cyclerecharge.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &CycleRecharge{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cyclerecharge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}
