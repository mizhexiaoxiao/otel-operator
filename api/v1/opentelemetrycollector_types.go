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

// OpenTelemetryCollectorSpec defines the desired state of OpenTelemetryCollector
type OpenTelemetryCollectorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image            string                        `json:"image"`
	Replicas         *int32                        `json:"replicas"`
	PodAnnotations   map[string]string             `json:"podAnnotations,omitempty"`
	ImagePullPolicy  corev1.PullPolicy             `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Config           string                        `json:"config"`
	DNSPolicy        corev1.DNSPolicy              `json:"dnsPolicy,omitempty"`
	HostNetwork      bool                          `json:"hostNetwork,omitempty"`
	Ports            []corev1.ContainerPort        `json:"ports,omitempty"`
	Resources        corev1.ResourceRequirements   `json:"resources,omitempty"`
}

// OpenTelemetryCollectorStatus defines the observed state of OpenTelemetryCollector
type OpenTelemetryCollectorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OpenTelemetryCollector is the Schema for the opentelemetrycollectors API
type OpenTelemetryCollector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenTelemetryCollectorSpec   `json:"spec,omitempty"`
	Status OpenTelemetryCollectorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OpenTelemetryCollectorList contains a list of OpenTelemetryCollector
type OpenTelemetryCollectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenTelemetryCollector `json:"items"`
}

func (o OpenTelemetryCollector) GetAnnotations() map[string]string {
	return o.Annotations
}

func (o OpenTelemetryCollector) GetConfig() string {
	return o.Spec.Config
}

func (o OpenTelemetryCollector) GetPodAnnotations() map[string]string {
	return o.Spec.PodAnnotations
}

func (o OpenTelemetryCollector) GetHostNetwork() bool {
	return o.Spec.HostNetwork
}

func init() {
	SchemeBuilder.Register(&OpenTelemetryCollector{}, &OpenTelemetryCollectorList{})
}
