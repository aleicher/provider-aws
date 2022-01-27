/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// StageParameters defines the desired state of Stage
type StageParameters struct {
	// Region is which region the Stage will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Whether cache clustering is enabled for the stage.
	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty"`
	// The stage's cache cluster size.
	CacheClusterSize *string `json:"cacheClusterSize,omitempty"`
	// The canary deployment settings of this stage.
	CanarySettings *CanarySettings `json:"canarySettings,omitempty"`
	// [Required] The identifier of the Deployment resource for the Stage resource.
	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentID"`
	// The description of the Stage resource.
	Description *string `json:"description,omitempty"`
	// The version of the associated API documentation.
	DocumentationVersion *string `json:"documentationVersion,omitempty"`
	// [Required] The string identifier of the associated RestApi.
	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restAPIID"`
	// [Required] The name for the Stage resource. Stage names can only contain
	// alphanumeric characters, hyphens, and underscores. Maximum length is 128
	// characters.
	// +kubebuilder:validation:Required
	StageName *string `json:"stageName"`
	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]*string `json:"tags,omitempty"`
	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool `json:"tracingEnabled,omitempty"`
	// A map that defines the stage variables for the new Stage resource. Variable
	// names can have alphanumeric and underscore characters, and the values must
	// match [A-Za-z0-9-._~:/?#&=,]+.
	Variables             map[string]*string `json:"variables,omitempty"`
	CustomStageParameters `json:",inline"`
}

// StageSpec defines the desired state of Stage
type StageSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StageParameters `json:"forProvider"`
}

// StageObservation defines the observed state of Stage
type StageObservation struct {
	// Settings for logging access in this stage.
	AccessLogSettings *AccessLogSettings `json:"accessLogSettings,omitempty"`
	// The status of the cache cluster for the stage, if enabled.
	CacheClusterStatus *string `json:"cacheClusterStatus,omitempty"`
	// The identifier of a client certificate for an API stage.
	ClientCertificateID *string `json:"clientCertificateID,omitempty"`
	// The timestamp when the stage was created.
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// The timestamp when the stage last updated.
	LastUpdatedDate *metav1.Time `json:"lastUpdatedDate,omitempty"`
	// A map that defines the method settings for a Stage resource. Keys (designated
	// as /{method_setting_key below) are method paths defined as {resource_path}/{http_method}
	// for an individual method override, or /\*/\* for overriding all methods in
	// the stage.
	MethodSettings map[string]*MethodSetting `json:"methodSettings,omitempty"`
	// The ARN of the WebAcl associated with the Stage.
	WebACLARN *string `json:"webACLARN,omitempty"`
}

// StageStatus defines the observed state of Stage.
type StageStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stage is the Schema for the Stages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StageSpec   `json:"spec"`
	Status            StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StageList contains a list of Stages
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stage `json:"items"`
}

// Repository type metadata.
var (
	StageKind             = "Stage"
	StageGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StageKind}.String()
	StageKindAPIVersion   = StageKind + "." + GroupVersion.String()
	StageGroupVersionKind = GroupVersion.WithKind(StageKind)
)

func init() {
	SchemeBuilder.Register(&Stage{}, &StageList{})
}
