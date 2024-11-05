package general

import "fmt"

type VersionIsInvalid struct {
    msg string
}

func NewVersionIsInvalidError(msg string) VersionIsInvalid {
    return VersionIsInvalid{msg: msg}
}

func (v VersionIsInvalid) Error() string {
    return fmt.Sprintf("version is invalid %s", v.msg)
}
