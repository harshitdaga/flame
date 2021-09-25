// Copyright (c) 2021 Cisco Systems, Inc. and its affiliates
// All rights reserved
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

package cmd

import (
	"github.com/spf13/cobra"

	"wwwin-github.cisco.com/eti/fledge/cmd/fledgectl/resources/design"
)

var createDesignCmd = &cobra.Command{
	Use:   "design <designId>",
	Short: "Create a new design template",
	Long:  "This command creates a new design template",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		designId := args[0]

		flags := cmd.Flags()
		description, err := flags.GetString("desc")
		if err != nil {
			return err
		}

		params := design.Params{}
		params.Endpoint = config.ApiServer.Endpoint
		params.User = config.User
		params.DesignId = designId
		params.Desc = description

		return design.Create(params)
	},
}

func init() {
	createDesignCmd.Flags().StringP("desc", "d", "", "Design description")
	createCmd.AddCommand(createDesignCmd)
}
