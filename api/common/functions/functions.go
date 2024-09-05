package functions

import "encoding/json"

func EncodeJSON(v any) string {
	b, e := json.MarshalIndent(v, "", "")
	if e != nil {
		return ""
	}
	return string(b)
}
