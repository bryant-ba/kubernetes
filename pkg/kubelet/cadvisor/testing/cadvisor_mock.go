/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by mockery v2.53.3. DO NOT EDIT.

package testing

import (
	context "context"

	v1 "github.com/google/cadvisor/info/v1"
	mock "github.com/stretchr/testify/mock"

	v2 "github.com/google/cadvisor/info/v2"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// ContainerFsInfo provides a mock function with given fields: _a0
func (_m *MockInterface) ContainerFsInfo(_a0 context.Context) (v2.FsInfo, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ContainerFsInfo")
	}

	var r0 v2.FsInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (v2.FsInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) v2.FsInfo); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(v2.FsInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_ContainerFsInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ContainerFsInfo'
type MockInterface_ContainerFsInfo_Call struct {
	*mock.Call
}

// ContainerFsInfo is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockInterface_Expecter) ContainerFsInfo(_a0 interface{}) *MockInterface_ContainerFsInfo_Call {
	return &MockInterface_ContainerFsInfo_Call{Call: _e.mock.On("ContainerFsInfo", _a0)}
}

func (_c *MockInterface_ContainerFsInfo_Call) Run(run func(_a0 context.Context)) *MockInterface_ContainerFsInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInterface_ContainerFsInfo_Call) Return(_a0 v2.FsInfo, _a1 error) *MockInterface_ContainerFsInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_ContainerFsInfo_Call) RunAndReturn(run func(context.Context) (v2.FsInfo, error)) *MockInterface_ContainerFsInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ContainerInfoV2 provides a mock function with given fields: name, options
func (_m *MockInterface) ContainerInfoV2(name string, options v2.RequestOptions) (map[string]v2.ContainerInfo, error) {
	ret := _m.Called(name, options)

	if len(ret) == 0 {
		panic("no return value specified for ContainerInfoV2")
	}

	var r0 map[string]v2.ContainerInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string, v2.RequestOptions) (map[string]v2.ContainerInfo, error)); ok {
		return rf(name, options)
	}
	if rf, ok := ret.Get(0).(func(string, v2.RequestOptions) map[string]v2.ContainerInfo); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]v2.ContainerInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string, v2.RequestOptions) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_ContainerInfoV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ContainerInfoV2'
type MockInterface_ContainerInfoV2_Call struct {
	*mock.Call
}

// ContainerInfoV2 is a helper method to define mock.On call
//   - name string
//   - options v2.RequestOptions
func (_e *MockInterface_Expecter) ContainerInfoV2(name interface{}, options interface{}) *MockInterface_ContainerInfoV2_Call {
	return &MockInterface_ContainerInfoV2_Call{Call: _e.mock.On("ContainerInfoV2", name, options)}
}

func (_c *MockInterface_ContainerInfoV2_Call) Run(run func(name string, options v2.RequestOptions)) *MockInterface_ContainerInfoV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(v2.RequestOptions))
	})
	return _c
}

func (_c *MockInterface_ContainerInfoV2_Call) Return(_a0 map[string]v2.ContainerInfo, _a1 error) *MockInterface_ContainerInfoV2_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_ContainerInfoV2_Call) RunAndReturn(run func(string, v2.RequestOptions) (map[string]v2.ContainerInfo, error)) *MockInterface_ContainerInfoV2_Call {
	_c.Call.Return(run)
	return _c
}

// GetDirFsInfo provides a mock function with given fields: path
func (_m *MockInterface) GetDirFsInfo(path string) (v2.FsInfo, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for GetDirFsInfo")
	}

	var r0 v2.FsInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (v2.FsInfo, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) v2.FsInfo); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(v2.FsInfo)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetDirFsInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDirFsInfo'
type MockInterface_GetDirFsInfo_Call struct {
	*mock.Call
}

// GetDirFsInfo is a helper method to define mock.On call
//   - path string
func (_e *MockInterface_Expecter) GetDirFsInfo(path interface{}) *MockInterface_GetDirFsInfo_Call {
	return &MockInterface_GetDirFsInfo_Call{Call: _e.mock.On("GetDirFsInfo", path)}
}

func (_c *MockInterface_GetDirFsInfo_Call) Run(run func(path string)) *MockInterface_GetDirFsInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockInterface_GetDirFsInfo_Call) Return(_a0 v2.FsInfo, _a1 error) *MockInterface_GetDirFsInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetDirFsInfo_Call) RunAndReturn(run func(string) (v2.FsInfo, error)) *MockInterface_GetDirFsInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetRequestedContainersInfo provides a mock function with given fields: containerName, options
func (_m *MockInterface) GetRequestedContainersInfo(containerName string, options v2.RequestOptions) (map[string]*v1.ContainerInfo, error) {
	ret := _m.Called(containerName, options)

	if len(ret) == 0 {
		panic("no return value specified for GetRequestedContainersInfo")
	}

	var r0 map[string]*v1.ContainerInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string, v2.RequestOptions) (map[string]*v1.ContainerInfo, error)); ok {
		return rf(containerName, options)
	}
	if rf, ok := ret.Get(0).(func(string, v2.RequestOptions) map[string]*v1.ContainerInfo); ok {
		r0 = rf(containerName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*v1.ContainerInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string, v2.RequestOptions) error); ok {
		r1 = rf(containerName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetRequestedContainersInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequestedContainersInfo'
type MockInterface_GetRequestedContainersInfo_Call struct {
	*mock.Call
}

// GetRequestedContainersInfo is a helper method to define mock.On call
//   - containerName string
//   - options v2.RequestOptions
func (_e *MockInterface_Expecter) GetRequestedContainersInfo(containerName interface{}, options interface{}) *MockInterface_GetRequestedContainersInfo_Call {
	return &MockInterface_GetRequestedContainersInfo_Call{Call: _e.mock.On("GetRequestedContainersInfo", containerName, options)}
}

func (_c *MockInterface_GetRequestedContainersInfo_Call) Run(run func(containerName string, options v2.RequestOptions)) *MockInterface_GetRequestedContainersInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(v2.RequestOptions))
	})
	return _c
}

func (_c *MockInterface_GetRequestedContainersInfo_Call) Return(_a0 map[string]*v1.ContainerInfo, _a1 error) *MockInterface_GetRequestedContainersInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetRequestedContainersInfo_Call) RunAndReturn(run func(string, v2.RequestOptions) (map[string]*v1.ContainerInfo, error)) *MockInterface_GetRequestedContainersInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ImagesFsInfo provides a mock function with given fields: _a0
func (_m *MockInterface) ImagesFsInfo(_a0 context.Context) (v2.FsInfo, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ImagesFsInfo")
	}

	var r0 v2.FsInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (v2.FsInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) v2.FsInfo); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(v2.FsInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_ImagesFsInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ImagesFsInfo'
type MockInterface_ImagesFsInfo_Call struct {
	*mock.Call
}

// ImagesFsInfo is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockInterface_Expecter) ImagesFsInfo(_a0 interface{}) *MockInterface_ImagesFsInfo_Call {
	return &MockInterface_ImagesFsInfo_Call{Call: _e.mock.On("ImagesFsInfo", _a0)}
}

func (_c *MockInterface_ImagesFsInfo_Call) Run(run func(_a0 context.Context)) *MockInterface_ImagesFsInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInterface_ImagesFsInfo_Call) Return(_a0 v2.FsInfo, _a1 error) *MockInterface_ImagesFsInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_ImagesFsInfo_Call) RunAndReturn(run func(context.Context) (v2.FsInfo, error)) *MockInterface_ImagesFsInfo_Call {
	_c.Call.Return(run)
	return _c
}

// MachineInfo provides a mock function with no fields
func (_m *MockInterface) MachineInfo() (*v1.MachineInfo, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MachineInfo")
	}

	var r0 *v1.MachineInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.MachineInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.MachineInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MachineInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_MachineInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MachineInfo'
type MockInterface_MachineInfo_Call struct {
	*mock.Call
}

// MachineInfo is a helper method to define mock.On call
func (_e *MockInterface_Expecter) MachineInfo() *MockInterface_MachineInfo_Call {
	return &MockInterface_MachineInfo_Call{Call: _e.mock.On("MachineInfo")}
}

func (_c *MockInterface_MachineInfo_Call) Run(run func()) *MockInterface_MachineInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_MachineInfo_Call) Return(_a0 *v1.MachineInfo, _a1 error) *MockInterface_MachineInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_MachineInfo_Call) RunAndReturn(run func() (*v1.MachineInfo, error)) *MockInterface_MachineInfo_Call {
	_c.Call.Return(run)
	return _c
}

// RootFsInfo provides a mock function with no fields
func (_m *MockInterface) RootFsInfo() (v2.FsInfo, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RootFsInfo")
	}

	var r0 v2.FsInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() (v2.FsInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() v2.FsInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v2.FsInfo)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_RootFsInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RootFsInfo'
type MockInterface_RootFsInfo_Call struct {
	*mock.Call
}

// RootFsInfo is a helper method to define mock.On call
func (_e *MockInterface_Expecter) RootFsInfo() *MockInterface_RootFsInfo_Call {
	return &MockInterface_RootFsInfo_Call{Call: _e.mock.On("RootFsInfo")}
}

func (_c *MockInterface_RootFsInfo_Call) Run(run func()) *MockInterface_RootFsInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_RootFsInfo_Call) Return(_a0 v2.FsInfo, _a1 error) *MockInterface_RootFsInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_RootFsInfo_Call) RunAndReturn(run func() (v2.FsInfo, error)) *MockInterface_RootFsInfo_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with no fields
func (_m *MockInterface) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockInterface_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockInterface_Expecter) Start() *MockInterface_Start_Call {
	return &MockInterface_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockInterface_Start_Call) Run(run func()) *MockInterface_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_Start_Call) Return(_a0 error) *MockInterface_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Start_Call) RunAndReturn(run func() error) *MockInterface_Start_Call {
	_c.Call.Return(run)
	return _c
}

// VersionInfo provides a mock function with no fields
func (_m *MockInterface) VersionInfo() (*v1.VersionInfo, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for VersionInfo")
	}

	var r0 *v1.VersionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.VersionInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.VersionInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.VersionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_VersionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VersionInfo'
type MockInterface_VersionInfo_Call struct {
	*mock.Call
}

// VersionInfo is a helper method to define mock.On call
func (_e *MockInterface_Expecter) VersionInfo() *MockInterface_VersionInfo_Call {
	return &MockInterface_VersionInfo_Call{Call: _e.mock.On("VersionInfo")}
}

func (_c *MockInterface_VersionInfo_Call) Run(run func()) *MockInterface_VersionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_VersionInfo_Call) Return(_a0 *v1.VersionInfo, _a1 error) *MockInterface_VersionInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_VersionInfo_Call) RunAndReturn(run func() (*v1.VersionInfo, error)) *MockInterface_VersionInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
