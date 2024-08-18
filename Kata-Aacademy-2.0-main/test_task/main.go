package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter an expression: ")
	expression, _ := reader.ReadString('\n')

	result := calculate(expression)
	fmt.Printf("Result: %s\n", result)

}

func calculate(expression string) string {
	expression = strings.TrimSuffix(expression, "\n")
	var result string
	var operator string
	temp := ""
	ArabicNumbers := "1234567890"
	Gtokens := make([]string, 3)

	//выделяем и подготавливаем к работе два операнда

	sum := strings.Contains(expression, " + ")
	subtract := strings.Contains(expression, " - ")
	multiply := strings.Contains(expression, " * ")
	divide := strings.Contains(expression, " / ")

	if sum {
		tokens := strings.Split(expression, " + ")
		operator = "+"
		tokens[0] = strings.ReplaceAll(tokens[0], "\"", "")
		tokens[0] = strings.TrimSpace(tokens[0])
		Gtokens[0] = tokens[0]
		tokens[1] = strings.ReplaceAll(tokens[1], "\"", "")
		tokens[1] = strings.TrimSpace(tokens[1])
		Gtokens[2] = tokens[1]
	} else if subtract {
		operator = "-"
		tokens := strings.Split(expression, " - ")
		tokens[0] = strings.ReplaceAll(tokens[0], "\"", "")
		tokens[0] = strings.TrimSpace(tokens[0])
		Gtokens[0] = tokens[0]
		tokens[1] = strings.ReplaceAll(tokens[1], "\"", "")
		tokens[1] = strings.TrimSpace(tokens[1])
		Gtokens[2] = tokens[1]
	} else if multiply {
		operator = "*"
		tokens := strings.Split(expression, " * ")
		if !(strings.ContainsAny(tokens[0], "\"")) {
			fmt.Println("Первым аргументом выражения, подаваемым на вход, должна быть строка")
			os.Exit(1)
		}
		tokens[0] = strings.ReplaceAll(tokens[0], "\"", "")
		tokens[0] = strings.TrimSpace(tokens[0])
		Gtokens[0] = tokens[0]
		tokens[1] = strings.ReplaceAll(tokens[1], "\"", "")
		tokens[1] = strings.TrimSpace(tokens[1])
		Gtokens[2] = tokens[1]
	} else if divide {
		operator = "/"
		tokens := strings.Split(expression, " / ")
		if !(strings.ContainsAny(tokens[0], "\"")) {
			fmt.Println("Первым аргументом выражения, подаваемым на вход, должна быть строка")
			os.Exit(1)
		}
		tokens[0] = strings.ReplaceAll(tokens[0], "\"", "")
		tokens[0] = strings.TrimSpace(tokens[0])
		Gtokens[0] = tokens[0]
		tokens[1] = strings.ReplaceAll(tokens[1], "\"", "")
		tokens[1] = strings.TrimSpace(tokens[1])
		Gtokens[2] = tokens[1]
	}

	if len(Gtokens[0]) > 10 || len(Gtokens[2]) > 10 {
		fmt.Println("строки длиной не более 10 символов")
		os.Exit(1)
	}

	switch operator {
	case "+":
		result = Gtokens[0] + Gtokens[2]
	case "-":
		result = strings.ReplaceAll(Gtokens[0], Gtokens[2], "")
	case "*":
		multiplication, err := strconv.Atoi(Gtokens[2])
		if err != nil {
			fmt.Println("Калькулятор умеет работать только с целыми числами")
			os.Exit(1)
		}
		if multiplication < 1 || multiplication > 10 {
			fmt.Println("Калькулятор должен принимать на вход числа от 1 до 10 включительно")
			os.Exit(1)
		} else {
			for i := 0; i < multiplication; i++ {
				temp = temp + Gtokens[0]
				//fmt.Println(temp)
			}
			result = temp
			temp = ""
		}
	case "/":
		if !(strings.ContainsAny(Gtokens[2], ArabicNumbers)) {
			fmt.Println("неподдерживаемых операций (деление строки на строку)")
			os.Exit(1)
		}
		divider, err := strconv.Atoi(Gtokens[2])
		if err != nil {
			fmt.Println("Калькулятор умеет работать только с целыми числами")
			os.Exit(1)
		}
		flag := divider
		divider = len(Gtokens[0]) / divider
		word := Gtokens[0]
		letters := strings.Split(word, "")
		if flag < 1 || flag > 10 {
			fmt.Println("Калькулятор должен принимать на вход числа от 1 до 10 включительно")
			os.Exit(1)
		} else {
			for i := 0; i < divider; i++ {
				temp += letters[i]
				result = temp
			}
		}
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}
	/* Если строка, полученная в результате работы приложения, длиннее 40 символов, то в выводе после 40 символа должны стоять три точки (...)*/
	if len(result) > 40 {
		limitation := strings.Split(result, "")
		for i := 0; i < 40; i++ {
			temp += limitation[i]
			result = temp
		}
		result = result + "..."
	}

	return `"` + result + `"`
}
