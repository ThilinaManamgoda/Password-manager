// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"github.com/password-manager/pkg/passwords"
	"github.com/password-manager/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ShowPassword flag
const (
	// ShowPassword flag
	ShowPassword = "show-pass"
	// ErrMSGCannotGetFlag message
	ErrMSGCannotGetFlag = "cannot get value of %s flag"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [ID]",
	Short: "Get a password",
	Long:  `Get a password`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ! utils.IsArgSValid(args) {
			return errors.New("please give a ID")
		}
		id := args[0]
		if ! utils.IsArgValid(id) {
			return errors.New(fmt.Sprintf("invalid argument: %s", id))
		}
		mPassword, err := utils.GetFlagStringVal(cmd, MasterPassword)
		if err != nil {
			return errors.Wrapf(err, ErrMSGCannotGetFlag, mPassword)
		}
		if mPassword == "" {
			mPassword, err = promptForMPassword()
			if err != nil {
				return errors.Wrap(err, "cannot prompt for Master password")
			}
		}
		showPass, err := utils.GetFlagBoolVal(cmd, ShowPassword)
		if err != nil {
			return errors.Wrapf(err, ErrMSGCannotGetFlag, Password)
		}

		passwordRepo, err := passwords.InitPasswordRepo(mPassword)
		if err != nil {
			return errors.Wrapf(err, "cannot initialize password repository")
		}

		err = passwordRepo.GetPassword(id, showPass)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().BoolP(ShowPassword, "s", false, "Print password to STDOUT")
}
