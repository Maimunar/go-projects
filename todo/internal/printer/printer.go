package printer

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/maimunar/todo/internal/csv"
	"github.com/mergestat/timediff"
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
	taskTime, err := time.Parse(time.RFC3339, task.CreatedAt)
	if err != nil {
		fmt.Errorf("Error	while parsing time %v", err)
		return ""
	}
	tDiff := timediff.TimeDiff(taskTime)
	if all {
		return fmt.Sprintf("%v\t%v\t%v\t%v\t", task.ID, task.Description, tDiff, task.IsComplete)
	}
	return fmt.Sprintf("%v\t%v\t%v\t", task.ID, task.Description, tDiff)
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
