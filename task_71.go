package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	// если введенное число больше 10, то выполняются команды внутри фигурных скобок. Команды выполняются только тогда, когда условие верно (истинно)
	if x > 10 {
		fmt.Println(x) // вывести х
	}
}