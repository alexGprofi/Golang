package main

import "fmt"

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	//fmt.Println(oldRune, newRune)
	uOld := fmt.Sprintf("%U", oldRune)
	fmt.Println("log: ", log)
	var myLog string
	//uNew := fmt.Sprintf("%U", newRune)
	for _, char := range log {
		a := fmt.Sprintf("%U", char)
		if a == uOld {
			//log[index] = fmt.Sprintf("%c",newRune)
			//fmt.Println("a: ", a)
			//fmt.Println("Index: ", index)
			//b := fmt.Sprintf("%T", newRune)
			//fmt.Println("b: ", b)
			//old := fmt.Sprintf("%T", oldRune)
			//fmt.Println("old: ", old)
			//log_index := fmt.Sprintf("Type: %T, value: %v", log[index], log[index])
			//fmt.Println("log[index]: ", log_index)
			myLog += fmt.Sprintf("%c", newRune)
		} else {
			myLog += fmt.Sprintf("%c", char)
		}

	}
	return myLog
}

func main() {
	myString := "❗hello"
	for index, char := range myString {
		fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U   Type: %T\n", index, char, char, char)
	}
	mylog := Replace(myString, '❗', '?')
	fmt.Println(mylog)
}
