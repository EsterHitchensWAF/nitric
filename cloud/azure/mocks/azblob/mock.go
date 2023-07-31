// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nitrictech/nitric/cloud/azure/runtime/storage/iface (interfaces: AzblobServiceUrlIface,AzblobContainerUrlIface,AzblobBlockBlobUrlIface,AzblobDownloadResponse)

// Package mock_iface is a generated GoMock package.
package mock_iface

import (
	context "context"
	io "io"
	url "net/url"
	reflect "reflect"

	azblob "github.com/Azure/azure-storage-blob-go/azblob"
	gomock "github.com/golang/mock/gomock"
	azblob_service_iface "github.com/nitrictech/nitric/cloud/azure/runtime/storage/iface"
)

// MockAzblobServiceUrlIface is a mock of AzblobServiceUrlIface interface.
type MockAzblobServiceUrlIface struct {
	ctrl     *gomock.Controller
	recorder *MockAzblobServiceUrlIfaceMockRecorder
}

// MockAzblobServiceUrlIfaceMockRecorder is the mock recorder for MockAzblobServiceUrlIface.
type MockAzblobServiceUrlIfaceMockRecorder struct {
	mock *MockAzblobServiceUrlIface
}

// NewMockAzblobServiceUrlIface creates a new mock instance.
func NewMockAzblobServiceUrlIface(ctrl *gomock.Controller) *MockAzblobServiceUrlIface {
	mock := &MockAzblobServiceUrlIface{ctrl: ctrl}
	mock.recorder = &MockAzblobServiceUrlIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzblobServiceUrlIface) EXPECT() *MockAzblobServiceUrlIfaceMockRecorder {
	return m.recorder
}

// GetUserDelegationCredential mocks base method.
func (m *MockAzblobServiceUrlIface) GetUserDelegationCredential(arg0 context.Context, arg1 azblob.KeyInfo, arg2 *int32, arg3 *string) (azblob.StorageAccountCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDelegationCredential", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(azblob.StorageAccountCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDelegationCredential indicates an expected call of GetUserDelegationCredential.
func (mr *MockAzblobServiceUrlIfaceMockRecorder) GetUserDelegationCredential(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDelegationCredential", reflect.TypeOf((*MockAzblobServiceUrlIface)(nil).GetUserDelegationCredential), arg0, arg1, arg2, arg3)
}

// NewContainerURL mocks base method.
func (m *MockAzblobServiceUrlIface) NewContainerURL(arg0 string) azblob_service_iface.AzblobContainerUrlIface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewContainerURL", arg0)
	ret0, _ := ret[0].(azblob_service_iface.AzblobContainerUrlIface)
	return ret0
}

// NewContainerURL indicates an expected call of NewContainerURL.
func (mr *MockAzblobServiceUrlIfaceMockRecorder) NewContainerURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewContainerURL", reflect.TypeOf((*MockAzblobServiceUrlIface)(nil).NewContainerURL), arg0)
}

// MockAzblobContainerUrlIface is a mock of AzblobContainerUrlIface interface.
type MockAzblobContainerUrlIface struct {
	ctrl     *gomock.Controller
	recorder *MockAzblobContainerUrlIfaceMockRecorder
}

// MockAzblobContainerUrlIfaceMockRecorder is the mock recorder for MockAzblobContainerUrlIface.
type MockAzblobContainerUrlIfaceMockRecorder struct {
	mock *MockAzblobContainerUrlIface
}

// NewMockAzblobContainerUrlIface creates a new mock instance.
func NewMockAzblobContainerUrlIface(ctrl *gomock.Controller) *MockAzblobContainerUrlIface {
	mock := &MockAzblobContainerUrlIface{ctrl: ctrl}
	mock.recorder = &MockAzblobContainerUrlIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzblobContainerUrlIface) EXPECT() *MockAzblobContainerUrlIfaceMockRecorder {
	return m.recorder
}

// ListBlobsFlatSegment mocks base method.
func (m *MockAzblobContainerUrlIface) ListBlobsFlatSegment(arg0 context.Context, arg1 azblob.Marker, arg2 azblob.ListBlobsSegmentOptions) (*azblob.ListBlobsFlatSegmentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBlobsFlatSegment", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azblob.ListBlobsFlatSegmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlobsFlatSegment indicates an expected call of ListBlobsFlatSegment.
func (mr *MockAzblobContainerUrlIfaceMockRecorder) ListBlobsFlatSegment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlobsFlatSegment", reflect.TypeOf((*MockAzblobContainerUrlIface)(nil).ListBlobsFlatSegment), arg0, arg1, arg2)
}

// NewBlockBlobURL mocks base method.
func (m *MockAzblobContainerUrlIface) NewBlockBlobURL(arg0 string) azblob_service_iface.AzblobBlockBlobUrlIface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBlockBlobURL", arg0)
	ret0, _ := ret[0].(azblob_service_iface.AzblobBlockBlobUrlIface)
	return ret0
}

// NewBlockBlobURL indicates an expected call of NewBlockBlobURL.
func (mr *MockAzblobContainerUrlIfaceMockRecorder) NewBlockBlobURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBlockBlobURL", reflect.TypeOf((*MockAzblobContainerUrlIface)(nil).NewBlockBlobURL), arg0)
}

// MockAzblobBlockBlobUrlIface is a mock of AzblobBlockBlobUrlIface interface.
type MockAzblobBlockBlobUrlIface struct {
	ctrl     *gomock.Controller
	recorder *MockAzblobBlockBlobUrlIfaceMockRecorder
}

// MockAzblobBlockBlobUrlIfaceMockRecorder is the mock recorder for MockAzblobBlockBlobUrlIface.
type MockAzblobBlockBlobUrlIfaceMockRecorder struct {
	mock *MockAzblobBlockBlobUrlIface
}

// NewMockAzblobBlockBlobUrlIface creates a new mock instance.
func NewMockAzblobBlockBlobUrlIface(ctrl *gomock.Controller) *MockAzblobBlockBlobUrlIface {
	mock := &MockAzblobBlockBlobUrlIface{ctrl: ctrl}
	mock.recorder = &MockAzblobBlockBlobUrlIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzblobBlockBlobUrlIface) EXPECT() *MockAzblobBlockBlobUrlIfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAzblobBlockBlobUrlIface) Delete(arg0 context.Context, arg1 azblob.DeleteSnapshotsOptionType, arg2 azblob.BlobAccessConditions) (*azblob.BlobDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azblob.BlobDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAzblobBlockBlobUrlIfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAzblobBlockBlobUrlIface)(nil).Delete), arg0, arg1, arg2)
}

// Download mocks base method.
func (m *MockAzblobBlockBlobUrlIface) Download(arg0 context.Context, arg1, arg2 int64, arg3 azblob.BlobAccessConditions, arg4 bool, arg5 azblob.ClientProvidedKeyOptions) (azblob_service_iface.AzblobDownloadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(azblob_service_iface.AzblobDownloadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockAzblobBlockBlobUrlIfaceMockRecorder) Download(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockAzblobBlockBlobUrlIface)(nil).Download), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetProperties mocks base method.
func (m *MockAzblobBlockBlobUrlIface) GetProperties(arg0 context.Context, arg1 azblob.BlobAccessConditions, arg2 azblob.ClientProvidedKeyOptions) (*azblob.BlobGetPropertiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperties", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azblob.BlobGetPropertiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProperties indicates an expected call of GetProperties.
func (mr *MockAzblobBlockBlobUrlIfaceMockRecorder) GetProperties(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperties", reflect.TypeOf((*MockAzblobBlockBlobUrlIface)(nil).GetProperties), arg0, arg1, arg2)
}

// Upload mocks base method.
func (m *MockAzblobBlockBlobUrlIface) Upload(arg0 context.Context, arg1 io.ReadSeeker, arg2 azblob.BlobHTTPHeaders, arg3 azblob.Metadata, arg4 azblob.BlobAccessConditions, arg5 azblob.AccessTierType, arg6 azblob.BlobTagsMap, arg7 azblob.ClientProvidedKeyOptions) (*azblob.BlockBlobUploadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(*azblob.BlockBlobUploadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockAzblobBlockBlobUrlIfaceMockRecorder) Upload(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockAzblobBlockBlobUrlIface)(nil).Upload), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// Url mocks base method.
func (m *MockAzblobBlockBlobUrlIface) Url() url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Url")
	ret0, _ := ret[0].(url.URL)
	return ret0
}

// Url indicates an expected call of Url.
func (mr *MockAzblobBlockBlobUrlIfaceMockRecorder) Url() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Url", reflect.TypeOf((*MockAzblobBlockBlobUrlIface)(nil).Url))
}

// MockAzblobDownloadResponse is a mock of AzblobDownloadResponse interface.
type MockAzblobDownloadResponse struct {
	ctrl     *gomock.Controller
	recorder *MockAzblobDownloadResponseMockRecorder
}

// MockAzblobDownloadResponseMockRecorder is the mock recorder for MockAzblobDownloadResponse.
type MockAzblobDownloadResponseMockRecorder struct {
	mock *MockAzblobDownloadResponse
}

// NewMockAzblobDownloadResponse creates a new mock instance.
func NewMockAzblobDownloadResponse(ctrl *gomock.Controller) *MockAzblobDownloadResponse {
	mock := &MockAzblobDownloadResponse{ctrl: ctrl}
	mock.recorder = &MockAzblobDownloadResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzblobDownloadResponse) EXPECT() *MockAzblobDownloadResponseMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockAzblobDownloadResponse) Body(arg0 azblob.RetryReaderOptions) io.ReadCloser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockAzblobDownloadResponseMockRecorder) Body(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockAzblobDownloadResponse)(nil).Body), arg0)
}
