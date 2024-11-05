package general

import "fmt"

type ValueIsRequired struct {
    msg string
}

func NewValueIsRequiredError(msg string) ValueIsRequired {
    return ValueIsRequired{msg: msg}
}

func (v ValueIsRequired) Error() string {
    return fmt.Sprintf("value is required %s", v.msg)
}
