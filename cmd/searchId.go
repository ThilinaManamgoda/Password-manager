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
	"github.com/ThilinaManamgoda/password-manager/pkg/passwords"
	"github.com/ThilinaManamgoda/password-manager/pkg/utils"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// searchIDCmd represents the searchId command
var searchIDCmd = &cobra.Command{
	Use:   "searchId [ID]",
	Short: "Search Password with ID",
	Long:  `You can use either complete or part of ID for searching`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !utils.IsArgSValid(args) {
			return errors.New("Please give a ID")
		}
		searchID := args[0]
		if !utils.IsArgValid(searchID) {
			return errors.New(fmt.Sprintf("Invalid argument: %s", searchID))
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

		if !utils.IsArgSValid(args) {
			return errors.New("Please give a ID")
		}

		passwordRepo, err := passwords.InitPasswordRepo(mPassword)
		if err != nil {
			return errors.Wrapf(err, "cannot initialize password repository")
		}

		passwordEntries, err := passwordRepo.SearchID(searchID, showPass)
		if err != nil {
			return errors.Wrapf(err, "cannot search ID")
		}

		if len(passwordEntries) != 0 {
			var idList []string
			for _, val := range passwordEntries {
				idList = append(idList, val.ID)
			}
			sID, _ := utils.PromptForSelect("Choose", idList)
			err := passwordRepo.GetPassword(sID, showPass)
			if err != nil {
				return errors.Wrapf(err, "cannot get password for ID: %s", sID)
			}
		} else {
			return errors.New("cannot find any match")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchIDCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchIDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchIDCmd.Flags().BoolP(ShowPassword, "s", false, "Print password to STDOUT")
}
