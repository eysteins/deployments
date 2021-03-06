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

package model

import (
	"errors"
	"os"
	"time"

	"github.com/mendersoftware/deployments/resources/images"
)

// Errors specific to interface
var (
	ErrFileStorageFileNotFound = errors.New("File not found")
)

// FileStorage allows to store and manage large files
type FileStorage interface {
	Delete(objectId string) error
	Exists(objectId string) (bool, error)
	LastModified(objectId string) (time.Time, error)
	PutRequest(objectId string, duration time.Duration) (*images.Link, error)
	GetRequest(objectId string, duration time.Duration, responseContentType string) (*images.Link, error)
	PutFile(objectId string, image *os.File, contentType string) error
}
