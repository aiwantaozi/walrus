// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// ResourceRun is the model entity for the ResourceRun schema.
type ResourceRun struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Status holds the value of the "status" field.
	Status status.Status `json:"status,omitempty"`
	// ID of the project to belong.
	ProjectID object.ID `json:"project_id,omitempty"`
	// ID of the environment to which the run belongs.
	EnvironmentID object.ID `json:"environment_id,omitempty"`
	// ID of the resource to which the run belongs.
	ResourceID object.ID `json:"resource_id,omitempty"`
	// Name of the template.
	TemplateName string `json:"template_name,omitempty"`
	// Version of the template.
	TemplateVersion string `json:"template_version,omitempty"`
	// ID of the template.
	TemplateID object.ID `json:"template_id,omitempty"`
	// Attributes to configure the template.
	Attributes property.Values `json:"attributes,omitempty"`
	// Computed attributes generated from attributes and schemas.
	ComputedAttributes property.Values `json:"computed_attributes,omitempty"`
	// Variables of the run.
	Variables crypto.Map[string, string] `json:"variables,omitempty"`
	// Input plan of the run.
	InputPlan string `json:"-"`
	// Output of the run.
	Output string `json:"-"`
	// Type of deployer.
	DeployerType string `json:"deployer_type,omitempty"`
	// Duration in seconds of the run deploying.
	Duration int `json:"duration,omitempty"`
	// Previous provider requirement of the run.
	PreviousRequiredProviders []types.ProviderRequirement `json:"previous_required_providers,omitempty"`
	// Record of the run.
	Record string `json:"record,omitempty"`
	// Change comment of the run.
	ChangeComment string `json:"change_comment,omitempty"`
	// User who created the run.
	CreatedBy string `json:"created_by,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceRunQuery when eager-loading is set.
	Edges        ResourceRunEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ResourceRunEdges holds the relations/edges for other nodes in the graph.
type ResourceRunEdges struct {
	// Project to which the run belongs.
	Project *Project `json:"project,omitempty"`
	// Environment to which the run deploys.
	Environment *Environment `json:"environment,omitempty"`
	// Resource to which the run belongs.
	Resource *Resource `json:"resource,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceRunEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// EnvironmentOrErr returns the Environment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceRunEdges) EnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.Environment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.Environment, nil
	}
	return nil, &NotLoadedError{edge: "environment"}
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceRunEdges) ResourceOrErr() (*Resource, error) {
	if e.loadedTypes[2] {
		if e.Resource == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resource.Label}
		}
		return e.Resource, nil
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResourceRun) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resourcerun.FieldStatus, resourcerun.FieldPreviousRequiredProviders:
			values[i] = new([]byte)
		case resourcerun.FieldVariables:
			values[i] = new(crypto.Map[string, string])
		case resourcerun.FieldID, resourcerun.FieldProjectID, resourcerun.FieldEnvironmentID, resourcerun.FieldResourceID, resourcerun.FieldTemplateID:
			values[i] = new(object.ID)
		case resourcerun.FieldAttributes, resourcerun.FieldComputedAttributes:
			values[i] = new(property.Values)
		case resourcerun.FieldDuration:
			values[i] = new(sql.NullInt64)
		case resourcerun.FieldTemplateName, resourcerun.FieldTemplateVersion, resourcerun.FieldInputPlan, resourcerun.FieldOutput, resourcerun.FieldDeployerType, resourcerun.FieldRecord, resourcerun.FieldChangeComment, resourcerun.FieldCreatedBy:
			values[i] = new(sql.NullString)
		case resourcerun.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourceRun fields.
func (rr *ResourceRun) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resourcerun.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rr.ID = *value
			}
		case resourcerun.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				rr.CreateTime = new(time.Time)
				*rr.CreateTime = value.Time
			}
		case resourcerun.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rr.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case resourcerun.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				rr.ProjectID = *value
			}
		case resourcerun.FieldEnvironmentID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value != nil {
				rr.EnvironmentID = *value
			}
		case resourcerun.FieldResourceID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value != nil {
				rr.ResourceID = *value
			}
		case resourcerun.FieldTemplateName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_name", values[i])
			} else if value.Valid {
				rr.TemplateName = value.String
			}
		case resourcerun.FieldTemplateVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_version", values[i])
			} else if value.Valid {
				rr.TemplateVersion = value.String
			}
		case resourcerun.FieldTemplateID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field template_id", values[i])
			} else if value != nil {
				rr.TemplateID = *value
			}
		case resourcerun.FieldAttributes:
			if value, ok := values[i].(*property.Values); !ok {
				return fmt.Errorf("unexpected type %T for field attributes", values[i])
			} else if value != nil {
				rr.Attributes = *value
			}
		case resourcerun.FieldComputedAttributes:
			if value, ok := values[i].(*property.Values); !ok {
				return fmt.Errorf("unexpected type %T for field computed_attributes", values[i])
			} else if value != nil {
				rr.ComputedAttributes = *value
			}
		case resourcerun.FieldVariables:
			if value, ok := values[i].(*crypto.Map[string, string]); !ok {
				return fmt.Errorf("unexpected type %T for field variables", values[i])
			} else if value != nil {
				rr.Variables = *value
			}
		case resourcerun.FieldInputPlan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input_plan", values[i])
			} else if value.Valid {
				rr.InputPlan = value.String
			}
		case resourcerun.FieldOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value.Valid {
				rr.Output = value.String
			}
		case resourcerun.FieldDeployerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployer_type", values[i])
			} else if value.Valid {
				rr.DeployerType = value.String
			}
		case resourcerun.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				rr.Duration = int(value.Int64)
			}
		case resourcerun.FieldPreviousRequiredProviders:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field previous_required_providers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rr.PreviousRequiredProviders); err != nil {
					return fmt.Errorf("unmarshal field previous_required_providers: %w", err)
				}
			}
		case resourcerun.FieldRecord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field record", values[i])
			} else if value.Valid {
				rr.Record = value.String
			}
		case resourcerun.FieldChangeComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field change_comment", values[i])
			} else if value.Valid {
				rr.ChangeComment = value.String
			}
		case resourcerun.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				rr.CreatedBy = value.String
			}
		default:
			rr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResourceRun.
// This includes values selected through modifiers, order, etc.
func (rr *ResourceRun) Value(name string) (ent.Value, error) {
	return rr.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the ResourceRun entity.
func (rr *ResourceRun) QueryProject() *ProjectQuery {
	return NewResourceRunClient(rr.config).QueryProject(rr)
}

// QueryEnvironment queries the "environment" edge of the ResourceRun entity.
func (rr *ResourceRun) QueryEnvironment() *EnvironmentQuery {
	return NewResourceRunClient(rr.config).QueryEnvironment(rr)
}

// QueryResource queries the "resource" edge of the ResourceRun entity.
func (rr *ResourceRun) QueryResource() *ResourceQuery {
	return NewResourceRunClient(rr.config).QueryResource(rr)
}

// Update returns a builder for updating this ResourceRun.
// Note that you need to call ResourceRun.Unwrap() before calling this method if this ResourceRun
// was returned from a transaction, and the transaction was committed or rolled back.
func (rr *ResourceRun) Update() *ResourceRunUpdateOne {
	return NewResourceRunClient(rr.config).UpdateOne(rr)
}

// Unwrap unwraps the ResourceRun entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rr *ResourceRun) Unwrap() *ResourceRun {
	_tx, ok := rr.config.driver.(*txDriver)
	if !ok {
		panic("model: ResourceRun is not a transactional entity")
	}
	rr.config.driver = _tx.drv
	return rr
}

// String implements the fmt.Stringer.
func (rr *ResourceRun) String() string {
	var builder strings.Builder
	builder.WriteString("ResourceRun(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rr.ID))
	if v := rr.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", rr.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.EnvironmentID))
	builder.WriteString(", ")
	builder.WriteString("resource_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.ResourceID))
	builder.WriteString(", ")
	builder.WriteString("template_name=")
	builder.WriteString(rr.TemplateName)
	builder.WriteString(", ")
	builder.WriteString("template_version=")
	builder.WriteString(rr.TemplateVersion)
	builder.WriteString(", ")
	builder.WriteString("template_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.TemplateID))
	builder.WriteString(", ")
	builder.WriteString("attributes=")
	builder.WriteString(fmt.Sprintf("%v", rr.Attributes))
	builder.WriteString(", ")
	builder.WriteString("computed_attributes=")
	builder.WriteString(fmt.Sprintf("%v", rr.ComputedAttributes))
	builder.WriteString(", ")
	builder.WriteString("variables=")
	builder.WriteString(fmt.Sprintf("%v", rr.Variables))
	builder.WriteString(", ")
	builder.WriteString("input_plan=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("output=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("deployer_type=")
	builder.WriteString(rr.DeployerType)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", rr.Duration))
	builder.WriteString(", ")
	builder.WriteString("previous_required_providers=")
	builder.WriteString(fmt.Sprintf("%v", rr.PreviousRequiredProviders))
	builder.WriteString(", ")
	builder.WriteString("record=")
	builder.WriteString(rr.Record)
	builder.WriteString(", ")
	builder.WriteString("change_comment=")
	builder.WriteString(rr.ChangeComment)
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(rr.CreatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// ResourceRuns is a parsable slice of ResourceRun.
type ResourceRuns []*ResourceRun