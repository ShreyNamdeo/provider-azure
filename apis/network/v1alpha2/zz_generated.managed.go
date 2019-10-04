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

// GetBindingPhase of this Subnet.
func (mg *Subnet) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this Subnet.
func (mg *Subnet) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetNonPortableClassReference of this Subnet.
func (mg *Subnet) GetNonPortableClassReference() *corev1.ObjectReference {
	return mg.Spec.NonPortableClassReference
}

// GetReclaimPolicy of this Subnet.
func (mg *Subnet) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this Subnet.
func (mg *Subnet) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this Subnet.
func (mg *Subnet) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this Subnet.
func (mg *Subnet) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetConditions of this Subnet.
func (mg *Subnet) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetNonPortableClassReference of this Subnet.
func (mg *Subnet) SetNonPortableClassReference(r *corev1.ObjectReference) {
	mg.Spec.NonPortableClassReference = r
}

// SetReclaimPolicy of this Subnet.
func (mg *Subnet) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this Subnet.
func (mg *Subnet) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this VirtualNetwork.
func (mg *VirtualNetwork) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this VirtualNetwork.
func (mg *VirtualNetwork) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetNonPortableClassReference of this VirtualNetwork.
func (mg *VirtualNetwork) GetNonPortableClassReference() *corev1.ObjectReference {
	return mg.Spec.NonPortableClassReference
}

// GetReclaimPolicy of this VirtualNetwork.
func (mg *VirtualNetwork) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this VirtualNetwork.
func (mg *VirtualNetwork) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this VirtualNetwork.
func (mg *VirtualNetwork) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this VirtualNetwork.
func (mg *VirtualNetwork) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetConditions of this VirtualNetwork.
func (mg *VirtualNetwork) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetNonPortableClassReference of this VirtualNetwork.
func (mg *VirtualNetwork) SetNonPortableClassReference(r *corev1.ObjectReference) {
	mg.Spec.NonPortableClassReference = r
}

// SetReclaimPolicy of this VirtualNetwork.
func (mg *VirtualNetwork) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this VirtualNetwork.
func (mg *VirtualNetwork) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
