// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Catalog is the predicate function for catalog builders.
type Catalog func(*sql.Selector)

// Connector is the predicate function for connector builders.
type Connector func(*sql.Selector)

// CostReport is the predicate function for costreport builders.
type CostReport func(*sql.Selector)

// DistributeLock is the predicate function for distributelock builders.
type DistributeLock func(*sql.Selector)

// Environment is the predicate function for environment builders.
type Environment func(*sql.Selector)

// EnvironmentConnectorRelationship is the predicate function for environmentconnectorrelationship builders.
type EnvironmentConnectorRelationship func(*sql.Selector)

// Perspective is the predicate function for perspective builders.
type Perspective func(*sql.Selector)

// Project is the predicate function for project builders.
type Project func(*sql.Selector)

// Role is the predicate function for role builders.
type Role func(*sql.Selector)

// Service is the predicate function for service builders.
type Service func(*sql.Selector)

// ServiceRelationship is the predicate function for servicerelationship builders.
type ServiceRelationship func(*sql.Selector)

// ServiceResource is the predicate function for serviceresource builders.
type ServiceResource func(*sql.Selector)

// ServiceResourceRelationship is the predicate function for serviceresourcerelationship builders.
type ServiceResourceRelationship func(*sql.Selector)

// ServiceRevision is the predicate function for servicerevision builders.
type ServiceRevision func(*sql.Selector)

// Setting is the predicate function for setting builders.
type Setting func(*sql.Selector)

// Subject is the predicate function for subject builders.
type Subject func(*sql.Selector)

// SubjectRoleRelationship is the predicate function for subjectrolerelationship builders.
type SubjectRoleRelationship func(*sql.Selector)

// Template is the predicate function for template builders.
type Template func(*sql.Selector)

// TemplateVersion is the predicate function for templateversion builders.
type TemplateVersion func(*sql.Selector)

// Token is the predicate function for token builders.
type Token func(*sql.Selector)

// Variable is the predicate function for variable builders.
type Variable func(*sql.Selector)
