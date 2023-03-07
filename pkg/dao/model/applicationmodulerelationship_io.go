// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types"
)

// ApplicationModuleRelationshipQueryInput is the input for the ApplicationModuleRelationship query.
type ApplicationModuleRelationshipQueryInput struct {
	// ID of the application to which the relationship connects.
	ApplicationID types.ID `json:"applicationId"`
	// ID of the module to which the relationship connects.
	ModuleID string `json:"moduleId"`
	// Name of the module customized to the application.
	Name string `json:"name"`
}

// Model converts the ApplicationModuleRelationshipQueryInput to ApplicationModuleRelationship.
func (in ApplicationModuleRelationshipQueryInput) Model() *ApplicationModuleRelationship {
	return &ApplicationModuleRelationship{
		ApplicationID: in.ApplicationID,
		ModuleID:      in.ModuleID,
		Name:          in.Name,
	}
}

// ApplicationModuleRelationshipCreateInput is the input for the ApplicationModuleRelationship creation.
type ApplicationModuleRelationshipCreateInput struct {
	// Name of the module customized to the application.
	Name string `json:"name"`
	// Attributes to configure the module.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Applications that connect to the relationship.
	Application ApplicationQueryInput `json:"application"`
	// Modules that connect to the relationship.
	Module ModuleQueryInput `json:"module"`
}

// Model converts the ApplicationModuleRelationshipCreateInput to ApplicationModuleRelationship.
func (in ApplicationModuleRelationshipCreateInput) Model() *ApplicationModuleRelationship {
	var entity = &ApplicationModuleRelationship{
		Name:       in.Name,
		Attributes: in.Attributes,
	}
	entity.ApplicationID = in.Application.ID
	entity.ModuleID = in.Module.ID
	return entity
}

// ApplicationModuleRelationshipUpdateInput is the input for the ApplicationModuleRelationship modification.
type ApplicationModuleRelationshipUpdateInput struct {
	// Name of the module customized to the application.
	Name string `json:"name"`
	// Attributes to configure the module.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Applications that connect to the relationship.
	Application ApplicationQueryInput `json:"application,omitempty"`
	// Modules that connect to the relationship.
	Module ModuleQueryInput `json:"module,omitempty"`
}

// Model converts the ApplicationModuleRelationshipUpdateInput to ApplicationModuleRelationship.
func (in ApplicationModuleRelationshipUpdateInput) Model() *ApplicationModuleRelationship {
	var entity = &ApplicationModuleRelationship{
		Name:       in.Name,
		Attributes: in.Attributes,
	}
	entity.ApplicationID = in.Application.ID
	entity.ModuleID = in.Module.ID
	return entity
}

// ApplicationModuleRelationshipOutput is the output for the ApplicationModuleRelationship.
type ApplicationModuleRelationshipOutput struct {
	// Name of the module customized to the application.
	Name string `json:"name,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Attributes to configure the module.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Applications that connect to the relationship.
	Application *ApplicationOutput `json:"application,omitempty"`
	// Modules that connect to the relationship.
	Module *ModuleOutput `json:"module,omitempty"`
}

// ExposeApplicationModuleRelationship converts the ApplicationModuleRelationship to ApplicationModuleRelationshipOutput.
func ExposeApplicationModuleRelationship(in *ApplicationModuleRelationship) *ApplicationModuleRelationshipOutput {
	if in == nil {
		return nil
	}
	var entity = &ApplicationModuleRelationshipOutput{
		Name:        in.Name,
		CreateTime:  in.CreateTime,
		UpdateTime:  in.UpdateTime,
		Attributes:  in.Attributes,
		Application: ExposeApplication(in.Edges.Application),
		Module:      ExposeModule(in.Edges.Module),
	}
	if entity.Application == nil {
		entity.Application = &ApplicationOutput{}
	}
	entity.Application.ID = in.ApplicationID
	if entity.Module == nil {
		entity.Module = &ModuleOutput{}
	}
	entity.Module.ID = in.ModuleID
	return entity
}

// ExposeApplicationModuleRelationships converts the ApplicationModuleRelationship slice to ApplicationModuleRelationshipOutput pointer slice.
func ExposeApplicationModuleRelationships(in []*ApplicationModuleRelationship) []*ApplicationModuleRelationshipOutput {
	var out = make([]*ApplicationModuleRelationshipOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeApplicationModuleRelationship(in[i])
		if o == nil {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}