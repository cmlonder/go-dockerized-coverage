// Package todo provides a simple todo application
package todo

import "errors"

// Todo representation of todo with id, title and completed status
type Todo struct {
	Id       int
	Title    string
	Completed bool
}

// TodoApp holds todo list and id counter
type TodoApp struct {
	todos map[int]Todo
	id int
}

// CreateTodoApp creates a new TodoApp by initializing todo list and id counter
func CreateTodoApp() TodoApp {
	todoApp := TodoApp{}
	todoApp.todos = make(map[int]Todo)
	todoApp.id = 0
	return todoApp
}

// Add adds a new todo with given title and returns created todo with a new id
func (app *TodoApp) Add(title string) Todo {
	app.id = app.id + 1
	app.todos[app.id] = Todo{app.id, title, false}
	return app.todos[app.id]
}

// Complete completes a todo with given id and
// returns todo not found exception if todo with given id is not found
func (app *TodoApp) Complete(id int) error {
	if todo, ok := app.todos[id]; ok {
		todo.Completed = true
		app.todos[id] = todo
		return nil
	} else {
		return errors.New("todo not found")
	}
}
