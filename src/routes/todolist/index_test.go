package todolist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"main/db/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetId(t *testing.T) {
	id := getId("/todolist/12")

	expected := "12"

	if id != expected {
		t.Errorf("Expected %s but got %s", expected, id)
	}
}

func Test_GetTodolistItemsById(t *testing.T) {
	assert.Equal(t, true, false)
}
func Test_GetTodoListItems(t *testing.T) {
	assert.Equal(t, true, false)
}

func Test_CreateTodoListItem(t *testing.T) {
	urlAnotherUrl := fmt.Sprintf("/todolist/%d", 12)

	req, err := http.NewRequest("GET", urlAnotherUrl, nil)

	if err != nil {
		t.Fail()
	}

	rr := httptest.NewRecorder()

	// create tested route
	handler := http.HandlerFunc(TodoList{
		Repository: repository.GetMockRepository(),
	}.createTodoListItem)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `TodoList 12`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	fmt.Print("Work")
}

func Test_UpdateTodolistItemsById(t *testing.T) {
	assert.Equal(t, true, false)
}
func Test_DeleteTodolistItemsById(t *testing.T) {
	assert.Equal(t, true, false)
}
