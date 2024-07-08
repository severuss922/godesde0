package arreglos_slices

import (
	"fmt"
)

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //Slice creado desde un vector desde la pos 3 hasta el final
	porcion2 := arreglo[:5]  //Slice creado desde un vector desde el inicio hasta el elemento 5
	porcion3 := arreglo[6:7] //Slice creado desde un vector desde el elemento 6 hasta el 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	for j := 0; j < 130; j++ {
		elementos = append(elementos, j)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0)
	for i := 0; i < 130; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
