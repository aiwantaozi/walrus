//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetCustomResourceDefinitions() map[string]*v1.CustomResourceDefinition {
	return map[string]*v1.CustomResourceDefinition{
		"Catalog":            crd_pkg_apis_walruscore_v1_Catalog(),
		"Connector":          crd_pkg_apis_walruscore_v1_Connector(),
		"Resource":           crd_pkg_apis_walruscore_v1_Resource(),
		"ResourceDefinition": crd_pkg_apis_walruscore_v1_ResourceDefinition(),
		"ResourceRun":        crd_pkg_apis_walruscore_v1_ResourceRun(),
		"Template":           crd_pkg_apis_walruscore_v1_Template(),
	}
}

func crd_pkg_apis_walruscore_v1_Catalog() *v1.CustomResourceDefinition {
	return &v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "catalogs.walruscore.seal.io",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "walruscore.seal.io",
			Names: v1.CustomResourceDefinitionNames{
				Plural:   "catalogs",
				Singular: "catalog",
				Kind:     "Catalog",
				ListKind: "CatalogList",
			},
			Scope: "Namespaced",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Description: "Catalog is the schema for the catalogs API.",
							Type:        "object",
							Required: []string{
								"metadata",
								"spec",
							},
							Properties: map[string]v1.JSONSchemaProps{
								"apiVersion": {
									Type: "string",
								},
								"kind": {
									Type: "string",
								},
								"metadata": {
									Type: "object",
								},
								"spec": {
									Type: "object",
									Required: []string{
										"templateFormat",
									},
									Properties: map[string]v1.JSONSchemaProps{
										"description": {
											Description: "Description of the catalog.",
											Type:        "string",
										},
										"filtering": {
											Description: "Filtering specifies the filtering rules for the catalog.",
											Type:        "object",
											Properties: map[string]v1.JSONSchemaProps{
												"excludeFilter": {
													Description: "ExcludeFilter specifies the regular expression for matching the exclude template name.",
													Type:        "string",
												},
												"includeFilter": {
													Description: "IncludeFilter specifies the regular expression for matching the include template name.",
													Type:        "string",
												},
											},
											Nullable: true,
										},
										"helmRepositorySource": {
											Description: "HelmRepositorySource specifies the helm repository configure.",
											Type:        "object",
											Required: []string{
												"url",
											},
											Properties: map[string]v1.JSONSchemaProps{
												"password": {
													Description: "Password for OCI registry.",
													Type:        "string",
												},
												"secretRef": {
													Description: "SecretRef specifies the Secret containing authentication credentials for helm repository.\nFor HTTP/S basic auth the secret must contain 'username' and 'password' fields.",
													Type:        "object",
													Required: []string{
														"name",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"name": {
															Description: "Name of the object.",
															Type:        "string",
														},
													},
													Nullable: true,
												},
												"url": {
													Description: "URL of the source address, a valid URL contains at least a protocol and host.",
													Type:        "string",
												},
												"username": {
													Description: "Username for OCI registry.",
													Type:        "string",
												},
											},
											Nullable: true,
										},
										"ociRegistrySource": {
											Description: "OCIRegistrySource specifies the OCI registry configure.",
											Type:        "object",
											Properties: map[string]v1.JSONSchemaProps{
												"domain": {
													Description: "Domain for OCI registry.",
													Type:        "string",
												},
												"password": {
													Description: "Password for OCI registry.",
													Type:        "string",
												},
												"secretRef": {
													Description: "SecretRef specifies the Secret containing authentication credentials for helm repository.\nFor HTTP/S basic auth the secret must contain 'username' and 'password' fields.",
													Type:        "object",
													Required: []string{
														"name",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"name": {
															Description: "Name of the object.",
															Type:        "string",
														},
													},
													Nullable: true,
												},
												"username": {
													Description: "Username for OCI registry.",
													Type:        "string",
												},
											},
											Nullable: true,
										},
										"templateFormat": {
											Description: "TemplateFormat of the catalog.",
											Type:        "string",
											XValidations: []v1.ValidationRule{
												{
													Rule:    "oldSelf == self",
													Message: "immutable field",
												},
											},
										},
										"vcsSource": {
											Description: "VCSSource specifies the vcs source configure.",
											Type:        "object",
											Required: []string{
												"platform",
												"url",
											},
											Properties: map[string]v1.JSONSchemaProps{
												"platform": {
													Type: "string",
													Enum: []v1.JSON{
														{
															Raw: []byte(`"github"`),
														},
														{
															Raw: []byte(`"gitlab"`),
														},
														{
															Raw: []byte(`"gitee"`),
														},
													},
												},
												"secretRef": {
													Description: "SecretRef specifies the Secret containing authentication credentials.\nFor HTTP/S bear token the secret must contain 'token' field.",
													Type:        "object",
													Required: []string{
														"name",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"name": {
															Description: "Name of the object.",
															Type:        "string",
														},
													},
													Nullable: true,
												},
												"token": {
													Description: "Token for HTTP/S bear token.",
													Type:        "string",
												},
												"url": {
													Description: "URL of the source address, a valid URL contains at least a protocol and host.",
													Type:        "string",
												},
											},
											Nullable: true,
										},
									},
								},
								"status": {
									Type: "object",
									Required: []string{
										"lastSyncTime",
										"templateCount",
									},
									Properties: map[string]v1.JSONSchemaProps{
										"conditions": {
											Description: "Conditions holds the conditions for the catalog.",
											Type:        "array",
											Items: &v1.JSONSchemaPropsOrArray{
												Schema: &v1.JSONSchemaProps{
													Type: "object",
													Required: []string{
														"type",
														"lastTransitionTime",
														"reason",
														"message",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"lastTransitionTime": {
															Description: "lastTransitionTime is the last time the condition transitioned from one status to another.\nThis should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
															Type:        "string",
															Format:      "date-time",
														},
														"message": {
															Description: "message is a human readable message indicating details about the transition.\nThis may be an empty string.",
															Type:        "string",
														},
														"observedGeneration": {
															Description: "observedGeneration represents the .metadata.generation that the condition was set based upon.\nFor instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date\nwith respect to the current state of the instance.",
															Type:        "integer",
															Format:      "int64",
														},
														"reason": {
															Description: "reason contains a programmatic identifier indicating the reason for the condition's last transition.\nProducers of specific condition types may define expected values and meanings for this field,\nand whether the values are considered a guaranteed API.\nThe value should be a CamelCase string.\nThis field may not be empty.",
															Type:        "string",
														},
														"status": {
															Description: "status of the condition, one of True, False, Unknown.",
															Type:        "number",
														},
														"type": {
															Description: "type of condition in CamelCase or in foo.example.com/CamelCase.\n---\nMany .condition.type values are consistent across resources like Available, but because arbitrary conditions can be\nuseful (see .node.status.conditions), the ability to deconflict is important.\nThe regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
															Type:        "string",
														},
													},
												},
											},
											Nullable: true,
										},
										"lastSyncTime": {
											Description: "LastSyncTime record the last auto/manual sync catalog time.",
											Type:        "string",
											Format:      "date-time",
										},
										"templateCount": {
											Description: "TemplateCount include template count.",
											Type:        "integer",
											Format:      "int64",
										},
									},
								},
							},
						},
					},
					Subresources: &v1.CustomResourceSubresources{
						Status: &v1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
}

func crd_pkg_apis_walruscore_v1_Connector() *v1.CustomResourceDefinition {
	return &v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "connectors.walruscore.seal.io",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "walruscore.seal.io",
			Names: v1.CustomResourceDefinitionNames{
				Plural:   "connectors",
				Singular: "connector",
				Kind:     "Connector",
				ListKind: "ConnectorList",
			},
			Scope: "Namespaced",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Description: "Connector is the schema for the connectors API.",
							Type:        "object",
							Required: []string{
								"metadata",
								"spec",
							},
							Properties: map[string]v1.JSONSchemaProps{
								"apiVersion": {
									Type: "string",
								},
								"kind": {
									Type: "string",
								},
								"metadata": {
									Type: "object",
								},
								"spec": {
									Type: "object",
								},
								"status": {
									Type: "object",
								},
							},
						},
					},
					Subresources: &v1.CustomResourceSubresources{
						Status: &v1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
}

func crd_pkg_apis_walruscore_v1_Resource() *v1.CustomResourceDefinition {
	return &v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "resources.walruscore.seal.io",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "walruscore.seal.io",
			Names: v1.CustomResourceDefinitionNames{
				Plural:   "resources",
				Singular: "resource",
				Kind:     "Resource",
				ListKind: "ResourceList",
			},
			Scope: "Namespaced",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Description: "Resource is the schema for the resources API.",
							Type:        "object",
							Required: []string{
								"metadata",
								"spec",
							},
							Properties: map[string]v1.JSONSchemaProps{
								"apiVersion": {
									Type: "string",
								},
								"kind": {
									Type: "string",
								},
								"metadata": {
									Type: "object",
								},
								"spec": {
									Type: "object",
								},
								"status": {
									Type: "object",
								},
							},
						},
					},
					Subresources: &v1.CustomResourceSubresources{
						Status: &v1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
}

func crd_pkg_apis_walruscore_v1_ResourceDefinition() *v1.CustomResourceDefinition {
	return &v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "resourcedefinitions.walruscore.seal.io",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "walruscore.seal.io",
			Names: v1.CustomResourceDefinitionNames{
				Plural:   "resourcedefinitions",
				Singular: "resourcedefinition",
				Kind:     "ResourceDefinition",
				ListKind: "ResourceDefinitionList",
			},
			Scope: "Namespaced",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Description: "ResourceDefinition is the schema for the resource definitions API.",
							Type:        "object",
							Required: []string{
								"metadata",
								"spec",
							},
							Properties: map[string]v1.JSONSchemaProps{
								"apiVersion": {
									Type: "string",
								},
								"kind": {
									Type: "string",
								},
								"metadata": {
									Type: "object",
								},
								"spec": {
									Type: "object",
								},
								"status": {
									Type: "object",
								},
							},
						},
					},
					Subresources: &v1.CustomResourceSubresources{
						Status: &v1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
}

func crd_pkg_apis_walruscore_v1_ResourceRun() *v1.CustomResourceDefinition {
	return &v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "resourceruns.walruscore.seal.io",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "walruscore.seal.io",
			Names: v1.CustomResourceDefinitionNames{
				Plural:   "resourceruns",
				Singular: "resourcerun",
				Kind:     "ResourceRun",
				ListKind: "ResourceRunList",
			},
			Scope: "Namespaced",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Description: "ResourceRun is the schema for the resource runs API.",
							Type:        "object",
							Required: []string{
								"metadata",
								"spec",
							},
							Properties: map[string]v1.JSONSchemaProps{
								"apiVersion": {
									Type: "string",
								},
								"kind": {
									Type: "string",
								},
								"metadata": {
									Type: "object",
								},
								"spec": {
									Type: "object",
								},
								"status": {
									Type: "object",
								},
							},
						},
					},
					Subresources: &v1.CustomResourceSubresources{
						Status: &v1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
}

func crd_pkg_apis_walruscore_v1_Template() *v1.CustomResourceDefinition {
	return &v1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "templates.walruscore.seal.io",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "walruscore.seal.io",
			Names: v1.CustomResourceDefinitionNames{
				Plural:   "templates",
				Singular: "template",
				Kind:     "Template",
				ListKind: "TemplateList",
			},
			Scope: "Namespaced",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Description: "Template is the schema for the templates API.",
							Type:        "object",
							Required: []string{
								"metadata",
								"spec",
							},
							Properties: map[string]v1.JSONSchemaProps{
								"apiVersion": {
									Type: "string",
								},
								"kind": {
									Type: "string",
								},
								"metadata": {
									Type: "object",
								},
								"spec": {
									Type: "object",
									Required: []string{
										"templateFormat",
									},
									Properties: map[string]v1.JSONSchemaProps{
										"description": {
											Description: "Description of the template.",
											Type:        "string",
										},
										"helmChart": {
											Description: "HelmChart specifies the tgz format helm chart configure.",
											Type:        "object",
											Required: []string{
												"url",
											},
											Properties: map[string]v1.JSONSchemaProps{
												"password": {
													Description: "Password for OCI registry.",
													Type:        "string",
												},
												"secretRef": {
													Description: "SecretRef specifies the Secret containing authentication credentials for helm repository.\nFor HTTP/S basic auth the secret must contain 'username' and 'password' fields.",
													Type:        "object",
													Required: []string{
														"name",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"name": {
															Description: "Name of the object.",
															Type:        "string",
														},
													},
													Nullable: true,
												},
												"url": {
													Description: "URL of download the template from chart tgz address, e.g. https://charts.bitnami.com/bitnami/phpbb-16.5.0.tgz.",
													Type:        "string",
												},
												"username": {
													Description: "Username for OCI registry.",
													Type:        "string",
												},
											},
											Nullable: true,
										},
										"helmOCIChart": {
											Description: "HelmOCIChart specifies the OCI format helm chart configure.",
											Type:        "object",
											Required: []string{
												"url",
											},
											Properties: map[string]v1.JSONSchemaProps{
												"password": {
													Description: "Password for OCI registry.",
													Type:        "string",
												},
												"secretRef": {
													Description: "SecretRef specifies the Secret containing authentication credentials for helm repository.\nFor HTTP/S basic auth the secret must contain 'username' and 'password' fields.",
													Type:        "object",
													Required: []string{
														"name",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"name": {
															Description: "Name of the object.",
															Type:        "string",
														},
													},
													Nullable: true,
												},
												"url": {
													Description: "URL of download the template from oci image with ref, e.g. oci://registry-1.docker.io/bitnamicharts/mysql?ref=10.1.0.",
													Type:        "string",
												},
												"username": {
													Description: "Username for OCI registry.",
													Type:        "string",
												},
											},
											Nullable: true,
										},
										"templateFormat": {
											Description: "TemplateFormat of the content.",
											Type:        "string",
										},
										"vcsRepository": {
											Description: "VCSRepository specifies the vcs repository configure.",
											Type:        "object",
											Required: []string{
												"url",
											},
											Properties: map[string]v1.JSONSchemaProps{
												"secretRef": {
													Description: "SecretRef specifies the Secret containing authentication credentials.\nFor HTTP/S bear token the secret must contain 'token' field.",
													Type:        "object",
													Required: []string{
														"name",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"name": {
															Description: "Name of the object.",
															Type:        "string",
														},
													},
													Nullable: true,
												},
												"token": {
													Description: "Token for HTTP/S bear token.",
													Type:        "string",
												},
												"url": {
													Description: "URL of download the template from vsc repository with ref and subpath, e.g. https://github.com/walrus-catalog/terraform-static-redis?ref=main.",
													Type:        "string",
												},
											},
											Nullable: true,
										},
									},
								},
								"status": {
									Type: "object",
									Required: []string{
										"originalName",
										"icon",
										"versions",
										"conditions",
									},
									Properties: map[string]v1.JSONSchemaProps{
										"conditions": {
											Description: "Conditions holds the conditions for the template.",
											Type:        "array",
											Items: &v1.JSONSchemaPropsOrArray{
												Schema: &v1.JSONSchemaProps{
													Type: "object",
													Required: []string{
														"type",
														"lastTransitionTime",
														"reason",
														"message",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"lastTransitionTime": {
															Description: "lastTransitionTime is the last time the condition transitioned from one status to another.\nThis should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
															Type:        "string",
															Format:      "date-time",
														},
														"message": {
															Description: "message is a human readable message indicating details about the transition.\nThis may be an empty string.",
															Type:        "string",
														},
														"observedGeneration": {
															Description: "observedGeneration represents the .metadata.generation that the condition was set based upon.\nFor instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date\nwith respect to the current state of the instance.",
															Type:        "integer",
															Format:      "int64",
														},
														"reason": {
															Description: "reason contains a programmatic identifier indicating the reason for the condition's last transition.\nProducers of specific condition types may define expected values and meanings for this field,\nand whether the values are considered a guaranteed API.\nThe value should be a CamelCase string.\nThis field may not be empty.",
															Type:        "string",
														},
														"status": {
															Description: "status of the condition, one of True, False, Unknown.",
															Type:        "number",
														},
														"type": {
															Description: "type of condition in CamelCase or in foo.example.com/CamelCase.\n---\nMany .condition.type values are consistent across resources like Available, but because arbitrary conditions can be\nuseful (see .node.status.conditions), the ability to deconflict is important.\nThe regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
															Type:        "string",
														},
													},
												},
											},
											Nullable: true,
										},
										"icon": {
											Description: "A URL to an SVG or PNG image to be used as an icon.",
											Type:        "string",
										},
										"originalName": {
											Description: "The original name of the template, name generate from chart.yaml, terraform git repo name.",
											Type:        "string",
										},
										"versions": {
											Description: "Versions holds the versions for the template.",
											Type:        "array",
											Items: &v1.JSONSchemaPropsOrArray{
												Schema: &v1.JSONSchemaProps{
													Type: "object",
													Required: []string{
														"version",
														"url",
														"schemaRef",
														"originalUISchemaRef",
														"uiSchemaRef",
													},
													Properties: map[string]v1.JSONSchemaProps{
														"originalUISchemaRef": {
															Description: "OriginalUISchema holds the original UI schema for the template version.",
															Type:        "object",
															Required: []string{
																"name",
															},
															Properties: map[string]v1.JSONSchemaProps{
																"name": {
																	Description: "Name of the object.",
																	Type:        "string",
																},
															},
															Nullable: true,
														},
														"schemaRef": {
															Description: "Schema holds the schema for the template version.",
															Type:        "object",
															Required: []string{
																"name",
															},
															Properties: map[string]v1.JSONSchemaProps{
																"name": {
																	Description: "Name of the object.",
																	Type:        "string",
																},
															},
															Nullable: true,
														},
														"uiSchemaRef": {
															Description: "UISchema holds the UI schema for the template version.",
															Type:        "object",
															Required: []string{
																"name",
															},
															Properties: map[string]v1.JSONSchemaProps{
																"name": {
																	Description: "Name of the object.",
																	Type:        "string",
																},
															},
															Nullable: true,
														},
														"url": {
															Description: "URL of downloading the version.\nhost.",
															Type:        "string",
														},
														"version": {
															Type: "string",
														},
													},
												},
											},
											Nullable: true,
										},
									},
								},
							},
						},
					},
					Subresources: &v1.CustomResourceSubresources{
						Status: &v1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
}
