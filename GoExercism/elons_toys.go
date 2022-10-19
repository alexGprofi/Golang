package main

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	time := 1
	if car.battery >= car.batteryDrain {
		car.battery = car.battery - (car.batteryDrain * time)
		car.distance += car.speed * time
	}

	return car
}

func (car *Car) Drive() {
	time := 1
	if car.battery >= car.batteryDrain {
		car.battery = car.battery - (car.batteryDrain * time)
		car.distance += car.speed * time
	}
	//return car
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(track int) bool {
	totalDrain := (car.batteryDrain * (track / car.speed))
	fmt.Println(totalDrain)
	if (car.battery - totalDrain) >= 0 {
		return true
	} else {
		return false
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	totalDrain := (car.batteryDrain * (track.distance / car.speed))
	fmt.Println(totalDrain)
	if (car.battery - totalDrain) >= 0 {
		return true
	} else {
		return false
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Println(car)
	car.Drive()
	fmt.Println(car)
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())

	trackDistance := 100

	fmt.Println(car.CanFinish(trackDistance))
	/*distance := 100
	track := NewTrack(distance)
	fmt.Println(track)
	//car = Drive(car)
	fmt.Println(car)

	fmt.Println(CanFinish(car, track))
	*/
}
