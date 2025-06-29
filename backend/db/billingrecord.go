// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chaitin/MonkeyCode/backend/db/billingrecord"
)

// BillingRecord is the model entity for the BillingRecord schema.
type BillingRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation string `json:"operation,omitempty"`
	// InputTokens holds the value of the "input_tokens" field.
	InputTokens int64 `json:"input_tokens,omitempty"`
	// OutputTokens holds the value of the "output_tokens" field.
	OutputTokens int64 `json:"output_tokens,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost int64 `json:"cost,omitempty"`
	// RequestTime holds the value of the "request_time" field.
	RequestTime time.Time `json:"request_time,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillingRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billingrecord.FieldMetadata:
			values[i] = new([]byte)
		case billingrecord.FieldInputTokens, billingrecord.FieldOutputTokens, billingrecord.FieldCost:
			values[i] = new(sql.NullInt64)
		case billingrecord.FieldID, billingrecord.FieldTenantID, billingrecord.FieldUserID, billingrecord.FieldModel, billingrecord.FieldOperation:
			values[i] = new(sql.NullString)
		case billingrecord.FieldRequestTime, billingrecord.FieldCreatedAt, billingrecord.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillingRecord fields.
func (br *BillingRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billingrecord.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				br.ID = value.String
			}
		case billingrecord.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				br.TenantID = value.String
			}
		case billingrecord.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				br.UserID = value.String
			}
		case billingrecord.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				br.Model = value.String
			}
		case billingrecord.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				br.Operation = value.String
			}
		case billingrecord.FieldInputTokens:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field input_tokens", values[i])
			} else if value.Valid {
				br.InputTokens = value.Int64
			}
		case billingrecord.FieldOutputTokens:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field output_tokens", values[i])
			} else if value.Valid {
				br.OutputTokens = value.Int64
			}
		case billingrecord.FieldCost:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cost", values[i])
			} else if value.Valid {
				br.Cost = value.Int64
			}
		case billingrecord.FieldRequestTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field request_time", values[i])
			} else if value.Valid {
				br.RequestTime = value.Time
			}
		case billingrecord.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &br.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case billingrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				br.CreatedAt = value.Time
			}
		case billingrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				br.UpdatedAt = value.Time
			}
		default:
			br.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillingRecord.
// This includes values selected through modifiers, order, etc.
func (br *BillingRecord) Value(name string) (ent.Value, error) {
	return br.selectValues.Get(name)
}

// Update returns a builder for updating this BillingRecord.
// Note that you need to call BillingRecord.Unwrap() before calling this method if this BillingRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (br *BillingRecord) Update() *BillingRecordUpdateOne {
	return NewBillingRecordClient(br.config).UpdateOne(br)
}

// Unwrap unwraps the BillingRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (br *BillingRecord) Unwrap() *BillingRecord {
	_tx, ok := br.config.driver.(*txDriver)
	if !ok {
		panic("db: BillingRecord is not a transactional entity")
	}
	br.config.driver = _tx.drv
	return br
}

// String implements the fmt.Stringer.
func (br *BillingRecord) String() string {
	var builder strings.Builder
	builder.WriteString("BillingRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", br.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(br.TenantID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(br.UserID)
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(br.Model)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(br.Operation)
	builder.WriteString(", ")
	builder.WriteString("input_tokens=")
	builder.WriteString(fmt.Sprintf("%v", br.InputTokens))
	builder.WriteString(", ")
	builder.WriteString("output_tokens=")
	builder.WriteString(fmt.Sprintf("%v", br.OutputTokens))
	builder.WriteString(", ")
	builder.WriteString("cost=")
	builder.WriteString(fmt.Sprintf("%v", br.Cost))
	builder.WriteString(", ")
	builder.WriteString("request_time=")
	builder.WriteString(br.RequestTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", br.Metadata))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(br.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(br.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BillingRecords is a parsable slice of BillingRecord.
type BillingRecords []*BillingRecord
