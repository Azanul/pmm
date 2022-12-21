// Code generated by mockery v1.0.0. DO NOT EDIT.

package dbaas

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"

	kubernetes "github.com/percona/pmm/managed/services/dbaas/kubernetes"
)

// mockKubernetesClient is an autogenerated mock type for the kubernetesClient type
type mockKubernetesClient struct {
	mock.Mock
}

// GetAllClusterResources provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockKubernetesClient) GetAllClusterResources(_a0 context.Context, _a1 kubernetes.ClusterType, _a2 *v1.PersistentVolumeList) (uint64, uint64, uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, kubernetes.ClusterType, *v1.PersistentVolumeList) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, kubernetes.ClusterType, *v1.PersistentVolumeList) uint64); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 uint64
	if rf, ok := ret.Get(2).(func(context.Context, kubernetes.ClusterType, *v1.PersistentVolumeList) uint64); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Get(2).(uint64)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, kubernetes.ClusterType, *v1.PersistentVolumeList) error); ok {
		r3 = rf(_a0, _a1, _a2)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetClusterType provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) GetClusterType(_a0 context.Context) (kubernetes.ClusterType, error) {
	ret := _m.Called(_a0)

	var r0 kubernetes.ClusterType
	if rf, ok := ret.Get(0).(func(context.Context) kubernetes.ClusterType); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(kubernetes.ClusterType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsumedCPUAndMemory provides a mock function with given fields: _a0, _a1
func (_m *mockKubernetesClient) GetConsumedCPUAndMemory(_a0 context.Context, _a1 string) (uint64, uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, string) uint64); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetConsumedDiskBytes provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockKubernetesClient) GetConsumedDiskBytes(_a0 context.Context, _a1 kubernetes.ClusterType, _a2 *v1.PersistentVolumeList) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, kubernetes.ClusterType, *v1.PersistentVolumeList) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kubernetes.ClusterType, *v1.PersistentVolumeList) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPersistentVolumes provides a mock function with given fields: ctx
func (_m *mockKubernetesClient) GetPersistentVolumes(ctx context.Context) (*v1.PersistentVolumeList, error) {
	ret := _m.Called(ctx)

	var r0 *v1.PersistentVolumeList
	if rf, ok := ret.Get(0).(func(context.Context) *v1.PersistentVolumeList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PersistentVolumeList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetKubeconfig provides a mock function with given fields: _a0
func (_m *mockKubernetesClient) SetKubeconfig(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
