package consts

type Status int

const StatusOK = Status(0)

func (s Status) IsOK() bool {
	return s == StatusOK
}
