package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/knetic/govaluate"
	"github.com/mthsrocha/mb/internal/models"
	"github.com/mthsrocha/mb/internal/server"
)

var Replist []string

func main() {

	softwareOption := softwareMenu()

	switch softwareOption {
	case 1:
		fmt.Println("Welcome to GO REPL!")

		commands := map[string]interface{}{
			"clear": clear,
			"add":   add,
			"rm":    rm,
			"done":  done,
			"list":  list,
		}
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("> ")
		cmd, args := getCl(reader)
		for ; !strings.EqualFold("quit", cmd); cmd, args = getCl(reader) {
			if value, exists := commands[cmd]; exists {
				if cmd == "list" || cmd == "clear" {
					value.(func())()
				} else if cmd == "add" || cmd == "rm" {
					value.(func(string))(args)
				} else {
					args, ok := strconv.Atoi(args)
					if ok != nil {
						log.Printf("invalid value", ok)
					}
					value.(func(int))(args)
				}
			} else {
				cmdValidate(cmd)
			}
			fmt.Print("> ")
		}

	case 2:
		var listId int64
		operation := operationMenu()

		fmt.Println("\nType ID of list:")
		fmt.Scanln(&listId)

		list := models.List{
			Id: listId,
		}

		for {
			var task string

			switch operation {
			case 1:
				fmt.Println("Type the task you want to put on list: ")
				fmt.Scanln(&task)
				addItem(&list, task)
				log.Println(&list, task)

			case 2:
				fmt.Println("Type the task you want to remove from list: ")
				fmt.Scanln(&task)
				removeItem(&list, task)

			case 3:
				fmt.Println("Type the task you want to mark as done: ")
				fmt.Scanln(&task)
				markItemDone(&list, task)
			case 4:
				removeAllItens(&list)
			case 5:
				listItens(&list)
			case 9:
				os.Exit(9)

			default:
				fmt.Println("Invalid option")
			}
		}

	case 3:
		server.HttpServer()
	case 9:
		os.Exit(99)

	default:
		fmt.Println("Invalid option")
	}

	fmt.Println("Starting web application on localhost:5012")

}

//User menu with database
func softwareMenu() int {
	var softwareOption int

	fmt.Println("Choose the way you want to use the software:")
	fmt.Println("1 - REPL")
	fmt.Println("2 - Using User menu")
	fmt.Println("3 - Using web server API")
	fmt.Println("9 - Leave program")
	fmt.Scanln(&softwareOption)

	return softwareOption
}

func operationMenu() int {
	var operation int

	fmt.Println("Choose an operation of the To Do List:")
	fmt.Println("1 - Add item to list")
	fmt.Println("2 - Remove item")
	fmt.Println("3 - Mark as done")
	fmt.Println("4 - Remove all done")
	fmt.Println("5 - List all")
	fmt.Println("9 - Leave program")
	fmt.Scanln(&operation)

	return operation
}

func addItem(l *models.List, task string) {
	err := l.InsertItem(task)
	if err != nil {
		log.Println("Insert error: ", err)
	}

}

func removeItem(l *models.List, task string) {
	err := l.RemoveItem(task)
	if err != nil {
		log.Println("Remove error: ", err)
	}
}

func markItemDone(l *models.List, task string) {
	item := models.Item{
		ListId: l.Id,
		Task:   task,
		Done:   false,
	}

	err := l.UpdateItemDone(item)
	if err != nil {
		log.Println("Update task error: ", err)
	}
}

func removeAllItens(l *models.List) {
	err := l.RemoveAllDoneItem()
	if err != nil {
		log.Println("Remove all done itens error: ", err)
	}
}

func listItens(l *models.List) {
	list, err := l.GetAllItens()
	if err != nil {
		log.Println("List all itens error: ", err)
	}

	for key, value := range list.Items {
		fmt.Sprintf("[%d]%s\n", key, value.Task)
	}
}

// REPL
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func recoverRepl(text string) {
	if r := recover(); r != nil {
		fmt.Println("> unknow command ", text)
	}
}

func cmdValidate(text string) {
	defer recoverRepl(text)

	t := strings.TrimSuffix(text, "\n")
	if t != "" {
		expression, errExp := govaluate.NewEvaluableExpression(text)
		result, errEval := expression.Evaluate(nil)
		if errExp == nil && errEval == nil {
			fmt.Println(">", result)
		} else {
			fmt.Println("> unknow command " + t)
		}
	}
}

func getCl(r *bufio.Reader) (string, string) {
	var command, argument string

	t, _ := r.ReadString('\n')
	fullString := strings.TrimSpace(t)

	ctr := false
	for _, value := range fullString {
		if !ctr && string(value) != " " {
			command += string(value)
		} else {
			ctr = true
			argument += string(value)
		}
	}
	argument = strings.TrimSpace(argument)
	return command, argument
}

func add(s string) {
	Replist = append(Replist, s)
	fmt.Printf("[%d] %s\n", len(Replist)-1, s)
}

func rm(arg string) {
	var index int
	if arg == "done" {
		var RemovedList []string
		for key, value := range Replist {
			i := 0
			if !strings.HasSuffix(value, "[DONE]") {
				RemovedList = append(RemovedList, value)
				i++
			} else {
				fmt.Printf("[%d] %s\n", key, value)
			}
		}
		Replist = RemovedList
		return
	} else {
		index, ok := strconv.Atoi(arg)
		if ok != nil {
			log.Println("Invalid value", index)
			return
		}
	}
	for key, value := range Replist {
		if key == index {
			Replist = append(Replist[:key], Replist[key+1:]...)
			fmt.Printf("[%d] %s\n", key, value)
			return
		}
	}
}

func done(index int) {
	for key, value := range Replist {
		if key == index && !strings.HasSuffix(value, "[DONE]") {
			Replist[key] = value + " [DONE]"
			fmt.Printf("[%d] %s\n", key, Replist[key])
			return
		}
	}
}

func list() {
	if len(Replist) > 0 {
		for key, value := range Replist {
			fmt.Printf("[%d] %s\n", key, value)
		}
	} else {
		fmt.Println("List is empty")
	}
}

func getPosition(s string) int {
	for key, value := range Replist {
		if value == s {
			return key
		}
	}
	return -1
}
