package consts

import (
	"slices"
	"strconv"
)

// TierLevel ...
type TierLevel int8

const (
	TierLevel1 = TierLevel(1)
	TierLevel2 = TierLevel(2)
)

var allTierLevels = []TierLevel{
	TierLevel1, TierLevel2,
}

func (c *TierLevel) UnmarshalJSON(d []byte) error {
	n, err := strconv.ParseInt(string(d), 10, 64)
	if err != nil {
		return err
	}
	level := TierLevel(n)
	if !slices.Contains(allTierLevels, level) {
		return ErrUnsupportedTierLevel
	}
	*c = level
	return nil
}
