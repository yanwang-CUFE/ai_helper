package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/config"
	"ai_helper/biz/domain/aggregate"
	"ai_helper/biz/domain/entity"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

// MessageService 消息领域上下文服务
type MessageService struct {
}

type SendMessageRequest struct {
	ConvID     int64           `json:"conv_id"`     // 会话id
	ReceiverID int64           `json:"receiver_id"` // 接收者id
	Role       string          `json:"role"`        // 用户身份
	Content    json.RawMessage `json:"content"`     // 消息内容
	Type       string          `json:"type"`        // 消息类型
	Status     string          `json:"status"`      // 消息状态
	Timestamp  int64           `json:"timestamp"`   // 消息时间戳
}

func (ss *MessageService) SendMessage(ctx context.Context, req SendMessageRequest) (*entity.MessageFrom, error) {
	// 构造出一个聚合根, 并将该聚合根持久化, 同时更新相关会话的时间戳
	msgAgg, err := ss.ConstructMessageAggregate(ctx, req)
	if err != nil {
		return nil, err
	}
	msgAgg.SetMessageID(config.GenerateIDInt64())
	err = msgAgg.MessageFrom.Persist(ctx)
	if err != nil {
		return nil, err
	}
	// 将发件箱的消息同步到收件箱
	err = msgAgg.SyncReceiverBox(ctx)
	if err != nil {
		return nil, err
	}
	return msgAgg.MessageFrom, nil
}

func (ss *MessageService) ConstructMessageAggregate(ctx context.Context, req SendMessageRequest) (*aggregate.MessageAggregate, error) {
	user := common.GetUser(ctx)
	msgAgg := aggregate.MessageAggregate{
		MessageFrom: &entity.MessageFrom{
			SenderID:   user.UserID,
			ConvID:     req.ConvID,
			Timestamp:  cast.ToTime(req.Timestamp),
			ReceiverID: req.ReceiverID,
			Type:       req.Type,
		},
	}
	// 以客服的身份发消息
	if req.Role == "helper" {
		msgAgg.MessageFrom.SenderID = common.HelperID
	}
	// 装载消息内容
	err := json.Unmarshal(req.Content, &msgAgg.MessageFrom.Content)
	if err != nil {
		return nil, err
	}
	return &msgAgg, nil
}
