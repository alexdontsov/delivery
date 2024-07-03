package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TwoStatusesShouldBeEqualWhenParamsIsEqual(t *testing.T) {
	// Arrange
	status1 := Created
	status2 := Created

	// Act
	result := status1.IsEqual(status2)

	// Assert
	assert.True(t, result)
}

func Test_TwoStatusesShouldNotBeEqualWhenParamsIsNotEqual(t *testing.T) {
	// Arrange
	status1 := Created
	status2 := Assigned

	// Act
	result := status1.IsEqual(status2)

	// Assert
	assert.False(t, result)
}

func Test_StatusShouldCanDoSelfCheckToEmpty(t *testing.T) {
	// Arrange
	status := Status{}

	// Act
	result := status.IsEmpty()

	// Assert
	assert.True(t, result)
}

func Test_StatusShouldCanFindByName(t *testing.T) {
	// Arrange
	status := Status{}

	// Act
	result, err := status.FromName("created")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, Created, result)
}
