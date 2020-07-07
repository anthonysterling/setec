package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "setec",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	rootCmd.AddCommand(decryptCmd)
	rootCmd.AddCommand(encryptCmd)
	return rootCmd.Execute()
}
