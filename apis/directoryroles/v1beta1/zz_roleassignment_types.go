// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RoleAssignmentInitParameters struct {

	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with directory_scope_id. See official documentation for example usage. Changing this forces a new resource to be created.
	// Identifier of the app-specific scope when the assignment scope is app-specific
	AppScopeID *string `json:"appScopeId,omitempty" tf:"app_scope_id,omitempty"`

	// Identifier of the app-specific scope when the assignment scope is app-specific
	AppScopeObjectID *string `json:"appScopeObjectId,omitempty" tf:"app_scope_object_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment. Cannot be used with app_scope_id. See official documentation for example usage. Changing this forces a new resource to be created.
	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeID *string `json:"directoryScopeId,omitempty" tf:"directory_scope_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeObjectID *string `json:"directoryScopeObjectId,omitempty" tf:"directory_scope_object_id,omitempty"`

	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	// The object ID of the member principal
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/users/v1beta1.User
	PrincipalObjectID *string `json:"principalObjectId,omitempty" tf:"principal_object_id,omitempty"`

	// Reference to a User in users to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDRef *v1.Reference `json:"principalObjectIdRef,omitempty" tf:"-"`

	// Selector for a User in users to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDSelector *v1.Selector `json:"principalObjectIdSelector,omitempty" tf:"-"`

	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	// The object ID of the directory role for this assignment
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/directoryroles/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("template_id",true)
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Reference to a Role in directoryroles to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDRef *v1.Reference `json:"roleIdRef,omitempty" tf:"-"`

	// Selector for a Role in directoryroles to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`
}

type RoleAssignmentObservation struct {

	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with directory_scope_id. See official documentation for example usage. Changing this forces a new resource to be created.
	// Identifier of the app-specific scope when the assignment scope is app-specific
	AppScopeID *string `json:"appScopeId,omitempty" tf:"app_scope_id,omitempty"`

	// Identifier of the app-specific scope when the assignment scope is app-specific
	AppScopeObjectID *string `json:"appScopeObjectId,omitempty" tf:"app_scope_object_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment. Cannot be used with app_scope_id. See official documentation for example usage. Changing this forces a new resource to be created.
	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeID *string `json:"directoryScopeId,omitempty" tf:"directory_scope_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeObjectID *string `json:"directoryScopeObjectId,omitempty" tf:"directory_scope_object_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	// The object ID of the member principal
	PrincipalObjectID *string `json:"principalObjectId,omitempty" tf:"principal_object_id,omitempty"`

	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	// The object ID of the directory role for this assignment
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`
}

type RoleAssignmentParameters struct {

	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with directory_scope_id. See official documentation for example usage. Changing this forces a new resource to be created.
	// Identifier of the app-specific scope when the assignment scope is app-specific
	// +kubebuilder:validation:Optional
	AppScopeID *string `json:"appScopeId,omitempty" tf:"app_scope_id,omitempty"`

	// Identifier of the app-specific scope when the assignment scope is app-specific
	// +kubebuilder:validation:Optional
	AppScopeObjectID *string `json:"appScopeObjectId,omitempty" tf:"app_scope_object_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment. Cannot be used with app_scope_id. See official documentation for example usage. Changing this forces a new resource to be created.
	// Identifier of the directory object representing the scope of the assignment
	// +kubebuilder:validation:Optional
	DirectoryScopeID *string `json:"directoryScopeId,omitempty" tf:"directory_scope_id,omitempty"`

	// Identifier of the directory object representing the scope of the assignment
	// +kubebuilder:validation:Optional
	DirectoryScopeObjectID *string `json:"directoryScopeObjectId,omitempty" tf:"directory_scope_object_id,omitempty"`

	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	// The object ID of the member principal
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/users/v1beta1.User
	// +kubebuilder:validation:Optional
	PrincipalObjectID *string `json:"principalObjectId,omitempty" tf:"principal_object_id,omitempty"`

	// Reference to a User in users to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDRef *v1.Reference `json:"principalObjectIdRef,omitempty" tf:"-"`

	// Selector for a User in users to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDSelector *v1.Selector `json:"principalObjectIdSelector,omitempty" tf:"-"`

	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	// The object ID of the directory role for this assignment
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/directoryroles/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("template_id",true)
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Reference to a Role in directoryroles to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDRef *v1.Reference `json:"roleIdRef,omitempty" tf:"-"`

	// Selector for a Role in directoryroles to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`
}

// RoleAssignmentSpec defines the desired state of RoleAssignment
type RoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAssignmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RoleAssignmentInitParameters `json:"initProvider,omitempty"`
}

// RoleAssignmentStatus defines the observed state of RoleAssignment.
type RoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RoleAssignment is the Schema for the RoleAssignments API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignmentSpec   `json:"spec"`
	Status            RoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentList contains a list of RoleAssignments
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	RoleAssignment_Kind             = "RoleAssignment"
	RoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAssignment_Kind}.String()
	RoleAssignment_KindAPIVersion   = RoleAssignment_Kind + "." + CRDGroupVersion.String()
	RoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(RoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
