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

// ExecutorSpec defines the desired state of Executor
type ExecutorSpec struct {
	// +kubebuilder:printcolumn:name="Server",type=string,JSONPath=`.spec.server`
	// +kubebuilder:validation:Pattern=`[a-z]+(?:\.[a-z]+)*(:\d+)`
	// Server is the URL of the Armada Server gRPC endpoint (format must be <host>:<port>)
	Server string `json:"server"`
	// ForceNoTLS enables gRPC connection over an unsecure connection. Should not be used in production environments.
	ForceNoTLS bool `json:"forceNoTLS,omitempty"`
}

type ExecutorKubernetesConfig struct {
	// ImpersonateUsers controls whether to impersonate users when creating Kubernetes objects.
	ImpersonateUsers bool `json:"impersonateUsers,omitempty"`
}

// ExecutorStatus defines the observed state of Executor
type ExecutorStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Executor is the Schema for the executors API
type Executor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExecutorSpec   `json:"spec,omitempty"`
	Status ExecutorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExecutorList contains a list of Executor
type ExecutorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Executor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Executor{}, &ExecutorList{})
}
