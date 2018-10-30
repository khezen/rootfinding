package rootfinding

import "errors"

var (
	// ErrRootIsNotBracketed - no root in the given interval
	ErrRootIsNotBracketed = errors.New("ErrRootIsNotBracketed - no root in the given interval")
)
