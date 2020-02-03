package db

import (
	"encoding/json"
	"fmt"
	"model"
	"utils/function"

	"github.com/go-redis/redis"
)

const (
	CACHE_USER_INFO         = "W_Redis_Cache_User_Info_Byte"
	CACHE_USER_LOGIN_ERROR  = "W_Redis_Cache_User_Login_Error_Cnt"
	CACHE_USER_LOGIN_FORBID = "W_Redis_Cache_User_Login_Forbid_Bool"
)

type UserInfo struct {
	Uid             uint32 `json:"uid"`
	UserName        string `json:"user_name"`
	NickName        string `json:"nick_name"`
	State           bool   `json:"state"`
	Power           uint8  `json:"power"`
	Salt            string `json:"salt"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	UserAgent       string `json:"user_agent"`
	CacheUpdateTime int64  `json:"cache_update_time"` //上次缓存更新时间
}

func GetUserInfo(uid uint32) (userinfo *UserInfo, err error) {
	info, e := RedisCache.Get(function.MakeKey(CACHE_USER_INFO, uid)).Bytes()
	if e == redis.Nil {
		//查数据库，设置redis
		return SetUserInfoCache(uid)
	} else if e != nil {
		return nil, e
	} else {
		if len(info) > 0 {
			var rm map[string]interface{}
			e := json.Unmarshal(info, &rm)
			if e != nil {
				//查数据库，设置redis
				return SetUserInfoCache(uid)
			}
			if e := function.MapToStruct(rm, &userinfo); e != nil {
				//查数据库，设置redis
				return SetUserInfoCache(uid)
			}
		} else {
			//查数据库，设置redis
			return SetUserInfoCache(uid)
		}
	}
	return
}

func SetUserInfoCache(uid uint32) (userinfo *UserInfo, err error) {
	userinfo, err = GetUserInfobyDB(uid)
	if err != nil {
		if uid == model.User_W_Uid {
			userinfo.Uid = model.User_W_Uid
			userinfo.UserName = model.User_W_UserName
			userinfo.NickName = model.User_W_NickName
			userinfo.State = model.User_W_State
			userinfo.Salt = model.User_W_Salt
			userinfo.Password = model.User_W_Password
			userinfo.Avatar = model.User_W_Avatar
			userinfo.UserAgent = model.User_W_UserAgent
		}
	}
	bts, err := json.Marshal(userinfo)
	if err != nil {
		fmt.Println("json.Marshal error " + err.Error())
		return
	}
	err = RedisCache.Set(function.MakeKey(CACHE_USER_INFO, uid), bts, model.User_Info_Persistence_Duration).Err()
	return
}
