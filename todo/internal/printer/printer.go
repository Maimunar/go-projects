package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/maimunar/todo/internal/csv"
)

const header = "ID\tTask\tCreated\t"
const headerAll = "ID\tTask\tCreated\tDone\t"

func getHeader(all bool) string {
	if all {
		return headerAll
	}
	return header
}

func formatTaskCsv(task csv.Task, all bool) string {
	if all {
		return fmt.Sprintf("%v\t%v\t%v\t%v\t", task.ID, task.Description, task.CreatedAt, task.IsComplete)
	}
	return fmt.Sprintf("%v\t%v\t%v\t", task.ID, task.Description, task.CreatedAt)
}

func PrintTasks(tasks []csv.Task, all bool) {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 16, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, getHeader(all))
	for _, task := range tasks {
		fmt.Fprintln(w, formatTaskCsv(task, all))
	}
	w.Flush()
}
