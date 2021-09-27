package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version for this command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0")
	},
}
