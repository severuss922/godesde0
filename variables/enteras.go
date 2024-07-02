package variables

import (
	"fmt"
)

func MuestroEnteros() {
	var enteroComun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("enteroComun = ", enteroComun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
