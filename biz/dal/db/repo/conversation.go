package repo

import (
	"context"
	"nearby/biz/dal/db"
	"nearby/biz/dal/db/po"

	"github.com/jinzhu/gorm"
)

type ConversationRepo struct {
	db *gorm.DB
}

func NewConversationRepo() *ConversationRepo {
	return &ConversationRepo{
		db: db.GetDB().Debug().LogMode(true),
	}
}

func (repo *ConversationRepo) CreateConversation(ctx context.Context, convPo po.Conversation) (*po.Conversation, error) {
	sql := repo.db.Model(po.Conversation{})
	if err := sql.Omit("id").Create(&convPo).Error; err != nil {
		return nil, err
	}
	return &convPo, nil
}

func (repo *ConversationRepo) CreateUserConvRelPo(ctx context.Context, userConvRelPo po.UserConvRel) (*po.UserConvRel, error) {
	sql := repo.db.Model(po.UserConvRel{})
	if err := sql.Omit("id").Create(&userConvRelPo).Error; err != nil {
		return nil, err
	}
	return &userConvRelPo, nil
}

type GetConvPoRequest struct {
	ConvID int64 `json:"conv_id"`
}

func (repo *ConversationRepo) GetConvPo(ctx context.Context, req GetConvPoRequest) (*po.Conversation, error) {
	sql := repo.db.Model(po.Conversation{})
	var conv po.Conversation
	if req.ConvID != 0 {
		sql = sql.Where("conv_id = ?", req.ConvID)
	}
	if err := sql.Find(&conv).Error; err != nil {
		return nil, err
	}
	return &conv, nil
}

type GetUserConvRelPosRequest struct {
	UserID        int64 `json:"user_id"`
	Limit         int64 `json:"limit"`
	TimestampFrom int64 `json:"timestamp_from"`
	TimestampTo   int64 `json:"timestamp_to"`
}

func (repo *ConversationRepo) GetUserConvRelPos(ctx context.Context, req GetUserConvRelPosRequest) (pos []*po.UserConvRel, total int64, err error) {
	sql := repo.db.Model(po.Conversation{})
	pos = make([]*po.UserConvRel, 0)
	sql = sql.Where("user_id = ?", req.UserID)
	sql = sql.Joins("left join conversation on conversation.conv_id = user_conv_rel.conv_id")
	if req.TimestampTo != 0 {
		sql = sql.Where("conversation.timestamp < ?", req.TimestampTo)
	}
	if req.TimestampFrom != 0 {
		sql = sql.Where("conversation.timestamp > ?", req.TimestampFrom)
	}
	sql = sql.Order("conversation.timestamp desc")
	sql = sql.Limit(req.Limit)
	err = sql.Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	err = sql.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return pos, total, nil
}
