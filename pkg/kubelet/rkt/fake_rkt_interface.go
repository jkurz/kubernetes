/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package rkt

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/coreos/go-systemd/dbus"
	rktapi "github.com/coreos/rkt/api/v1alpha"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// fakeRktInterface mocks the rktapi.PublicAPIClient interface for testing purpose.
type fakeRktInterface struct {
	sync.Mutex
	info   rktapi.Info
	called []string
	err    error
}

func newFakeRktInterface() *fakeRktInterface {
	return &fakeRktInterface{}
}

func (f *fakeRktInterface) CleanCalls() {
	f.Lock()
	defer f.Unlock()
	f.called = nil
}

func (f *fakeRktInterface) GetInfo(ctx context.Context, in *rktapi.GetInfoRequest, opts ...grpc.CallOption) (*rktapi.GetInfoResponse, error) {
	f.Lock()
	defer f.Unlock()

	f.called = append(f.called, "GetInfo")
	return &rktapi.GetInfoResponse{&f.info}, f.err
}

func (f *fakeRktInterface) ListPods(ctx context.Context, in *rktapi.ListPodsRequest, opts ...grpc.CallOption) (*rktapi.ListPodsResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (f *fakeRktInterface) InspectPod(ctx context.Context, in *rktapi.InspectPodRequest, opts ...grpc.CallOption) (*rktapi.InspectPodResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (f *fakeRktInterface) ListImages(ctx context.Context, in *rktapi.ListImagesRequest, opts ...grpc.CallOption) (*rktapi.ListImagesResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (f *fakeRktInterface) InspectImage(ctx context.Context, in *rktapi.InspectImageRequest, opts ...grpc.CallOption) (*rktapi.InspectImageResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (f *fakeRktInterface) ListenEvents(ctx context.Context, in *rktapi.ListenEventsRequest, opts ...grpc.CallOption) (rktapi.PublicAPI_ListenEventsClient, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (f *fakeRktInterface) GetLogs(ctx context.Context, in *rktapi.GetLogsRequest, opts ...grpc.CallOption) (rktapi.PublicAPI_GetLogsClient, error) {
	return nil, fmt.Errorf("Not implemented")
}

// fakeSystemd mocks the systemdInterface for testing purpose.
// TODO(yifan): Remove this once we have a package for launching rkt pods.
// See https://github.com/coreos/rkt/issues/1769.
type fakeSystemd struct {
	sync.Mutex
	called  []string
	version string
	err     error
}

func newFakeSystemd() *fakeSystemd {
	return &fakeSystemd{}
}

func (f *fakeSystemd) CleanCalls() {
	f.Lock()
	defer f.Unlock()
	f.called = nil
}

func (f *fakeSystemd) Version() (systemdVersion, error) {
	f.Lock()
	defer f.Unlock()

	f.called = append(f.called, "Version")
	v, _ := strconv.Atoi(f.version)
	return systemdVersion(v), f.err
}

func (f *fakeSystemd) ListUnits() ([]dbus.UnitStatus, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (f *fakeSystemd) StopUnit(name, mode string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func (f *fakeSystemd) RestartUnit(name, mode string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func (f *fakeSystemd) Reload() error {
	return fmt.Errorf("Not implemented")
}
