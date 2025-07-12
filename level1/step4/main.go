package main

import (
	"fmt"
)

const (
	freePlace   = "-"
	totalPlaces = 5
)

func viewQueue(queue [totalPlaces]string) {
	for i, name := range queue {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}

func getStateQueue(queue [totalPlaces]string) int {
	empty := 0
	for _, name := range queue {
		if name == freePlace {
			empty++
		}
	}
	return empty
}

func isFreePlace(queue [totalPlaces]string, place int) bool {
	if queue[place] == freePlace {
		return true
	}
	return false
}

func main() {
	var (
		command string
		queue   [totalPlaces]string = [totalPlaces]string{"-", "-", "-", "-", "-"}
		name    string
		place   int
	)
	for {
		fmt.Scan(&command)
		switch command {
		case "очередь":
			viewQueue(queue)
		case "количество":
			emptyPlaces := getStateQueue(queue)
			fmt.Printf("Осталось свободных мест: %d\n", emptyPlaces)
			fmt.Printf("Всего человек в очереди: %d\n", totalPlaces-emptyPlaces)
		case "конец":
			viewQueue(queue)
			return
		case "":
			continue
		default:
			name = command
			fmt.Scan(&place)
			switch {
			case place < 1 || place > 5:
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)
			case getStateQueue(queue) == 0:
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", place)
			case !isFreePlace(queue, place-1):
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
			default:
				queue[place-1] = name
				//fmt.Printf("%d. %s\n", place, name)
			}
		}
	}
}
