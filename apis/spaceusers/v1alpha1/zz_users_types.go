/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UsersObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UsersParameters struct {

	// +kubebuilder:validation:Optional
	Auditors []*string `json:"auditors,omitempty" tf:"auditors,omitempty"`

	// +kubebuilder:validation:Optional
	Developers []*string `json:"developers,omitempty" tf:"developers,omitempty"`

	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// +kubebuilder:validation:Optional
	Managers []*string `json:"managers,omitempty" tf:"managers,omitempty"`

	// +kubebuilder:validation:Required
	Space *string `json:"space" tf:"space,omitempty"`
}

// UsersSpec defines the desired state of Users
type UsersSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsersParameters `json:"forProvider"`
}

// UsersStatus defines the observed state of Users.
type UsersStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Users is the Schema for the Userss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type Users struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UsersSpec   `json:"spec"`
	Status            UsersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsersList contains a list of Userss
type UsersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Users `json:"items"`
}

// Repository type metadata.
var (
	Users_Kind             = "Users"
	Users_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Users_Kind}.String()
	Users_KindAPIVersion   = Users_Kind + "." + CRDGroupVersion.String()
	Users_GroupVersionKind = CRDGroupVersion.WithKind(Users_Kind)
)

func init() {
	SchemeBuilder.Register(&Users{}, &UsersList{})
}
