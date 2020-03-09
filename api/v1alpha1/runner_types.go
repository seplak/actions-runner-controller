/*
Copyright 2020 The actions-runner-controller authors.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KeyRunnerName       = "actions.summerwind.dev/runner-name"
	KeyRunnerRepository = "actions.summerwind.dev/runner-repository"

	EnvRunnerName       = "RUNNER_NAME"
	EnvRunnerRepository = "RUNNER_REPO"
	EnvRunnerToken      = "RUNNER_TOKEN"

	ContainerName = "runner"
)

// RunnerSpec defines the desired state of Runner
type RunnerSpec struct {
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:Pattern=`^[^/]+/[^/]+$`
	Repository string `json:"repository"`

	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// +optional
	Image string `json:"image,omitempty"`

	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`
}

// RunnerStatus defines the observed state of Runner
type RunnerStatus struct {
	// +optional
	Registration *RunnerStatusRegistration `json:"registration,omitempty"`
	Phase        string                    `json:"phase"`
	Reason       string                    `json:"reason"`
	Message      string                    `json:"message"`

	// +optional
	Replicas int32 `json:"replicas,omitempty"`

	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// +optional
	CurrentReplicas int32 `json:"currentReplicas,omitempty"`

	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`
}

type RunnerStatusRegistration struct {
	Repository string      `json:"repository"`
	Token      string      `json:"token"`
	ExpiresAt  metav1.Time `json:"expiresAt"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.replicas",name="Desired",type="integer",description="Desired Replicas"
// +kubebuilder:printcolumn:JSONPath=".status.replicas",name="Current",type="integer",description="Current Replicas"
// +kubebuilder:printcolumn:JSONPath=".status.readyReplicas",name="Ready",type="integer",description="Ready Replicas"

// Runner is the Schema for the runners API
type Runner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunnerSpec   `json:"spec,omitempty"`
	Status RunnerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RunnerList contains a list of Runner
type RunnerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Runner `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.replicas",name=Desired,type=number
// +kubebuilder:printcolumn:JSONPath=".status.availableReplicas",name=Current,type=number
// +kubebuilder:printcolumn:JSONPath=".status.readyReplicas",name=Ready,type=number

// RunnerSet is the Schema for the runnersets API
type RunnerSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunnerSetSpec   `json:"spec,omitempty"`
	Status RunnerSetStatus `json:"status,omitempty"`
}

// RunnerSetSpec defines the desired state of RunnerSet
type RunnerSetSpec struct {
	Replicas *int `json:"replicas"`

	Template RunnerTemplate `json:"template"`
}

type RunnerSetStatus struct {
	AvailableReplicas int `json:"availableReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
}

type RunnerTemplate struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RunnerSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// RunnerList contains a list of Runner
type RunnerSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunnerSet `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.replicas",name=Desired,type=number
// +kubebuilder:printcolumn:JSONPath=".status.availableReplicas",name=Current,type=number
// +kubebuilder:printcolumn:JSONPath=".status.readyReplicas",name=Ready,type=number

// RunnerSet is the Schema for the runnersets API
type RunnerDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunnerDeploymentSpec   `json:"spec,omitempty"`
	Status RunnerDeploymentStatus `json:"status,omitempty"`
}

// RunnerSetSpec defines the desired state of RunnerDeployment
type RunnerDeploymentSpec struct {
	Replicas *int `json:"replicas"`

	Template RunnerTemplate `json:"template"`
}

type RunnerDeploymentStatus struct {
	AvailableReplicas int `json:"availableReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
}

// +kubebuilder:object:root=true

// RunnerList contains a list of Runner
type RunnerDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunnerDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Runner{}, &RunnerList{}, &RunnerSet{}, &RunnerSetList{}, &RunnerDeployment{}, &RunnerDeploymentList{})
}
