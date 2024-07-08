package cmd

import (
	"github.com/iarsham/yoo/internal"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask \"question\"",
	Short: "ask a question and yoo will answer it for you",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pterm.DefaultSpinner.Start("Thinking...📩")
		resp, err := internal.RspOpenAI(args[0])
		if err != nil {
			return err
		}
		pterm.DefaultBox.WithTitle("Answer").WithTextStyle(pterm.NewStyle(pterm.FgRed, pterm.BgWhite, pterm.Bold)).Println(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
