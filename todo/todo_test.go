package todo

import (
	"testing"
)

func TestCompleteForFoundTodo(t *testing.T) {
	app := CreateTodoApp()
	todo := app.Add("Wake up early")
	if todo.Completed {
		t.Error("Expected completed is false, but it is true")
	}

	err := app.Complete(todo.Id)

	if err != nil {
		t.Errorf("Expected error is nil, but it is %v", err)
	}

	if todo.Completed {
		t.Error("Expected completed is true, but it is false")
	}
}

func TestCompleteForNotFoundTodo(t *testing.T) {
	app := CreateTodoApp()
	err := app.Complete(1)
	if err == nil {
		t.Errorf("Expected error is not nil, but it is %v", err)
	}

}

