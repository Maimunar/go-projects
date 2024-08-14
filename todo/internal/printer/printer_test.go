package printer

import (
	"fmt"
	"testing"
	"time"

	"github.com/maimunar/todo/internal/csv"
	"github.com/mergestat/timediff"
)

func Test_getHeader(t *testing.T) {
	type args struct {
		all bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all is true",
			args: args{all: true},
			want: "ID\tTask\tCreated\tDone\t",
		},
		{
			name: "all is false",
			args: args{all: false},
			want: "ID\tTask\tCreated\t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHeader(tt.args.all); got != tt.want {
				t.Errorf("getHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatTaskCsv(t *testing.T) {
	taskCreated := time.Now().Format(time.RFC3339)
	exampleTask := csv.Task{
		ID:          1,
		Description: "task",
		CreatedAt:   taskCreated,
		IsComplete:  false,
	}
	taskWantedTime, _ := time.Parse(time.RFC3339, taskCreated)

	type args struct {
		task csv.Task
		all  bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal task, all is true",
			args: args{task: exampleTask, all: true},
			want: fmt.Sprintf("1\ttask\t%v\tfalse\t", timediff.TimeDiff(taskWantedTime)),
		},
		{
			name: "normal task, all is false",
			args: args{task: exampleTask, all: false},
			want: fmt.Sprintf("1\ttask\t%v\t", timediff.TimeDiff(taskWantedTime)),
		},
		{
			name: "empty task, all is true",
			args: args{task: csv.Task{}, all: true},
			want: "",
		},
		{
			name: "empty task, all is false",
			args: args{task: csv.Task{}, all: false},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTaskCsv(tt.args.task, tt.args.all); got != tt.want {
				t.Errorf("formatTaskCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}
