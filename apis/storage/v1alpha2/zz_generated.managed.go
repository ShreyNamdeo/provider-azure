/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import (
	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// GetBindingPhase of this Account.
func (mg *Account) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this Account.
func (mg *Account) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetNonPortableClassReference of this Account.
func (mg *Account) GetNonPortableClassReference() *corev1.ObjectReference {
	return mg.Spec.NonPortableClassReference
}

// GetReclaimPolicy of this Account.
func (mg *Account) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this Account.
func (mg *Account) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this Account.
func (mg *Account) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this Account.
func (mg *Account) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetConditions of this Account.
func (mg *Account) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetNonPortableClassReference of this Account.
func (mg *Account) SetNonPortableClassReference(r *corev1.ObjectReference) {
	mg.Spec.NonPortableClassReference = r
}

// SetReclaimPolicy of this Account.
func (mg *Account) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this Account.
func (mg *Account) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
