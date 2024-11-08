// Code generated by mockery. DO NOT EDIT.

package automock

import (
	structs "awesomeProject/todo-app/structs"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ListManager is an autogenerated mock type for the ListManager type
type ListManager struct {
	mock.Mock
}

type ListManager_Expecter struct {
	mock *mock.Mock
}

func (_m *ListManager) EXPECT() *ListManager_Expecter {
	return &ListManager_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *ListManager) Create(_a0 structs.List) (structs.List, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 structs.List
	var r1 error
	if rf, ok := ret.Get(0).(func(structs.List) (structs.List, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(structs.List) structs.List); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.List)
	}

	if rf, ok := ret.Get(1).(func(structs.List) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManager_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ListManager_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 structs.List
func (_e *ListManager_Expecter) Create(_a0 interface{}) *ListManager_Create_Call {
	return &ListManager_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *ListManager_Create_Call) Run(run func(_a0 structs.List)) *ListManager_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(structs.List))
	})
	return _c
}

func (_c *ListManager_Create_Call) Return(_a0 structs.List, _a1 error) *ListManager_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ListManager_Create_Call) RunAndReturn(run func(structs.List) (structs.List, error)) *ListManager_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *ListManager) Delete(_a0 uuid.UUID) (structs.List, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 structs.List
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (structs.List, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) structs.List); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.List)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManager_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ListManager_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 uuid.UUID
func (_e *ListManager_Expecter) Delete(_a0 interface{}) *ListManager_Delete_Call {
	return &ListManager_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *ListManager_Delete_Call) Run(run func(_a0 uuid.UUID)) *ListManager_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *ListManager_Delete_Call) Return(_a0 structs.List, _a1 error) *ListManager_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ListManager_Delete_Call) RunAndReturn(run func(uuid.UUID) (structs.List, error)) *ListManager_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0
func (_m *ListManager) Get(_a0 uuid.UUID) (structs.List, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 structs.List
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (structs.List, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) structs.List); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.List)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManager_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ListManager_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 uuid.UUID
func (_e *ListManager_Expecter) Get(_a0 interface{}) *ListManager_Get_Call {
	return &ListManager_Get_Call{Call: _e.mock.On("Get", _a0)}
}

func (_c *ListManager_Get_Call) Run(run func(_a0 uuid.UUID)) *ListManager_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *ListManager_Get_Call) Return(_a0 structs.List, _a1 error) *ListManager_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ListManager_Get_Call) RunAndReturn(run func(uuid.UUID) (structs.List, error)) *ListManager_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *ListManager) GetAll() []structs.List {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []structs.List
	if rf, ok := ret.Get(0).(func() []structs.List); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]structs.List)
		}
	}

	return r0
}

// ListManager_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type ListManager_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *ListManager_Expecter) GetAll() *ListManager_GetAll_Call {
	return &ListManager_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *ListManager_GetAll_Call) Run(run func()) *ListManager_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ListManager_GetAll_Call) Return(_a0 []structs.List) *ListManager_GetAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ListManager_GetAll_Call) RunAndReturn(run func() []structs.List) *ListManager_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *ListManager) Update(_a0 structs.List) (structs.List, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 structs.List
	var r1 error
	if rf, ok := ret.Get(0).(func(structs.List) (structs.List, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(structs.List) structs.List); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.List)
	}

	if rf, ok := ret.Get(1).(func(structs.List) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManager_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ListManager_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 structs.List
func (_e *ListManager_Expecter) Update(_a0 interface{}) *ListManager_Update_Call {
	return &ListManager_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *ListManager_Update_Call) Run(run func(_a0 structs.List)) *ListManager_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(structs.List))
	})
	return _c
}

func (_c *ListManager_Update_Call) Return(_a0 structs.List, _a1 error) *ListManager_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ListManager_Update_Call) RunAndReturn(run func(structs.List) (structs.List, error)) *ListManager_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewListManager creates a new instance of ListManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewListManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *ListManager {
	mock := &ListManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
