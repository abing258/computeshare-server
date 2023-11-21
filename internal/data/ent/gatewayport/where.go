// Code generated by ent, DO NOT EDIT.

package gatewayport

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldLTE(FieldID, id))
}

// FkGatewayID applies equality check predicate on the "fk_gateway_id" field. It's identical to FkGatewayIDEQ.
func FkGatewayID(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldFkGatewayID, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldPort, v))
}

// IsUse applies equality check predicate on the "is_use" field. It's identical to IsUseEQ.
func IsUse(v bool) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldIsUse, v))
}

// FkGatewayIDEQ applies the EQ predicate on the "fk_gateway_id" field.
func FkGatewayIDEQ(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldFkGatewayID, v))
}

// FkGatewayIDNEQ applies the NEQ predicate on the "fk_gateway_id" field.
func FkGatewayIDNEQ(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldNEQ(FieldFkGatewayID, v))
}

// FkGatewayIDIn applies the In predicate on the "fk_gateway_id" field.
func FkGatewayIDIn(vs ...string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldIn(FieldFkGatewayID, vs...))
}

// FkGatewayIDNotIn applies the NotIn predicate on the "fk_gateway_id" field.
func FkGatewayIDNotIn(vs ...string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldNotIn(FieldFkGatewayID, vs...))
}

// FkGatewayIDGT applies the GT predicate on the "fk_gateway_id" field.
func FkGatewayIDGT(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldGT(FieldFkGatewayID, v))
}

// FkGatewayIDGTE applies the GTE predicate on the "fk_gateway_id" field.
func FkGatewayIDGTE(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldGTE(FieldFkGatewayID, v))
}

// FkGatewayIDLT applies the LT predicate on the "fk_gateway_id" field.
func FkGatewayIDLT(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldLT(FieldFkGatewayID, v))
}

// FkGatewayIDLTE applies the LTE predicate on the "fk_gateway_id" field.
func FkGatewayIDLTE(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldLTE(FieldFkGatewayID, v))
}

// FkGatewayIDContains applies the Contains predicate on the "fk_gateway_id" field.
func FkGatewayIDContains(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldContains(FieldFkGatewayID, v))
}

// FkGatewayIDHasPrefix applies the HasPrefix predicate on the "fk_gateway_id" field.
func FkGatewayIDHasPrefix(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldHasPrefix(FieldFkGatewayID, v))
}

// FkGatewayIDHasSuffix applies the HasSuffix predicate on the "fk_gateway_id" field.
func FkGatewayIDHasSuffix(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldHasSuffix(FieldFkGatewayID, v))
}

// FkGatewayIDEqualFold applies the EqualFold predicate on the "fk_gateway_id" field.
func FkGatewayIDEqualFold(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEqualFold(FieldFkGatewayID, v))
}

// FkGatewayIDContainsFold applies the ContainsFold predicate on the "fk_gateway_id" field.
func FkGatewayIDContainsFold(v string) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldContainsFold(FieldFkGatewayID, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v int64) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldLTE(FieldPort, v))
}

// IsUseEQ applies the EQ predicate on the "is_use" field.
func IsUseEQ(v bool) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldEQ(FieldIsUse, v))
}

// IsUseNEQ applies the NEQ predicate on the "is_use" field.
func IsUseNEQ(v bool) predicate.GatewayPort {
	return predicate.GatewayPort(sql.FieldNEQ(FieldIsUse, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GatewayPort) predicate.GatewayPort {
	return predicate.GatewayPort(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GatewayPort) predicate.GatewayPort {
	return predicate.GatewayPort(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GatewayPort) predicate.GatewayPort {
	return predicate.GatewayPort(func(s *sql.Selector) {
		p(s.Not())
	})
}
