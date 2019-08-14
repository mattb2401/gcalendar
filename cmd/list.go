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
	"time"

	"github.com/mattb2401/gcalendar/src"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var date string

// initCmd represents the login command
var listCmd = &cobra.Command{
	Use:   "all",
	Short: "List all events that are on my calendar",
	RunE: func(cmd *cobra.Command, args []string) error {
		if date != "" {
			switch date {
			case "today":
				err := src.ListEventsByDate(time.Now().Format("2006-01-02"))
				if err != nil {
					return err
				}
			case "tomorrow":
				today := time.Now()
				tomorrow := today.AddDate(0, 0, 1)
				err := src.ListEventsByDate(tomorrow.Format("2006-01-02"))
				if err != nil {
					return err
				}
			default:
				err := src.ListEventsByDate(date)
				if err != nil {
					return err
				}
			}
		} else {
			err := src.ListAllEvents()
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVarP(&date, "date", "d", "", "Checks for specified day events")
	viper.BindPFlag("date", listCmd.PersistentFlags().Lookup("date"))
}
