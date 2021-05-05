// +build !ignore_autogenerated

/*
Copyright Red Hat, Inc.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	operators "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*OperatorGroup)(nil), (*operators.OperatorGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OperatorGroup_To_operators_OperatorGroup(a.(*OperatorGroup), b.(*operators.OperatorGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.OperatorGroup)(nil), (*OperatorGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_OperatorGroup_To_v1_OperatorGroup(a.(*operators.OperatorGroup), b.(*OperatorGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OperatorGroupList)(nil), (*operators.OperatorGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OperatorGroupList_To_operators_OperatorGroupList(a.(*OperatorGroupList), b.(*operators.OperatorGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.OperatorGroupList)(nil), (*OperatorGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_OperatorGroupList_To_v1_OperatorGroupList(a.(*operators.OperatorGroupList), b.(*OperatorGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OperatorGroupSpec)(nil), (*operators.OperatorGroupSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OperatorGroupSpec_To_operators_OperatorGroupSpec(a.(*OperatorGroupSpec), b.(*operators.OperatorGroupSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.OperatorGroupSpec)(nil), (*OperatorGroupSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_OperatorGroupSpec_To_v1_OperatorGroupSpec(a.(*operators.OperatorGroupSpec), b.(*OperatorGroupSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OperatorGroupStatus)(nil), (*operators.OperatorGroupStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OperatorGroupStatus_To_operators_OperatorGroupStatus(a.(*OperatorGroupStatus), b.(*operators.OperatorGroupStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.OperatorGroupStatus)(nil), (*OperatorGroupStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_OperatorGroupStatus_To_v1_OperatorGroupStatus(a.(*operators.OperatorGroupStatus), b.(*OperatorGroupStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_OperatorGroup_To_operators_OperatorGroup(in *OperatorGroup, out *operators.OperatorGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_OperatorGroupSpec_To_operators_OperatorGroupSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_OperatorGroupStatus_To_operators_OperatorGroupStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_OperatorGroup_To_operators_OperatorGroup is an autogenerated conversion function.
func Convert_v1_OperatorGroup_To_operators_OperatorGroup(in *OperatorGroup, out *operators.OperatorGroup, s conversion.Scope) error {
	return autoConvert_v1_OperatorGroup_To_operators_OperatorGroup(in, out, s)
}

func autoConvert_operators_OperatorGroup_To_v1_OperatorGroup(in *operators.OperatorGroup, out *OperatorGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_operators_OperatorGroupSpec_To_v1_OperatorGroupSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_operators_OperatorGroupStatus_To_v1_OperatorGroupStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_operators_OperatorGroup_To_v1_OperatorGroup is an autogenerated conversion function.
func Convert_operators_OperatorGroup_To_v1_OperatorGroup(in *operators.OperatorGroup, out *OperatorGroup, s conversion.Scope) error {
	return autoConvert_operators_OperatorGroup_To_v1_OperatorGroup(in, out, s)
}

func autoConvert_v1_OperatorGroupList_To_operators_OperatorGroupList(in *OperatorGroupList, out *operators.OperatorGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]operators.OperatorGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_OperatorGroupList_To_operators_OperatorGroupList is an autogenerated conversion function.
func Convert_v1_OperatorGroupList_To_operators_OperatorGroupList(in *OperatorGroupList, out *operators.OperatorGroupList, s conversion.Scope) error {
	return autoConvert_v1_OperatorGroupList_To_operators_OperatorGroupList(in, out, s)
}

func autoConvert_operators_OperatorGroupList_To_v1_OperatorGroupList(in *operators.OperatorGroupList, out *OperatorGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]OperatorGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_operators_OperatorGroupList_To_v1_OperatorGroupList is an autogenerated conversion function.
func Convert_operators_OperatorGroupList_To_v1_OperatorGroupList(in *operators.OperatorGroupList, out *OperatorGroupList, s conversion.Scope) error {
	return autoConvert_operators_OperatorGroupList_To_v1_OperatorGroupList(in, out, s)
}

func autoConvert_v1_OperatorGroupSpec_To_operators_OperatorGroupSpec(in *OperatorGroupSpec, out *operators.OperatorGroupSpec, s conversion.Scope) error {
	out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.TargetNamespaces = *(*[]string)(unsafe.Pointer(&in.TargetNamespaces))
	out.ServiceAccountName = in.ServiceAccountName
	out.StaticProvidedAPIs = in.StaticProvidedAPIs
	return nil
}

// Convert_v1_OperatorGroupSpec_To_operators_OperatorGroupSpec is an autogenerated conversion function.
func Convert_v1_OperatorGroupSpec_To_operators_OperatorGroupSpec(in *OperatorGroupSpec, out *operators.OperatorGroupSpec, s conversion.Scope) error {
	return autoConvert_v1_OperatorGroupSpec_To_operators_OperatorGroupSpec(in, out, s)
}

func autoConvert_operators_OperatorGroupSpec_To_v1_OperatorGroupSpec(in *operators.OperatorGroupSpec, out *OperatorGroupSpec, s conversion.Scope) error {
	out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.TargetNamespaces = *(*[]string)(unsafe.Pointer(&in.TargetNamespaces))
	out.ServiceAccountName = in.ServiceAccountName
	out.StaticProvidedAPIs = in.StaticProvidedAPIs
	return nil
}

// Convert_operators_OperatorGroupSpec_To_v1_OperatorGroupSpec is an autogenerated conversion function.
func Convert_operators_OperatorGroupSpec_To_v1_OperatorGroupSpec(in *operators.OperatorGroupSpec, out *OperatorGroupSpec, s conversion.Scope) error {
	return autoConvert_operators_OperatorGroupSpec_To_v1_OperatorGroupSpec(in, out, s)
}

func autoConvert_v1_OperatorGroupStatus_To_operators_OperatorGroupStatus(in *OperatorGroupStatus, out *operators.OperatorGroupStatus, s conversion.Scope) error {
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.ServiceAccountRef = (*corev1.ObjectReference)(unsafe.Pointer(in.ServiceAccountRef))
	out.LastUpdated = (*metav1.Time)(unsafe.Pointer(in.LastUpdated))
	return nil
}

// Convert_v1_OperatorGroupStatus_To_operators_OperatorGroupStatus is an autogenerated conversion function.
func Convert_v1_OperatorGroupStatus_To_operators_OperatorGroupStatus(in *OperatorGroupStatus, out *operators.OperatorGroupStatus, s conversion.Scope) error {
	return autoConvert_v1_OperatorGroupStatus_To_operators_OperatorGroupStatus(in, out, s)
}

func autoConvert_operators_OperatorGroupStatus_To_v1_OperatorGroupStatus(in *operators.OperatorGroupStatus, out *OperatorGroupStatus, s conversion.Scope) error {
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.ServiceAccountRef = (*corev1.ObjectReference)(unsafe.Pointer(in.ServiceAccountRef))
	out.LastUpdated = (*metav1.Time)(unsafe.Pointer(in.LastUpdated))
	return nil
}

// Convert_operators_OperatorGroupStatus_To_v1_OperatorGroupStatus is an autogenerated conversion function.
func Convert_operators_OperatorGroupStatus_To_v1_OperatorGroupStatus(in *operators.OperatorGroupStatus, out *OperatorGroupStatus, s conversion.Scope) error {
	return autoConvert_operators_OperatorGroupStatus_To_v1_OperatorGroupStatus(in, out, s)
}