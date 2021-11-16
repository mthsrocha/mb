package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"log"

	"github.com/mthsrocha/mb/internal/models"
)


func Handlers() {
	http.HandleFunc("/list/", getAllList)
	http.HandleFunc("/insert/", postItem)
	http.HandleFunc("/remove/", removeItem)
	http.HandleFunc("/remove/all/", removeAllItem)
	http.HandleFunc("/update/", patchItemTask)
	http.HandleFunc("update/done/", patchItemDone)
}

func getAllList(w http.ResponseWriter, r *http.Request) {
	queryParmId := r.URL.Query().Get("listId")
	listId, err := strconv.ParseInt(queryParmId, 10, 64)
	if err != nil {
		log.Println("ListID cant be converted")
	}
	list := models.List{
		Id: listId,
	}
	responseAllItens, err := list.GetAllItens()
	if err != nil {
		log.Println("Error: ", err)
	}

	for key, value := range responseAllItens.Items {
		fmt.Sprintf("[%d]%s\n", key, value.Task)
	}
}

func postItem(w http.ResponseWriter, r *http.Request) {
	queryParmId := r.URL.Query().Get("listId")
	queryParmItem := r.URL.Query().Get("item")
	listId, err := strconv.ParseInt(queryParmId, 10, 64)
	if err != nil {
		log.Println("ListID cant be converted")
	}

	list := models.List{
		Id: listId,
	}
	ok := list.InsertItem(queryParmItem)
	if ok != nil {
		log.Fatalf("Error: ", ok)
	}
}

func removeItem(w http.ResponseWriter, r *http.Request) {
	queryParmId := r.URL.Query().Get("listId")
	queryParmItem := r.URL.Query().Get("item")
	listId, err := strconv.ParseInt(queryParmId, 10, 64)
	if err != nil {
		log.Println("ListID cant be converted")
	}

	list := models.List{
		Id: listId,
	}
	ok := list.RemoveItem(queryParmItem)
	if ok != nil {
		log.Fatalf("Error: ", ok)
	}
}

func removeAllItem(w http.ResponseWriter, r *http.Request) {
	queryParmId := r.URL.Query().Get("listId")
	listId, err := strconv.ParseInt(queryParmId, 10, 64)
	if err != nil {
		log.Println("ListID cant be converted")
	}

	list := models.List{
		Id: listId,
	}
	ok := list.RemoveAllDoneItem()
	if ok != nil {
		log.Fatalf("Error: ", ok)
	}
}

func patchItemDone(w http.ResponseWriter, r *http.Request) {
	queryParmId := r.URL.Query().Get("listId")
	queryParmItem := r.URL.Query().Get("item")
	listId, err := strconv.ParseInt(queryParmId, 10, 64)
	if err != nil {
		log.Println("ListID cant be converted")
	}

	list := models.List{
		Id: listId,
	}
	item := models.Item{
		ListId: listId,
		Task: queryParmItem,
		Done: false,
	}

	ok := list.UpdateItemDone(item)
	if ok != nil {
		log.Fatalf("Error: ", ok)
	}
}

func patchItemTask(w http.ResponseWriter, r *http.Request) {
	queryParmId := r.URL.Query().Get("listId")
	queryParmItem := r.URL.Query().Get("item")
	listId, err := strconv.ParseInt(queryParmId, 10, 64)
	if err != nil {
		log.Println("ListID cant be converted")
	}

	list := models.List{
		Id: listId,
	}
	item := models.Item{
		ListId: listId,
		Task: queryParmItem,
		Done: false,
	}

	ok := list.UpdateItemTask(item, queryParmItem)
	if ok != nil {
		log.Fatalf("Error: ", ok)
	}
}