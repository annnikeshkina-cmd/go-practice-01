package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	count := 0
	for _, statuses := range statuses {
		if statuses == target {
			count++
		}
	}
	return count
}

func findStatus(statuses []string, target string) int {
	for i, statuses := range statuses {
		if statuses == target {
			return i
		}
	}
	return -1
}

func completionPercent(statuses []string) float64 {
	all := float64(len(statuses))
	done := float64(countStatus(statuses, "done"))
	switch all {
	case 0:
		return 0
	default:
		return done / all
	}
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}

	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])
	fmt.Printf("Индекс первого done: %d\n", findStatus(statuses, "done"))
	fmt.Printf("Доля выполненных задач: %f\n", completionPercent(statuses))
}
