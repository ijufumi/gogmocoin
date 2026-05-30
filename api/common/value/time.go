package value

import (
	"bytes"
	"strconv"
	"time"
)

// TimeInMillis wraps time.Time and unmarshals GMO Coin timestamps expressed as a
// millisecond epoch, accepting both quoted ("1700000000000") and bare numeric
// JSON representations.
type TimeInMillis struct {
	time.Time
}

func (d *TimeInMillis) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// Trim surrounding quotes if present without assuming a minimum length, so a
	// short or unquoted value cannot trigger an out-of-range slice panic.
	data = bytes.Trim(data, `"`)
	if len(data) == 0 {
		return nil
	}
	longData, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.UnixMilli(longData)

	return nil
}
