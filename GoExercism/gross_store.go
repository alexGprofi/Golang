package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var bill = make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unit_exists := units[unit]
	if unit_exists > 0 {
		bill[item] += unit_exists
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	numItems := bill[item]
	numUnits := units[unit]
	fmt.Println("Parameters: ", numItems, numUnits)
	if numItems-numUnits == 0 {
		fmt.Println("Bill: ", bill)
		delete(bill, item)
		return true
	}
	if bill[item] == 0 {
		return false
	}
	if numItems-numUnits > 0 && numUnits > 0 {
		bill[item] = numItems - numUnits
		return true
	} else {
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

	if bill[item] > 0 {
		return bill[item], true
	} else {
		return 0, false
	}
}

func main() {
	units := Units()
	fmt.Println(units)
	bill := NewBill()
	fmt.Println(bill)

	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	fmt.Println(bill)
	ok = RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	fmt.Println("1 Removed", bill)
	ok = RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	fmt.Println("2 Removed", bill)

	bill = map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(bill, "carrota")
	fmt.Println(bill)
	fmt.Println(qty)
	// Output: 12
	fmt.Println(ok)
	// Output: true
}
