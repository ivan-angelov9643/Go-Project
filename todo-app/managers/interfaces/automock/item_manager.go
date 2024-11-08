// Code generated by mockery. DO NOT EDIT.

package automock

import (
	structs "awesomeProject/todo-app/structs"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ItemManager is an autogenerated mock type for the ItemManager type
type ItemManager struct {
	mock.Mock
}

type ItemManager_Expecter struct {
	mock *mock.Mock
}

func (_m *ItemManager) EXPECT() *ItemManager_Expecter {
	return &ItemManager_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *ItemManager) Create(_a0 structs.Item) (structs.Item, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 structs.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(structs.Item) (structs.Item, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(structs.Item) structs.Item); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.Item)
	}

	if rf, ok := ret.Get(1).(func(structs.Item) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ItemManager_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ItemManager_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 structs.Item
func (_e *ItemManager_Expecter) Create(_a0 interface{}) *ItemManager_Create_Call {
	return &ItemManager_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *ItemManager_Create_Call) Run(run func(_a0 structs.Item)) *ItemManager_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(structs.Item))
	})
	return _c
}

func (_c *ItemManager_Create_Call) Return(_a0 structs.Item, _a1 error) *ItemManager_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ItemManager_Create_Call) RunAndReturn(run func(structs.Item) (structs.Item, error)) *ItemManager_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *ItemManager) Delete(_a0 uuid.UUID) (structs.Item, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 structs.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (structs.Item, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) structs.Item); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.Item)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ItemManager_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ItemManager_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 uuid.UUID
func (_e *ItemManager_Expecter) Delete(_a0 interface{}) *ItemManager_Delete_Call {
	return &ItemManager_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *ItemManager_Delete_Call) Run(run func(_a0 uuid.UUID)) *ItemManager_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *ItemManager_Delete_Call) Return(_a0 structs.Item, _a1 error) *ItemManager_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ItemManager_Delete_Call) RunAndReturn(run func(uuid.UUID) (structs.Item, error)) *ItemManager_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0
func (_m *ItemManager) Get(_a0 uuid.UUID) (structs.Item, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 structs.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (structs.Item, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) structs.Item); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.Item)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ItemManager_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ItemManager_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 uuid.UUID
func (_e *ItemManager_Expecter) Get(_a0 interface{}) *ItemManager_Get_Call {
	return &ItemManager_Get_Call{Call: _e.mock.On("Get", _a0)}
}

func (_c *ItemManager_Get_Call) Run(run func(_a0 uuid.UUID)) *ItemManager_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *ItemManager_Get_Call) Return(_a0 structs.Item, _a1 error) *ItemManager_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ItemManager_Get_Call) RunAndReturn(run func(uuid.UUID) (structs.Item, error)) *ItemManager_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *ItemManager) GetAll() []structs.Item {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []structs.Item
	if rf, ok := ret.Get(0).(func() []structs.Item); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]structs.Item)
		}
	}

	return r0
}

// ItemManager_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type ItemManager_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *ItemManager_Expecter) GetAll() *ItemManager_GetAll_Call {
	return &ItemManager_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *ItemManager_GetAll_Call) Run(run func()) *ItemManager_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ItemManager_GetAll_Call) Return(_a0 []structs.Item) *ItemManager_GetAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ItemManager_GetAll_Call) RunAndReturn(run func() []structs.Item) *ItemManager_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *ItemManager) Update(_a0 structs.Item) (structs.Item, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 structs.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(structs.Item) (structs.Item, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(structs.Item) structs.Item); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(structs.Item)
	}

	if rf, ok := ret.Get(1).(func(structs.Item) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ItemManager_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ItemManager_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 structs.Item
func (_e *ItemManager_Expecter) Update(_a0 interface{}) *ItemManager_Update_Call {
	return &ItemManager_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *ItemManager_Update_Call) Run(run func(_a0 structs.Item)) *ItemManager_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(structs.Item))
	})
	return _c
}

func (_c *ItemManager_Update_Call) Return(_a0 structs.Item, _a1 error) *ItemManager_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ItemManager_Update_Call) RunAndReturn(run func(structs.Item) (structs.Item, error)) *ItemManager_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewItemManager creates a new instance of ItemManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewItemManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *ItemManager {
	mock := &ItemManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
