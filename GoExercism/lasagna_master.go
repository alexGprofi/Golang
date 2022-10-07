package main

import "fmt"

func PreparationTime(layer []string, num int) int {
	if num == 0 {
		num = 2
	}
	prep := len(layer) * num
	return prep
}

func Quantities(layer []string) (int, float64) {
	var num_sauce, num_noodles int
	for i := 0; i < len(layer); i++ {
		if layer[i] == "sauce" {
			num_sauce++
		}
		if layer[i] == "noodles" {
			num_noodles++
		}

	}
	return num_noodles * 50, float64(num_sauce) * 0.2
}

func AddSecretIngredient(friendsList, myList []string) {
	myIndex := len(myList) - 1
	fIndex := len(friendsList) - 1
	myList[myIndex] = friendsList[fIndex]
}

func ScaleRecipe(quantity []float64, portions int) []float64 {
	qty := make([]float64, len(quantity))
	for i := 0; i < len(quantity); i++ {
		qty[i] = quantity[i] * (float64(portions) / 2)
	}
	return qty
}
func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	fmt.Println(PreparationTime(layers, 0))
	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)

	quantities := []float64{0.5, 250, 150, 3, 0.5}
	scaledQuantities := ScaleRecipe(quantities, 6)
	fmt.Println(scaledQuantities)
}
