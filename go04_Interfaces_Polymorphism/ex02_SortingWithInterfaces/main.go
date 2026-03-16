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
}

type Person struct {
	Name string
	Age  int
}

type People []Person

func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
