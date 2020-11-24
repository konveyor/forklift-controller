// +build !ignore_autogenerated

/*
Copyright 2019 Red Hat Inc.

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

// Code generated by controller-gen. DO NOT EDIT.

package mapped

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationNetwork) DeepCopyInto(out *DestinationNetwork) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationNetwork.
func (in *DestinationNetwork) DeepCopy() *DestinationNetwork {
	if in == nil {
		return nil
	}
	out := new(DestinationNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationStorage) DeepCopyInto(out *DestinationStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationStorage.
func (in *DestinationStorage) DeepCopy() *DestinationStorage {
	if in == nil {
		return nil
	}
	out := new(DestinationStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPair) DeepCopyInto(out *NetworkPair) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPair.
func (in *NetworkPair) DeepCopy() *NetworkPair {
	if in == nil {
		return nil
	}
	out := new(NetworkPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePair) DeepCopyInto(out *StoragePair) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePair.
func (in *StoragePair) DeepCopy() *StoragePair {
	if in == nil {
		return nil
	}
	out := new(StoragePair)
	in.DeepCopyInto(out)
	return out
}
