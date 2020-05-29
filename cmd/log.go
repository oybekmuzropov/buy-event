package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show all logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		logs, err := logService.GetAll(context.Background())

		if err != nil {
			return err
		}

		for _, log := range logs {
			cmd.Println(`----------**********----------`)
			cmd.Println("ID: ", log.ID)
			cmd.Println("Error: ", log.Error)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

}
