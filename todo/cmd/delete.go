/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/maimunar/todo/internal/csv"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the list of tasks.",
	Long: `Delete a task from the list of tasks.
	The task will be removed from the csv file.
	Example: todo delete <taskid>`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}

			err = csv.DeleteTask(csv.CsvFileLocation, id)
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}

			fmt.Println("Task deleted successfully.")

		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
