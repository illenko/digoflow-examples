package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/illenko/digoflow"
	"github.com/illenko/digoflow/container"
	"github.com/illenko/digoflow/task"
)

func main() {

	app, err := digoflow.NewApp("flows", "migrations")

	if err != nil {
		fmt.Printf("Error loading app: %v", err)
		return
	}

	app.RegisterTask("uuidGenerator", &UUIDGenerator{})

	err = app.Start()

	if err != nil {
		fmt.Printf("Error starting app: %v", err)
		return
	}
}

type UUIDGenerator struct{}

func (t *UUIDGenerator) Execute(_ *container.Container, _ task.Input) (task.Output, error) {
	return task.Output{"result": uuid.New().String()}, nil
}
