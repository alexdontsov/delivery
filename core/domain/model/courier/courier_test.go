package courier

import (
	"delivery/core/domain/model/sharedkernel"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenCreatedCourierWillHaveStatusFree(t *testing.T) {
	// Arrange
	location, _ := sharedkernel.CreateLocation(10, 10)
	courier := NewCourier("Vasiliy", Transport{id: uuid.New(), name: "Bicycle", speed: Bicycle}, location)

	// Act
	result := courier.status.Value == Free.Value

	// Assert
	assert.True(t, result)
}
