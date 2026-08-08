package domain

import "errors"

var ErrNotFound = errors.New("not found")
var ErrItemExist = errors.New("item with such Id already exists ")
var ErrIdNotFound = errors.New("id not found")
var ErrForbidden = errors.New("forbidden")
var ErrInvalidStatus = errors.New("invalid trip status")
