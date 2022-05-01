package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed int, batteryDrain int) Car {
	newCar := Car{speed: speed,
		batteryDrain: batteryDrain,
	}
	return newCar
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
	}
	return track
	//panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	result := Car{
		battery:      car.battery,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     car.distance,
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if car.distance >= track.distance {
		return true
	} else {
		return false
	}
}
