package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ahmedMunna1767/tasks_db"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "task",
	Short: `
-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	    --------------------------------------
	  ||πππTask is a CLI task managerπππ||
	    --------------------------------------
-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
`,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "β½β½β½ Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		_, err = tasks_db.CreateTask(task, taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		fmt.Printf("\n	Added \"%s\" to your task list πππ \n\n", task)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "β½β½β½ Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("")
			fmt.Println("	πππππ You have no tasks to complete! Why not take a vacation? πππππ")
			fmt.Println("")
			return
		}
		fmt.Println("\n    You have the following tasks:")
		for i, task := range tasks {
			taskFields := strings.Split(task.Value, "\n")
			var completed string
			if taskFields[2] == "false" {
				completed = "Hurry Up π Nay π"
			} else {
				completed = "Completed π Yay π"
			}
			fmt.Printf("\tπππ %d. %s (%s) (%s)\n", i+1, taskFields[0], taskFields[1][0:11], completed)
		}
		fmt.Println("")
	},
}

var deleteCmd = &cobra.Command{

	Use:   "delete",
	Short: "β½β½β½ Delete Task/s",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg, "π₯π₯π₯")
			} else {
				ids = append(ids, id)
			}
		}

		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("	Invalid task number: ", id, "π₯π₯π₯")
				continue
			}
			task := tasks[id-1]
			err := tasks_db.DeleteTask(task.Key, taskBucket, db)
			if err != nil {
				fmt.Printf("	Failed to delete Task: \"%d\" . Error: %s, π₯π₯π₯\n\n", id, err)
			} else {
				fmt.Printf("	Deleted Task No: \"%d\" πππ\n\n", id)
			}
		}
	},
}

var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "β½β½β½ See details of task/s",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg, "π₯π₯π₯")
			} else {
				ids = append(ids, id)
			}
		}

		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("	Invalid task number:", id, "π₯π₯π₯")
				continue
			}
			task := tasks[id-1]
			taskFields := strings.Split(task.Value, "\n")
			var completed string
			if taskFields[2] == "false" {
				completed = "Hurry Up π Nay π"
			} else {
				completed = "Completed π Yay ππππ"
			}
			fmt.Printf("\n   Showing details for task %d πππ \n", id)
			fmt.Printf("\tCreated at : %s\n", taskFields[1])
			fmt.Printf("\tTask Details : %s\n", taskFields[0])
			fmt.Printf("\tStatus : %s\n", completed)
			if taskFields[2] == "true" {
				fmt.Printf("\tCompleted At : %s (ππππ) \n", taskFields[4])
			}
			fmt.Println()
		}
	},
}

var chStatusCmd = &cobra.Command{
	Use:   "done",
	Short: "β½β½β½ Change Completion Status of task/s ",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg, "π₯π₯π₯")
			} else {
				ids = append(ids, id)
			}
		}

		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id, "π₯π₯π₯")
				continue
			}
			task := tasks[id-1]
			taskFields := strings.Split(task.Value, "\n")
			// var completed string
			if taskFields[2] == "false" {
				taskFields[2] = "true"
			} else {
				taskFields[2] = "false"
			}
			_, err = tasks_db.UpdateTask(taskFields[0]+"\n"+taskFields[1]+"\n"+taskFields[2]+"\n", task.Key, taskBucket, db)
			if err != nil {
				fmt.Println("Something wrong Happened", "π₯π₯π₯")
			} else {
				fmt.Printf("Successfully changed Status of task %d πππ", id)
			}
			fmt.Println()
		}
	},
}

var showCompleted = &cobra.Command{
	Use:   "completed",
	Short: "β½β½β½ List of Completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		if len(tasks) == 0 {
			fmt.Println("")
			fmt.Println("	πππππ You have no tasks to complete! Why not take a vacation? πππππ")
			fmt.Println("")
			return
		}
		taskNotCompleted := 0
		for id, task := range tasks {
			taskFields := strings.Split(task.Value, "\n")
			var completed string
			if taskFields[2] == "false" {
				completed = "Hurry Up π Nay π"
				continue
			} else {
				completed = "Completed π Yay ππππ"
				taskNotCompleted++
			}
			fmt.Printf("\n   Showing details for task %d πππ \n", id+1)
			fmt.Printf("\tCreated at : %s\n", taskFields[1])
			fmt.Printf("\tTask Details : %s\n", taskFields[0])
			fmt.Printf("\tStatus : %s\n", completed)
			if taskFields[2] == "true" {
				fmt.Printf("\tCompleted At : %s (ππππ) \n", taskFields[4])
			}
			fmt.Println()
		}

		if taskNotCompleted == 0 {
			fmt.Printf("\n\t(ππππ) None of the Tasks are Completed (ππππ) \n\n")
		} else {

			fmt.Printf("\n\tYou have Completed %d of %d tasks Today.\n\t(ππππ) Congratulations (ππππ) \n\n", taskNotCompleted, len(tasks))
		}

	},
}

var showNotCompleted = &cobra.Command{
	Use:   "!completed",
	Short: "β½β½β½ List of UnCompleted tasks",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		if len(tasks) == 0 {
			fmt.Println("")
			fmt.Println("	πππππ You have no tasks to complete! Why not take a vacation? πππππ")
			fmt.Println("")
			return
		}
		taskNotCompleted := 0
		for id, task := range tasks {
			taskFields := strings.Split(task.Value, "\n")
			var completed string
			if taskFields[2] == "false" {
				completed = "Hurry Up π Nay π"
				taskNotCompleted++
			} else {
				completed = "Completed π Yay ππππ"
				continue
			}
			fmt.Printf("\n   Showing details for task %d πππ \n", id+1)
			fmt.Printf("\tCreated at : %s\n", taskFields[1])
			fmt.Printf("\tTask Details : %s\n", taskFields[0])
			fmt.Printf("\tStatus : %s\n", completed)
			if taskFields[2] == "true" {
				fmt.Printf("\tCompleted At : %s (ππππ) \n", taskFields[4])
			}
			fmt.Println()
		}
		if taskNotCompleted == 0 {
			fmt.Printf("\n\t(ππππ) Congratulations. All of Your Tasks are Completed (ππππ) \n\n")
		} else {
			fmt.Printf("\n\tYou have %d of %d tasks left to Complete Today.\n\t(ππππ) Hurry Up (ππππ) \n\n", taskNotCompleted, len(tasks))
		}

	},
}

var doneForTheDayCmd = &cobra.Command{
	Use:   "donefortheday",
	Short: "β½β½β½ Delete all Task/s Completed or not ",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		if len(tasks) == 0 {
			fmt.Println("")
			fmt.Println("	πππππ You have no tasks to complete / delete! Why not take a vacation? πππππ")
			fmt.Println("")
			return
		}
		fmt.Println("")
		for id, task := range tasks {
			err := tasks_db.DeleteTask(task.Key, taskBucket, db)
			if err != nil {
				fmt.Printf("	Failed to delete Task: \"%d\" . Error: %s, π₯π₯π₯\n", id+1, err)
			} else {
				fmt.Printf("	Deleted Task No: \"%d\" πππ\n", id+1)
			}
		}
		fmt.Println("	All Tasks are Deleted \n	(πππ) See You Tomorrow (πππ)")
		fmt.Println("")
	},
}

var nextUnCompleted = &cobra.Command{
	Use:   "next",
	Short: "β½β½β½ Show next uncompleted task ",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "tasks.db")
		taskBucket, db, err := tasks_db.Init(dbPath)
		if err != nil {
			panic(err)
		}
		tasks, err := tasks_db.AllTasks(taskBucket, db)
		if err != nil {
			fmt.Println("Something went wrong:", err, "π₯π₯π₯")
			return
		}
		if len(tasks) == 0 {
			fmt.Println("")
			fmt.Println("	πππππ You have no tasks to complete! Why not take a vacation? πππππ")
			fmt.Println("")
			return
		}
		taskNotCompleted := 0
		for _, task := range tasks {
			taskFields := strings.Split(task.Value, "\n")
			var completed string
			if taskFields[2] == "false" {
				completed = "Hurry Up π Nay π"
				taskNotCompleted++
			} else {
				completed = "Completed π Yay ππππ"
				continue
			}
			fmt.Printf("\n   πππ Next task to complete in your list πππ \n")
			fmt.Printf("\tCreated at : %s\n", taskFields[1])
			fmt.Printf("\tTask Details : %s\n", taskFields[0])
			fmt.Printf("\tStatus : %s\n", completed)
			if taskFields[2] == "true" {
				fmt.Printf("\tCompleted At : %s (ππππ) \n", taskFields[4])
			}
			fmt.Println()
			break
		}
		if taskNotCompleted == 0 {
			fmt.Printf("\n\t(ππππ) Congratulations. All of Your Tasks are Completed (ππππ) \n\n")
		}
		// } else {
		// 	fmt.Printf("\n\tYou have %d of %d tasks left to Complete Today.\n\t(ππππ) Hurry Up (ππππ) \n\n", taskNotCompleted, len(tasks))
		// }

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(detailsCmd)
	RootCmd.AddCommand(chStatusCmd)
	RootCmd.AddCommand(showCompleted)
	RootCmd.AddCommand(showNotCompleted)
	RootCmd.AddCommand(doneForTheDayCmd)
	RootCmd.AddCommand(nextUnCompleted)
}
