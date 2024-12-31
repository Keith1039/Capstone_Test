package main

import (
	"github.com/Keith1039/Capstone_Test/parameters"
)

func main() {
	//m := map[string]int{"apples": 1, "oranges": 2, "bananas": 3}
	//fruits := make([]string, 0, len(m))
	//for key, _ := range m {
	//	fruits = append(fruits, key)
	//}
	//sort.SliceStable(fruits, func(i, j int) bool {
	//	return m[fruits[i]] > m[fruits[j]]
	//})
	//fmt.Println(fruits)

	writer := parameters.QueryWriter{TableName: "purchase"}
	err := writer.Init()
	if err != nil {
		panic(err)
	}
	writer.CreateTableOrder()
	writer.ProcessTables()
}
