package main

import (
	"fmt"
	"sort"
)

// Опишем склад с товарами
type Goods struct {
	Names []string // Название товара
	Boxes []int    // Коробка, в которой лежит товар
}

/*
	type Interface interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}
*/

func (g Goods) Len() int {
	return len(g.Names)
}
func (g Goods) Swap(i, j int) {
	g.Names[i], g.Names[j] = g.Names[j], g.Names[i] // Меняем местами как название товара,
	g.Boxes[i], g.Boxes[j] = g.Boxes[j], g.Boxes[i] // так и коробки, где товар хранится
}
func (g Goods) Less(i, j int) bool {
	return g.Boxes[i] < g.Boxes[j]
}

func main() {
	goods := Goods{
		Names: []string{"phone", "tablet", "earphones"},
		Boxes: []int{3, 1, 2},
	}
	sort.Sort(goods)
	fmt.Println(goods) // {[tablet earphones phone] [1 2 3]}
}
