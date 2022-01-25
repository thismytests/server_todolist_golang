package todolist

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
	"main/db/repository"
	"net/http"
	"strings"
)

type TodoList struct {
	Repository repository.Repository
}

func getId(str string) string {
	result := strings.Split(str, "/")
	id := result[len(result)-1]
	return id
}

func (todoList TodoList) getTodolistItemById(writer http.ResponseWriter, request *http.Request) {
	id := getId(request.URL.Path)

	item, err := todoList.Repository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}

	json := simplejson.New()
	json.Set("item", item)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payload)
}

func (todoList TodoList) createTodoListItem(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")

	var todoListItem *repository.TodoListModel
	_ = json.NewDecoder(request.Body).Decode(&todoListItem)

	err := todoList.Repository.Create(todoListItem)

	if err != nil {
		fmt.Println(err)
	}

	// todo ... Mykolai Lytvyn ... will be updated
	json.NewEncoder(writer).Encode(todoListItem)
}

func (todoList TodoList) updateTodolistItemsById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")

	var todoListItem *repository.TodoListModel
	_ = json.NewDecoder(request.Body).Decode(&todoListItem)

	err := todoList.Repository.Update(todoListItem)

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(writer).Encode(todoListItem)
}

func (todoList TodoList) deleteTodolistItemsById(writer http.ResponseWriter, request *http.Request) {
	id := getId(request.URL.Path)

	err := todoList.Repository.Delete(id)

	if err != nil {
		fmt.Println("err", err)
	}

	// todo ... Mykolai Lytvyn ... will be updated
	//json.NewEncoder(writer).Encode(id)

	json := simplejson.New()
	json.Set("foo", "bar")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payload)
}

func (todoList TodoList) GetTodoListItems(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")

	if request.Method == "GET" {
		items, err := todoList.Repository.Find()

		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(writer).Encode(items)

	}
}

func (todoList TodoList) CreateDeleteUpdateItemsById(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "PUT":
		todoList.createTodoListItem(writer, request)
	case "GET":
		todoList.getTodolistItemById(writer, request)
	case "POST":
		todoList.updateTodolistItemsById(writer, request)
	case "DELETE":
		todoList.deleteTodolistItemsById(writer, request)
	default:
		fmt.Println("route was not found")
		// todo ... Mykolai Lytvyn ... will be updated with status
		writer.Write([]byte(`{"message": "Not Found"}`))
	}

}
