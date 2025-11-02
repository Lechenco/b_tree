package main

import (
	"fmt"

	"github.com/Lechenco/b_tree/internal/models"
)

type Data struct {
	Name string
	Job  string
}

func main() {
	tree := models.NewBTree[Data]()

	fmt.Printf("Tree: %v\n", tree)
}
