package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var um int

func main() {
	// use for loops as while loops
	// em go ao apenas declarar a condição, irá se comportar como um while
	// for um != "sair" {
	// 	fmt.Scan(&um)
	// 	fmt.Println("Iteration:", i)
	// 	i++
	// }
	// Outra forma de usar como while:
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum:", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	// ++ operador que incrementa +1
	// -- operador que decrementa -1

	// num := 1213
	// for {
	// 	fmt.Print("Try to guess the number: ")
	// 	fmt.Scan(&um)
	// 	if um < num {
	// 		fmt.Println("the number is higher")
	// 		continue
	// 	}
	// 	if um > num {
	// 		fmt.Println("the number is lower")
	// 		continue
	// 	}
	// 	if um == num {
	// 		fmt.Println("you found it!")
	// 		fmt.Print(num)
	// 		break
	// 	}

	// criando uma seed para gerar um numero aleatorio usando o tempo atual:
	source := rand.NewSource(time.Now().UnixNano())
	// criando o gerador de numeros aleatorios usando a seed acima
	random := rand.New(source)
	// gerando numeros aleatorios de 0 a 99
	// random.Intn(100) + 1: gera números aleatorios de a 0 99
	// + 1 aumenta o intervalo em 1 e gera números de 1 a 100
	target := random.Intn(100) + 1

	fmt.Println("Welcome to the guessing game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can u guess it?")
	var guess int
	fmt.Println("Endereço da variavel guess:", &guess)
	for {
		fmt.Println("Enter yor gues: ")
		// passando o endereço da memoria
		// ... se for passado a variavel diretamente criará uma cópia
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println("You found it!")
			fmt.Println("The number was:", target)
			break
		} else if guess < target {
			fmt.Println("The number is higher")
		} else {
			fmt.Println("The number is lower")
		}

	}
}
