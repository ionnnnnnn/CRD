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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RollingMonitor12Spec defines the desired state of RollingMonitor12
type RollingMonitor12Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RollingMonitor12. Edit rollingmonitor12_types.go to remove/update
	DeploymentName string `json:"deploymentName,omitempty"`
}

// RollingMonitor12Status defines the observed state of RollingMonitor12
type RollingMonitor12Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RollingMonitor12 is the Schema for the rollingmonitor12s API
type RollingMonitor12 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RollingMonitor12Spec   `json:"spec,omitempty"`
	Status RollingMonitor12Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RollingMonitor12List contains a list of RollingMonitor12
type RollingMonitor12List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RollingMonitor12 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RollingMonitor12{}, &RollingMonitor12List{})
}