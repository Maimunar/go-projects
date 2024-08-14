/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/maimunar/todo/internal/csv"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list of tasks.",
	Long: `Add a new task to the list of tasks. 
	The task will be saved in the csv file.
	Format: todo add <description>,
	Example: todo add "Buy milk"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, desc := range args {
			if desc == "" {
				continue
			}
			err := csv.AddTask(csv.CsvFileLocation, desc)
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}
			fmt.Println("Task added successfully.")
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
