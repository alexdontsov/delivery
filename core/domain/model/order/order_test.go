package order

import (
    "delivery/core/domain/model/sharedkernel"
    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_OrderShouldCanHaveStatusCreated(t *testing.T) {
    // Arrange
    location, _ := sharedkernel.CreateLocation(1, 1)
    order, _ := NewOrder(uuid.New(), location)

    // Act
    result := order.status.Value == Created.Value

    assert.True(t, result)
}

func Test_OrderAssignOrderToCourier(t *testing.T) {
    // Arrange
    location, _ := sharedkernel.CreateLocation(1, 1)
    order, _ := NewOrder(uuid.New(), location)
    assignOrder, _ := AssignOrderToCourier(order, uuid.New())

    // Act
    result := assignOrder.status.Value == Assigned.Value

    // Assert
    assert.True(t, result)
}

func Test_OrderComplete(t *testing.T) {
    // Arrange
    location, _ := sharedkernel.CreateLocation(1, 1)
    order, _ := NewOrder(uuid.New(), location)
    completedOrder, _ := CompleteOrder(order)

    // Act
    result := completedOrder.status.Value == Created.Value

    // Assert
    assert.True(t, result)
}
