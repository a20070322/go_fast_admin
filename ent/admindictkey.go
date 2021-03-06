// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/a20070322/go_fast_admin/ent/admindict"
	"github.com/a20070322/go_fast_admin/ent/admindictkey"
)

// AdminDictKey is the model entity for the AdminDictKey schema.
type AdminDictKey struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"-"`
	// DictLabel holds the value of the "dict_label" field.
	// 字典标签
	DictLabel string `json:"dict_label,omitempty"`
	// DictCode holds the value of the "dict_code" field.
	// 字典键值
	DictCode string `json:"dict_code,omitempty"`
	// Sort holds the value of the "sort" field.
	// 排序
	Sort int `json:"sort,omitempty"`
	// Remarks holds the value of the "remarks" field.
	// 备注
	Remarks string `json:"remarks,omitempty"`
	// IsEnable holds the value of the "is_enable" field.
	// 状态
	IsEnable bool `json:"is_enable,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdminDictKeyQuery when eager-loading is set.
	Edges          AdminDictKeyEdges `json:"edges"`
	admin_dict_key *int
}

// AdminDictKeyEdges holds the relations/edges for other nodes in the graph.
type AdminDictKeyEdges struct {
	// P holds the value of the P edge.
	P *AdminDict `json:"P,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// POrErr returns the P value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdminDictKeyEdges) POrErr() (*AdminDict, error) {
	if e.loadedTypes[0] {
		if e.P == nil {
			// The edge P was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: admindict.Label}
		}
		return e.P, nil
	}
	return nil, &NotLoadedError{edge: "P"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AdminDictKey) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case admindictkey.FieldIsEnable:
			values[i] = new(sql.NullBool)
		case admindictkey.FieldID, admindictkey.FieldSort:
			values[i] = new(sql.NullInt64)
		case admindictkey.FieldDictLabel, admindictkey.FieldDictCode, admindictkey.FieldRemarks:
			values[i] = new(sql.NullString)
		case admindictkey.FieldCreatedAt, admindictkey.FieldUpdatedAt, admindictkey.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case admindictkey.ForeignKeys[0]: // admin_dict_key
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AdminDictKey", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdminDictKey fields.
func (adk *AdminDictKey) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case admindictkey.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			adk.ID = int(value.Int64)
		case admindictkey.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				adk.CreatedAt = value.Time
			}
		case admindictkey.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				adk.UpdatedAt = value.Time
			}
		case admindictkey.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				adk.DeletedAt = value.Time
			}
		case admindictkey.FieldDictLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dict_label", values[i])
			} else if value.Valid {
				adk.DictLabel = value.String
			}
		case admindictkey.FieldDictCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dict_code", values[i])
			} else if value.Valid {
				adk.DictCode = value.String
			}
		case admindictkey.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				adk.Sort = int(value.Int64)
			}
		case admindictkey.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				adk.Remarks = value.String
			}
		case admindictkey.FieldIsEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_enable", values[i])
			} else if value.Valid {
				adk.IsEnable = value.Bool
			}
		case admindictkey.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field admin_dict_key", value)
			} else if value.Valid {
				adk.admin_dict_key = new(int)
				*adk.admin_dict_key = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryP queries the "P" edge of the AdminDictKey entity.
func (adk *AdminDictKey) QueryP() *AdminDictQuery {
	return (&AdminDictKeyClient{config: adk.config}).QueryP(adk)
}

// Update returns a builder for updating this AdminDictKey.
// Note that you need to call AdminDictKey.Unwrap() before calling this method if this AdminDictKey
// was returned from a transaction, and the transaction was committed or rolled back.
func (adk *AdminDictKey) Update() *AdminDictKeyUpdateOne {
	return (&AdminDictKeyClient{config: adk.config}).UpdateOne(adk)
}

// Unwrap unwraps the AdminDictKey entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (adk *AdminDictKey) Unwrap() *AdminDictKey {
	tx, ok := adk.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdminDictKey is not a transactional entity")
	}
	adk.config.driver = tx.drv
	return adk
}

// String implements the fmt.Stringer.
func (adk *AdminDictKey) String() string {
	var builder strings.Builder
	builder.WriteString("AdminDictKey(")
	builder.WriteString(fmt.Sprintf("id=%v", adk.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(adk.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(adk.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(adk.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", dict_label=")
	builder.WriteString(adk.DictLabel)
	builder.WriteString(", dict_code=")
	builder.WriteString(adk.DictCode)
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", adk.Sort))
	builder.WriteString(", remarks=")
	builder.WriteString(adk.Remarks)
	builder.WriteString(", is_enable=")
	builder.WriteString(fmt.Sprintf("%v", adk.IsEnable))
	builder.WriteByte(')')
	return builder.String()
}

// AdminDictKeys is a parsable slice of AdminDictKey.
type AdminDictKeys []*AdminDictKey

func (adk AdminDictKeys) config(cfg config) {
	for _i := range adk {
		adk[_i].config = cfg
	}
}
