// Copyright © 2019 The Vultr-cli Authors
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
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vultr/vultr-cli/cmd/printer"
)

// applicationCmd represents the application command
var applicationCmd = &cobra.Command{
	Use:     "apps",
	Aliases: []string{"a"},
	Short:   "Display all available applications",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		apps, err := client.Application.List(context.TODO())

		if err != nil {
			fmt.Println("error getting available applications")
			os.Exit(1)
		}

		printer.Application(apps)
	},
}
