/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/maimunar/todo/internal/csv"
	"github.com/maimunar/todo/internal/printer"
	"github.com/spf13/cobra"
)

var all bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List uncompleted tasks from the list of tasks.",
	Long: `List uncompleted tasks from the list of tasks.
	The List is taken from the csv file.
	Example: todo list
	Can also use -a or --all tag to also show the completed tasks.
	Example: todo list -a`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := csv.GetTasks(csv.CsvFileLocation, all)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}

		printer.PrintTasks(tasks, all)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all tasks including completed tasks.")
}
