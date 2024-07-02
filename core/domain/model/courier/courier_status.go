package courier

import (
	"errors"
)

var (
	ErrStatusIsWrong = errors.New("status.is.wrong")
)

var (
	Free         = Status{"free"}
	NotAvailable = Status{"notAvailable"}
	Ready        = Status{"ready"}
	Busy         = Status{"busy"}
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
	case NotAvailable.Value:
		return NotAvailable, nil
	case Ready.Value:
		return Ready, nil
	case Busy.Value:
		return Busy, nil
	default:
		return Status{}, ErrStatusIsWrong
	}
}
