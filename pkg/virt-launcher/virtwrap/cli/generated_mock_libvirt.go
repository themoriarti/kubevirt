// Code generated by MockGen. DO NOT EDIT.
// Source: libvirt.go
//
// Generated by this command:
//
//	mockgen -source libvirt.go -imports libvirt=libvirt.org/go/libvirt -package=cli -destination=generated_mock_libvirt.go
//

// Package cli is a generated GoMock package.
package cli

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
	libvirt "libvirt.org/go/libvirt"

	api "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
	stats "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/stats"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
	isgomock struct{}
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// AgentEventLifecycleRegister mocks base method.
func (m *MockConnection) AgentEventLifecycleRegister(callback libvirt.DomainEventAgentLifecycleCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentEventLifecycleRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// AgentEventLifecycleRegister indicates an expected call of AgentEventLifecycleRegister.
func (mr *MockConnectionMockRecorder) AgentEventLifecycleRegister(callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentEventLifecycleRegister", reflect.TypeOf((*MockConnection)(nil).AgentEventLifecycleRegister), callback)
}

// Close mocks base method.
func (m *MockConnection) Close() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// DomainDefineXML mocks base method.
func (m *MockConnection) DomainDefineXML(xml string) (VirDomain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainDefineXML", xml)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainDefineXML indicates an expected call of DomainDefineXML.
func (mr *MockConnectionMockRecorder) DomainDefineXML(xml any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainDefineXML", reflect.TypeOf((*MockConnection)(nil).DomainDefineXML), xml)
}

// DomainEventDeregister mocks base method.
func (m *MockConnection) DomainEventDeregister(registrationID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainEventDeregister", registrationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DomainEventDeregister indicates an expected call of DomainEventDeregister.
func (mr *MockConnectionMockRecorder) DomainEventDeregister(registrationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainEventDeregister", reflect.TypeOf((*MockConnection)(nil).DomainEventDeregister), registrationID)
}

// DomainEventDeviceAddedRegister mocks base method.
func (m *MockConnection) DomainEventDeviceAddedRegister(callback libvirt.DomainEventDeviceAddedCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainEventDeviceAddedRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// DomainEventDeviceAddedRegister indicates an expected call of DomainEventDeviceAddedRegister.
func (mr *MockConnectionMockRecorder) DomainEventDeviceAddedRegister(callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainEventDeviceAddedRegister", reflect.TypeOf((*MockConnection)(nil).DomainEventDeviceAddedRegister), callback)
}

// DomainEventDeviceRemovedRegister mocks base method.
func (m *MockConnection) DomainEventDeviceRemovedRegister(callback libvirt.DomainEventDeviceRemovedCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainEventDeviceRemovedRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// DomainEventDeviceRemovedRegister indicates an expected call of DomainEventDeviceRemovedRegister.
func (mr *MockConnectionMockRecorder) DomainEventDeviceRemovedRegister(callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainEventDeviceRemovedRegister", reflect.TypeOf((*MockConnection)(nil).DomainEventDeviceRemovedRegister), callback)
}

// DomainEventLifecycleRegister mocks base method.
func (m *MockConnection) DomainEventLifecycleRegister(callback libvirt.DomainEventLifecycleCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainEventLifecycleRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// DomainEventLifecycleRegister indicates an expected call of DomainEventLifecycleRegister.
func (mr *MockConnectionMockRecorder) DomainEventLifecycleRegister(callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainEventLifecycleRegister", reflect.TypeOf((*MockConnection)(nil).DomainEventLifecycleRegister), callback)
}

// DomainEventMemoryDeviceSizeChangeRegister mocks base method.
func (m *MockConnection) DomainEventMemoryDeviceSizeChangeRegister(callback libvirt.DomainEventMemoryDeviceSizeChangeCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainEventMemoryDeviceSizeChangeRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// DomainEventMemoryDeviceSizeChangeRegister indicates an expected call of DomainEventMemoryDeviceSizeChangeRegister.
func (mr *MockConnectionMockRecorder) DomainEventMemoryDeviceSizeChangeRegister(callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainEventMemoryDeviceSizeChangeRegister", reflect.TypeOf((*MockConnection)(nil).DomainEventMemoryDeviceSizeChangeRegister), callback)
}

// GetAllDomainStats mocks base method.
func (m *MockConnection) GetAllDomainStats(statsTypes libvirt.DomainStatsTypes, flags libvirt.ConnectGetAllDomainStatsFlags) ([]libvirt.DomainStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDomainStats", statsTypes, flags)
	ret0, _ := ret[0].([]libvirt.DomainStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDomainStats indicates an expected call of GetAllDomainStats.
func (mr *MockConnectionMockRecorder) GetAllDomainStats(statsTypes, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDomainStats", reflect.TypeOf((*MockConnection)(nil).GetAllDomainStats), statsTypes, flags)
}

// GetDomainDirtyRate mocks base method.
func (m *MockConnection) GetDomainDirtyRate(calculationDuration time.Duration, flags libvirt.DomainDirtyRateCalcFlags) ([]*stats.DomainStatsDirtyRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainDirtyRate", calculationDuration, flags)
	ret0, _ := ret[0].([]*stats.DomainStatsDirtyRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainDirtyRate indicates an expected call of GetDomainDirtyRate.
func (mr *MockConnectionMockRecorder) GetDomainDirtyRate(calculationDuration, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainDirtyRate", reflect.TypeOf((*MockConnection)(nil).GetDomainDirtyRate), calculationDuration, flags)
}

// GetDomainStats mocks base method.
func (m *MockConnection) GetDomainStats(statsTypes libvirt.DomainStatsTypes, l *stats.DomainJobInfo, flags libvirt.ConnectGetAllDomainStatsFlags) ([]*stats.DomainStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainStats", statsTypes, l, flags)
	ret0, _ := ret[0].([]*stats.DomainStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainStats indicates an expected call of GetDomainStats.
func (mr *MockConnectionMockRecorder) GetDomainStats(statsTypes, l, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainStats", reflect.TypeOf((*MockConnection)(nil).GetDomainStats), statsTypes, l, flags)
}

// GetQemuVersion mocks base method.
func (m *MockConnection) GetQemuVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQemuVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQemuVersion indicates an expected call of GetQemuVersion.
func (mr *MockConnectionMockRecorder) GetQemuVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQemuVersion", reflect.TypeOf((*MockConnection)(nil).GetQemuVersion))
}

// GetSEVInfo mocks base method.
func (m *MockConnection) GetSEVInfo() (*api.SEVNodeParameters, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSEVInfo")
	ret0, _ := ret[0].(*api.SEVNodeParameters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSEVInfo indicates an expected call of GetSEVInfo.
func (mr *MockConnectionMockRecorder) GetSEVInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSEVInfo", reflect.TypeOf((*MockConnection)(nil).GetSEVInfo))
}

// ListAllDomains mocks base method.
func (m *MockConnection) ListAllDomains(flags libvirt.ConnectListAllDomainsFlags) ([]VirDomain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllDomains", flags)
	ret0, _ := ret[0].([]VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllDomains indicates an expected call of ListAllDomains.
func (mr *MockConnectionMockRecorder) ListAllDomains(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDomains", reflect.TypeOf((*MockConnection)(nil).ListAllDomains), flags)
}

// LookupDomainByName mocks base method.
func (m *MockConnection) LookupDomainByName(name string) (VirDomain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupDomainByName", name)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupDomainByName indicates an expected call of LookupDomainByName.
func (mr *MockConnectionMockRecorder) LookupDomainByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupDomainByName", reflect.TypeOf((*MockConnection)(nil).LookupDomainByName), name)
}

// QemuAgentCommand mocks base method.
func (m *MockConnection) QemuAgentCommand(command, domainName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QemuAgentCommand", command, domainName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QemuAgentCommand indicates an expected call of QemuAgentCommand.
func (mr *MockConnectionMockRecorder) QemuAgentCommand(command, domainName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QemuAgentCommand", reflect.TypeOf((*MockConnection)(nil).QemuAgentCommand), command, domainName)
}

// SetReconnectChan mocks base method.
func (m *MockConnection) SetReconnectChan(reconnect chan bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReconnectChan", reconnect)
}

// SetReconnectChan indicates an expected call of SetReconnectChan.
func (mr *MockConnectionMockRecorder) SetReconnectChan(reconnect any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReconnectChan", reflect.TypeOf((*MockConnection)(nil).SetReconnectChan), reconnect)
}

// VolatileDomainEventDeviceRemovedRegister mocks base method.
func (m *MockConnection) VolatileDomainEventDeviceRemovedRegister(domain VirDomain, callback libvirt.DomainEventDeviceRemovedCallback) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolatileDomainEventDeviceRemovedRegister", domain, callback)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolatileDomainEventDeviceRemovedRegister indicates an expected call of VolatileDomainEventDeviceRemovedRegister.
func (mr *MockConnectionMockRecorder) VolatileDomainEventDeviceRemovedRegister(domain, callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolatileDomainEventDeviceRemovedRegister", reflect.TypeOf((*MockConnection)(nil).VolatileDomainEventDeviceRemovedRegister), domain, callback)
}

// MockStream is a mock of Stream interface.
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
	isgomock struct{}
}

// MockStreamMockRecorder is the mock recorder for MockStream.
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance.
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStream) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStreamMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close))
}

// Read mocks base method.
func (m *MockStream) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStreamMockRecorder) Read(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStream)(nil).Read), p)
}

// UnderlyingStream mocks base method.
func (m *MockStream) UnderlyingStream() *libvirt.Stream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnderlyingStream")
	ret0, _ := ret[0].(*libvirt.Stream)
	return ret0
}

// UnderlyingStream indicates an expected call of UnderlyingStream.
func (mr *MockStreamMockRecorder) UnderlyingStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnderlyingStream", reflect.TypeOf((*MockStream)(nil).UnderlyingStream))
}

// Write mocks base method.
func (m *MockStream) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStreamMockRecorder) Write(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStream)(nil).Write), p)
}

// MockVirDomain is a mock of VirDomain interface.
type MockVirDomain struct {
	ctrl     *gomock.Controller
	recorder *MockVirDomainMockRecorder
	isgomock struct{}
}

// MockVirDomainMockRecorder is the mock recorder for MockVirDomain.
type MockVirDomainMockRecorder struct {
	mock *MockVirDomain
}

// NewMockVirDomain creates a new mock instance.
func NewMockVirDomain(ctrl *gomock.Controller) *MockVirDomain {
	mock := &MockVirDomain{ctrl: ctrl}
	mock.recorder = &MockVirDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirDomain) EXPECT() *MockVirDomainMockRecorder {
	return m.recorder
}

// AbortJob mocks base method.
func (m *MockVirDomain) AbortJob() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortJob")
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortJob indicates an expected call of AbortJob.
func (mr *MockVirDomainMockRecorder) AbortJob() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortJob", reflect.TypeOf((*MockVirDomain)(nil).AbortJob))
}

// AttachDeviceFlags mocks base method.
func (m *MockVirDomain) AttachDeviceFlags(xml string, flags libvirt.DomainDeviceModifyFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachDeviceFlags", xml, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachDeviceFlags indicates an expected call of AttachDeviceFlags.
func (mr *MockVirDomainMockRecorder) AttachDeviceFlags(xml, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachDeviceFlags", reflect.TypeOf((*MockVirDomain)(nil).AttachDeviceFlags), xml, flags)
}

// AuthorizedSSHKeysSet mocks base method.
func (m *MockVirDomain) AuthorizedSSHKeysSet(user string, keys []string, flags libvirt.DomainAuthorizedSSHKeysFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizedSSHKeysSet", user, keys, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// AuthorizedSSHKeysSet indicates an expected call of AuthorizedSSHKeysSet.
func (mr *MockVirDomainMockRecorder) AuthorizedSSHKeysSet(user, keys, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizedSSHKeysSet", reflect.TypeOf((*MockVirDomain)(nil).AuthorizedSSHKeysSet), user, keys, flags)
}

// BlockResize mocks base method.
func (m *MockVirDomain) BlockResize(disk string, size uint64, flags libvirt.DomainBlockResizeFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockResize", disk, size, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// BlockResize indicates an expected call of BlockResize.
func (mr *MockVirDomainMockRecorder) BlockResize(disk, size, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockResize", reflect.TypeOf((*MockVirDomain)(nil).BlockResize), disk, size, flags)
}

// CoreDumpWithFormat mocks base method.
func (m *MockVirDomain) CoreDumpWithFormat(to string, format libvirt.DomainCoreDumpFormat, flags libvirt.DomainCoreDumpFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoreDumpWithFormat", to, format, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// CoreDumpWithFormat indicates an expected call of CoreDumpWithFormat.
func (mr *MockVirDomainMockRecorder) CoreDumpWithFormat(to, format, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoreDumpWithFormat", reflect.TypeOf((*MockVirDomain)(nil).CoreDumpWithFormat), to, format, flags)
}

// CreateWithFlags mocks base method.
func (m *MockVirDomain) CreateWithFlags(flags libvirt.DomainCreateFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWithFlags indicates an expected call of CreateWithFlags.
func (mr *MockVirDomainMockRecorder) CreateWithFlags(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithFlags", reflect.TypeOf((*MockVirDomain)(nil).CreateWithFlags), flags)
}

// DestroyFlags mocks base method.
func (m *MockVirDomain) DestroyFlags(flags libvirt.DomainDestroyFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyFlags indicates an expected call of DestroyFlags.
func (mr *MockVirDomainMockRecorder) DestroyFlags(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyFlags", reflect.TypeOf((*MockVirDomain)(nil).DestroyFlags), flags)
}

// DetachDeviceFlags mocks base method.
func (m *MockVirDomain) DetachDeviceFlags(xml string, flags libvirt.DomainDeviceModifyFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachDeviceFlags", xml, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachDeviceFlags indicates an expected call of DetachDeviceFlags.
func (mr *MockVirDomainMockRecorder) DetachDeviceFlags(xml, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachDeviceFlags", reflect.TypeOf((*MockVirDomain)(nil).DetachDeviceFlags), xml, flags)
}

// FSFreeze mocks base method.
func (m *MockVirDomain) FSFreeze(mounts []string, flags uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FSFreeze", mounts, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// FSFreeze indicates an expected call of FSFreeze.
func (mr *MockVirDomainMockRecorder) FSFreeze(mounts, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FSFreeze", reflect.TypeOf((*MockVirDomain)(nil).FSFreeze), mounts, flags)
}

// FSThaw mocks base method.
func (m *MockVirDomain) FSThaw(mounts []string, flags uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FSThaw", mounts, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// FSThaw indicates an expected call of FSThaw.
func (mr *MockVirDomainMockRecorder) FSThaw(mounts, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FSThaw", reflect.TypeOf((*MockVirDomain)(nil).FSThaw), mounts, flags)
}

// Free mocks base method.
func (m *MockVirDomain) Free() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Free")
	ret0, _ := ret[0].(error)
	return ret0
}

// Free indicates an expected call of Free.
func (mr *MockVirDomainMockRecorder) Free() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockVirDomain)(nil).Free))
}

// GetBlockInfo mocks base method.
func (m *MockVirDomain) GetBlockInfo(disk string, flags uint32) (*libvirt.DomainBlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockInfo", disk, flags)
	ret0, _ := ret[0].(*libvirt.DomainBlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockInfo indicates an expected call of GetBlockInfo.
func (mr *MockVirDomainMockRecorder) GetBlockInfo(disk, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockInfo", reflect.TypeOf((*MockVirDomain)(nil).GetBlockInfo), disk, flags)
}

// GetDiskErrors mocks base method.
func (m *MockVirDomain) GetDiskErrors(flags uint32) ([]libvirt.DomainDiskError, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskErrors", flags)
	ret0, _ := ret[0].([]libvirt.DomainDiskError)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskErrors indicates an expected call of GetDiskErrors.
func (mr *MockVirDomainMockRecorder) GetDiskErrors(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskErrors", reflect.TypeOf((*MockVirDomain)(nil).GetDiskErrors), flags)
}

// GetGuestInfo mocks base method.
func (m *MockVirDomain) GetGuestInfo(types libvirt.DomainGuestInfoTypes, flags uint32) (*libvirt.DomainGuestInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuestInfo", types, flags)
	ret0, _ := ret[0].(*libvirt.DomainGuestInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGuestInfo indicates an expected call of GetGuestInfo.
func (mr *MockVirDomainMockRecorder) GetGuestInfo(types, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuestInfo", reflect.TypeOf((*MockVirDomain)(nil).GetGuestInfo), types, flags)
}

// GetJobInfo mocks base method.
func (m *MockVirDomain) GetJobInfo() (*libvirt.DomainJobInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobInfo")
	ret0, _ := ret[0].(*libvirt.DomainJobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobInfo indicates an expected call of GetJobInfo.
func (mr *MockVirDomainMockRecorder) GetJobInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobInfo", reflect.TypeOf((*MockVirDomain)(nil).GetJobInfo))
}

// GetJobStats mocks base method.
func (m *MockVirDomain) GetJobStats(flags libvirt.DomainGetJobStatsFlags) (*libvirt.DomainJobInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobStats", flags)
	ret0, _ := ret[0].(*libvirt.DomainJobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobStats indicates an expected call of GetJobStats.
func (mr *MockVirDomainMockRecorder) GetJobStats(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobStats", reflect.TypeOf((*MockVirDomain)(nil).GetJobStats), flags)
}

// GetLaunchSecurityInfo mocks base method.
func (m *MockVirDomain) GetLaunchSecurityInfo(flags uint32) (*libvirt.DomainLaunchSecurityParameters, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLaunchSecurityInfo", flags)
	ret0, _ := ret[0].(*libvirt.DomainLaunchSecurityParameters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLaunchSecurityInfo indicates an expected call of GetLaunchSecurityInfo.
func (mr *MockVirDomainMockRecorder) GetLaunchSecurityInfo(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLaunchSecurityInfo", reflect.TypeOf((*MockVirDomain)(nil).GetLaunchSecurityInfo), flags)
}

// GetName mocks base method.
func (m *MockVirDomain) GetName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetName indicates an expected call of GetName.
func (mr *MockVirDomainMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockVirDomain)(nil).GetName))
}

// GetState mocks base method.
func (m *MockVirDomain) GetState() (libvirt.DomainState, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].(libvirt.DomainState)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetState indicates an expected call of GetState.
func (mr *MockVirDomainMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockVirDomain)(nil).GetState))
}

// GetUUIDString mocks base method.
func (m *MockVirDomain) GetUUIDString() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUUIDString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUUIDString indicates an expected call of GetUUIDString.
func (mr *MockVirDomainMockRecorder) GetUUIDString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUUIDString", reflect.TypeOf((*MockVirDomain)(nil).GetUUIDString))
}

// GetXMLDesc mocks base method.
func (m *MockVirDomain) GetXMLDesc(flags libvirt.DomainXMLFlags) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetXMLDesc", flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetXMLDesc indicates an expected call of GetXMLDesc.
func (mr *MockVirDomainMockRecorder) GetXMLDesc(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXMLDesc", reflect.TypeOf((*MockVirDomain)(nil).GetXMLDesc), flags)
}

// MemoryStats mocks base method.
func (m *MockVirDomain) MemoryStats(nrStats, flags uint32) ([]libvirt.DomainMemoryStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemoryStats", nrStats, flags)
	ret0, _ := ret[0].([]libvirt.DomainMemoryStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemoryStats indicates an expected call of MemoryStats.
func (mr *MockVirDomainMockRecorder) MemoryStats(nrStats, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemoryStats", reflect.TypeOf((*MockVirDomain)(nil).MemoryStats), nrStats, flags)
}

// MigrateStartPostCopy mocks base method.
func (m *MockVirDomain) MigrateStartPostCopy(flags uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateStartPostCopy", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateStartPostCopy indicates an expected call of MigrateStartPostCopy.
func (mr *MockVirDomainMockRecorder) MigrateStartPostCopy(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateStartPostCopy", reflect.TypeOf((*MockVirDomain)(nil).MigrateStartPostCopy), flags)
}

// MigrateToURI3 mocks base method.
func (m *MockVirDomain) MigrateToURI3(arg0 string, arg1 *libvirt.DomainMigrateParameters, arg2 libvirt.DomainMigrateFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateToURI3", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateToURI3 indicates an expected call of MigrateToURI3.
func (mr *MockVirDomainMockRecorder) MigrateToURI3(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateToURI3", reflect.TypeOf((*MockVirDomain)(nil).MigrateToURI3), arg0, arg1, arg2)
}

// PinEmulator mocks base method.
func (m *MockVirDomain) PinEmulator(cpumap []bool, flags libvirt.DomainModificationImpact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinEmulator", cpumap, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// PinEmulator indicates an expected call of PinEmulator.
func (mr *MockVirDomainMockRecorder) PinEmulator(cpumap, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinEmulator", reflect.TypeOf((*MockVirDomain)(nil).PinEmulator), cpumap, flags)
}

// PinVcpuFlags mocks base method.
func (m *MockVirDomain) PinVcpuFlags(vcpu uint, cpuMap []bool, flags libvirt.DomainModificationImpact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinVcpuFlags", vcpu, cpuMap, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// PinVcpuFlags indicates an expected call of PinVcpuFlags.
func (mr *MockVirDomainMockRecorder) PinVcpuFlags(vcpu, cpuMap, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinVcpuFlags", reflect.TypeOf((*MockVirDomain)(nil).PinVcpuFlags), vcpu, cpuMap, flags)
}

// Reboot mocks base method.
func (m *MockVirDomain) Reboot(flags libvirt.DomainRebootFlagValues) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reboot", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reboot indicates an expected call of Reboot.
func (mr *MockVirDomainMockRecorder) Reboot(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reboot", reflect.TypeOf((*MockVirDomain)(nil).Reboot), flags)
}

// Reset mocks base method.
func (m *MockVirDomain) Reset(flags uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockVirDomainMockRecorder) Reset(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockVirDomain)(nil).Reset), flags)
}

// Resume mocks base method.
func (m *MockVirDomain) Resume() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resume")
	ret0, _ := ret[0].(error)
	return ret0
}

// Resume indicates an expected call of Resume.
func (mr *MockVirDomainMockRecorder) Resume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resume", reflect.TypeOf((*MockVirDomain)(nil).Resume))
}

// SetLaunchSecurityState mocks base method.
func (m *MockVirDomain) SetLaunchSecurityState(params *libvirt.DomainLaunchSecurityStateParameters, flags uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLaunchSecurityState", params, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLaunchSecurityState indicates an expected call of SetLaunchSecurityState.
func (mr *MockVirDomainMockRecorder) SetLaunchSecurityState(params, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLaunchSecurityState", reflect.TypeOf((*MockVirDomain)(nil).SetLaunchSecurityState), params, flags)
}

// SetTime mocks base method.
func (m *MockVirDomain) SetTime(secs int64, nsecs uint, flags libvirt.DomainSetTimeFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTime", secs, nsecs, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTime indicates an expected call of SetTime.
func (mr *MockVirDomainMockRecorder) SetTime(secs, nsecs, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTime", reflect.TypeOf((*MockVirDomain)(nil).SetTime), secs, nsecs, flags)
}

// SetUserPassword mocks base method.
func (m *MockVirDomain) SetUserPassword(user, password string, flags libvirt.DomainSetUserPasswordFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserPassword", user, password, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserPassword indicates an expected call of SetUserPassword.
func (mr *MockVirDomainMockRecorder) SetUserPassword(user, password, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserPassword", reflect.TypeOf((*MockVirDomain)(nil).SetUserPassword), user, password, flags)
}

// SetVcpusFlags mocks base method.
func (m *MockVirDomain) SetVcpusFlags(vcpu uint, flags libvirt.DomainVcpuFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVcpusFlags", vcpu, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVcpusFlags indicates an expected call of SetVcpusFlags.
func (mr *MockVirDomainMockRecorder) SetVcpusFlags(vcpu, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVcpusFlags", reflect.TypeOf((*MockVirDomain)(nil).SetVcpusFlags), vcpu, flags)
}

// ShutdownFlags mocks base method.
func (m *MockVirDomain) ShutdownFlags(flags libvirt.DomainShutdownFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShutdownFlags indicates an expected call of ShutdownFlags.
func (mr *MockVirDomainMockRecorder) ShutdownFlags(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownFlags", reflect.TypeOf((*MockVirDomain)(nil).ShutdownFlags), flags)
}

// Suspend mocks base method.
func (m *MockVirDomain) Suspend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspend")
	ret0, _ := ret[0].(error)
	return ret0
}

// Suspend indicates an expected call of Suspend.
func (mr *MockVirDomainMockRecorder) Suspend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspend", reflect.TypeOf((*MockVirDomain)(nil).Suspend))
}

// UndefineFlags mocks base method.
func (m *MockVirDomain) UndefineFlags(flags libvirt.DomainUndefineFlagsValues) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndefineFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// UndefineFlags indicates an expected call of UndefineFlags.
func (mr *MockVirDomainMockRecorder) UndefineFlags(flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndefineFlags", reflect.TypeOf((*MockVirDomain)(nil).UndefineFlags), flags)
}

// UpdateDeviceFlags mocks base method.
func (m *MockVirDomain) UpdateDeviceFlags(xml string, flags libvirt.DomainDeviceModifyFlags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeviceFlags", xml, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeviceFlags indicates an expected call of UpdateDeviceFlags.
func (mr *MockVirDomainMockRecorder) UpdateDeviceFlags(xml, flags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeviceFlags", reflect.TypeOf((*MockVirDomain)(nil).UpdateDeviceFlags), xml, flags)
}
