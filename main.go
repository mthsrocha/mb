package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/mthsrocha/mb/internal/server"
)

type ToDoList struct {
	Items []Item
}

type Item struct {
	Item string
	Done bool
}

type Database struct {
	Host string
	dbname string
	user string
	passw string
}

func main() {

	fmt.Println("Starting web application")



//	server.HttpServer()
}


func (l *ToDoList) get_AllList() *ToDoList {
	toDo_list := &ToDoList{}

	for _, value := range l.Items {

	}

	return toDo_list
}

func (l *ToDoList) insertItem(s string) (error) {
	item := Item{
		Item: s,
		Done: false,
	}
	l.Items = append(l.Items, item)
	log.Println("Task added to ToDo List.")
	return nil
}

func (l *ToDoList) itemDone(s string, done bool) (error) {

	for _, item := range l.Items {
		if item.Item == s && done == false {
			item.Done = true
			log.Println("Item setted to Done successfully")
			return nil
		}
	}
	return errors.New("Task already done.")

}

func (l *ToDoList) deleteItem() (error) {


	return nil
}

func (l *ToDoList) deleteAllDoneItens() (error) {


	return nil
}
