package main

import (
	"fmt"

	"github.com/Lechenco/b_tree/internal/models"
	"github.com/Lechenco/b_tree/internal/services"
)

type Data struct {
	Name string
	Job  string
}

func main() {
	treeService := services.BTreeService[Data]{}
	treeService.InitService(models.BTreeConfig{
		MaxKeysPerNode: 1,
	})

	fmt.Printf("service: %v\n", treeService)
}
