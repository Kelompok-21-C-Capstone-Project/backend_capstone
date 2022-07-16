package errordriver

import "errors"

var (
	ErrUnauthorized   error = errors.New("unauthorized")
	ErrDataNotFound   error = errors.New("data not found")
	ErrDataUsed       error = errors.New("the data is already used")
	ErrBadRequest     error = errors.New("bad request")
	ErrInternalServer error = errors.New("internal server error")
	ErrWrongPass      error = errors.New("wrong password/identifier")
	ErrUserNoAccess   error = errors.New("user have no access")
	ErrItemIsOut      error = errors.New("item is currently sold out")
)
