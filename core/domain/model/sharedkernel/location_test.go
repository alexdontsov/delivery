package sharedkernel

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_CreateLocationShouldBeCreateWhenParamsIsCorrect(t *testing.T) {
    // Arrange
    x := 1
    y := 1

    // Act
    location, err := CreateLocation(x, y)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, x, location.X)
    assert.Equal(t, y, location.Y)
}

func Test_NewLocationShouldBeReturnErrorWhenXIsIncorrect(t *testing.T) {
    // Arrange
    x := 0
    y := 1

    // Act
    location, err := CreateLocation(x, y)

    // Assert
    assert.Equal(t, Location{}, location)
    assert.Error(t, err)
    assert.Equal(t, "value 0 of x is out of range , min value is 1, max value is 10", err.Error())
}

func Test_TwoLocationsShouldBeEqualWhenParamsIsEqual(t *testing.T) {
    // Arrange
    location1, err := CreateLocation(1, 1)
    assert.NoError(t, err)
    location2, err := CreateLocation(1, 1)
    assert.NoError(t, err)

    // Act
    result := location1.IsEqual(location2)

    // Assert
    assert.NoError(t, err)
    assert.True(t, result)
}

func Test_TwoLocationsShouldNotBeEqualWhenParamsIsNotEqual(t *testing.T) {
    // Arrange
    location1, err := CreateLocation(1, 1)
    assert.NoError(t, err)
    location2, err := CreateLocation(2, 2)
    assert.NoError(t, err)

    // Act
    result := location1.IsEqual(location2)

    // Assert
    assert.NoError(t, err)
    assert.False(t, result)
}

func Test_LocationShouldCanDoSelfCheckToEmpty(t *testing.T) {
    // Arrange
    location := Location{}

    // Act
    result := location.IsEmpty()

    // Assert
    assert.True(t, result)
}

func Test_LocationShouldCanCalculateDistanceToAnotherLocation(t *testing.T) {
    // Arrange
    location1, err := CreateLocation(1, 1)
    assert.NoError(t, err)
    location2, err := CreateLocation(2, 2)
    assert.NoError(t, err)

    // Act
    result := location1.DistanceTo(location2)

    // Assert
    assert.Equal(t, 2, result)
}
