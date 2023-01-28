package homework3

import (
	"fmt"
	"strings"
)

// Написать функцию, принимающую аргумент типа int. Выводить на экран "Да", если аргумент больше 10 и меньше 20,
// "Нет" если аргумент меньше 10, "Наверное" если аргумент больше 20.
func TaskOne(val int) {
	if 10 < val && val < 20 {
		fmt.Println("Yes")
	} else if val < 10 {
		fmt.Println("No")
	} else if val > 20 {
		fmt.Println("Probably")
	}
}

// Написать функцию, принимающую string в качестве аргумента. Значение аргумента это написанные на русском языке числа
// "один", "два" и так далее, до "десять". Выводить на экран соответствующее число. аргумент не совпадает с описанными
// значениями, выводить "Не знаю"
func taskTwo(val string) {
	switch strings.Trim(val, " ") {
	case "one":
		fmt.Println(1)
	case "two":
		fmt.Println(2)
	case "three":
		fmt.Println(3)
	case "four":
		fmt.Println(4)
	case "five":
		fmt.Println(5)
	case "six":
		fmt.Println(6)
	case "seven":
		fmt.Println(7)
	case "eight":
		fmt.Println(8)
	case "nine":
		fmt.Println(9)
	case "ten":
		fmt.Println(10)
	default:
		fmt.Println("Don't know")
	}
}

// Написать функцию, выводящую на экран все нечетные числа от 0 до аргумента типа int, переданного в функцию.
func taskThird(val int) {
	if val < 0 {
		return
	}
	for i := 0; i < val; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")

		}
	}
}

// Написать функцию, принимающие 2 аргумента типа int. Вевести на экран все числа от A до B если A<B, иначе вывести
// все числа от B до A. Использовать рекурсию.
func taskForth(a int, b int) {
	if a < b {
		if a >= b {
			return
		}
		fmt.Println(a)
		a++
		taskForth(a, b)
	} else if a > b {
		if a <= b {
			return
		}
		fmt.Println(b)
		b++
		taskForth(a, b)
	}
}
