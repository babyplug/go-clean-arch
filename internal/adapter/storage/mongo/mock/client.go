// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -source=client.go -destination=mock/client.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	mongo "go-hexagonal-architecture/internal/adapter/storage/mongo"
	reflect "reflect"
	time "time"

	bson "go.mongodb.org/mongo-driver/bson"
	mongo0 "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	readconcern "go.mongodb.org/mongo-driver/mongo/readconcern"
	readpref "go.mongodb.org/mongo-driver/mongo/readpref"
	writeconcern "go.mongodb.org/mongo-driver/mongo/writeconcern"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockClient) Connect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockClientMockRecorder) Connect(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClient)(nil).Connect), ctx)
}

// Database mocks base method.
func (m *MockClient) Database(name string, opts ...*options.DatabaseOptions) mongo.Database {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Database", varargs...)
	ret0, _ := ret[0].(mongo.Database)
	return ret0
}

// Database indicates an expected call of Database.
func (mr *MockClientMockRecorder) Database(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Database", reflect.TypeOf((*MockClient)(nil).Database), varargs...)
}

// Disconnect mocks base method.
func (m *MockClient) Disconnect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockClientMockRecorder) Disconnect(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockClient)(nil).Disconnect), ctx)
}

// ListDatabaseNames mocks base method.
func (m *MockClient) ListDatabaseNames(ctx context.Context, filter any, opts ...*options.ListDatabasesOptions) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDatabaseNames", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDatabaseNames indicates an expected call of ListDatabaseNames.
func (mr *MockClientMockRecorder) ListDatabaseNames(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabaseNames", reflect.TypeOf((*MockClient)(nil).ListDatabaseNames), varargs...)
}

// ListDatabases mocks base method.
func (m *MockClient) ListDatabases(ctx context.Context, filter any, opts ...*options.ListDatabasesOptions) (mongo0.ListDatabasesResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDatabases", varargs...)
	ret0, _ := ret[0].(mongo0.ListDatabasesResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDatabases indicates an expected call of ListDatabases.
func (mr *MockClientMockRecorder) ListDatabases(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabases", reflect.TypeOf((*MockClient)(nil).ListDatabases), varargs...)
}

// NumberSessionsInProgress mocks base method.
func (m *MockClient) NumberSessionsInProgress() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumberSessionsInProgress")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumberSessionsInProgress indicates an expected call of NumberSessionsInProgress.
func (mr *MockClientMockRecorder) NumberSessionsInProgress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumberSessionsInProgress", reflect.TypeOf((*MockClient)(nil).NumberSessionsInProgress))
}

// Ping mocks base method.
func (m *MockClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, rp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockClientMockRecorder) Ping(ctx, rp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockClient)(nil).Ping), ctx, rp)
}

// StartSession mocks base method.
func (m *MockClient) StartSession(opts ...*options.SessionOptions) (mongo0.Session, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartSession", varargs...)
	ret0, _ := ret[0].(mongo0.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSession indicates an expected call of StartSession.
func (mr *MockClientMockRecorder) StartSession(opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSession", reflect.TypeOf((*MockClient)(nil).StartSession), opts...)
}

// Timeout mocks base method.
func (m *MockClient) Timeout() *time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(*time.Duration)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MockClientMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockClient)(nil).Timeout))
}

// UseSession mocks base method.
func (m *MockClient) UseSession(ctx context.Context, fn func(mongo0.SessionContext) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseSession", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseSession indicates an expected call of UseSession.
func (mr *MockClientMockRecorder) UseSession(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseSession", reflect.TypeOf((*MockClient)(nil).UseSession), ctx, fn)
}

// UseSessionWithOptions mocks base method.
func (m *MockClient) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo0.SessionContext) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseSessionWithOptions", ctx, opts, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseSessionWithOptions indicates an expected call of UseSessionWithOptions.
func (mr *MockClientMockRecorder) UseSessionWithOptions(ctx, opts, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseSessionWithOptions", reflect.TypeOf((*MockClient)(nil).UseSessionWithOptions), ctx, opts, fn)
}

// Watch mocks base method.
func (m *MockClient) Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo0.ChangeStream, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(*mongo0.ChangeStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockClientMockRecorder) Watch(ctx, pipeline any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClient)(nil).Watch), varargs...)
}

// MockCursor is a mock of Cursor interface.
type MockCursor struct {
	ctrl     *gomock.Controller
	recorder *MockCursorMockRecorder
	isgomock struct{}
}

// MockCursorMockRecorder is the mock recorder for MockCursor.
type MockCursorMockRecorder struct {
	mock *MockCursor
}

// NewMockCursor creates a new mock instance.
func NewMockCursor(ctrl *gomock.Controller) *MockCursor {
	mock := &MockCursor{ctrl: ctrl}
	mock.recorder = &MockCursorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCursor) EXPECT() *MockCursorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCursor) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCursorMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCursor)(nil).Close), arg0)
}

// Decode mocks base method.
func (m *MockCursor) Decode(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockCursorMockRecorder) Decode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockCursor)(nil).Decode), arg0)
}

// Next mocks base method.
func (m *MockCursor) Next(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockCursorMockRecorder) Next(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockCursor)(nil).Next), arg0)
}

// MockSingleResult is a mock of SingleResult interface.
type MockSingleResult struct {
	ctrl     *gomock.Controller
	recorder *MockSingleResultMockRecorder
	isgomock struct{}
}

// MockSingleResultMockRecorder is the mock recorder for MockSingleResult.
type MockSingleResultMockRecorder struct {
	mock *MockSingleResult
}

// NewMockSingleResult creates a new mock instance.
func NewMockSingleResult(ctrl *gomock.Controller) *MockSingleResult {
	mock := &MockSingleResult{ctrl: ctrl}
	mock.recorder = &MockSingleResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSingleResult) EXPECT() *MockSingleResultMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockSingleResult) Decode(v any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockSingleResultMockRecorder) Decode(v any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockSingleResult)(nil).Decode), v)
}

// DecodeBytes mocks base method.
func (m *MockSingleResult) DecodeBytes() (bson.Raw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeBytes")
	ret0, _ := ret[0].(bson.Raw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeBytes indicates an expected call of DecodeBytes.
func (mr *MockSingleResultMockRecorder) DecodeBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeBytes", reflect.TypeOf((*MockSingleResult)(nil).DecodeBytes))
}

// Err mocks base method.
func (m *MockSingleResult) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockSingleResultMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSingleResult)(nil).Err))
}

// Raw mocks base method.
func (m *MockSingleResult) Raw() (bson.Raw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(bson.Raw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Raw indicates an expected call of Raw.
func (mr *MockSingleResultMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockSingleResult)(nil).Raw))
}

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
	isgomock struct{}
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockDatabase) Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Aggregate", varargs...)
	ret0, _ := ret[0].(mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockDatabaseMockRecorder) Aggregate(ctx, pipeline any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockDatabase)(nil).Aggregate), varargs...)
}

// Client mocks base method.
func (m *MockDatabase) Client() *mongo0.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*mongo0.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockDatabaseMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockDatabase)(nil).Client))
}

// Collection mocks base method.
func (m *MockDatabase) Collection(name string, opts ...*options.CollectionOptions) mongo.Collection {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Collection", varargs...)
	ret0, _ := ret[0].(mongo.Collection)
	return ret0
}

// Collection indicates an expected call of Collection.
func (mr *MockDatabaseMockRecorder) Collection(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockDatabase)(nil).Collection), varargs...)
}

// CreateCollection mocks base method.
func (m *MockDatabase) CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCollection", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCollection indicates an expected call of CreateCollection.
func (mr *MockDatabaseMockRecorder) CreateCollection(ctx, name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCollection", reflect.TypeOf((*MockDatabase)(nil).CreateCollection), varargs...)
}

// CreateView mocks base method.
func (m *MockDatabase) CreateView(ctx context.Context, viewName, viewOn string, pipeline any, opts ...*options.CreateViewOptions) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, viewName, viewOn, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateView", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateView indicates an expected call of CreateView.
func (mr *MockDatabaseMockRecorder) CreateView(ctx, viewName, viewOn, pipeline any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, viewName, viewOn, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateView", reflect.TypeOf((*MockDatabase)(nil).CreateView), varargs...)
}

// Drop mocks base method.
func (m *MockDatabase) Drop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockDatabaseMockRecorder) Drop(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockDatabase)(nil).Drop), ctx)
}

// ListCollectionNames mocks base method.
func (m *MockDatabase) ListCollectionNames(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCollectionNames", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCollectionNames indicates an expected call of ListCollectionNames.
func (mr *MockDatabaseMockRecorder) ListCollectionNames(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCollectionNames", reflect.TypeOf((*MockDatabase)(nil).ListCollectionNames), varargs...)
}

// ListCollectionSpecifications mocks base method.
func (m *MockDatabase) ListCollectionSpecifications(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) ([]*mongo0.CollectionSpecification, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCollectionSpecifications", varargs...)
	ret0, _ := ret[0].([]*mongo0.CollectionSpecification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCollectionSpecifications indicates an expected call of ListCollectionSpecifications.
func (mr *MockDatabaseMockRecorder) ListCollectionSpecifications(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCollectionSpecifications", reflect.TypeOf((*MockDatabase)(nil).ListCollectionSpecifications), varargs...)
}

// ListCollections mocks base method.
func (m *MockDatabase) ListCollections(ctx context.Context, filter any, opts ...*options.ListCollectionsOptions) (mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCollections", varargs...)
	ret0, _ := ret[0].(mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCollections indicates an expected call of ListCollections.
func (mr *MockDatabaseMockRecorder) ListCollections(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCollections", reflect.TypeOf((*MockDatabase)(nil).ListCollections), varargs...)
}

// Name mocks base method.
func (m *MockDatabase) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockDatabaseMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDatabase)(nil).Name))
}

// ReadConcern mocks base method.
func (m *MockDatabase) ReadConcern() *readconcern.ReadConcern {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadConcern")
	ret0, _ := ret[0].(*readconcern.ReadConcern)
	return ret0
}

// ReadConcern indicates an expected call of ReadConcern.
func (mr *MockDatabaseMockRecorder) ReadConcern() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConcern", reflect.TypeOf((*MockDatabase)(nil).ReadConcern))
}

// ReadPreference mocks base method.
func (m *MockDatabase) ReadPreference() *readpref.ReadPref {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPreference")
	ret0, _ := ret[0].(*readpref.ReadPref)
	return ret0
}

// ReadPreference indicates an expected call of ReadPreference.
func (mr *MockDatabaseMockRecorder) ReadPreference() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPreference", reflect.TypeOf((*MockDatabase)(nil).ReadPreference))
}

// RunCommand mocks base method.
func (m *MockDatabase) RunCommand(ctx context.Context, runCommand any, opts ...*options.RunCmdOptions) *mongo0.SingleResult {
	m.ctrl.T.Helper()
	varargs := []any{ctx, runCommand}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommand", varargs...)
	ret0, _ := ret[0].(*mongo0.SingleResult)
	return ret0
}

// RunCommand indicates an expected call of RunCommand.
func (mr *MockDatabaseMockRecorder) RunCommand(ctx, runCommand any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, runCommand}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockDatabase)(nil).RunCommand), varargs...)
}

// RunCommandCursor mocks base method.
func (m *MockDatabase) RunCommandCursor(ctx context.Context, runCommand any, opts ...*options.RunCmdOptions) (mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, runCommand}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommandCursor", varargs...)
	ret0, _ := ret[0].(mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCommandCursor indicates an expected call of RunCommandCursor.
func (mr *MockDatabaseMockRecorder) RunCommandCursor(ctx, runCommand any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, runCommand}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommandCursor", reflect.TypeOf((*MockDatabase)(nil).RunCommandCursor), varargs...)
}

// Watch mocks base method.
func (m *MockDatabase) Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo0.ChangeStream, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(*mongo0.ChangeStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockDatabaseMockRecorder) Watch(ctx, pipeline any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockDatabase)(nil).Watch), varargs...)
}

// WriteConcern mocks base method.
func (m *MockDatabase) WriteConcern() *writeconcern.WriteConcern {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteConcern")
	ret0, _ := ret[0].(*writeconcern.WriteConcern)
	return ret0
}

// WriteConcern indicates an expected call of WriteConcern.
func (mr *MockDatabaseMockRecorder) WriteConcern() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteConcern", reflect.TypeOf((*MockDatabase)(nil).WriteConcern))
}

// MockCollection is a mock of Collection interface.
type MockCollection struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionMockRecorder
	isgomock struct{}
}

// MockCollectionMockRecorder is the mock recorder for MockCollection.
type MockCollectionMockRecorder struct {
	mock *MockCollection
}

// NewMockCollection creates a new mock instance.
func NewMockCollection(ctrl *gomock.Controller) *MockCollection {
	mock := &MockCollection{ctrl: ctrl}
	mock.recorder = &MockCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollection) EXPECT() *MockCollectionMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockCollection) Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Aggregate", varargs...)
	ret0, _ := ret[0].(mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockCollectionMockRecorder) Aggregate(ctx, pipeline any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockCollection)(nil).Aggregate), varargs...)
}

// BulkWrite mocks base method.
func (m *MockCollection) BulkWrite(ctx context.Context, models []mongo0.WriteModel, opts ...*options.BulkWriteOptions) (*mongo0.BulkWriteResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, models}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkWrite", varargs...)
	ret0, _ := ret[0].(*mongo0.BulkWriteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkWrite indicates an expected call of BulkWrite.
func (mr *MockCollectionMockRecorder) BulkWrite(ctx, models any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, models}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkWrite", reflect.TypeOf((*MockCollection)(nil).BulkWrite), varargs...)
}

// CountDocuments mocks base method.
func (m *MockCollection) CountDocuments(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountDocuments", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDocuments indicates an expected call of CountDocuments.
func (mr *MockCollectionMockRecorder) CountDocuments(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDocuments", reflect.TypeOf((*MockCollection)(nil).CountDocuments), varargs...)
}

// DeleteMany mocks base method.
func (m *MockCollection) DeleteMany(ctx context.Context, filter any, opts ...*options.DeleteOptions) (*mongo0.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMany", varargs...)
	ret0, _ := ret[0].(*mongo0.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMany indicates an expected call of DeleteMany.
func (mr *MockCollectionMockRecorder) DeleteMany(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMany", reflect.TypeOf((*MockCollection)(nil).DeleteMany), varargs...)
}

// DeleteOne mocks base method.
func (m *MockCollection) DeleteOne(ctx context.Context, filter any, opts ...*options.DeleteOptions) (*mongo0.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOne", varargs...)
	ret0, _ := ret[0].(*mongo0.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockCollectionMockRecorder) DeleteOne(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockCollection)(nil).DeleteOne), varargs...)
}

// Drop mocks base method.
func (m *MockCollection) Drop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockCollectionMockRecorder) Drop(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockCollection)(nil).Drop), ctx)
}

// EstimatedDocumentCount mocks base method.
func (m *MockCollection) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EstimatedDocumentCount", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimatedDocumentCount indicates an expected call of EstimatedDocumentCount.
func (mr *MockCollectionMockRecorder) EstimatedDocumentCount(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimatedDocumentCount", reflect.TypeOf((*MockCollection)(nil).EstimatedDocumentCount), varargs...)
}

// Find mocks base method.
func (m *MockCollection) Find(ctx context.Context, filter any, opts ...*options.FindOptions) (mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCollectionMockRecorder) Find(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollection)(nil).Find), varargs...)
}

// FindOne mocks base method.
func (m *MockCollection) FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) *mongo0.SingleResult {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*mongo0.SingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCollectionMockRecorder) FindOne(ctx, filter any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockCollection)(nil).FindOne), varargs...)
}

// Indexes mocks base method.
func (m *MockCollection) Indexes() mongo0.IndexView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].(mongo0.IndexView)
	return ret0
}

// Indexes indicates an expected call of Indexes.
func (mr *MockCollectionMockRecorder) Indexes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexes", reflect.TypeOf((*MockCollection)(nil).Indexes))
}

// InsertMany mocks base method.
func (m *MockCollection) InsertMany(ctx context.Context, documents []any, opts ...*options.InsertManyOptions) (*mongo0.InsertManyResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, documents}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertMany", varargs...)
	ret0, _ := ret[0].(*mongo0.InsertManyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMany indicates an expected call of InsertMany.
func (mr *MockCollectionMockRecorder) InsertMany(ctx, documents any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, documents}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMany", reflect.TypeOf((*MockCollection)(nil).InsertMany), varargs...)
}

// InsertOne mocks base method.
func (m *MockCollection) InsertOne(ctx context.Context, document any, opts ...*options.InsertOneOptions) (*mongo0.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, document}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo0.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockCollectionMockRecorder) InsertOne(ctx, document any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, document}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockCollection)(nil).InsertOne), varargs...)
}

// UpdateMany mocks base method.
func (m *MockCollection) UpdateMany(ctx context.Context, filter, update any, opts ...*options.UpdateOptions) (*mongo0.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMany", varargs...)
	ret0, _ := ret[0].(*mongo0.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMany indicates an expected call of UpdateMany.
func (mr *MockCollectionMockRecorder) UpdateMany(ctx, filter, update any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMany", reflect.TypeOf((*MockCollection)(nil).UpdateMany), varargs...)
}

// UpdateOne mocks base method.
func (m *MockCollection) UpdateOne(ctx context.Context, filter, update any, opts ...*options.UpdateOptions) (*mongo0.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOne", varargs...)
	ret0, _ := ret[0].(*mongo0.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockCollectionMockRecorder) UpdateOne(ctx, filter, update any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockCollection)(nil).UpdateOne), varargs...)
}

// Watch mocks base method.
func (m *MockCollection) Watch(ctx context.Context, pipeline any, opts ...*options.ChangeStreamOptions) (*mongo0.ChangeStream, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(*mongo0.ChangeStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockCollectionMockRecorder) Watch(ctx, pipeline any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockCollection)(nil).Watch), varargs...)
}
