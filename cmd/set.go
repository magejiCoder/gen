// Copyright © 2020 NAME HERE scnace.me
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
	"errors"
	"os"

	"github.com/magejiCoder/gen/internal/generator"
	"github.com/spf13/cobra"
)

var dest string
var pkg string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "generate the customized set data structure",
	Long: `sub-commnad set is a generator wrapper of the https://github.com/scylladb/go-set for supporting the customized set data structure,
	specifically, it is a Go template implementation of its gen_set.sh, use the same 'gen set <SETNAME> <SETTYPE>'.
	`,
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("The set name and set type are required")
		}
		setName, setType := args[0], args[1]
		if err := generator.NewGenSet(pkg, setName, setType).GenerateTo(dest); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
	pwd, _ := os.Getwd()
	setCmd.Flags().StringVarP(&dest, "out", "o", pwd, "the directory generated the set files in")
	setCmd.Flags().StringVarP(&pkg, "gopackage", "p", "main", "the generated set package name")
}
