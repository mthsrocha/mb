package models

import (
	"errors"
	"log"

	"github.com/mthsrocha/mb/internal/database"
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

func (l *List) GetAllItens() (*List, error) {
	db := database.Connect()
	defer db.Close()

	select_sql := "SELECT * FROM to_do_list WHERE ListId=$1"

	queryset, err := db.Query(select_sql, l.Id)
	if err != nil {
		log.Println("Query error: ", err)
		return nil, err
	}

	toDo_list := &List{}
	toDo_item := &Item{}

	for queryset.Next() {
		var item string
		var done bool

		err = queryset.Scan(&item, &done)
		if err != nil {
			log.Println("Query error: ", err)
			return nil, err
		}

		toDo_item.Item = item
		toDo_item.Done = done

		toDo_list.Items = append(toDo_list.Items, *toDo_item)
	}
	return toDo_list, nil
}

func (l *List) InsertItem(s string) error {
	db := database.Connect()
	defer db.Close()

	insert_sql := "INSERT INTO to_do_list(ListId, Item, Done) VALUES ($1, $2, $3)"

	item := Item{
		ListId: l.Id,
		Item: s,
		Done: false,
	}

	queryset, err := db.Prepare(insert_sql)
	if err != nil {
		log.Println("Query error: ", err)
		return err
	}

	queryset.Exec(item.ListId, item.Item, item.Done)
	log.Println("Task added to ToDo List.")
	return nil
}

func (l *List) RemoveItem(s string) error {
	db := database.Connect()
	defer db.Close()

	sql_delete := "DELETE FROM to_do_list WHERE ListId=$1 and Item=$2"
	queryset, err := db.Prepare(sql_delete)
	if err != nil {
		log.Println("Query error: ", err)
		return err
	}

	queryset.Exec(l.Id, s)
	return nil
}

func (l *List) RemoveAllDoneItem() error {
	db := database.Connect()
	defer db.Close()

	delete_all_sql := "DELETE FROM to_do_list WHERE ListId=$1 and Done=$2"

	queryset, err := db.Prepare(delete_all_sql)
	if err != nil {
		log.Println("Query error: ", err)
		return err
	}

	queryset.Exec(l.Id, true)
	return nil
}

func (l *List) UpdateItemDone(item Item) error {
	db := database.Connect()
	defer db.Close()

	if item.Done == true {
		return errors.New("Task already done.")
	}
	sql_update := "UPDATE to_do_list SET Done=$1 WHERE ListId=$2"

	queryset, err := db.Prepare(sql_update)
	if err != nil {
		log.Println("Query error: ", err)
		return err
	}

	queryset.Exec(item.Done, l.Id)
	return nil
}

func (l *List) UpdateItemTask(item Item, s string) error {
	db := database.Connect()
	defer db.Close()

	for _, item := range l.Items{
		if item.Item == s {
			return errors.New("Task already Created.")
		}
	}

	sql_update := "UPDATE to_do_list SET Item=$1 WHERE ListId=$2"
	queryset, err := db.Prepare(sql_update)
	if err != nil {
		log.Println("Query error: ", err)
		return err
	}
	queryset.Exec(item.Item, l.Id)
	return nil
}
