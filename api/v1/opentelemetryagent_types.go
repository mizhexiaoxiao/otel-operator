/*
Copyright 2023.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OpenTelemetryAgentSpec defines the desired state of OpenTelemetryAgent
type OpenTelemetryAgentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image            string                        `json:"image"`
	PodAnnotations   map[string]string             `json:"podAnnotations,omitempty"`
	ImagePullPolicy  corev1.PullPolicy             `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Config           string                        `json:"config"`
	DNSPolicy        corev1.DNSPolicy              `json:"dnsPolicy,omitempty"`
	HostNetwork      bool                          `json:"hostNetwork,omitempty"`
	Ports            []corev1.ContainerPort        `json:"ports,omitempty"`
	Resources        corev1.ResourceRequirements   `json:"resources,omitempty"`
}

// OpenTelemetryAgentStatus defines the observed state of OpenTelemetryAgent
type OpenTelemetryAgentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OpenTelemetryAgent is the Schema for the opentelemetryagents API
type OpenTelemetryAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenTelemetryAgentSpec   `json:"spec,omitempty"`
	Status OpenTelemetryAgentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OpenTelemetryAgentList contains a list of OpenTelemetryAgent
type OpenTelemetryAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenTelemetryAgent `json:"items"`
}

func (o OpenTelemetryAgent) GetAnnotations() map[string]string {
	return o.Annotations
}

func (o OpenTelemetryAgent) GetConfig() string {
	return o.Spec.Config
}

func (o OpenTelemetryAgent) GetPodAnnotations() map[string]string {
	return o.Spec.PodAnnotations
}

func (o OpenTelemetryAgent) GetHostNetwork() bool {
	return o.Spec.HostNetwork
}

func init() {
	SchemeBuilder.Register(&OpenTelemetryAgent{}, &OpenTelemetryAgentList{})
}
