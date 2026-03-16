package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("PEOPLE")

	people := []Person{
		{"Verstappen", 28},
		{"Norris", 26},
		{"Perez", 36},
		{"Alonso", 44},
		{"Antonelli", 19},
	}
	fmt.Println(people)
	fmt.Println("len:", People(people).Len())
	sort.Sort(People(people))
	fmt.Println("Sorted:", people)

	fmt.Println("\nPRODUCTS")

	products := []Product{
		{"Milk", 0.95, 20},
		{"Meat", 3.99, 10},
		{"Macbook Air M4", 999.99, 3},
		{"Bag of Pens", 3, 12},
	}
	fmt.Println(products)
	fmt.Println(Products(products).Len())
	sort.Sort(Products(products))
	fmt.Println(products)
}

type Person struct {
	Name string
	Age  int
}

type People []Person

func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Product struct {
	Name  string
	Price float64
	Stock int
}

type Products []Product

func (p Products) Len() int           { return len(p) }
func (p Products) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p Products) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
