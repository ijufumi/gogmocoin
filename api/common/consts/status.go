package consts

// Status is the numeric status code returned in every GMO Coin REST response.
// StatusOK (0) indicates success; any other value indicates an error.
type Status int

// StatusOK is the success status code.
const StatusOK = Status(0)

// IsOK reports whether the status indicates a successful response.
func (s Status) IsOK() bool {
	return s == StatusOK
}
