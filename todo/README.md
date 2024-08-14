# Todo App

## Goal

A CLI-based todo list application that allows you to add, list, complete and delete tasks from a data store.

## Getting Started

To get started, clone the repository and build the application:

```
$ go build -o todo
```

Then you can run the application and understand how to use it by running the following:

```
$ ./todo
```

If you do not have a go environment set up, you can use the pre-built binary in the repository by using the command above and skipping the build step.

## How to use

### Add

The add method is used to create new tasks in the underlying data store. It takes a list of arguments with the task description

```
$ tasks add <description>
```

for example:

```
$ tasks add "Tidy my desk"
$ tasks add "Fix my schedule" "Write up documentation for new project feature" "Clean shoes"
```

### List

This method returns a list of all of the **uncompleted** tasks, with the option to return all tasks regardless of whether or not they are completed.

for example:

```
$ tasks list
ID    Task                                                Created
1     Tidy up my desk                                     a minute ago
3     Change my keyboard mapping to use escape/control    a few seconds ago
```

or for showing all tasks, using a flag -a or --all

```
$ tasks list -a
ID    Task                                                Created          Done
1     Tidy up my desk                                     2 minutes ago    false
2     Write up documentation for new project feature      a minute ago     true
3     Change my keyboard mapping to use escape/control    a minute ago     false
```

### Complete

To mark a task as done, use the following method

```
$ tasks complete <taskid>
```

for example

```
$ tasks complete 1
$ tasks complete 2 3
```

### Delete

The following method deletes a task from the data store

```
$ tasks delete <taskid>
```

## Notable Packages Used

- `encoding/csv` for writing out as a csv file
- `strconv` for turning types into strings and visa versa
- `text/tabwriter` for writing out tab aligned output
- `os` for opening and reading files
- `github.com/spf13/cobra` for the command line interface
- `github.com/mergestat/timediff` for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)

### Data File

The Data File used is a CSV file with the following columns:

```
ID,Description,CreatedAt,IsComplete
1,My new task,2024-07-27T16:45:19-05:00,true
2,Finish this video,2024-07-27T16:45:26-05:00,true
3,Find a video editor,2024-07-27T16:45:31-05:00,false
```

## Technical Considerations

### Stderr vs Stdout

Make sure to write any diagnostics or errors to stderr stream and write output to stdout.

### File Locking

One major consideration is that the underlying data file should be locked by the process to prevent concurrent read/writes. This can
be achieved using the flock system call in unix like systems to obtain an exclusive lock on the file.

You can achieve this in go using the following code:

```go
func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

    // Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}
```

Then to unlock the file, use the following:

```go
func closeFile(f *os.File) error {
 syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
 return f.Close()
}
```

## Extra Features for the Future

- Change the IsComplete property of the Task data model to use a timestamp instead, which gives further information.
- Change from CSV to JSON, JSONL or SQLite
- Add in an optional due date to the tasks

## Credit

This project is a modified version of 1 of the projects proposed in this repository by dreamsofcode: https://github.com/dreamsofcode-io/goprojects/tree/main
