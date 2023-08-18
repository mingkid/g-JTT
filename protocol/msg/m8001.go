package msg

type M8001 struct {
}

type M8001Result uint16

const (
	M8001Success M8001Result = iota
	M8001Fail
)
