/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NetworkConfigStatusApplyConfiguration represents an declarative configuration of the NetworkConfigStatus type for use
// with apply.
type NetworkConfigStatusApplyConfiguration struct {
	MACAddress  *string `json:"macaddress,omitempty"`
	NetworkName *string `json:"networkname,omitempty"`
	Status      *string `json:"status,omitempty"`
	Message     *string `json:"message,omitempty"`
}

// NetworkConfigStatusApplyConfiguration constructs an declarative configuration of the NetworkConfigStatus type for use with
// apply.
func NetworkConfigStatus() *NetworkConfigStatusApplyConfiguration {
	return &NetworkConfigStatusApplyConfiguration{}
}

// WithMACAddress sets the MACAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MACAddress field is set to the value of the last call.
func (b *NetworkConfigStatusApplyConfiguration) WithMACAddress(value string) *NetworkConfigStatusApplyConfiguration {
	b.MACAddress = &value
	return b
}

// WithNetworkName sets the NetworkName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkName field is set to the value of the last call.
func (b *NetworkConfigStatusApplyConfiguration) WithNetworkName(value string) *NetworkConfigStatusApplyConfiguration {
	b.NetworkName = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *NetworkConfigStatusApplyConfiguration) WithStatus(value string) *NetworkConfigStatusApplyConfiguration {
	b.Status = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *NetworkConfigStatusApplyConfiguration) WithMessage(value string) *NetworkConfigStatusApplyConfiguration {
	b.Message = &value
	return b
}
