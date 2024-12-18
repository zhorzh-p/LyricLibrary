// Code generated by mockery v2.49.1. DO NOT EDIT.

package music

import (
	mock "github.com/stretchr/testify/mock"
	music "github.com/zhorzh-p/LyricLibrary/internal/domain/repositories/music"
)

// VerseDatabaseRepository is an autogenerated mock type for the VerseDatabaseRepository type
type VerseDatabaseRepository struct {
	mock.Mock
}

type VerseDatabaseRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *VerseDatabaseRepository) EXPECT() *VerseDatabaseRepository_Expecter {
	return &VerseDatabaseRepository_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: songId, offset, limit
func (_m *VerseDatabaseRepository) Get(songId uint, offset uint, limit uint) ([]music.VerseEntity, error) {
	ret := _m.Called(songId, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []music.VerseEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint) ([]music.VerseEntity, error)); ok {
		return rf(songId, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, uint) []music.VerseEntity); ok {
		r0 = rf(songId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]music.VerseEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, uint) error); ok {
		r1 = rf(songId, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerseDatabaseRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type VerseDatabaseRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - songId uint
//   - offset uint
//   - limit uint
func (_e *VerseDatabaseRepository_Expecter) Get(songId interface{}, offset interface{}, limit interface{}) *VerseDatabaseRepository_Get_Call {
	return &VerseDatabaseRepository_Get_Call{Call: _e.mock.On("Get", songId, offset, limit)}
}

func (_c *VerseDatabaseRepository_Get_Call) Run(run func(songId uint, offset uint, limit uint)) *VerseDatabaseRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(uint), args[2].(uint))
	})
	return _c
}

func (_c *VerseDatabaseRepository_Get_Call) Return(_a0 []music.VerseEntity, _a1 error) *VerseDatabaseRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VerseDatabaseRepository_Get_Call) RunAndReturn(run func(uint, uint, uint) ([]music.VerseEntity, error)) *VerseDatabaseRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewVerseDatabaseRepository creates a new instance of VerseDatabaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVerseDatabaseRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *VerseDatabaseRepository {
	mock := &VerseDatabaseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
