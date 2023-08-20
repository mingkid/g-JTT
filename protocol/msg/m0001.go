package msg

// M0001 终端通用应答
type M0001 struct {
	Head
	AnswerSerialNo uint16      // 应答流水号
	AnswerMsgID    MsgID       // 应答消息ID
	Result         M0001Result // 处理结果
	ErrorCode      uint16      // 错误代码
}

type M0001Result uint8

const (
	M0001ResultOK          M0001Result = iota // 成功/确认
	M0001ResulFAIL                            // 失败
	M0001ResulMSG_ERROR                       // 消息有误
	M0001ResulUNKNOWN                         // 不支持
	M0001ResulWARNING_DEAL                    // 报警处理确认
)
