package main

import (
	"fmt"
	"os"

	"github.com/mthsrocha/mb/internal/models"
	"github.com/mthsrocha/mb/internal/server"
)

type List struct {
	Id    int64
	Items []Item
}

type Item struct {
	ListId int64
	Item   string
	Done   bool
}


func main() {
	
	softwareOption := softwareMenu()
	
	switch softwareOption {
	case 1:
		var listId int64
		operation := operationMenu()

		fmt.Println("\nType ID of list:")
		fmt.Scanln(&listId)

		list := List{
			Id: listId,
		}
		
		for {
			switch operation {
			case 1:
				// add an item
				list.addItem()

			case 2:
				// remove an item

			case 3:
				// mark as done

			case 4:
				// remove all done

			case 5: 
				// list all

			case 9:
				os.Exit(9)

			default:
				fmt.Println("Invalid option")
			}
		}

	case 2:
		server.HttpServer()
	case 9:
		os.Exit(99)

	default:
		fmt.Println("Invalid option")
	}

	fmt.Println("Starting web application on localhost:5012")

}

func softwareMenu() int {
	var softwareOption int

	fmt.Println("Welcome... Choose the way you want to use the software:")
	fmt.Println("1 - Using terminal")
	fmt.Println("2 - Using web server API")
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

func (l *List) addItem(listId int64, task string) error {

	return nil
}

func (l *List) removeItem(listId int64, task string) error {


	return nil
}

func (l *List) markItemDone(listId int64, task string) error {


	return nil
}

func (l *List) removeAllItens() error {


	return nil
}

func (l *List) listItens() {

}