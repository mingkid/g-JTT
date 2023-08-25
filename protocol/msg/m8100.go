package msg

type M8100 struct {
	AnswerSerialNo uint16
	Result         M8100Result
	Token          string
}

type M8100Result uint8

const (
	// M8100ResultSuccess 成功
	M8100ResultSuccess M8100Result = iota

	// M8100ResultCarRegistered 车辆已被注册
	M8100ResultCarRegistered

	// M8100ResultCarNotInDB 数据库中无该车辆
	M8100ResultCarNotInDB

	// M8100ResultTermRegistered 终端已被注册
	M8100ResultTermRegistered

	// M8100TermNotInDB 数据库中无该终端
	M8100TermNotInDB
)
