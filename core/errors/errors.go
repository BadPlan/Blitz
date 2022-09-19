package errors

import "errors"

var (
	ErrEmptySpriteComponent = errors.New("empty sprite component")
	ErrNoLocalToWorldSystem = errors.New("local_to_world_system is not defined")
	ErrRotationIsNotDefined = errors.New("rotation is not  defined")
)
