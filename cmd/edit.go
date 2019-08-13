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
	"gcalendar/src"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reoccurance string
var attendies []string
var eventID string

// editCmd represents the login command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits a already existing event on the calendar",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := src.UpdateEventOnCalendar(reoccurance, attendies, eventID)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.PersistentFlags().StringSliceVarP(&attendies, "attendies", "a", []string{}, "Add attendies to the event.")
	editCmd.PersistentFlags().StringVarP(&reoccurance, "reoccurance", "r", "", "Is the event a reoccuraning event.")
	editCmd.PersistentFlags().StringVarP(&eventID, "eventId", "e", "", "EventID on the calendar")
	viper.BindPFlag("reoccurance", editCmd.PersistentFlags().Lookup("reoccurance"))
	viper.BindPFlag("attendies", editCmd.PersistentFlags().Lookup("attendies"))
	viper.BindPFlag("eventId", editCmd.PersistentFlags().Lookup("eventId"))
	editCmd.MarkFlagRequired("eventId")
}
