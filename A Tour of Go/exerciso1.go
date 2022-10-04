package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
  z:=1.0
  for i:=0; i<10;i++ {
     z -=(z*z-x)/(2*z)
	 if z == math.Sqrt(x) {
	    return z
	 }
	 fmt.Println(i)
  }
  return z
}

func main() {
	fmt.Println(Sqrt(35))
	fmt.Println(math.Sqrt(35))
}
