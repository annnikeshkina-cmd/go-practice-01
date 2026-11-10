package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Введите три целых значения времени: ")

	_, err1 := fmt.Scan(&n1)
	if err1 != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	var n2 int

	_, err2 := fmt.Scan(&n2)
	if err2 != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	var n3 int

	_, err3 := fmt.Scan(&n3)
	if err3 != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	sum := float64(n1 + n2 + n3)
	arif := sum / 3
	fmt.Printf("Среднее время выполнения: %.2f. \n", arif)
}
