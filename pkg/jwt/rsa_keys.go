// Copyright 2021 EMQ Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jwt

import (
	"crypto/rsa"
	"os"
	"sync"

	"github.com/golang-jwt/jwt"
	"github.com/yuxi311/webService/pkg/utils"
)

var (
	repositoryLock sync.Mutex
)

const (
	RSAPrivateKeyPath = "etc/webservice_key"
	RSAPublicKeyPath  = "etc/webservice_key.pub"
)

func GetPrivateKey() (*rsa.PrivateKey, error) {
	repositoryLock.Lock()
	defer repositoryLock.Unlock()

	keyBytes, err := os.ReadFile(utils.ToFilePath(RSAPrivateKeyPath))
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	return signKey, nil
}

func GetPublicKey() (*rsa.PublicKey, error) {
	repositoryLock.Lock()
	defer repositoryLock.Unlock()

	keyBytes, err := os.ReadFile(utils.ToFilePath(RSAPublicKeyPath))
	if err != nil {
		return nil, err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}
