package sharedkernel

import (
	"delivery/pkg/general"
	"math"
	"math/rand/v2"
)

const minX = 1
const maxX = 10
const minY = 1
const maxY = 10

type Location struct {
	X int
	Y int
}

func CreateLocation(x int, y int) (Location, error) {
	if x < minX || x > maxX {
		return Location{}, general.NewValueIsOutOfRange("x", x, minX, maxX)
	}

	if y < minY || y > maxY {
		return Location{}, general.NewValueIsOutOfRange("y", y, minY, maxY)
	}

	return Location{
		X: x,
		Y: y,
	}, nil
}

func (location *Location) IsEqual(other Location) bool {
	return location.X == other.X && location.Y == other.Y
}

func (location *Location) IsEmpty() bool {
	return *location == Location{}
}

func (location *Location) DistanceTo(targetLocation Location) int {
	diffX := math.Abs(float64(location.X - targetLocation.X))
	diffY := math.Abs(float64(location.Y - targetLocation.Y))

	return int(diffX + diffY)
}

func CreateRandomLocation() (Location, error) {
	return CreateLocation(rand.IntN(maxX), rand.IntN(maxY))
}
