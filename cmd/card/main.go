package main

import "fmt"

func main() {
	const course = "Бэкенд-разработка на языке Go"
	var name string = "Арарат"
	var group string = "Учебная-01"
	completed := 67
	var interested bool = true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Println("Курс: ", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", interested)
}
