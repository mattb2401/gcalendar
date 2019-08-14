package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use: "welcome",
	Short: `
.------..------..------..------..------..------..------..------..------.
|G.--. ||C.--. ||A.--. ||L.--. ||E.--. ||N.--. ||D.--. ||A.--. ||R.--. |
| :/\: || :/\: || (\/) || :/\: || (\/) || :(): || :/\: || (\/) || :(): |
| :\/: || :\/: || :\/: || (__) || :\/: || ()() || (__) || :\/: || ()() |
| '--'G|| '--'C|| '--'A|| '--'L|| '--'E|| '--'N|| '--'D|| '--'A|| '--'R|
.------..------..------..------..------..------..------..------..------.   
Google calendar on command-line`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
