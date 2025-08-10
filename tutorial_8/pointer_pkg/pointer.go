package pointerpkg

import "fmt"

func Pointer_Practice() {
	var Thing1 = [5]float64{0, 1, 2, 3, 4}
	
	for i := 0;i < 5; i++ {
		fmt.Printf("Thing1[%v]: %.2f\n",i,Thing1[i])
	}

	var Result [5]float64 = Square(&Thing1)

	for i := 0;i < 5; i++ {
		fmt.Printf("Result[%v]: %.2f\n",i,Result[i])
	}
}

func Square(thing *[5]float64) [5]float64 {
	for i := range thing{
		thing[i] = (thing[i] * thing[i])
	}
	return *thing
}