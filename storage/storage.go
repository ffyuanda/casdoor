// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import "github.com/qor/oss"

func GetStorageProvider(providerType string, clientId string, clientSecret string, region string, bucket string, endpoint string) oss.StorageInterface {
	switch providerType {
	case "Local File System":
		return NewLocalFileSystemStorageProvider(clientId, clientSecret, region, bucket, endpoint)
	case "AWS S3":
		return NewAwsS3StorageProvider(clientId, clientSecret, region, bucket, endpoint)
	case "Aliyun OSS":
		return NewAliyunStorageProvider(clientId, clientSecret, region, bucket, endpoint)
	}

	return nil
}
