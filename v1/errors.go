package coc

import "errors"

var (
	ErrClanNotFound = errors.New("clan not found")
	ErrNotInWar     = errors.New("clan is not in a war")
	ErrTagMissing   = errors.New("no tag provided")
)
