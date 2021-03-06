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
	"github.com/ThilinaManamgoda/password-manager/pkg/config"
	"github.com/ThilinaManamgoda/password-manager/pkg/inputs"
	"github.com/ThilinaManamgoda/password-manager/pkg/passwords"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export password repository to a file",
	Long:  `This command can be used to export password repository to file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mPassword, err := inputs.GetFlagStringVal(cmd, inputs.FlagMasterPassword)
		if err != nil {
			return errors.Wrapf(err, inputs.ErrMsgCannotGetFlag, inputs.FlagMasterPassword)
		}
		if mPassword == "" {
			mPassword, err = inputs.PromptForMPassword()
			if err != nil {
				return errors.Wrap(err, "cannot prompt for Master password")
			}
		}
		csvFile, err := inputs.GetFlagStringVal(cmd, config.FlagCSVFile)
		if err != nil {
			return errors.Wrapf(err, inputs.ErrMsgCannotGetFlag, config.FlagCSVFile)
		}
		htmlFile, err := inputs.GetFlagStringVal(cmd, config.FlagHTMLFile)
		if err != nil {
			return errors.Wrapf(err, inputs.ErrMsgCannotGetFlag, config.FlagHTMLFile)
		}
		if csvFile == "" && htmlFile == "" {
			return errors.New("must provide a medium to export")
		}
		if csvFile != "" && htmlFile != "" {
			return errors.New("must provide a single medium to export")
		}
		passwordRepo, err := passwords.LoadRepo(mPassword, false)
		if err != nil {
			return errors.Wrap(err, "couldn't initialize password repository")
		}
		if csvFile != "" {
			err = passwordRepo.Export(passwords.CSVExporterID, map[string]string{passwords.ConfKeyCSVFilePath: csvFile})
			if err != nil {
				return errors.Wrap(err, "couldn't export password repository to the CSV file")
			}
		} else {
			err = passwordRepo.Export(passwords.HTMLExporterID, map[string]string{passwords.ConfKeyHTMLFilePath: htmlFile})
			if err != nil {
				return errors.Wrap(err, "couldn't export password repository to the HTML file")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP(config.FlagCSVFile, "c", "", "export passwords to a csv file")
	exportCmd.Flags().StringP(config.FlagHTMLFile, "y", "", "export passwords to a HTML file")
}
