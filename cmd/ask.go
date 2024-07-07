package cmd

import (
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "ask a question and yoo will answer it for you",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(args[0])
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
