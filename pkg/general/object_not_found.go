package general

import "fmt"

type ObjectNotFound struct {
    msg string
}

func NewObjectNotFoundError(msg string) ObjectNotFound {
    return ObjectNotFound{msg: msg}
}

func (v ObjectNotFound) Error() string {
    return fmt.Sprintf("object not found %s", v.msg)
}
