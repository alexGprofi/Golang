package main

import "golang.org/x/tour/pic"



func Pic(dx, dy int) [][]uint8 {


	rpic :=  make([][]uint8, dy)

	for i, _ := range rpic {
		rpic[i] = make([]uint8, dx)
		for j, _ := range rpic[i] {
			rpic[i][j] = uint8((i+j)/2)
		}
	}

	return rpic


}

func main() {
	pic.Show(Pic)
}
