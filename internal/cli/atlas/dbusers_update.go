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

package atlas

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type DBUsersUpdateOpts struct {
	cli.GlobalOpts
	username string
	password string
	roles    []string
	store    store.DatabaseUserUpdater
}

func (opts *DBUsersUpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DBUsersUpdateOpts) Run() error {
	current := new(atlas.DatabaseUser)

	opts.update(current)

	result, err := opts.store.UpdateDatabaseUser(current)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *DBUsersUpdateOpts) update(out *atlas.DatabaseUser) {
	out.GroupID = opts.ConfigProjectID()
	out.Username = opts.username
	if opts.password != "" {
		out.Password = opts.password
	}

	out.Roles = convert.BuildAtlasRoles(opts.roles)
}

// mongocli atlas dbuser(s) update <username> [--password password] [--role roleName@dbName] [--projectId projectId]
func DBUsersUpdateBuilder() *cobra.Command {
	opts := &DBUsersUpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <username>",
		Short: description.UpdateDBUser,
		Example: `
  Update roles for a user
  $ mongocli atlas dbuser update <username> --role readWriteAnyDatabase --projectId <projectId>`,
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.username = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.username, flag.Username, flag.UsernameShort, "", usage.Username)
	cmd.Flags().StringVarP(&opts.password, flag.Password, flag.PasswordShort, "", usage.Password)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.Roles)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
