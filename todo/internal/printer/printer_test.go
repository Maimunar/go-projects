package printer

import (
	"testing"

	"github.com/maimunar/todo/internal/csv"
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
	exampleTask := csv.Task{
		ID:          1,
		Description: "task",
		CreatedAt:   "2021-01-01",
		IsComplete:  false,
	}

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
			want: "1\ttask\t2021-01-01\tfalse\t",
		},
		{
			name: "normal task, all is false",
			args: args{task: exampleTask, all: false},
			want: "1\ttask\t2021-01-01\t",
		},
		{
			name: "empty task, all is true",
			args: args{task: csv.Task{}, all: true},
			want: "0\t\t\tfalse\t",
		},
		{
			name: "empty task, all is false",
			args: args{task: csv.Task{}, all: false},
			want: "0\t\t\t",
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
