package mocks

import (
	http "net/http"
	reflect "reflect"
	domain "utest/core/domain"
	dto "utest/core/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockFeiraService mock FeiraService interface.
type MockFeiraService struct {
	ctrl     *gomock.Controller
	recorder *MockFeiraServiceMockRecorder
}

// MockFeiraServiceMockRecorder mock for MockFeiraService.
type MockFeiraServiceMockRecorder struct {
	mock *MockFeiraService
}

// NewMockFeiraService nova istancia.
func NewMockFeiraService(ctrl *gomock.Controller) *MockFeiraService {
	mock := &MockFeiraService{ctrl: ctrl}
	mock.recorder = &MockFeiraServiceMockRecorder{mock}
	return mock
}

// EXPECT retorna um objeto.
func (m *MockFeiraService) EXPECT() *MockFeiraServiceMockRecorder {
	return m.recorder
}

// Create mocks.
func (m *MockFeiraService) Create(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", response, request)
}

// Create indicates an expected call of Create.
func (mr *MockFeiraServiceMockRecorder) Create(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFeiraService)(nil).Create), response, request)
}

// Update mocks.
func (m *MockFeiraService) Update(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", response, request)
}

// Update indicates an expected call of Update.
func (mr *MockFeiraServiceMockRecorder) Update(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFeiraService)(nil).Update), response, request)
}

// Delete mocks.
func (m *MockFeiraService) Delete(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", response, request)
}

// Delete indicates an expected call of Delete.
func (mr *MockFeiraServiceMockRecorder) Delete(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeiraService)(nil).Delete), response, request)
}

// MockFeiraUseCase is a mock of FeiraUseCase interface.
type MockFeiraUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockFeiraUseCaseMockRecorder
}

// MockFeiraUseCaseMockRecorder is the mock recorder for MockFeiraUseCase.
type MockFeiraUseCaseMockRecorder struct {
	mock *MockFeiraUseCase
}

// NewMockFeiraUseCase creates a new mock instance.
func NewMockFeiraUseCase(ctrl *gomock.Controller) *MockFeiraUseCase {
	mock := &MockFeiraUseCase{ctrl: ctrl}
	mock.recorder = &MockFeiraUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeiraUseCase) EXPECT() *MockFeiraUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFeiraUseCase) Create(FeiraRequest *dto.CreateFeiraRequest) (*domain.Feira, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", FeiraRequest)
	ret0, _ := ret[0].(*domain.Feira)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFeiraUseCaseMockRecorder) Create(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFeiraUseCase)(nil).Create), FeiraRequest)
}


// Update mocks base method.
func (m *MockFeiraUseCase) Update(FeiraRequest *dto.UpdateFeiraRequest) (*domain.Feira, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", FeiraRequest)
	ret0, _ := ret[0].(*domain.Feira)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Create.
func (mr *MockFeiraUseCaseMockRecorder) Update(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFeiraUseCase)(nil).Update), FeiraRequest)
}


// GetNome mocks base method.
func (m *MockFeiraUseCase) GetNome(FeiraRequest *dto.GetNomeRequest) (*domain.Feira, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNome", FeiraRequest)
	ret0, _ := ret[0].(*domain.Feira)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNome indicates an expected call of Create.
func (mr *MockFeiraUseCaseMockRecorder) GetNome(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNome", reflect.TypeOf((*MockFeiraUseCase)(nil).Update), FeiraRequest)
}

// Delete mocks base method.
func (m *MockFeiraUseCase) Delete(FeiraRequest *dto.DeleteFeiraRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", FeiraRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Create.
func (mr *MockFeiraUseCaseMockRecorder) Delete(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeiraUseCase)(nil).Delete), FeiraRequest)
}


// MockFeiraRepository is a mock of FeiraRepository interface.
type MockFeiraRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFeiraRepositoryMockRecorder
}

// MockFeiraRepositoryMockRecorder is the mock recorder for MockFeiraRepository.
type MockFeiraRepositoryMockRecorder struct {
	mock *MockFeiraRepository
}

// NewMockFeiraRepository creates a new mock instance.
func NewMockFeiraRepository(ctrl *gomock.Controller) *MockFeiraRepository {
	mock := &MockFeiraRepository{ctrl: ctrl}
	mock.recorder = &MockFeiraRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeiraRepository) EXPECT() *MockFeiraRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFeiraRepository) Create(FeiraRequest *dto.CreateFeiraRequest) (*domain.Feira, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", FeiraRequest)
	ret0, _ := ret[0].(*domain.Feira)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFeiraRepositoryMockRecorder) Create(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFeiraRepository)(nil).Create), FeiraRequest)
}

// Create mocks base method.
func (m *MockFeiraRepository) GetNome(FeiraRequest *dto.GetNomeRequest) (*domain.Feira, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNome", FeiraRequest)
	ret0, _ := ret[0].(*domain.Feira)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFeiraRepositoryMockRecorder) GetNome(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNome", reflect.TypeOf((*MockFeiraRepository)(nil).GetNome), FeiraRequest)
}



// Update mocks base method.
func (m *MockFeiraRepository) Update(FeiraRequest *dto.UpdateFeiraRequest) (*domain.Feira, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", FeiraRequest)
	ret0, _ := ret[0].(*domain.Feira)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockFeiraRepositoryMockRecorder) Update(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFeiraRepository)(nil).Update), FeiraRequest)
}

// Delete mocks base method.
func (m *MockFeiraRepository) Delete(FeiraRequest *dto.DeleteFeiraRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", FeiraRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFeiraRepositoryMockRecorder) Delete(FeiraRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeiraRepository)(nil).Delete), FeiraRequest)
}