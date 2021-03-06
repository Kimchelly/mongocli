// Copyright 2020 MongoDB Inc
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
// +build e2e cloudmanager

package cloud_manager_test

import (
	"os"
	"path/filepath"
)

const (
	entity       = "cloud-manager"
	mongoCliPath = "../../bin/mongocli"
)

func cli() (string, error) {
	cliPath, err := filepath.Abs(mongoCliPath)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(cliPath); err != nil {
		return "", err
	}
	return cliPath, nil
}
