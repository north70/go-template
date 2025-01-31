// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock

import (
	context "context"

	domain "github.com/north70/go-template/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// FooServiceMock is an autogenerated mock type for the FooService type
type FooServiceMock struct {
	mock.Mock
}

type FooServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *FooServiceMock) EXPECT() *FooServiceMock_Expecter {
	return &FooServiceMock_Expecter{mock: &_m.Mock}
}

// GetFoo provides a mock function with given fields: ctx, id
func (_m *FooServiceMock) GetFoo(ctx context.Context, id string) (*domain.Foo, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Foo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Foo, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Foo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Foo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FooServiceMock_GetFoo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFoo'
type FooServiceMock_GetFoo_Call struct {
	*mock.Call
}

// GetFoo is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *FooServiceMock_Expecter) GetFoo(ctx interface{}, id interface{}) *FooServiceMock_GetFoo_Call {
	return &FooServiceMock_GetFoo_Call{Call: _e.mock.On("GetFoo", ctx, id)}
}

func (_c *FooServiceMock_GetFoo_Call) Run(run func(ctx context.Context, id string)) *FooServiceMock_GetFoo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FooServiceMock_GetFoo_Call) Return(_a0 *domain.Foo, _a1 error) *FooServiceMock_GetFoo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FooServiceMock_GetFoo_Call) RunAndReturn(run func(context.Context, string) (*domain.Foo, error)) *FooServiceMock_GetFoo_Call {
	_c.Call.Return(run)
	return _c
}

// NewFooServiceMock creates a new instance of FooServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFooServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *FooServiceMock {
	mock := &FooServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
