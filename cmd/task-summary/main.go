package main

import "fmt"

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

func countStatus(statuses []string, target string) int {
	cnt := 0
	for _, status := range statuses {
		if status == target {
			cnt += 1
		}
	}
	return cnt
}
func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}
func completionPercent(statuses []string) float64 {
	if len(statuses) != 0 {
		return float64(countStatus(statuses, "done")) / float64(len(statuses)) * 100
	}
	return 0
}
func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")
	fmt.Printf("Задач всего: %d\n", len(statuses))

	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))
	cnt := make(map[string]int)
	for _, status := range statuses {
		cnt[status]++
	}
	order := []string{"new", "in_progress", "done"}
	fmt.Print("Статистика: ")
	for _, status := range order {
		fmt.Printf("%s=%d ", status, cnt[status])
	}
	fmt.Printf("\nИндекс первой задачи 'in_progress': %d\n", findStatus(statuses, "in_progress"))
	fmt.Printf("Индекс первой задачи 'blocked': %d\n", findStatus(statuses, "blocked"))
	fmt.Printf("Процент завершенных задач: %.2f\n", completionPercent(statuses))
}
