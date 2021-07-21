package common

import (
	"errors"
	"strconv"
)

func StringToFloat64(ctx *Context, s string) (float64, error) {
	y, err := strconv.ParseFloat(s, 64)
	if !errors.Is(err, nil) {
		ctx.Logger.Error(err)
		return 0.0, err
	}

	return y, nil
}
