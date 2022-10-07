package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	if week == 1 {
		sum = TotalBirdCount(birdsPerDay[:7])
	} else if week == 2 {
		sum = TotalBirdCount(birdsPerDay[8:14])
	} else {
		sum = 0
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	birdCounts := []int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0, 8, 0}
	fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println(BirdsInWeek(birdsPerDay, 2))
	fmt.Println(birdsPerDay)
	fmt.Println(FixBirdCountLog(birdsPerDay))

	fmt.Println(TotalBirdCount(birdCounts))
	fmt.Println(BirdsInWeek(birdCounts, 1))
	fmt.Println(birdCounts)
	fmt.Println(FixBirdCountLog(birdCounts))
}
