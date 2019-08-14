// Copyright © 2019 SEBUUMA MATT <mttsebuuma@gmail.com>
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
	"github.com/mattb2401/gcalendar/src"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the login command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add new note to your calendar",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := src.AddEventToCalendar(reoccurance, attendies)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().StringSliceVarP(&attendies, "attendies", "a", []string{}, "Add attendies to the event.")
	newCmd.PersistentFlags().StringVarP(&reoccurance, "reoccurance", "r", "", "Is the event a reoccuraning event.")
	viper.BindPFlag("reoccurance", newCmd.PersistentFlags().Lookup("reoccurance"))
	viper.BindPFlag("attendies", newCmd.PersistentFlags().Lookup("attendies"))
}
