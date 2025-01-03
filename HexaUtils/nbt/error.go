package nbt

import "fmt"

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(message string) error {
	return Error{Message: message}
}

func NoRootCompound(tagID uint8) error {
	return NewError(fmt.Sprintf("nbt: root tag is not a compound, id is %d", tagID))
}

func InvalidJavaString() error {
	return NewError("nbt: invalid java string")
}

func UnknownTagId(tagId uint8) error {
	return NewError(fmt.Sprintf("nbt: unknown tag id %d", tagId))
}
