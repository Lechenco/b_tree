package main

import (
	"fmt"
	"math/rand"

	"github.com/Lechenco/b_tree/internal/models"
	"github.com/Lechenco/b_tree/internal/services"
)

type Data struct {
	Name string
	Job  string
}

func generateElement(key int) models.Element[Data] {
	return models.Element[Data]{
		Key: key,
		Data: Data{
			Name: "teste",
			Job:  "teste",
		},
	}
}

func main() {
	treeService := services.BTreeService[Data]{}

	// keysToAdd := []int{1, 2, 3, 4, 5, 6, 7, 12, 13, 14, 8, 10, 11, 9}
	keysToAdd := []int{}
	randomKeysToAdd := 1000

	treeService.InitService(models.BTreeConfig{
		MaxElementsPerNode: 2,
	})

	for _, key := range keysToAdd {
		treeService.Add(generateElement(key))
	}

	for i := 0; i < randomKeysToAdd; i++ {
		treeService.Add(generateElement(rand.Intn(10000)))
	}

	fmt.Println("service: ", treeService)
}
