package statics

import "errors"

var (
	InternalError    = errors.New("InternalError")
	NotEnoughBalance = errors.New("NotEnoughBalance")
	NotExistsBalance = errors.New("NotExistsBalance")
)
