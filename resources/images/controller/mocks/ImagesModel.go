// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package mocks

import (
	"time"

	"github.com/mendersoftware/deployments/resources/images"
	"github.com/stretchr/testify/mock"
)

// ImagesModel is an autogenerated mock type for the ImagesModel type
type ImagesModel struct {
	mock.Mock
}

// CreateImage provides a mock function with given fields: constructorData
func (_m *ImagesModel) CreateImage(constructorData *images.SoftwareImageConstructor) (string, error) {
	ret := _m.Called(constructorData)

	var r0 string
	if rf, ok := ret.Get(0).(func(*images.SoftwareImageConstructor) string); ok {
		r0 = rf(constructorData)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*images.SoftwareImageConstructor) error); ok {
		r1 = rf(constructorData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImage provides a mock function with given fields: imageID
func (_m *ImagesModel) DeleteImage(imageID string) error {
	ret := _m.Called(imageID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(imageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DownloadLink provides a mock function with given fields: imageID, expire
func (_m *ImagesModel) DownloadLink(imageID string, expire time.Duration) (*images.Link, error) {
	ret := _m.Called(imageID, expire)

	var r0 *images.Link
	if rf, ok := ret.Get(0).(func(string, time.Duration) *images.Link); ok {
		r0 = rf(imageID, expire)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*images.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Duration) error); ok {
		r1 = rf(imageID, expire)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditImage provides a mock function with given fields: id, constructorData
func (_m *ImagesModel) EditImage(id string, constructorData *images.SoftwareImageConstructor) (bool, error) {
	ret := _m.Called(id, constructorData)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, *images.SoftwareImageConstructor) bool); ok {
		r0 = rf(id, constructorData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *images.SoftwareImageConstructor) error); ok {
		r1 = rf(id, constructorData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImage provides a mock function with given fields: id
func (_m *ImagesModel) GetImage(id string) (*images.SoftwareImage, error) {
	ret := _m.Called(id)

	var r0 *images.SoftwareImage
	if rf, ok := ret.Get(0).(func(string) *images.SoftwareImage); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*images.SoftwareImage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImages provides a mock function with given fields: filters
func (_m *ImagesModel) ListImages(filters map[string]string) ([]*images.SoftwareImage, error) {
	ret := _m.Called(filters)

	var r0 []*images.SoftwareImage
	if rf, ok := ret.Get(0).(func(map[string]string) []*images.SoftwareImage); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*images.SoftwareImage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadLink provides a mock function with given fields: imageID, expire
func (_m *ImagesModel) UploadLink(imageID string, expire time.Duration) (*images.Link, error) {
	ret := _m.Called(imageID, expire)

	var r0 *images.Link
	if rf, ok := ret.Get(0).(func(string, time.Duration) *images.Link); ok {
		r0 = rf(imageID, expire)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*images.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Duration) error); ok {
		r1 = rf(imageID, expire)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
