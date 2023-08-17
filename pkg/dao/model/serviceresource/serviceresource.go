// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package serviceresource

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"golang.org/x/exp/slices"
)

const (
	// Label holds the string label denoting the serviceresource type in the database.
	Label = "service_resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldEnvironmentID holds the string denoting the environment_id field in the database.
	FieldEnvironmentID = "environment_id"
	// FieldServiceID holds the string denoting the service_id field in the database.
	FieldServiceID = "service_id"
	// FieldConnectorID holds the string denoting the connector_id field in the database.
	FieldConnectorID = "connector_id"
	// FieldCompositionID holds the string denoting the composition_id field in the database.
	FieldCompositionID = "composition_id"
	// FieldClassID holds the string denoting the class_id field in the database.
	FieldClassID = "class_id"
	// FieldMode holds the string denoting the mode field in the database.
	FieldMode = "mode"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDeployerType holds the string denoting the deployer_type field in the database.
	FieldDeployerType = "deployer_type"
	// FieldShape holds the string denoting the shape field in the database.
	FieldShape = "shape"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeEnvironment holds the string denoting the environment edge name in mutations.
	EdgeEnvironment = "environment"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// EdgeConnector holds the string denoting the connector edge name in mutations.
	EdgeConnector = "connector"
	// EdgeComposition holds the string denoting the composition edge name in mutations.
	EdgeComposition = "composition"
	// EdgeComponents holds the string denoting the components edge name in mutations.
	EdgeComponents = "components"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeInstances holds the string denoting the instances edge name in mutations.
	EdgeInstances = "instances"
	// EdgeDependencies holds the string denoting the dependencies edge name in mutations.
	EdgeDependencies = "dependencies"
	// Table holds the table name of the serviceresource in the database.
	Table = "service_resources"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "service_resources"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// EnvironmentTable is the table that holds the environment relation/edge.
	EnvironmentTable = "service_resources"
	// EnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	EnvironmentInverseTable = "environments"
	// EnvironmentColumn is the table column denoting the environment relation/edge.
	EnvironmentColumn = "environment_id"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "service_resources"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "service_id"
	// ConnectorTable is the table that holds the connector relation/edge.
	ConnectorTable = "service_resources"
	// ConnectorInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorInverseTable = "connectors"
	// ConnectorColumn is the table column denoting the connector relation/edge.
	ConnectorColumn = "connector_id"
	// CompositionTable is the table that holds the composition relation/edge.
	CompositionTable = "service_resources"
	// CompositionColumn is the table column denoting the composition relation/edge.
	CompositionColumn = "composition_id"
	// ComponentsTable is the table that holds the components relation/edge.
	ComponentsTable = "service_resources"
	// ComponentsColumn is the table column denoting the components relation/edge.
	ComponentsColumn = "composition_id"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "service_resources"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_id"
	// InstancesTable is the table that holds the instances relation/edge.
	InstancesTable = "service_resources"
	// InstancesColumn is the table column denoting the instances relation/edge.
	InstancesColumn = "class_id"
	// DependenciesTable is the table that holds the dependencies relation/edge.
	DependenciesTable = "service_resource_relationships"
	// DependenciesInverseTable is the table name for the ServiceResourceRelationship entity.
	// It exists in this package in order to avoid circular dependency with the "serviceresourcerelationship" package.
	DependenciesInverseTable = "service_resource_relationships"
	// DependenciesColumn is the table column denoting the dependencies relation/edge.
	DependenciesColumn = "service_resource_id"
)

// Columns holds all SQL columns for serviceresource fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldProjectID,
	FieldEnvironmentID,
	FieldServiceID,
	FieldConnectorID,
	FieldCompositionID,
	FieldClassID,
	FieldMode,
	FieldType,
	FieldName,
	FieldDeployerType,
	FieldShape,
	FieldStatus,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/seal-io/walrus/pkg/dao/model/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// ProjectIDValidator is a validator for the "project_id" field. It is called by the builders before save.
	ProjectIDValidator func(string) error
	// EnvironmentIDValidator is a validator for the "environment_id" field. It is called by the builders before save.
	EnvironmentIDValidator func(string) error
	// ServiceIDValidator is a validator for the "service_id" field. It is called by the builders before save.
	ServiceIDValidator func(string) error
	// ConnectorIDValidator is a validator for the "connector_id" field. It is called by the builders before save.
	ConnectorIDValidator func(string) error
	// ModeValidator is a validator for the "mode" field. It is called by the builders before save.
	ModeValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DeployerTypeValidator is a validator for the "deployer_type" field. It is called by the builders before save.
	DeployerTypeValidator func(string) error
	// ShapeValidator is a validator for the "shape" field. It is called by the builders before save.
	ShapeValidator func(string) error
)

// OrderOption defines the ordering options for the ServiceResource queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByEnvironmentID orders the results by the environment_id field.
func ByEnvironmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnvironmentID, opts...).ToFunc()
}

// ByServiceID orders the results by the service_id field.
func ByServiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServiceID, opts...).ToFunc()
}

// ByConnectorID orders the results by the connector_id field.
func ByConnectorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConnectorID, opts...).ToFunc()
}

// ByCompositionID orders the results by the composition_id field.
func ByCompositionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompositionID, opts...).ToFunc()
}

// ByClassID orders the results by the class_id field.
func ByClassID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClassID, opts...).ToFunc()
}

// ByMode orders the results by the mode field.
func ByMode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMode, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDeployerType orders the results by the deployer_type field.
func ByDeployerType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeployerType, opts...).ToFunc()
}

// ByShape orders the results by the shape field.
func ByShape(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShape, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByEnvironmentField orders the results by environment field.
func ByEnvironmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEnvironmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByServiceField orders the results by service field.
func ByServiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByConnectorField orders the results by connector field.
func ByConnectorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConnectorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCompositionField orders the results by composition field.
func ByCompositionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompositionStep(), sql.OrderByField(field, opts...))
	}
}

// ByComponentsCount orders the results by components count.
func ByComponentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newComponentsStep(), opts...)
	}
}

// ByComponents orders the results by components terms.
func ByComponents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newComponentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByClassField orders the results by class field.
func ByClassField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassStep(), sql.OrderByField(field, opts...))
	}
}

// ByInstancesCount orders the results by instances count.
func ByInstancesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInstancesStep(), opts...)
	}
}

// ByInstances orders the results by instances terms.
func ByInstances(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstancesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDependenciesCount orders the results by dependencies count.
func ByDependenciesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDependenciesStep(), opts...)
	}
}

// ByDependencies orders the results by dependencies terms.
func ByDependencies(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDependenciesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newEnvironmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EnvironmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EnvironmentTable, EnvironmentColumn),
	)
}
func newServiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
	)
}
func newConnectorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConnectorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ConnectorTable, ConnectorColumn),
	)
}
func newCompositionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompositionTable, CompositionColumn),
	)
}
func newComponentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ComponentsTable, ComponentsColumn),
	)
}
func newClassStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
	)
}
func newInstancesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
	)
}
func newDependenciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DependenciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, DependenciesTable, DependenciesColumn),
	)
}

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return slices.Clone(Columns)
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}
