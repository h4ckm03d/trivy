// Code generated by mockery v1.0.0. DO NOT EDIT.

package db

import (
	pkgdb "github.com/aquasecurity/trivy-db/pkg/db"
	mock "github.com/stretchr/testify/mock"
)

// mockDbOperation is an autogenerated mock type for the dbOperation type
type mockDbOperation struct {
	mock.Mock
}

type dbOperationGetMetadataReturns struct {
	Metadata pkgdb.Metadata
	Err      error
}

type dbOperationGetMetadataExpectation struct {
	Returns dbOperationGetMetadataReturns
}

func (_m *mockDbOperation) ApplyGetMetadataExpectation(e dbOperationGetMetadataExpectation) {
	var args []interface{}
	_m.On("GetMetadata", args...).Return(e.Returns.Metadata, e.Returns.Err)
}

func (_m *mockDbOperation) ApplyGetMetadataExpectations(expectations []dbOperationGetMetadataExpectation) {
	for _, e := range expectations {
		_m.ApplyGetMetadataExpectation(e)
	}
}

// GetMetadata provides a mock function with given fields:
func (_m *mockDbOperation) GetMetadata() (pkgdb.Metadata, error) {
	ret := _m.Called()

	var r0 pkgdb.Metadata
	if rf, ok := ret.Get(0).(func() pkgdb.Metadata); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pkgdb.Metadata)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
