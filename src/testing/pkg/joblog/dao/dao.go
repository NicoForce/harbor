// Code generated by mockery v2.42.2. DO NOT EDIT.

package dao

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/pkg/joblog/models"

	time "time"
)

// DAO is an autogenerated mock type for the DAO type
type DAO struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, jobLog
func (_m *DAO) Create(ctx context.Context, jobLog *models.JobLog) (int64, error) {
	ret := _m.Called(ctx, jobLog)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.JobLog) (int64, error)); ok {
		return rf(ctx, jobLog)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.JobLog) int64); ok {
		r0 = rf(ctx, jobLog)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.JobLog) error); ok {
		r1 = rf(ctx, jobLog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBefore provides a mock function with given fields: ctx, t
func (_m *DAO) DeleteBefore(ctx context.Context, t time.Time) (int64, error) {
	ret := _m.Called(ctx, t)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBefore")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) (int64, error)); ok {
		return rf(ctx, t)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) int64); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, uuid
func (_m *DAO) Get(ctx context.Context, uuid string) (*models.JobLog, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.JobLog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.JobLog, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.JobLog); ok {
		r0 = rf(ctx, uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JobLog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDAO creates a new instance of DAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *DAO {
	mock := &DAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
