/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package v1beta1

import (
	apiextv1b1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubeFedConfigSpec defines the desired state of KubeFedConfig
type KubeFedConfigSpec struct {
	// The scope of the KubeFed control plane should be either
	// `Namespaced` or `Cluster`. `Namespaced` indicates that the
	// KubeFed namespace will be the only target of the control plane.
	Scope apiextv1b1.ResourceScope `json:"scope"`
	// +optional
	ControllerDuration *DurationConfig `json:"controllerDuration,omitempty"`
	// +optional
	LeaderElect *LeaderElectConfig `json:"leaderElect,omitempty"`
	// +optional
	FeatureGates []FeatureGatesConfig `json:"featureGates,omitempty"`
	// +optional
	ClusterHealthCheck *ClusterHealthCheckConfig `json:"clusterHealthCheck,omitempty"`
	// +optional
	SyncController *SyncControllerConfig `json:"syncController,omitempty"`
}

// DurationConfig defines the duration of the KubeFed control plane.
type DurationConfig struct {
	// Time to wait before reconciling on a healthy cluster.
	// +optional
	AvailableDelay *metav1.Duration `json:"availableDelay,omitempty"`
	// Time to wait before giving up on an unhealthy cluster.
	// +optional
	UnavailableDelay *metav1.Duration `json:"unavailableDelay,omitempty"`
}

// LeaderElectConfig defines the leader election configuration of the KubeFed control plane.
type LeaderElectConfig struct {
	// The duration that non-leader candidates will wait after observing a leadership
	// renewal until attempting to acquire leadership of a led but unrenewed leader
	// slot. This is effectively the maximum duration that a leader can be stopped
	// before it is replaced by another candidate. This is only applicable if leader
	// election is enabled.
	// +optional
	LeaseDuration *metav1.Duration `json:"leaseDuration,omitempty"`
	// The interval between attempts by the acting master to renew a leadership slot
	// before it stops leading. This must be less than or equal to the lease duration.
	// This is only applicable if leader election is enabled.
	// +optional
	RenewDeadline *metav1.Duration `json:"renewDeadline,omitempty"`
	// The duration the clients should wait between attempting acquisition and renewal
	// of a leadership. This is only applicable if leader election is enabled.
	// +optional
	RetryPeriod *metav1.Duration `json:"retryPeriod,omitempty"`
	// The type of resource object that is used for locking during
	// leader election. Supported options are `configmaps` (default) and `endpoints`.
	// +optional
	ResourceLock *ResourceLockType `json:"resourceLock,omitempty"`
}

// ResourceLockType string
type ResourceLockType string

const (
	ConfigMapsResourceLock ResourceLockType = "configmaps"
	EndpointsResourceLock  ResourceLockType = "endpoints"
)

// FeatureGatesConfig defines the feature gates of the KubeFed control plane.
type FeatureGatesConfig struct {
	Name          string            `json:"name"`
	Configuration ConfigurationMode `json:"configuration"`
}

// ConfigurationMode string
type ConfigurationMode string

const (
	ConfigurationEnabled  ConfigurationMode = "Enabled"
	ConfigurationDisabled ConfigurationMode = "Disabled"
)

// ClusterHealthCheckConfig defines the cluster health check configuration of the KubeFed control plane.
type ClusterHealthCheckConfig struct {
	// How often to monitor the cluster health.
	// +optional
	Period *metav1.Duration `json:"period,omitempty"`
	// Minimum consecutive failures for the cluster health to be considered failed after having succeeded.
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty"`
	// Minimum consecutive successes for the cluster health to be considered successful after having failed.
	// +optional
	SuccessThreshold *int64 `json:"successThreshold,omitempty"`
	// Duration after which the cluster health check times out.
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// SyncControllerConfig defines the sync controller configuration of the KubeFed control plane.
type SyncControllerConfig struct {
	// Whether to adopt pre-existing resources in member clusters. Defaults to
	// "Enabled".
	// +optional
	AdoptResources *ResourceAdoption `json:"adoptResources,omitempty"`
}

// ResourceAdoption string
type ResourceAdoption string

const (
	AdoptResourcesEnabled  ResourceAdoption = "Enabled"
	AdoptResourcesDisabled ResourceAdoption = "Disabled"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeFedConfig
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=kubefedconfigs
type KubeFedConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubeFedConfigSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeFedConfigList contains a list of KubeFedConfig
type KubeFedConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeFedConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeFedConfig{}, &KubeFedConfigList{})
}
