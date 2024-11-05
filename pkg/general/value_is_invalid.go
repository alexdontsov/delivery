package general

import "fmt"

type ValueIsInvalid struct {
    msg string
}

func NewValueIsInvalidError(msg string) ValueIsInvalid {
    return ValueIsInvalid{msg: msg}
}

func (v ValueIsInvalid) Error() string {
    return fmt.Sprintf("value is invalid %s", v.msg)
}
