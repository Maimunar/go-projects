package csv

import (
	"os"
	"reflect"
	"testing"
	"time"
)

var exampleTask = Task{
	ID:          1,
	Description: "task",
	CreatedAt:   "2021-01-01",
	IsComplete:  false,
}

const csvFileLocationUnit = "testunit.csv"
const csvFileLocation = "test.csv"

func Test_parseTask(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name    string
		args    args
		want    Task
		wantErr bool
	}{
		{
			name: "parse example task",
			args: args{data: []string{"1", "task", "2021-01-01", "false"}},
			want: exampleTask,
		},
		{
			name:    "error on parse empty task",
			args:    args{data: []string{}},
			wantErr: true,
		},
		{
			name:    "error on parse more fields than expected",
			args:    args{data: []string{"1", "task", "2021-01-01", "false", "extra"}},
			wantErr: true,
		},
		{
			name:    "error on parse invalid id",
			args:    args{data: []string{"a", "task", "2021-01-01", "false"}},
			wantErr: true,
		},
		{
			name:    "error on parse invalid isComplete",
			args:    args{data: []string{"1", "task", "2021-01-01", "invalid"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTask(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createTask(t *testing.T) {
	type args struct {
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    Task
		wantErr bool
	}{
		{
			name: "create task",
			args: args{description: "task"},
			want: Task{
				ID:          0,
				Description: "task",
				CreatedAt:   time.Now().Format(time.DateOnly),
				IsComplete:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createTask(csvFileLocationUnit, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("createTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTask() = %v, want %v", got, tt.want)
			}
		})
	}
	err := os.Remove(csvFileLocationUnit)
	if err != nil {
		t.Errorf("createTask() error on cleanup, couldnt delete file: %v", err)
	}
}

func Test_convertTaskForCSV(t *testing.T) {
	type args struct {
		task Task
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "convert example task",
			args: args{task: exampleTask},
			want: []string{"1", "task", "2021-01-01", "false"},
		},
		{
			name: "convert empty task",
			args: args{task: Task{}},
			want: []string{"0", "", "", "false"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTaskForCSV(tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTaskForCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FullIntegrationTest(t *testing.T) {
	var wantedTask = Task{
		ID:          0,
		Description: "task",
		CreatedAt:   time.Now().Format(time.DateOnly),
		IsComplete:  false,
	}
	// Create a task
	err := AddTask(csvFileLocation, wantedTask.Description)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on createTask: %v", err)
		return
	}

	tasks, err := GetTasks(csvFileLocation, true)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on GetTasks: %v", err)
		return
	}

	if len(tasks) != 1 {
		t.Errorf("FullIntegrationTest() error on GetTasks, expected 1 task, got %v", len(tasks))
		return
	}

	if !reflect.DeepEqual(tasks[0], wantedTask) {
		t.Errorf("FullintegrationTest() error on GetTasks %v, wanted %v", tasks[0], wantedTask)
		return
	}

	tasks, err = GetTasks(csvFileLocation, false)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on GetTasks: %v", err)
		return
	}

	if len(tasks) != 1 {
		t.Errorf("FullIntegrationTest() error on GetTasks without completed, expected 1 task, got %v", len(tasks))
		return
	}

	// Mark task as complete
	err = MarkTaskAsComplete(csvFileLocation, wantedTask.ID)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on MarkTaskAsComplete: %v", err)
		return
	}

	wantedTask.IsComplete = true

	tasks, err = GetTasks(csvFileLocation, true)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on GetTasks: %v", err)
		return
	}

	if len(tasks) != 1 {
		t.Errorf("FullIntegrationTest() error on GetTasks, expected 1 task, got %v", len(tasks))
		return
	}

	if !reflect.DeepEqual(tasks[0], wantedTask) {
		t.Errorf("FullintegrationTest() error on GetTasks, %v, wanted %v", tasks[0], wantedTask)
		return
	}

	tasks, err = GetTasks(csvFileLocation, false)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on GetTasks: %v", err)
		return
	}

	if len(tasks) != 0 {
		t.Errorf("FullIntegrationTest() error on GetTasks without completed, expected 0 tasks, got %v", len(tasks))
		return
	}
	// Delete task
	err = DeleteTask(csvFileLocation, wantedTask.ID)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on DeleteTask: %v", err)
		return
	}

	tasks, err = GetTasks(csvFileLocation, true)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on GetTasks: %v", err)
		return
	}

	if len(tasks) != 0 {
		t.Errorf("FullIntegrationTest() error on GetTasks, expected 0 tasks, got %v", len(tasks))
		return
	}

	tasks, err = GetTasks(csvFileLocation, false)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on GetTasks: %v", err)
		return
	}

	if len(tasks) != 0 {
		t.Errorf("FullIntegrationTest() error on GetTasks without completed, expected 0 task, got %v", len(tasks))
		return
	}

	// Cleanup
	err = os.Remove(csvFileLocation)
	if err != nil {
		t.Errorf("FullIntegrationTest() error on cleanup, couldnt delete file: %v", err)
	}
}
