package main

import "fmt"

func main() {
	// for i (inicializador):= 0 ; (condição)
	// ... i <= 5; i++ (acrescenta mais um)
	// simple iteration over a range
	/*for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	//a
	// iterate over collection
	// slice:
	/*numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}*/
	// for index(indice que itera sobre os elementos)
	//... , value (valor do elemento na slice)
	// := range numbers {} no range da quantidade de valores dentro da slice
	// sempre começa com 0 representando o primeiro elemento da slice
	// % place holders (%d para decimais, %s para strings, %f para floats, %v general values)
	/*for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}*/

	// break and continue statments
	/*for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println("Even number:", i)
			continue
		}

		fmt.Println("Odd number:", i)

		if i == 5 {
			break
		}

	}*/
	// ASTERISC LAYOUT
	// rows := 5
	// // outer loop
	// for i:=1; i<=rows; i++ {
	// 	// inner loop
	// 	for j:=1; j<=rows-i;j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loops for stars
	// 	for k:=1; k<=2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// for i:=range 10 {
	// 	fmt.Println(i)
	// } 
	// GO TAMBÉM PERMITE DECLARAR LOOPS DA SEGUINTE FORMA:
	for i := range 20{
		fmt.Println(i)
	}
}
