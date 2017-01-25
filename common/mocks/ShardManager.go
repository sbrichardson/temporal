package mocks

import mock "github.com/stretchr/testify/mock"
import persistence "code.uber.internal/devexp/minions/common/persistence"

// ShardManager is an autogenerated mock type for the ShardManager type
type ShardManager struct {
	mock.Mock
}

// CreateShard provides a mock function with given fields: request
func (_m *ShardManager) CreateShard(request *persistence.CreateShardRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.CreateShardRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetShard provides a mock function with given fields: request
func (_m *ShardManager) GetShard(request *persistence.GetShardRequest) (*persistence.GetShardResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetShardResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetShardRequest) *persistence.GetShardResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetShardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetShardRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateShard provides a mock function with given fields: request
func (_m *ShardManager) UpdateShard(request *persistence.UpdateShardRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.UpdateShardRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ persistence.ShardManager = (*ShardManager)(nil)