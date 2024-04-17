// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1alpha1

import (
	workflowv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	v1 "github.com/seal-io/walrus/pkg/clients/applyconfiguration/core/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// TemplateApplyConfiguration represents an declarative configuration of the Template type for use
// with apply.
type TemplateApplyConfiguration struct {
	Name                         *string                                  `json:"name,omitempty"`
	Inputs                       *InputsApplyConfiguration                `json:"inputs,omitempty"`
	Outputs                      *OutputsApplyConfiguration               `json:"outputs,omitempty"`
	NodeSelector                 map[string]string                        `json:"nodeSelector,omitempty"`
	Affinity                     *v1.AffinityApplyConfiguration           `json:"affinity,omitempty"`
	Metadata                     *MetadataApplyConfiguration              `json:"metadata,omitempty"`
	Daemon                       *bool                                    `json:"daemon,omitempty"`
	Steps                        []workflowv1alpha1.ParallelSteps         `json:"steps,omitempty"`
	Container                    *v1.ContainerApplyConfiguration          `json:"container,omitempty"`
	ContainerSet                 *ContainerSetTemplateApplyConfiguration  `json:"containerSet,omitempty"`
	Script                       *ScriptTemplateApplyConfiguration        `json:"script,omitempty"`
	Resource                     *ResourceTemplateApplyConfiguration      `json:"resource,omitempty"`
	DAG                          *DAGTemplateApplyConfiguration           `json:"dag,omitempty"`
	Suspend                      *SuspendTemplateApplyConfiguration       `json:"suspend,omitempty"`
	Data                         *DataApplyConfiguration                  `json:"data,omitempty"`
	HTTP                         *HTTPApplyConfiguration                  `json:"http,omitempty"`
	Plugin                       *PluginApplyConfiguration                `json:"plugin,omitempty"`
	Volumes                      []v1.VolumeApplyConfiguration            `json:"volumes,omitempty"`
	InitContainers               []UserContainerApplyConfiguration        `json:"initContainers,omitempty"`
	Sidecars                     []UserContainerApplyConfiguration        `json:"sidecars,omitempty"`
	ArchiveLocation              *ArtifactLocationApplyConfiguration      `json:"archiveLocation,omitempty"`
	ActiveDeadlineSeconds        *intstr.IntOrString                      `json:"activeDeadlineSeconds,omitempty"`
	RetryStrategy                *RetryStrategyApplyConfiguration         `json:"retryStrategy,omitempty"`
	Parallelism                  *int64                                   `json:"parallelism,omitempty"`
	FailFast                     *bool                                    `json:"failFast,omitempty"`
	Tolerations                  []v1.TolerationApplyConfiguration        `json:"tolerations,omitempty"`
	SchedulerName                *string                                  `json:"schedulerName,omitempty"`
	PriorityClassName            *string                                  `json:"priorityClassName,omitempty"`
	Priority                     *int32                                   `json:"priority,omitempty"`
	ServiceAccountName           *string                                  `json:"serviceAccountName,omitempty"`
	AutomountServiceAccountToken *bool                                    `json:"automountServiceAccountToken,omitempty"`
	Executor                     *ExecutorConfigApplyConfiguration        `json:"executor,omitempty"`
	HostAliases                  []v1.HostAliasApplyConfiguration         `json:"hostAliases,omitempty"`
	SecurityContext              *v1.PodSecurityContextApplyConfiguration `json:"securityContext,omitempty"`
	PodSpecPatch                 *string                                  `json:"podSpecPatch,omitempty"`
	Metrics                      *MetricsApplyConfiguration               `json:"metrics,omitempty"`
	Synchronization              *SynchronizationApplyConfiguration       `json:"synchronization,omitempty"`
	Memoize                      *MemoizeApplyConfiguration               `json:"memoize,omitempty"`
	Timeout                      *string                                  `json:"timeout,omitempty"`
}

// TemplateApplyConfiguration constructs an declarative configuration of the Template type for use with
// apply.
func Template() *TemplateApplyConfiguration {
	return &TemplateApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithName(value string) *TemplateApplyConfiguration {
	b.Name = &value
	return b
}

// WithInputs sets the Inputs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Inputs field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithInputs(value *InputsApplyConfiguration) *TemplateApplyConfiguration {
	b.Inputs = value
	return b
}

// WithOutputs sets the Outputs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Outputs field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithOutputs(value *OutputsApplyConfiguration) *TemplateApplyConfiguration {
	b.Outputs = value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *TemplateApplyConfiguration) WithNodeSelector(entries map[string]string) *TemplateApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithAffinity(value *v1.AffinityApplyConfiguration) *TemplateApplyConfiguration {
	b.Affinity = value
	return b
}

// WithMetadata sets the Metadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Metadata field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithMetadata(value *MetadataApplyConfiguration) *TemplateApplyConfiguration {
	b.Metadata = value
	return b
}

// WithDaemon sets the Daemon field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Daemon field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithDaemon(value bool) *TemplateApplyConfiguration {
	b.Daemon = &value
	return b
}

// WithSteps adds the given value to the Steps field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Steps field.
func (b *TemplateApplyConfiguration) WithSteps(values ...workflowv1alpha1.ParallelSteps) *TemplateApplyConfiguration {
	for i := range values {
		b.Steps = append(b.Steps, values[i])
	}
	return b
}

// WithContainer sets the Container field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Container field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithContainer(value *v1.ContainerApplyConfiguration) *TemplateApplyConfiguration {
	b.Container = value
	return b
}

// WithContainerSet sets the ContainerSet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerSet field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithContainerSet(value *ContainerSetTemplateApplyConfiguration) *TemplateApplyConfiguration {
	b.ContainerSet = value
	return b
}

// WithScript sets the Script field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Script field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithScript(value *ScriptTemplateApplyConfiguration) *TemplateApplyConfiguration {
	b.Script = value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithResource(value *ResourceTemplateApplyConfiguration) *TemplateApplyConfiguration {
	b.Resource = value
	return b
}

// WithDAG sets the DAG field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DAG field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithDAG(value *DAGTemplateApplyConfiguration) *TemplateApplyConfiguration {
	b.DAG = value
	return b
}

// WithSuspend sets the Suspend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Suspend field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithSuspend(value *SuspendTemplateApplyConfiguration) *TemplateApplyConfiguration {
	b.Suspend = value
	return b
}

// WithData sets the Data field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Data field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithData(value *DataApplyConfiguration) *TemplateApplyConfiguration {
	b.Data = value
	return b
}

// WithHTTP sets the HTTP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTP field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithHTTP(value *HTTPApplyConfiguration) *TemplateApplyConfiguration {
	b.HTTP = value
	return b
}

// WithPlugin sets the Plugin field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Plugin field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithPlugin(value *PluginApplyConfiguration) *TemplateApplyConfiguration {
	b.Plugin = value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *TemplateApplyConfiguration) WithVolumes(values ...*v1.VolumeApplyConfiguration) *TemplateApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumes")
		}
		b.Volumes = append(b.Volumes, *values[i])
	}
	return b
}

// WithInitContainers adds the given value to the InitContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InitContainers field.
func (b *TemplateApplyConfiguration) WithInitContainers(values ...*UserContainerApplyConfiguration) *TemplateApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithInitContainers")
		}
		b.InitContainers = append(b.InitContainers, *values[i])
	}
	return b
}

// WithSidecars adds the given value to the Sidecars field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Sidecars field.
func (b *TemplateApplyConfiguration) WithSidecars(values ...*UserContainerApplyConfiguration) *TemplateApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSidecars")
		}
		b.Sidecars = append(b.Sidecars, *values[i])
	}
	return b
}

// WithArchiveLocation sets the ArchiveLocation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ArchiveLocation field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithArchiveLocation(value *ArtifactLocationApplyConfiguration) *TemplateApplyConfiguration {
	b.ArchiveLocation = value
	return b
}

// WithActiveDeadlineSeconds sets the ActiveDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveDeadlineSeconds field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithActiveDeadlineSeconds(value intstr.IntOrString) *TemplateApplyConfiguration {
	b.ActiveDeadlineSeconds = &value
	return b
}

// WithRetryStrategy sets the RetryStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RetryStrategy field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithRetryStrategy(value *RetryStrategyApplyConfiguration) *TemplateApplyConfiguration {
	b.RetryStrategy = value
	return b
}

// WithParallelism sets the Parallelism field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Parallelism field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithParallelism(value int64) *TemplateApplyConfiguration {
	b.Parallelism = &value
	return b
}

// WithFailFast sets the FailFast field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailFast field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithFailFast(value bool) *TemplateApplyConfiguration {
	b.FailFast = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *TemplateApplyConfiguration) WithTolerations(values ...*v1.TolerationApplyConfiguration) *TemplateApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}

// WithSchedulerName sets the SchedulerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchedulerName field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithSchedulerName(value string) *TemplateApplyConfiguration {
	b.SchedulerName = &value
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithPriorityClassName(value string) *TemplateApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithPriority(value int32) *TemplateApplyConfiguration {
	b.Priority = &value
	return b
}

// WithServiceAccountName sets the ServiceAccountName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountName field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithServiceAccountName(value string) *TemplateApplyConfiguration {
	b.ServiceAccountName = &value
	return b
}

// WithAutomountServiceAccountToken sets the AutomountServiceAccountToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AutomountServiceAccountToken field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithAutomountServiceAccountToken(value bool) *TemplateApplyConfiguration {
	b.AutomountServiceAccountToken = &value
	return b
}

// WithExecutor sets the Executor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Executor field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithExecutor(value *ExecutorConfigApplyConfiguration) *TemplateApplyConfiguration {
	b.Executor = value
	return b
}

// WithHostAliases adds the given value to the HostAliases field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HostAliases field.
func (b *TemplateApplyConfiguration) WithHostAliases(values ...*v1.HostAliasApplyConfiguration) *TemplateApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHostAliases")
		}
		b.HostAliases = append(b.HostAliases, *values[i])
	}
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithSecurityContext(value *v1.PodSecurityContextApplyConfiguration) *TemplateApplyConfiguration {
	b.SecurityContext = value
	return b
}

// WithPodSpecPatch sets the PodSpecPatch field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodSpecPatch field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithPodSpecPatch(value string) *TemplateApplyConfiguration {
	b.PodSpecPatch = &value
	return b
}

// WithMetrics sets the Metrics field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Metrics field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithMetrics(value *MetricsApplyConfiguration) *TemplateApplyConfiguration {
	b.Metrics = value
	return b
}

// WithSynchronization sets the Synchronization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Synchronization field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithSynchronization(value *SynchronizationApplyConfiguration) *TemplateApplyConfiguration {
	b.Synchronization = value
	return b
}

// WithMemoize sets the Memoize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Memoize field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithMemoize(value *MemoizeApplyConfiguration) *TemplateApplyConfiguration {
	b.Memoize = value
	return b
}

// WithTimeout sets the Timeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Timeout field is set to the value of the last call.
func (b *TemplateApplyConfiguration) WithTimeout(value string) *TemplateApplyConfiguration {
	b.Timeout = &value
	return b
}