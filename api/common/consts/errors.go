package consts

import "errors"

// ErrUnsupportedTierLevel is returned when a tier level outside the known set is decoded.
var ErrUnsupportedTierLevel = errors.New("unsupported tierLevel")

// ErrUnsupportedSymbol is returned when a symbol outside the known set is decoded.
var ErrUnsupportedSymbol = errors.New("unsupported symbol")
