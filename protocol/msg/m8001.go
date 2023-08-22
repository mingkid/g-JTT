package msg

type M8001 struct {
	Head
	AnswerSerialNo uint16      // 应答流水号，对应的终端消息的流水号
	AnswerMsgID    MsgID       // 应答ID，对应的终端消息的ID
	Result         M8001Result // 结果
}

type M8001Result byte

const (
	// M8001ResultSuccess 成功
	M8001ResultSuccess M8001Result = iota

	// M8001ResultFail 失败
	M8001ResultFail

	// M8001ResultMsgErr 消息有误
	M8001ResultMsgErr

	// M8001ResultUnsupported 不支持
	M8001ResultUnsupported

	// M8001WarnConfirm 报警处理确认
	M8001WarnConfirm
)
