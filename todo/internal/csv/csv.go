package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/maimunar/todo/internal/files"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   string
	IsComplete  bool
}

const CsvFileLocation = "tasks.csv"

func parseTask(data []string) (Task, error) {
	if len(data) != 4 {
		return Task{}, fmt.Errorf("invalid data")
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		return Task{}, fmt.Errorf("invalid id")
	}

	isComplete, err := strconv.ParseBool(data[3])
	if err != nil {
		return Task{}, fmt.Errorf("invalid isComplete")
	}

	return Task{
		ID:          id,
		Description: data[1],
		CreatedAt:   data[2],
		IsComplete:  isComplete,
	}, nil
}

func createTask(csvFile string, description string) (Task, error) {
	tasks, err := GetTasks(csvFile, true)
	if err != nil {
		return Task{}, err
	}
	id := 0
	for _, task := range tasks {
		if task.ID >= id {
			id = task.ID + 1
		}
	}

	createdAt := time.Now().Format(time.DateOnly)

	return Task{
		ID:          id,
		Description: description,
		CreatedAt:   createdAt,
		IsComplete:  false,
	}, nil
}

func convertTaskForCSV(task Task) []string {
	return []string{strconv.Itoa(task.ID), task.Description, task.CreatedAt, strconv.FormatBool(task.IsComplete)}
}

func GetTasks(csvFile string, all bool) ([]Task, error) {
	file, err := files.OpenFile(csvFile, 0)

	if err != nil {
		return nil, err
	}

	defer files.CloseFile(file)

	csvReader := csv.NewReader(file)
	tasks := []Task{}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		task, err := parseTask(record)
		if err != nil {
			return nil, err
		}

		// If with all tag, add all, otherwise only incomplete
		if all || !task.IsComplete {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func AddTask(csvFile string, description string) error {
	task, err := createTask(csvFile, description)
	if err != nil {
		return err
	}

	file, err := files.OpenFile(csvFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE)

	if err != nil {
		return err
	}

	defer files.CloseFile(file)

	taskCsv := convertTaskForCSV(task)

	csvWriter := csv.NewWriter(file)
	err = csvWriter.Write(taskCsv)
	csvWriter.Flush()

	return err
}

func DeleteTask(csvFile string, id int) error {
	tasks, err := GetTasks(csvFile, true)
	if err != nil {
		return err
	}

	file, err := files.OpenFile(csvFile, os.O_WRONLY|os.O_TRUNC)

	if err != nil {
		return err
	}

	defer files.CloseFile(file)

	csvWriter := csv.NewWriter(file)
	for _, task := range tasks {
		if task.ID == id {
			continue
		}
		csvWriter.Write(convertTaskForCSV(task))
	}
	csvWriter.Flush()

	return err
}

func MarkTaskAsComplete(csvFile string, id int) error {
	tasks, err := GetTasks(csvFile, true)
	if err != nil {
		return err
	}

	file, err := files.OpenFile(csvFile, os.O_WRONLY|os.O_TRUNC)

	if err != nil {
		return err
	}

	defer files.CloseFile(file)

	tasksCSV := [][]string{}

	for _, task := range tasks {
		if task.ID == id {
			task.IsComplete = true
		}
		tasksCSV = append(tasksCSV, convertTaskForCSV(task))
	}

	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(tasksCSV)
	return err
}
