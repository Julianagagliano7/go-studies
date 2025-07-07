package main

import "fmt"

//for inicializacao; condicao; fim iteracao{

//}

func main() {

	//For comum
	// sum := 0

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	sum += i
	// }

	// fmt.Println(sum)

	//Simulando o while
	// sum := 0

	// for sum < 20 {
	// 	fmt.Println("loop")
	// 	sum += 2
	// }

	// fmt.Println(sum)

	//For de Slices
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

}
