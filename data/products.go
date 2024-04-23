package data

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

type Fruit struct {
	Name        string  `fake:"{fruit}"`
	Description string  `fake:"{sentence:3}"`
	Price       float64 `fake:"{price:1,10}"`
}

func generateFruit() []string {
	var f Fruit
	gofakeit.Struct(&f)

	froot := []string{}
	froot = append(froot, f.Name)
	froot = append(froot, f.Description)
	froot = append(froot, fmt.Sprintf("%.2f", f.Price))

	return froot
}

func FruitList(length int) (fruits [][]string) {
	for i := 0; i < length; i++ {
		onefruit := generateFruit()
		fruits = append(fruits, onefruit)
	}
	return
}
