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
		MaxElementsPerNode: 2,
	})

	treeService.Add(models.Element[Data]{
		Key: 1,
		Data: Data{
			Name: "teste",
			Job:  "ifood",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 2,
		Data: Data{
			Name: "teste2",
			Job:  "ifood2",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 3,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 4,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 5,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 6,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 7,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 12,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})
	treeService.Add(models.Element[Data]{
		Key: 13,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})
	treeService.Add(models.Element[Data]{
		Key: 14,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 8,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 10,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 11,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	treeService.Add(models.Element[Data]{
		Key: 9,
		Data: Data{
			Name: "teste3",
			Job:  "99",
		},
	})

	fmt.Println("service: ", treeService)
}
