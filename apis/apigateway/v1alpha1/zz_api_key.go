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

// APIKeyParameters defines the desired state of APIKey
type APIKeyParameters struct {
	// Region is which region the APIKey will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerID *string `json:"customerID,omitempty"`
	// The description of the ApiKey.
	Description *string `json:"description,omitempty"`
	// Specifies whether the ApiKey can be used by callers.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies whether (true) or not (false) the key identifier is distinct from
	// the created API key value. This parameter is deprecated and should not be
	// used.
	GenerateDistinctID *bool `json:"generateDistinctID,omitempty"`
	// The name of the ApiKey.
	Name *string `json:"name,omitempty"`
	// DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.
	StageKeys []*StageKey `json:"stageKeys,omitempty"`
	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]*string `json:"tags,omitempty"`
	// Specifies a value of the API key.
	Value                  *string `json:"value,omitempty"`
	CustomAPIKeyParameters `json:",inline"`
}

// APIKeySpec defines the desired state of APIKey
type APIKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       APIKeyParameters `json:"forProvider"`
}

// APIKeyObservation defines the observed state of APIKey
type APIKeyObservation struct {
	// The timestamp when the API Key was created.
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// The identifier of the API Key.
	ID *string `json:"id,omitempty"`
	// The timestamp when the API Key was last updated.
	LastUpdatedDate *metav1.Time `json:"lastUpdatedDate,omitempty"`
}

// APIKeyStatus defines the observed state of APIKey.
type APIKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          APIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIKey is the Schema for the APIKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type APIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIKeySpec   `json:"spec"`
	Status            APIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIKeyList contains a list of APIKeys
type APIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIKey `json:"items"`
}

// Repository type metadata.
var (
	APIKeyKind             = "APIKey"
	APIKeyGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIKeyKind}.String()
	APIKeyKindAPIVersion   = APIKeyKind + "." + GroupVersion.String()
	APIKeyGroupVersionKind = GroupVersion.WithKind(APIKeyKind)
)

func init() {
	SchemeBuilder.Register(&APIKey{}, &APIKeyList{})
}
