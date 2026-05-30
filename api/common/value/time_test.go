package value

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeInMillis_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		isZero  bool
		want    time.Time
	}{
		{name: "quoted millis", input: `"1700000000000"`, want: time.UnixMilli(1700000000000)},
		{name: "bare numeric millis", input: `1700000000000`, want: time.UnixMilli(1700000000000)},
		{name: "null is zero", input: `null`, isZero: true},
		{name: "empty string is zero", input: `""`, isZero: true},
		{name: "non-numeric is error", input: `"not-a-number"`, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v TimeInMillis
			err := json.Unmarshal([]byte(tt.input), &v)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if tt.isZero {
				assert.True(t, v.IsZero())
				return
			}
			assert.True(t, tt.want.Equal(v.Time), "expected %v, got %v", tt.want, v.Time)
		})
	}
}

// TestTimeInMillis_ShortInputDoesNotPanic guards the slice-bounds fix: a short
// or unquoted value must not panic during unmarshaling.
func TestTimeInMillis_ShortInputDoesNotPanic(t *testing.T) {
	for _, input := range []string{`"`, `5`, `"5"`, ``} {
		assert.NotPanics(t, func() {
			var v TimeInMillis
			_ = v.UnmarshalJSON([]byte(input))
		})
	}
}
