package value

import (
	"strconv"
	"time"
)

type TimeInMillis struct {
	time.Time
}

func (d *TimeInMillis) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	longData, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.UnixMilli(longData)

	return nil
}
