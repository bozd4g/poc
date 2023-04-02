package main

import (
	"context"
	"fmt"
	"log"
	"poc-testing/client/internal/todo"
)

func main()  {
	ctx := context.Background()

	todoClient := todo.NewClient("http://localhost:3000")
	todos, err := todoClient.GetTodos(ctx)
	if err != nil {
		log.Fatal("failed to get todos", err)
	}

	fmt.Println(todos)
}