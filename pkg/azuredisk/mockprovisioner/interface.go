/*
Copyright The Kubernetes Authors.

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

// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/azuredisk/interface_v2.go

// Package mockprovisioner is a generated GoMock package.
package mockprovisioner

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	azdiskv1beta2 "sigs.k8s.io/azuredisk-csi-driver/pkg/apis/azuredisk/v1beta2"
	azdisk "sigs.k8s.io/azuredisk-csi-driver/pkg/apis/client/clientset/versioned"
	azureconstants "sigs.k8s.io/azuredisk-csi-driver/pkg/azureconstants"
)

// MockCrdProvisioner is a mock of CrdProvisioner interface
type MockCrdProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockCrdProvisionerMockRecorder
}

// MockCrdProvisionerMockRecorder is the mock recorder for MockCrdProvisioner
type MockCrdProvisionerMockRecorder struct {
	mock *MockCrdProvisioner
}

// NewMockCrdProvisioner creates a new mock instance
func NewMockCrdProvisioner(ctrl *gomock.Controller) *MockCrdProvisioner {
	mock := &MockCrdProvisioner{ctrl: ctrl}
	mock.recorder = &MockCrdProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrdProvisioner) EXPECT() *MockCrdProvisionerMockRecorder {
	return m.recorder
}

// CreateVolume mocks base method
func (m *MockCrdProvisioner) CreateVolume(ctx context.Context, volumeName string, capacityRange *azdiskv1beta2.CapacityRange, volumeCapabilities []azdiskv1beta2.VolumeCapability, parameters, secrets map[string]string, volumeContentSource *azdiskv1beta2.ContentVolumeSource, accessibilityReq *azdiskv1beta2.TopologyRequirement) (*azdiskv1beta2.AzVolumeStatusDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", ctx, volumeName, capacityRange, volumeCapabilities, parameters, secrets, volumeContentSource, accessibilityReq)
	ret0, _ := ret[0].(*azdiskv1beta2.AzVolumeStatusDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockCrdProvisionerMockRecorder) CreateVolume(ctx, volumeName, capacityRange, volumeCapabilities, parameters, secrets, volumeContentSource, accessibilityReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockCrdProvisioner)(nil).CreateVolume), ctx, volumeName, capacityRange, volumeCapabilities, parameters, secrets, volumeContentSource, accessibilityReq)
}

// DeleteVolume mocks base method
func (m *MockCrdProvisioner) DeleteVolume(ctx context.Context, volumeID string, secrets map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", ctx, volumeID, secrets)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockCrdProvisionerMockRecorder) DeleteVolume(ctx, volumeID, secrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockCrdProvisioner)(nil).DeleteVolume), ctx, volumeID, secrets)
}

// PublishVolume mocks base method
func (m *MockCrdProvisioner) PublishVolume(ctx context.Context, volumeID, nodeID string, volumeCapability *azdiskv1beta2.VolumeCapability, readOnly bool, secrets, volumeContext map[string]string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishVolume", ctx, volumeID, nodeID, volumeCapability, readOnly, secrets, volumeContext)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishVolume indicates an expected call of PublishVolume
func (mr *MockCrdProvisionerMockRecorder) PublishVolume(ctx, volumeID, nodeID, volumeCapability, readOnly, secrets, volumeContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishVolume", reflect.TypeOf((*MockCrdProvisioner)(nil).PublishVolume), ctx, volumeID, nodeID, volumeCapability, readOnly, secrets, volumeContext)
}

// WaitForAttach mocks base method
func (m *MockCrdProvisioner) WaitForAttach(ctx context.Context, volume, node string) (*azdiskv1beta2.AzVolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAttach", ctx, volume, node)
	ret0, _ := ret[0].(*azdiskv1beta2.AzVolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForAttach indicates an expected call of WaitForAttach
func (mr *MockCrdProvisionerMockRecorder) WaitForAttach(ctx, volume, node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAttach", reflect.TypeOf((*MockCrdProvisioner)(nil).WaitForAttach), ctx, volume, node)
}

// UnpublishVolume mocks base method
func (m *MockCrdProvisioner) UnpublishVolume(ctx context.Context, volumeID, nodeID string, secrets map[string]string, mode azureconstants.UnpublishMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpublishVolume", ctx, volumeID, nodeID, secrets, mode)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpublishVolume indicates an expected call of UnpublishVolume
func (mr *MockCrdProvisionerMockRecorder) UnpublishVolume(ctx, volumeID, nodeID, secrets, mode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishVolume", reflect.TypeOf((*MockCrdProvisioner)(nil).UnpublishVolume), ctx, volumeID, nodeID, secrets, mode)
}

// WaitForDetach mocks base method
func (m *MockCrdProvisioner) WaitForDetach(ctx context.Context, volume, node string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDetach", ctx, volume, node)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForDetach indicates an expected call of WaitForDetach
func (mr *MockCrdProvisionerMockRecorder) WaitForDetach(ctx, volume, node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDetach", reflect.TypeOf((*MockCrdProvisioner)(nil).WaitForDetach), ctx, volume, node)
}

// GetAzVolumeAttachment mocks base method
func (m *MockCrdProvisioner) GetAzVolumeAttachment(ctx context.Context, volumeID, nodeID string) (*azdiskv1beta2.AzVolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzVolumeAttachment", ctx, volumeID, nodeID)
	ret0, _ := ret[0].(*azdiskv1beta2.AzVolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzVolumeAttachment indicates an expected call of GetAzVolumeAttachment
func (mr *MockCrdProvisionerMockRecorder) GetAzVolumeAttachment(ctx, volumeID, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzVolumeAttachment", reflect.TypeOf((*MockCrdProvisioner)(nil).GetAzVolumeAttachment), ctx, volumeID, nodeID)
}

// ExpandVolume mocks base method
func (m *MockCrdProvisioner) ExpandVolume(ctx context.Context, volumeID string, capacityRange *azdiskv1beta2.CapacityRange, secrets map[string]string) (*azdiskv1beta2.AzVolumeStatusDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpandVolume", ctx, volumeID, capacityRange, secrets)
	ret0, _ := ret[0].(*azdiskv1beta2.AzVolumeStatusDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpandVolume indicates an expected call of ExpandVolume
func (mr *MockCrdProvisionerMockRecorder) ExpandVolume(ctx, volumeID, capacityRange, secrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandVolume", reflect.TypeOf((*MockCrdProvisioner)(nil).ExpandVolume), ctx, volumeID, capacityRange, secrets)
}

// GetDiskClientSet mocks base method
func (m *MockCrdProvisioner) GetDiskClientSet() azdisk.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskClientSet")
	ret0, _ := ret[0].(azdisk.Interface)
	return ret0
}

// GetDiskClientSet indicates an expected call of GetDiskClientSet
func (mr *MockCrdProvisionerMockRecorder) GetDiskClientSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskClientSet", reflect.TypeOf((*MockCrdProvisioner)(nil).GetDiskClientSet))
}

// IsDriverUninstall mocks base method
func (m *MockCrdProvisioner) IsDriverUninstall() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDriverUninstall")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDriverUninstall indicates an expected call of GetDiskClientSet
func (mr *MockCrdProvisionerMockRecorder) IsDriverUninstall() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDriverUninstall", reflect.TypeOf((*MockCrdProvisioner)(nil).IsDriverUninstall))
}

// MockNodeProvisioner is a mock of NodeProvisioner interface
type MockNodeProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockNodeProvisionerMockRecorder
}

// MockNodeProvisionerMockRecorder is the mock recorder for MockNodeProvisioner
type MockNodeProvisionerMockRecorder struct {
	mock *MockNodeProvisioner
}

// NewMockNodeProvisioner creates a new mock instance
func NewMockNodeProvisioner(ctrl *gomock.Controller) *MockNodeProvisioner {
	mock := &MockNodeProvisioner{ctrl: ctrl}
	mock.recorder = &MockNodeProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeProvisioner) EXPECT() *MockNodeProvisionerMockRecorder {
	return m.recorder
}

// GetDevicePathWithLUN mocks base method
func (m *MockNodeProvisioner) GetDevicePathWithLUN(ctx context.Context, lun int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicePathWithLUN", ctx, lun)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicePathWithLUN indicates an expected call of GetDevicePathWithLUN
func (mr *MockNodeProvisionerMockRecorder) GetDevicePathWithLUN(ctx, lun interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicePathWithLUN", reflect.TypeOf((*MockNodeProvisioner)(nil).GetDevicePathWithLUN), ctx, lun)
}

// GetDevicePathWithMountPath mocks base method
func (m *MockNodeProvisioner) GetDevicePathWithMountPath(mountPath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicePathWithMountPath", mountPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicePathWithMountPath indicates an expected call of GetDevicePathWithMountPath
func (mr *MockNodeProvisionerMockRecorder) GetDevicePathWithMountPath(mountPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicePathWithMountPath", reflect.TypeOf((*MockNodeProvisioner)(nil).GetDevicePathWithMountPath), mountPath)
}

// IsBlockDevicePath mocks base method
func (m *MockNodeProvisioner) IsBlockDevicePath(path string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBlockDevicePath", path)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsBlockDevicePath indicates an expected call of IsBlockDevicePath
func (mr *MockNodeProvisionerMockRecorder) IsBlockDevicePath(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBlockDevicePath", reflect.TypeOf((*MockNodeProvisioner)(nil).IsBlockDevicePath), path)
}

// PreparePublishPath mocks base method
func (m *MockNodeProvisioner) PreparePublishPath(target string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreparePublishPath", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreparePublishPath indicates an expected call of PreparePublishPath
func (mr *MockNodeProvisionerMockRecorder) PreparePublishPath(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreparePublishPath", reflect.TypeOf((*MockNodeProvisioner)(nil).PreparePublishPath), target)
}

// EnsureMountPointReady mocks base method
func (m *MockNodeProvisioner) EnsureMountPointReady(target string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureMountPointReady", target)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureMountPointReady indicates an expected call of EnsureMountPointReady
func (mr *MockNodeProvisionerMockRecorder) EnsureMountPointReady(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureMountPointReady", reflect.TypeOf((*MockNodeProvisioner)(nil).EnsureMountPointReady), target)
}

// EnsureBlockTargetReady mocks base method
func (m *MockNodeProvisioner) EnsureBlockTargetReady(target string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureBlockTargetReady", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureBlockTargetReady indicates an expected call of EnsureBlockTargetReady
func (mr *MockNodeProvisionerMockRecorder) EnsureBlockTargetReady(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureBlockTargetReady", reflect.TypeOf((*MockNodeProvisioner)(nil).EnsureBlockTargetReady), target)
}

// FormatAndMount mocks base method
func (m *MockNodeProvisioner) FormatAndMount(source, target, fstype string, options []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatAndMount", source, target, fstype, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// FormatAndMount indicates an expected call of FormatAndMount
func (mr *MockNodeProvisionerMockRecorder) FormatAndMount(source, target, fstype, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatAndMount", reflect.TypeOf((*MockNodeProvisioner)(nil).FormatAndMount), source, target, fstype, options)
}

// Mount mocks base method
func (m *MockNodeProvisioner) Mount(source, target, fstype string, options []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", source, target, fstype, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount
func (mr *MockNodeProvisionerMockRecorder) Mount(source, target, fstype, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockNodeProvisioner)(nil).Mount), source, target, fstype, options)
}

// Unmount mocks base method
func (m *MockNodeProvisioner) Unmount(target string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount
func (mr *MockNodeProvisionerMockRecorder) Unmount(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockNodeProvisioner)(nil).Unmount), target)
}

// CleanupMountPoint mocks base method
func (m *MockNodeProvisioner) CleanupMountPoint(path string, extensiveCheck bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupMountPoint", path, extensiveCheck)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupMountPoint indicates an expected call of CleanupMountPoint
func (mr *MockNodeProvisionerMockRecorder) CleanupMountPoint(path, extensiveCheck interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupMountPoint", reflect.TypeOf((*MockNodeProvisioner)(nil).CleanupMountPoint), path, extensiveCheck)
}

// RescanVolume mocks base method
func (m *MockNodeProvisioner) RescanVolume(devicePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RescanVolume", devicePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// RescanVolume indicates an expected call of RescanVolume
func (mr *MockNodeProvisionerMockRecorder) RescanVolume(devicePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RescanVolume", reflect.TypeOf((*MockNodeProvisioner)(nil).RescanVolume), devicePath)
}

// Resize mocks base method
func (m *MockNodeProvisioner) Resize(source, target string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", source, target)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resize indicates an expected call of Resize
func (mr *MockNodeProvisionerMockRecorder) Resize(source, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockNodeProvisioner)(nil).Resize), source, target)
}

// GetBlockSizeBytes mocks base method
func (m *MockNodeProvisioner) GetBlockSizeBytes(devicePath string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockSizeBytes", devicePath)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockSizeBytes indicates an expected call of GetBlockSizeBytes
func (mr *MockNodeProvisionerMockRecorder) GetBlockSizeBytes(devicePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockSizeBytes", reflect.TypeOf((*MockNodeProvisioner)(nil).GetBlockSizeBytes), devicePath)
}
