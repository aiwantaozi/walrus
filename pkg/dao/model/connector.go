// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// Connector is the model entity for the Connector schema.
type Connector struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]string `json:"annotations,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Status holds the value of the "status" field.
	Status status.Status `json:"status,omitempty"`
	// ID of the project to belong, empty means for all projects.
	ProjectID object.ID `json:"project_id,omitempty"`
	// Category of the connector.
	Category string `json:"category,omitempty"`
	// Type of the connector.
	Type string `json:"type,omitempty"`
	// Connector config version.
	ConfigVersion string `json:"config_version,omitempty"`
	// Connector config data.
	ConfigData crypto.Properties `json:"config_data,omitempty"`
	// Config whether enable finOps, will install prometheus and opencost while enable.
	EnableFinOps bool `json:"enable_fin_ops,omitempty"`
	// Custom pricing user defined.
	FinOpsCustomPricing *types.FinOpsCustomPricing `json:"fin_ops_custom_pricing,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConnectorQuery when eager-loading is set.
	Edges        ConnectorEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ConnectorEdges holds the relations/edges for other nodes in the graph.
type ConnectorEdges struct {
	// Project to which the connector belongs.
	Project *Project `json:"project,omitempty"`
	// Environments holds the value of the environments edge.
	Environments []*EnvironmentConnectorRelationship `json:"environments,omitempty"`
	// ServiceResources that use the connector.
	Resources []*ServiceResource `json:"resources,omitempty"`
	// CostReports that linked to the connection.
	CostReports []*CostReport `json:"cost_reports,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConnectorEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// EnvironmentsOrErr returns the Environments value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectorEdges) EnvironmentsOrErr() ([]*EnvironmentConnectorRelationship, error) {
	if e.loadedTypes[1] {
		return e.Environments, nil
	}
	return nil, &NotLoadedError{edge: "environments"}
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectorEdges) ResourcesOrErr() ([]*ServiceResource, error) {
	if e.loadedTypes[2] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// CostReportsOrErr returns the CostReports value or an error if the edge
// was not loaded in eager-loading.
func (e ConnectorEdges) CostReportsOrErr() ([]*CostReport, error) {
	if e.loadedTypes[3] {
		return e.CostReports, nil
	}
	return nil, &NotLoadedError{edge: "cost_reports"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Connector) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case connector.FieldLabels, connector.FieldAnnotations, connector.FieldStatus, connector.FieldFinOpsCustomPricing:
			values[i] = new([]byte)
		case connector.FieldConfigData:
			values[i] = new(crypto.Properties)
		case connector.FieldID, connector.FieldProjectID:
			values[i] = new(object.ID)
		case connector.FieldEnableFinOps:
			values[i] = new(sql.NullBool)
		case connector.FieldName, connector.FieldDescription, connector.FieldCategory, connector.FieldType, connector.FieldConfigVersion:
			values[i] = new(sql.NullString)
		case connector.FieldCreateTime, connector.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Connector fields.
func (c *Connector) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case connector.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case connector.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case connector.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case connector.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case connector.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case connector.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = new(time.Time)
				*c.CreateTime = value.Time
			}
		case connector.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = new(time.Time)
				*c.UpdateTime = value.Time
			}
		case connector.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case connector.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				c.ProjectID = *value
			}
		case connector.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				c.Category = value.String
			}
		case connector.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case connector.FieldConfigVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config_version", values[i])
			} else if value.Valid {
				c.ConfigVersion = value.String
			}
		case connector.FieldConfigData:
			if value, ok := values[i].(*crypto.Properties); !ok {
				return fmt.Errorf("unexpected type %T for field config_data", values[i])
			} else if value != nil {
				c.ConfigData = *value
			}
		case connector.FieldEnableFinOps:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable_fin_ops", values[i])
			} else if value.Valid {
				c.EnableFinOps = value.Bool
			}
		case connector.FieldFinOpsCustomPricing:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field fin_ops_custom_pricing", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.FinOpsCustomPricing); err != nil {
					return fmt.Errorf("unmarshal field fin_ops_custom_pricing: %w", err)
				}
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Connector.
// This includes values selected through modifiers, order, etc.
func (c *Connector) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the Connector entity.
func (c *Connector) QueryProject() *ProjectQuery {
	return NewConnectorClient(c.config).QueryProject(c)
}

// QueryEnvironments queries the "environments" edge of the Connector entity.
func (c *Connector) QueryEnvironments() *EnvironmentConnectorRelationshipQuery {
	return NewConnectorClient(c.config).QueryEnvironments(c)
}

// QueryResources queries the "resources" edge of the Connector entity.
func (c *Connector) QueryResources() *ServiceResourceQuery {
	return NewConnectorClient(c.config).QueryResources(c)
}

// QueryCostReports queries the "cost_reports" edge of the Connector entity.
func (c *Connector) QueryCostReports() *CostReportQuery {
	return NewConnectorClient(c.config).QueryCostReports(c)
}

// Update returns a builder for updating this Connector.
// Note that you need to call Connector.Unwrap() before calling this method if this Connector
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Connector) Update() *ConnectorUpdateOne {
	return NewConnectorClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Connector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Connector) Unwrap() *Connector {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("model: Connector is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Connector) String() string {
	var builder strings.Builder
	builder.WriteString("Connector(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", c.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", c.Annotations))
	builder.WriteString(", ")
	if v := c.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(c.Category)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteString(", ")
	builder.WriteString("config_version=")
	builder.WriteString(c.ConfigVersion)
	builder.WriteString(", ")
	builder.WriteString("config_data=")
	builder.WriteString(fmt.Sprintf("%v", c.ConfigData))
	builder.WriteString(", ")
	builder.WriteString("enable_fin_ops=")
	builder.WriteString(fmt.Sprintf("%v", c.EnableFinOps))
	builder.WriteString(", ")
	builder.WriteString("fin_ops_custom_pricing=")
	builder.WriteString(fmt.Sprintf("%v", c.FinOpsCustomPricing))
	builder.WriteByte(')')
	return builder.String()
}

// Connectors is a parsable slice of Connector.
type Connectors []*Connector
