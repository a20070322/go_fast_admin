// Code generated by entc, DO NOT EDIT.

package adminrole

import (
	"time"
)

const (
	// Label holds the string label denoting the adminrole type in the database.
	Label = "admin_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsEnable holds the string denoting the is_enable field in the database.
	FieldIsEnable = "is_enable"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMenu holds the string denoting the menu edge name in mutations.
	EdgeMenu = "menu"
	// Table holds the table name of the adminrole in the database.
	Table = "admin_roles"
	// UserTable is the table the holds the user relation/edge. The primary key declared below.
	UserTable = "admin_role_user"
	// UserInverseTable is the table name for the AdminUser entity.
	// It exists in this package in order to avoid circular dependency with the "adminuser" package.
	UserInverseTable = "admin_users"
	// MenuTable is the table the holds the menu relation/edge. The primary key declared below.
	MenuTable = "admin_role_menu"
	// MenuInverseTable is the table name for the AdminMenus entity.
	// It exists in this package in order to avoid circular dependency with the "adminmenus" package.
	MenuInverseTable = "admin_menus"
)

// Columns holds all SQL columns for adminrole fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldIsEnable,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"admin_role_id", "admin_user_id"}
	// MenuPrimaryKey and MenuColumn2 are the table columns denoting the
	// primary key for the menu relation (M2M).
	MenuPrimaryKey = []string{"admin_role_id", "admin_menus_id"}
)

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsEnable holds the default value on creation for the "is_enable" field.
	DefaultIsEnable bool
)
