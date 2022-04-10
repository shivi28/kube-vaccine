//go:build !ignore_autogenerated_conversions
// +build !ignore_autogenerated_conversions

/*
Copyright The Kubernetes Authors.

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

package v3

import (
	v2 "kube-vaccine/api/v2"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Registration)(nil), (*v2.Registration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v3_Registration_To_v2_Registration(a.(*Registration), b.(*v2.Registration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2.Registration)(nil), (*Registration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_Registration_To_v3_Registration(a.(*v2.Registration), b.(*Registration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RegistrationList)(nil), (*v2.RegistrationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v3_RegistrationList_To_v2_RegistrationList(a.(*RegistrationList), b.(*v2.RegistrationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2.RegistrationList)(nil), (*RegistrationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_RegistrationList_To_v3_RegistrationList(a.(*v2.RegistrationList), b.(*RegistrationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RegistrationSpec)(nil), (*v2.RegistrationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v3_RegistrationSpec_To_v2_RegistrationSpec(a.(*RegistrationSpec), b.(*v2.RegistrationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2.RegistrationSpec)(nil), (*RegistrationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_RegistrationSpec_To_v3_RegistrationSpec(a.(*v2.RegistrationSpec), b.(*RegistrationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RegistrationStatus)(nil), (*v2.RegistrationStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v3_RegistrationStatus_To_v2_RegistrationStatus(a.(*RegistrationStatus), b.(*v2.RegistrationStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2.RegistrationStatus)(nil), (*RegistrationStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_RegistrationStatus_To_v3_RegistrationStatus(a.(*v2.RegistrationStatus), b.(*RegistrationStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v3_Registration_To_v2_Registration(in *Registration, out *v2.Registration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v3_RegistrationSpec_To_v2_RegistrationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v3_RegistrationStatus_To_v2_RegistrationStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v3_Registration_To_v2_Registration is an autogenerated conversion function.
func Convert_v3_Registration_To_v2_Registration(in *Registration, out *v2.Registration, s conversion.Scope) error {
	return autoConvert_v3_Registration_To_v2_Registration(in, out, s)
}

func autoConvert_v2_Registration_To_v3_Registration(in *v2.Registration, out *Registration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2_RegistrationSpec_To_v3_RegistrationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2_RegistrationStatus_To_v3_RegistrationStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2_Registration_To_v3_Registration is an autogenerated conversion function.
func Convert_v2_Registration_To_v3_Registration(in *v2.Registration, out *Registration, s conversion.Scope) error {
	return autoConvert_v2_Registration_To_v3_Registration(in, out, s)
}

func autoConvert_v3_RegistrationList_To_v2_RegistrationList(in *RegistrationList, out *v2.RegistrationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v2.Registration, len(*in))
		for i := range *in {
			if err := Convert_v3_Registration_To_v2_Registration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v3_RegistrationList_To_v2_RegistrationList is an autogenerated conversion function.
func Convert_v3_RegistrationList_To_v2_RegistrationList(in *RegistrationList, out *v2.RegistrationList, s conversion.Scope) error {
	return autoConvert_v3_RegistrationList_To_v2_RegistrationList(in, out, s)
}

func autoConvert_v2_RegistrationList_To_v3_RegistrationList(in *v2.RegistrationList, out *RegistrationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registration, len(*in))
		for i := range *in {
			if err := Convert_v2_Registration_To_v3_Registration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v2_RegistrationList_To_v3_RegistrationList is an autogenerated conversion function.
func Convert_v2_RegistrationList_To_v3_RegistrationList(in *v2.RegistrationList, out *RegistrationList, s conversion.Scope) error {
	return autoConvert_v2_RegistrationList_To_v3_RegistrationList(in, out, s)
}

func autoConvert_v3_RegistrationSpec_To_v2_RegistrationSpec(in *RegistrationSpec, out *v2.RegistrationSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.VerifiedID = in.VerifiedID
	// WARNING: in.VaccineDetails requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2_RegistrationSpec_To_v3_RegistrationSpec(in *v2.RegistrationSpec, out *RegistrationSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.VerifiedID = in.VerifiedID
	// WARNING: in.RegistrationDate requires manual conversion: does not exist in peer-type
	// WARNING: in.VaccineName requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v3_RegistrationStatus_To_v2_RegistrationStatus(in *RegistrationStatus, out *v2.RegistrationStatus, s conversion.Scope) error {
	return nil
}

// Convert_v3_RegistrationStatus_To_v2_RegistrationStatus is an autogenerated conversion function.
func Convert_v3_RegistrationStatus_To_v2_RegistrationStatus(in *RegistrationStatus, out *v2.RegistrationStatus, s conversion.Scope) error {
	return autoConvert_v3_RegistrationStatus_To_v2_RegistrationStatus(in, out, s)
}

func autoConvert_v2_RegistrationStatus_To_v3_RegistrationStatus(in *v2.RegistrationStatus, out *RegistrationStatus, s conversion.Scope) error {
	return nil
}

// Convert_v2_RegistrationStatus_To_v3_RegistrationStatus is an autogenerated conversion function.
func Convert_v2_RegistrationStatus_To_v3_RegistrationStatus(in *v2.RegistrationStatus, out *RegistrationStatus, s conversion.Scope) error {
	return autoConvert_v2_RegistrationStatus_To_v3_RegistrationStatus(in, out, s)
}
