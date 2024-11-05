package sharedkernel

import "delivery/pkg/general"

type Weight struct {
    Value int
}

func NewWeight(value int) (Weight, error) {
    const minWeight = 0
    const maxWeight = 1000

    if value <= minWeight {
        return Weight{}, general.NewValueIsOutOfRange("value", value, minWeight, maxWeight)
    }

    return Weight{
        Value: value,
    }, nil
}

func (weight *Weight) IsEqual(other Weight) bool {
    return weight.Value == other.Value
}

func (weight *Weight) IsEmpty() bool {
    return *weight == Weight{}
}
