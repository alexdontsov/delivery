package courier

import (
	"delivery/core/domain/model/sharedkernel"
	"github.com/google/uuid"
)

type Courier struct {
	id        uuid.UUID
	name      string
	status    Status
	transport Transport
	location  sharedkernel.Location
}

func NewCourier(name string, transport Transport, location sharedkernel.Location) Courier {
	return Courier{
		id:        uuid.New(),
		name:      name,
		status:    Status{Value: Free.Value},
		transport: transport,
		location:  location,
	}
}

func BusyCourier(courier Courier) Courier {
	courier.status.Value = Busy.Value

	return courier
}

func FreeCourier(courier Courier) Courier {
	courier.status.Value = Free.Value

	return courier
}

func DistanceSpeedCalculation(courier Courier, locationFrom *sharedkernel.Location, locationTo sharedkernel.Location) int {
	distanceTo := (*sharedkernel.Location).DistanceTo(locationFrom, locationTo)
	return distanceTo / courier.transport.speed
}
