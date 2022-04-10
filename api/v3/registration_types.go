/*
Copyright 2022.

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

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RegistrationSpec defines the desired state of Registration
type RegistrationSpec struct {
	Name           string           `json:"name"`
	VerifiedID     string           `json:"verified_id"`
	VaccineDetails []*VaccineDetail `json:"vaccine_details"`
}

// VaccineDetail defines vaccine details like name and registration date
type VaccineDetail struct {
	RegistrationDate string `json:"appointment_date"`
	VaccineName      string `json:"vaccine_name"`
}

// RegistrationStatus defines the observed state of Registration
type RegistrationStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Registration is the Schema for the registrations API
type Registration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegistrationSpec   `json:"spec,omitempty"`
	Status RegistrationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RegistrationList contains a list of Registration
type RegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Registration{}, &RegistrationList{})
}
