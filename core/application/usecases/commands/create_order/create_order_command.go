package create_order

import "github.com/google/uuid"

type CreateOrderCommand struct {
    Id        uuid.UUID `validate:"required" json:"id,omitempty"`
    LocationX int       `json:"location_x,omitempty"`
    LocationY int       `json:"location_y,omitempty"`
}
