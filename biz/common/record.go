package common

import (
	"ai_helper/biz/domain/val_obj"
)

var Record = map[string]*val_obj.UserClaims{
	"aaa": {
		UserID:   HelperID,
		Nickname: "aaa",
		HeadURL:  "",
		IsHelper: true,
	},
	"bbb": {
		UserID:   HelperID,
		Nickname: "bbb",
		HeadURL:  "",
		IsHelper: true,
	},
	"ccc": {
		UserID:   9999997,
		Nickname: "ccc",
		IsHelper: true,
	},
}

func IsHelper(userID string) bool {
	return Record[userID] != nil && Record[userID].IsHelper
}
