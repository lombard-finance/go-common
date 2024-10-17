package utils

import (
	"math"

	"github.com/pkg/errors"
)

func SafeInt64ToUint64(i int64) (uint64, error) {
	if i < 0 {
		return 0, errors.New("cannot convert negative int64 to uint64")
	}

	return uint64(i), nil
}

func SafeUint64ToInt64(u uint64) (int64, error) {
	if u > math.MaxInt64 {
		return 0, errors.New("uint64 value too large to convert to int64")
	}

	return int64(u), nil
}
