//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package walruscore

import (
	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func GetWebhookConfigurations(np string, c v1.WebhookClientConfig) (*v1.ValidatingWebhookConfiguration, *v1.MutatingWebhookConfiguration) {
	vwc := GetValidatingWebhookConfiguration(np+"-validation", c)
	mwc := GetMutatingWebhookConfiguration(np+"-mutation", c)
	return vwc, mwc
}

func GetValidatingWebhookConfiguration(n string, c v1.WebhookClientConfig) *v1.ValidatingWebhookConfiguration {
	return &v1.ValidatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admissionregistration.k8s.io/v1",
			Kind:       "ValidatingWebhookConfiguration",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: n,
		},
		Webhooks: []v1.ValidatingWebhook{
			vwh_walrus_pkg_webhooks_walruscore_CatalogWebhook(c),
			vwh_walrus_pkg_webhooks_walruscore_ConnectorWebhook(c),
			vwh_walrus_pkg_webhooks_walruscore_TemplateWebhook(c),
		},
	}
}

func GetMutatingWebhookConfiguration(n string, c v1.WebhookClientConfig) *v1.MutatingWebhookConfiguration {
	return &v1.MutatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admissionregistration.k8s.io/v1",
			Kind:       "MutatingWebhookConfiguration",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: n,
		},
		Webhooks: []v1.MutatingWebhook{
			mwh_walrus_pkg_webhooks_walruscore_ResourceDefinitionWebhook(c),
		},
	}
}

func (*CatalogWebhook) ValidatePath() string {
	return "/validate-walruscore-seal-io-v1-catalog"
}

func vwh_walrus_pkg_webhooks_walruscore_CatalogWebhook(c v1.WebhookClientConfig) v1.ValidatingWebhook {
	path := "/validate-walruscore-seal-io-v1-catalog"

	cc := c.DeepCopy()
	if cc.Service != nil {
		cc.Service.Path = &path
	} else if c.URL != nil {
		cc.URL = ptr.To(*c.URL + path)
	}

	return v1.ValidatingWebhook{
		Name:         "validate.walruscore.seal.io.v1.catalog",
		ClientConfig: *cc,
		Rules: []v1.RuleWithOperations{
			{
				Rule: v1.Rule{
					APIGroups: []string{
						"walruscore.seal.io",
					},
					APIVersions: []string{
						"v1",
					},
					Resources: []string{
						"catalogs",
					},
					Scope: ptr.To[v1.ScopeType]("Namespaced"),
				},
				Operations: []v1.OperationType{
					"CREATE",
					"UPDATE",
					"DELETE",
				},
			},
		},
		FailurePolicy:  ptr.To[v1.FailurePolicyType]("Fail"),
		MatchPolicy:    ptr.To[v1.MatchPolicyType]("Equivalent"),
		SideEffects:    ptr.To[v1.SideEffectClass]("None"),
		TimeoutSeconds: ptr.To[int32](10),
		AdmissionReviewVersions: []string{
			"v1",
		},
	}
}

func (*ConnectorWebhook) ValidatePath() string {
	return "/validate-walruscore-seal-io-v1-connector"
}

func vwh_walrus_pkg_webhooks_walruscore_ConnectorWebhook(c v1.WebhookClientConfig) v1.ValidatingWebhook {
	path := "/validate-walruscore-seal-io-v1-connector"

	cc := c.DeepCopy()
	if cc.Service != nil {
		cc.Service.Path = &path
	} else if c.URL != nil {
		cc.URL = ptr.To(*c.URL + path)
	}

	return v1.ValidatingWebhook{
		Name:         "validate.walruscore.seal.io.v1.connector",
		ClientConfig: *cc,
		Rules: []v1.RuleWithOperations{
			{
				Rule: v1.Rule{
					APIGroups: []string{
						"walruscore.seal.io",
					},
					APIVersions: []string{
						"v1",
					},
					Resources: []string{
						"connectors",
					},
					Scope: ptr.To[v1.ScopeType]("Namespaced"),
				},
				Operations: []v1.OperationType{
					"CREATE",
					"UPDATE",
					"DELETE",
				},
			},
		},
		FailurePolicy:  ptr.To[v1.FailurePolicyType]("Fail"),
		MatchPolicy:    ptr.To[v1.MatchPolicyType]("Equivalent"),
		SideEffects:    ptr.To[v1.SideEffectClass]("None"),
		TimeoutSeconds: ptr.To[int32](10),
		AdmissionReviewVersions: []string{
			"v1",
		},
	}
}

func (*ResourceDefinitionWebhook) DefaultPath() string {
	return "/mutate-walruscore-seal-io-v1-resourcedefinition"
}

func mwh_walrus_pkg_webhooks_walruscore_ResourceDefinitionWebhook(c v1.WebhookClientConfig) v1.MutatingWebhook {
	path := "/mutate-walruscore-seal-io-v1-resourcedefinition"

	cc := c.DeepCopy()
	if cc.Service != nil {
		cc.Service.Path = &path
	} else if c.URL != nil {
		cc.URL = ptr.To(*c.URL + path)
	}

	return v1.MutatingWebhook{
		Name:         "mutate.walruscore.seal.io.v1.resourcedefinition",
		ClientConfig: *cc,
		Rules: []v1.RuleWithOperations{
			{
				Rule: v1.Rule{
					APIGroups: []string{
						"walruscore.seal.io",
					},
					APIVersions: []string{
						"v1",
					},
					Resources: []string{
						"resourcedefinitions",
					},
					Scope: ptr.To[v1.ScopeType]("Namespaced"),
				},
				Operations: []v1.OperationType{
					"CREATE",
					"UPDATE",
				},
			},
		},
		FailurePolicy:  ptr.To[v1.FailurePolicyType]("Fail"),
		MatchPolicy:    ptr.To[v1.MatchPolicyType]("Equivalent"),
		SideEffects:    ptr.To[v1.SideEffectClass]("None"),
		TimeoutSeconds: ptr.To[int32](10),
		AdmissionReviewVersions: []string{
			"v1",
		},
	}
}

func (*TemplateWebhook) ValidatePath() string {
	return "/validate-walruscore-seal-io-v1-template"
}

func vwh_walrus_pkg_webhooks_walruscore_TemplateWebhook(c v1.WebhookClientConfig) v1.ValidatingWebhook {
	path := "/validate-walruscore-seal-io-v1-template"

	cc := c.DeepCopy()
	if cc.Service != nil {
		cc.Service.Path = &path
	} else if c.URL != nil {
		cc.URL = ptr.To(*c.URL + path)
	}

	return v1.ValidatingWebhook{
		Name:         "validate.walruscore.seal.io.v1.template",
		ClientConfig: *cc,
		Rules: []v1.RuleWithOperations{
			{
				Rule: v1.Rule{
					APIGroups: []string{
						"walruscore.seal.io",
					},
					APIVersions: []string{
						"v1",
					},
					Resources: []string{
						"templates",
					},
					Scope: ptr.To[v1.ScopeType]("Namespaced"),
				},
				Operations: []v1.OperationType{
					"CREATE",
					"DELETE",
				},
			},
		},
		FailurePolicy:  ptr.To[v1.FailurePolicyType]("Fail"),
		MatchPolicy:    ptr.To[v1.MatchPolicyType]("Equivalent"),
		SideEffects:    ptr.To[v1.SideEffectClass]("None"),
		TimeoutSeconds: ptr.To[int32](10),
		AdmissionReviewVersions: []string{
			"v1",
		},
	}
}
