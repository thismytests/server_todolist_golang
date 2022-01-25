package postgressql

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	repo "main/db/repository"
	"testing"
)

// todo ... Mykolai Lytvyn ... will be updated "github.com/google/uuid"
var todolist = &repo.TodoListModel{
	Id:         "TestID",
	Text:       "TestText",
	IsSelected: "TestFalse",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()

	repo := &repository{db}

	defer func() {
		repo.Close()
	}()

	query := "SELECT id, text, isSelected  FROM todolist WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "text", "isSelected"}).
		AddRow(todolist.Id, todolist.Text, todolist.IsSelected)

	mock.ExpectQuery(query).WithArgs(todolist.Id).WillReturnRows(rows)

	user, err := repo.FindByID(todolist.Id)

	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestFindByIDError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, text, isSelected  FROM todolist WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "text"})

	mock.ExpectQuery(query).WithArgs(todolist.Id).WillReturnRows(rows)

	user, err := repo.FindByID(todolist.Id)
	fmt.Println("", err)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestFind(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, text, isSelected FROM todolist"

	rows := sqlmock.NewRows([]string{"id", "text", "isSelected"}).
		AddRow(todolist.Id, todolist.Text, todolist.IsSelected)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.Find()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO todolist \\(text, isSelected\\) VALUES \\(\\$1, \\$2\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(todolist.Text, todolist.IsSelected).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(todolist)
	assert.NoError(t, err)
}

func TestCreateError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO todolist \\(text, isSelected\\) VALUES \\(\\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(todolist.Id, todolist.Text, todolist.IsSelected).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Create(todolist)
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE todolist SET text  = \\?, isSelected = \\?,  WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(todolist.Text, todolist.IsSelected).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(todolist)
	assert.NoError(t, err)
}

func TestUpdateErr(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE user SET name = \\?, email = \\?, phone = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(todolist.Text, todolist.IsSelected).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Update(todolist)
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE  FROM todolist  WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(todolist.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(todolist.Id)
	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE  FROM todolist WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(todolist.Id).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete("notExistId")
	assert.Error(t, err)
}
