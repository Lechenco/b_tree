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

func addKeys(service *services.BTreeService[Data]) {
	keysToAdd := []int{1, 2, 3, 4, 5, 6, 7, 12, 13, 14, 8, 10, 11, 9}
	for _, key := range keysToAdd {
		service.Add(generateElement(key))
	}
}

func addKeysRandom(service *services.BTreeService[Data]) {
	randomKeysToAdd := 1000
	for i := 0; i < randomKeysToAdd; i++ {
		service.Add(generateElement(rand.Intn(10000)))
	}
}

func main() {
	treeService := services.BTreeService[Data]{}

	treeService.InitService(models.BTreeConfig{
		MaxElementsPerNode: 2,
	})

	addKeys(&treeService)
	// addKeysRandom(&treeService)

	fmt.Println("service: ", treeService)

	fmt.Println("15: ", treeService.Get(15))
}
