package order

import (
	"errors"
)

var (
	ErrStatusIsWrong = errors.New("status.is.wrong")
)

var (
	Created   = Status{"created"}
	Assigned  = Status{"assigned"}
	Completed = Status{"completed"}
)

type Status struct {
	Value string
}

func (status *Status) IsEqual(other Status) bool {
	return status.Value == other.Value
}

func (status *Status) IsEmpty() bool {
	return *status == Status{}
}

func (status *Status) FromName(name string) (Status, error) {
	switch name {
	case Created.Value:
		return Created, nil
	case Assigned.Value:
		return Assigned, nil
	case Completed.Value:
		return Completed, nil
	default:
		return Status{}, ErrStatusIsWrong
	}
}
