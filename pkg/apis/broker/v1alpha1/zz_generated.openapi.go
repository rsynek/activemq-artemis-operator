// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/broker/v1alpha1.ActiveMQArtemisSecurity":       schema_pkg_apis_broker_v1alpha1_ActiveMQArtemisSecurity(ref),
		"./pkg/apis/broker/v1alpha1.ActiveMQArtemisSecuritySpec":   schema_pkg_apis_broker_v1alpha1_ActiveMQArtemisSecuritySpec(ref),
		"./pkg/apis/broker/v1alpha1.ActiveMQArtemisSecurityStatus": schema_pkg_apis_broker_v1alpha1_ActiveMQArtemisSecurityStatus(ref),
	}
}

func schema_pkg_apis_broker_v1alpha1_ActiveMQArtemisSecurity(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ActiveMQArtemisSecurity is the Schema for the activemqartemissecurities API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/broker/v1alpha1.ActiveMQArtemisSecuritySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/broker/v1alpha1.ActiveMQArtemisSecurityStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/broker/v1alpha1.ActiveMQArtemisSecuritySpec", "./pkg/apis/broker/v1alpha1.ActiveMQArtemisSecurityStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_broker_v1alpha1_ActiveMQArtemisSecuritySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ActiveMQArtemisSecuritySpec defines the desired state of ActiveMQArtemisSecurity",
				Properties: map[string]spec.Schema{
					"domain": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("./pkg/apis/broker/v1alpha1.DomainType"),
						},
					},
					"certDomain": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/broker/v1alpha1.CertDomainType"),
						},
					},
				},
				Required: []string{"domain", "certDomain"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/broker/v1alpha1.CertDomainType", "./pkg/apis/broker/v1alpha1.DomainType"},
	}
}

func schema_pkg_apis_broker_v1alpha1_ActiveMQArtemisSecurityStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ActiveMQArtemisSecurityStatus defines the observed state of ActiveMQArtemisSecurity",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}