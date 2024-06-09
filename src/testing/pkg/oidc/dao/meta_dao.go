// Code generated by mockery v2.42.2. DO NOT EDIT.

package dao

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/common/models"

	q "github.com/goharbor/harbor/src/lib/q"
)

// MetaDAO is an autogenerated mock type for the MetaDAO type
type MetaDAO struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, oidcUser
func (_m *MetaDAO) Create(ctx context.Context, oidcUser *models.OIDCUser) (int, error) {
	ret := _m.Called(ctx, oidcUser)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.OIDCUser) (int, error)); ok {
		return rf(ctx, oidcUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.OIDCUser) int); ok {
		r0 = rf(ctx, oidcUser)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.OIDCUser) error); ok {
		r1 = rf(ctx, oidcUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByUserID provides a mock function with given fields: ctx, uid
func (_m *MetaDAO) DeleteByUserID(ctx context.Context, uid int) error {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByUserID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *MetaDAO) GetByUsername(ctx context.Context, username string) (*models.OIDCUser, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 *models.OIDCUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.OIDCUser, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.OIDCUser); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OIDCUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *MetaDAO) List(ctx context.Context, query *q.Query) ([]*models.OIDCUser, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*models.OIDCUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*models.OIDCUser, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*models.OIDCUser); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.OIDCUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, oidcUser, props
func (_m *MetaDAO) Update(ctx context.Context, oidcUser *models.OIDCUser, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, oidcUser)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.OIDCUser, ...string) error); ok {
		r0 = rf(ctx, oidcUser, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMetaDAO creates a new instance of MetaDAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetaDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *MetaDAO {
	mock := &MetaDAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
