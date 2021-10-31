package model

import "nearby/biz/model/vo"

type SendMessageRequest struct {
	ReceiverID int64 `form:"receiver_id" json:"receiver_id"` // 接收方id, 可不传;
	ConvID     int64 `form:"conv_id" json:"conv_id"`         // 会话id
	/*
		枚举值[
		"visitor": 游客;
		"sys_helper": 系统客服;
		"be_helper": 后台客服;
		"user": 普通用户;
		]
	*/
	Role    string        `form:"role" json:"role"`       // 用户身份
	Content vo.MsgContent `form:"content" json:"content"` // 消息内容
	/*
		枚举值[
		"text": 文本;
		"rich_text": 富文本;
		"image": 图片;
		"audio": 语音;
		"video": 视频;
		]
	*/
	Type      string `form:"type" json:"type"`           // 消息类型
	Status    string `form:"status" json:"status"`       // 消息状态
	Timestamp int64  `form:"timestamp" json:"timestamp"` // 消息时间戳
}

type SendMessageData struct {
	MessageID int64 `json:"message_id"` // 该消息的唯一标识id
	ConvID    int64 `json:"conv_id"`    // 该消息所属的会话id
}

type SendMessageResponse struct {
	Meta Meta            `json:"meta"`
	Data SendMessageData `json:"data"`
}
