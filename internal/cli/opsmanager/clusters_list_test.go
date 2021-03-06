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

package opsmanager

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/fixture"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestCloudManagerClustersList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudManagerClustersLister(ctrl)

	defer ctrl.Finish()

	t.Run("ConfigProjectID is given", func(t *testing.T) {
		expected := fixture.AutomationConfig()

		listOpts := &ClustersListOpts{
			store: mockStore,
		}

		listOpts.ProjectID = "1"
		mockStore.
			EXPECT().
			GetAutomationConfig(listOpts.ProjectID).
			Return(expected, nil).
			Times(1)

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("No ConfigProjectID is given", func(t *testing.T) {
		expected := fixture.AllClusters()
		config.SetService(config.OpsManagerService)
		mockStore.
			EXPECT().
			ListAllProjectClusters().
			Return(expected, nil).
			Times(1)

		listOpts := &ClustersListOpts{
			store: mockStore,
		}

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
