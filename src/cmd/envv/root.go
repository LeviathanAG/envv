package envv

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "envv",
	Short: "envv is a tool to manage .env files using MongoDB",
}

func Execute() error {
	return rootCmd.Execute()
}
