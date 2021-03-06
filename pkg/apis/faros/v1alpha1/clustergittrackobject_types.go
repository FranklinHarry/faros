/*
Copyright 2018 Pusher Ltd.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ClusterGitTrackObject is the Schema for the clustergittrackobjects API
// +k8s:openapi-gen=true
type ClusterGitTrackObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitTrackObjectSpec   `json:"spec,omitempty"`
	Status GitTrackObjectStatus `json:"status,omitempty"`
}

// GetNamespacedName implementes the GitTrackObject interface
func (g *ClusterGitTrackObject) GetNamespacedName() string {
	return g.Name
}

// GetSpec implements the GitTrackObject interface
func (g *ClusterGitTrackObject) GetSpec() GitTrackObjectSpec {
	return g.Spec
}

// SetSpec implements the GitTrackObject interface
func (g *ClusterGitTrackObject) SetSpec(s GitTrackObjectSpec) {
	g.Spec = s
}

// GetStatus implements the GitTrackObject interface
func (g *ClusterGitTrackObject) GetStatus() GitTrackObjectStatus {
	return g.Status
}

// SetStatus implements the GitTrackObject interface
func (g *ClusterGitTrackObject) SetStatus(s GitTrackObjectStatus) {
	g.Status = s
}

// DeepCopyInterface implements the GitTrackObject interface
func (g *ClusterGitTrackObject) DeepCopyInterface() GitTrackObjectInterface {
	return g.DeepCopy()
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ClusterGitTrackObjectList contains a list of ClusterGitTrackObject
type ClusterGitTrackObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterGitTrackObject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterGitTrackObject{}, &ClusterGitTrackObjectList{})
}
