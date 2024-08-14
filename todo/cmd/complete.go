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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task from the list of tasks.",
	Long: `Complete a task from the list of tasks.
	The task will be marked as completed in the csv file.
Example: todo complete <taskid>`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}

			err = csv.MarkTaskAsComplete(csv.CsvFileLocation, id)
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}

			fmt.Println("Task completed successfully.")
		}

	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
