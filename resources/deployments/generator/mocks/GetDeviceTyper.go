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
	"context"
	"github.com/stretchr/testify/mock"
)

// GetDeviceTyper is an autogenerated mock type for the GetDeviceTyper type
type GetDeviceTyper struct {
	mock.Mock
}

// GetDeviceType provides a mock function with given fields: ctx, deviceID
func (_m *GetDeviceTyper) GetDeviceType(ctx context.Context, deviceID string) (string, error) {
	ret := _m.Called(ctx, deviceID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, deviceID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
