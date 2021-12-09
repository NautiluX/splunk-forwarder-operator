// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarder":       schema_pkg_apis_splunkforwarder_v1alpha1_SplunkForwarder(ref),
		"github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderSpec":   schema_pkg_apis_splunkforwarder_v1alpha1_SplunkForwarderSpec(ref),
		"github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderStatus": schema_pkg_apis_splunkforwarder_v1alpha1_SplunkForwarderStatus(ref),
	}
}

func schema_pkg_apis_splunkforwarder_v1alpha1_SplunkForwarder(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SplunkForwarder is the Schema for the splunkforwarders API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
							Ref: ref("github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderSpec", "github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_splunkforwarder_v1alpha1_SplunkForwarderSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SplunkForwarderSpec defines the desired state of SplunkForwarder",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"splunkLicenseAccepted": {
						SchemaProps: spec.SchemaProps{
							Description: "Adds an --accept-license flag to automatically accept the Splunk License Agreement. Must be true for the Red Hat provided Splunk Forwarder image. Optional: Defaults to false.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Container image path to the Splunk Forwarder",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imageTag": {
						SchemaProps: spec.SchemaProps{
							Description: "The container image tag of the Splunk Forwarder image. Is not used if ImageDigest is supplied. Optional: Defaults to latest",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imageDigest": {
						SchemaProps: spec.SchemaProps{
							Description: "Container image digest of the Splunk Forwarder image. Has precedence and is recommended over ImageTag. Optional: Defaults to latest",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clusterID": {
						SchemaProps: spec.SchemaProps{
							Description: "Unique cluster name. Optional: Looked up on the cluster if not provided, default to openshift",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"splunkInputs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderInputs"),
									},
								},
							},
						},
					},
					"useHeavyForwarder": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether an additional Splunk Heavy Forwarder should be deployed. Optional: Defaults to false.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"heavyForwarderImage": {
						SchemaProps: spec.SchemaProps{
							Description: "Container image path to the Splunk Heavy Forwarder image. Required when UseHeavyForwarder is true.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"heavyForwarderDigest": {
						SchemaProps: spec.SchemaProps{
							Description: "Container image digest of the container image defined in HeavyForwarderImage. Optional: Defaults to latest",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"heavyForwarderReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of desired Splunk Heavy Forwarder pods. Optional: Defaults to 2",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"heavyForwarderSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the value of the NodeSelector for the Splunk Heavy Forwarder pods with key: \"node-role.kubernetes.io\" Optional: Defaults to an empty value.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"filters": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "List of additional filters supplied to configure the Splunk Heavy Forwarder Optional: Defaults to no additional filters (no transforms.conf).",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkFilter"),
									},
								},
							},
						},
					},
				},
				Required: []string{"image", "splunkInputs"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkFilter", "github.com/openshift/splunk-forwarder-operator/pkg/apis/splunkforwarder/v1alpha1.SplunkForwarderInputs"},
	}
}

func schema_pkg_apis_splunkforwarder_v1alpha1_SplunkForwarderStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SplunkForwarderStatus defines the observed state of SplunkForwarder",
				Type:        []string{"object"},
			},
		},
	}
}
