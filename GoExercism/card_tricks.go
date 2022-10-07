package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	fav := []int{2, 6, 9}
	return fav
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < len(slice) && index >= 0 {
		return slice[index]
	} else {
		return -1
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index >= 0 {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	slice = append(values, slice...)
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < len(slice) && index >= 0 {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice
	}
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(FavoriteCards())
	fmt.Println(GetItem(s, 5))
	fmt.Println(SetItem(s, 5, 10))
	fmt.Println(PrependItems(s, 13, 14))
	fmt.Println(RemoveItem(s, 2))
}
