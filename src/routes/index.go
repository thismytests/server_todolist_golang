package routes

import (
	"github.com/gorilla/mux"
	postgresSql "main/db/postgressql"
	"main/db/repository"
	"main/routes/Auth"
	"main/routes/todolist"
	"net/http"
)

// DI
var repo repository.Repository = postgresSql.DBRepository()

type RegisterMuxRoutes struct {
}

func (registerRoutes RegisterMuxRoutes) Init() {
	router := mux.NewRouter()

	http.Handle("/", router)

	// unprotected route
	todoListInst := todolist.TodoList{
		Repository: repo,
	}
	router.HandleFunc("/todolist", todoListInst.GetTodoListItems)
	// Todolist CRUD
	router.HandleFunc("/todolist/{id:[0-9]+}", todoListInst.CreateDeleteUpdateItemsById)

	// JWT Auth
	router.HandleFunc("/authenticate", Auth.Authenticate)
}
