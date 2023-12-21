// Code generated by ent, DO NOT EDIT.

package s3user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the s3user type in the database.
	Label = "s3user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFkUserID holds the string denoting the fk_user_id field in the database.
	FieldFkUserID = "fk_user_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAccessKey holds the string denoting the access_key field in the database.
	FieldAccessKey = "access_key"
	// FieldSecretKey holds the string denoting the secret_key field in the database.
	FieldSecretKey = "secret_key"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the s3user in the database.
	Table = "s3users"
)

// Columns holds all SQL columns for s3user fields.
var Columns = []string{
	FieldID,
	FieldFkUserID,
	FieldType,
	FieldAccessKey,
	FieldSecretKey,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AccessKeyValidator is a validator for the "access_key" field. It is called by the builders before save.
	AccessKeyValidator func(string) error
	// SecretKeyValidator is a validator for the "secret_key" field. It is called by the builders before save.
	SecretKeyValidator func(string) error
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the S3User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFkUserID orders the results by the fk_user_id field.
func ByFkUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFkUserID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAccessKey orders the results by the access_key field.
func ByAccessKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessKey, opts...).ToFunc()
}

// BySecretKey orders the results by the secret_key field.
func BySecretKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecretKey, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
