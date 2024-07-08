package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yoo",
	Short: "yoo is a OpenAI CLI tool",
	Long: `  __  ______  ____ 
  / / / / __ \/ __ \
 / /_/ / /_/ / /_/ /
 \__, /\____/\____/ 		
/____/
is a OpenAI CLI tool that allows you to interact with OpenAI's API.
yoo is easy and fast!🚀`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
