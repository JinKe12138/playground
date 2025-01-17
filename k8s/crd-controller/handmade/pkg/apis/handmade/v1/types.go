/*
Copyright 2015 The Kubernetes Authors.

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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoddessMoment is a top-level type. A client is created for it.
type GoddessMoment struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec GoddessMomentSpec `json:"spec,omitempty"`
	// +optional
	Status GoddessMomentStatus `json:"status,omitempty"`
}

//GoddessMomentSpec .
type GoddessMomentSpec struct {
	FoodDemand []FoodDemand `json:"foodDemand,omitempty"`
}

//FoodDemand .
type FoodDemand struct {
	Name string `json:"name,omitempty"`
}

//GoddessMomentStatus .
type GoddessMomentStatus struct {
	FoodDemand []FoodDemandStatus `json:"foodDemand,omitempty"`
}

//FoodDemandStatus .
type FoodDemandStatus struct {
	Name        string      `json:"name,omitempty"`
	Status      FoodStatus  `json:"status,omitempty"`
	ClaimTime   metav1.Time `json:"claimTime,omitempty"`
	ArrivalTime metav1.Time `json:"arrivalTime,omitempty"`
	ClaimBy     string      `json:"claimBy,omitempty"`
}

//FoodStatus .
type FoodStatus string

const (
	//FoodStatusPending .
	FoodStatusPending = "Pending"
	//FoodStatusPendingArrival .
	FoodStatusPendingArrival = "PendingArrival"
	//FoodStatusArrived .
	FoodStatusArrived = "Arrived"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoddessMomentList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type GoddessMomentList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []GoddessMoment `json:"items"`
}
