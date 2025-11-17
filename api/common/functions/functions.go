package functions

import (
	"encoding/json"
	"fmt"
)

func EncodeJSON(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("<error marshaling JSON: %v>", err)
	}
	return string(b)
}
