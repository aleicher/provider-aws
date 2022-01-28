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

// TransitGatewayRouteTableParameters defines the desired state of TransitGatewayRouteTable
type TransitGatewayRouteTableParameters struct {
	// Region is which region the TransitGatewayRouteTable will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The tags to apply to the transit gateway route table.
	TagSpecifications                        []*TagSpecification `json:"tagSpecifications,omitempty"`
	CustomTransitGatewayRouteTableParameters `json:",inline"`
}

// TransitGatewayRouteTableSpec defines the desired state of TransitGatewayRouteTable
type TransitGatewayRouteTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TransitGatewayRouteTableParameters `json:"forProvider"`
}

// TransitGatewayRouteTableObservation defines the observed state of TransitGatewayRouteTable
type TransitGatewayRouteTableObservation struct {
	// The creation time.
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// Indicates whether this is the default association route table for the transit
	// gateway.
	DefaultAssociationRouteTable *bool `json:"defaultAssociationRouteTable,omitempty"`
	// Indicates whether this is the default propagation route table for the transit
	// gateway.
	DefaultPropagationRouteTable *bool `json:"defaultPropagationRouteTable,omitempty"`
	// The state of the transit gateway route table.
	State *string `json:"state,omitempty"`
	// Any tags assigned to the route table.
	Tags []*Tag `json:"tags,omitempty"`
	// The ID of the transit gateway.
	TransitGatewayID *string `json:"transitGatewayID,omitempty"`
	// The ID of the transit gateway route table.
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableID,omitempty"`
}

// TransitGatewayRouteTableStatus defines the observed state of TransitGatewayRouteTable.
type TransitGatewayRouteTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TransitGatewayRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTable is the Schema for the TransitGatewayRouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteTableSpec   `json:"spec"`
	Status            TransitGatewayRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTableList contains a list of TransitGatewayRouteTables
type TransitGatewayRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRouteTable `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRouteTableKind             = "TransitGatewayRouteTable"
	TransitGatewayRouteTableGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRouteTableKind}.String()
	TransitGatewayRouteTableKindAPIVersion   = TransitGatewayRouteTableKind + "." + GroupVersion.String()
	TransitGatewayRouteTableGroupVersionKind = GroupVersion.WithKind(TransitGatewayRouteTableKind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRouteTable{}, &TransitGatewayRouteTableList{})
}