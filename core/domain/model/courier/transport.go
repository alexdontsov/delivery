package courier

import "github.com/google/uuid"

var (
	Pedestrian = 1
	Bicycle    = 2
	Car        = 3
)

type Transport struct {
	id    uuid.UUID
	name  string
	speed int
}
