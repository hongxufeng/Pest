package db

import (
	"errors"
	"fmt"
	"model"
	"time"
	"utils/function"

	"github.com/go-redis/redis"
)

const (
	userlist = "user_list"
)

func UserLoginErrCnt(uid uint32) (cnt int64, err error) {
	cnt, err = RedisCache.Incr(function.MakeKey(CACHE_USER_LOGIN_ERROR, uid)).Result()
	if err != redis.Nil && err != nil {
		return
	} else {
		//cnt++
		err = RedisCache.Set(function.MakeKey(CACHE_USER_LOGIN_ERROR, uid), cnt, model.User_Forfid_Expiration_Duration).Err()
	}
	//fmt.Println(cnt)
	return
}
func SetUserForbid(uid uint32) (err error) {
	err = RedisCache.Set(function.MakeKey(CACHE_USER_LOGIN_FORBID, uid), true, model.User_Forfid_Expiration_Duration).Err()
	return
}
func CheckUserForbid(uid uint32) (forbid bool, err error) {
	//验证
	_, err = RedisCache.Get(function.MakeKey(CACHE_USER_LOGIN_FORBID, uid)).Result()
	if err == redis.Nil {
		forbid = false
	} else if err != nil {
		forbid = true
	}
	return
}
func GetUidbyName(name string) (uid uint32, err error) {
	uid = 0
	queryStr := "SELECT uid FROM " + userlist + " WHERE username=? limit 0,1"
	result, err := MysqlMain.Query(queryStr, name)
	if err != nil {
		return
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&uid)
	}
	return
}
func GetUserInfobyDB(uid uint32) (*UserInfo, error) {
	userinfo := new(UserInfo)
	query := "SELECT uid,username,nickname,password,salt,state,power,avatar,limit_name,limit_id,user_agent FROM " + userlist + " WHERE uid=? limit 0,1"
	result, e := MysqlMain.Query(query, uid)
	if e != nil {
		return userinfo, e
	}
	defer result.Close()
	if result.Next() {
		e = result.Scan(&userinfo.Uid, &userinfo.UserName, &userinfo.NickName, &userinfo.Password, &userinfo.Salt, &userinfo.State, &userinfo.Power, &userinfo.Avatar, &userinfo.LimitName, &userinfo.LimitID, &userinfo.UserAgent)
	} else {
		e = errors.New("您输入的用户未找到呢！")
	}
	return userinfo, e
}
func GetCheckUser(param string, name string) (bool, error) {
	var count int
	query := "SELECT Count(*) FROM " + userlist + " WHERE " + param + "=?"
	result, e := MysqlMain.Query(query, name)
	if e != nil {
		return true, e
	}
	defer result.Close()
	if result.Next() {
		e = result.Scan(&count)
	} else {
		return true, nil
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func ExecCreateUser(registerData *model.RegisterData) error {
	salt := function.GetSixRandom()
	md5password := function.Md5String(fmt.Sprintf("%s_%s", registerData.Password, salt))
	query := "Insert into " + userlist + " (username,password,nickname,salt,state,power,avatar,limit_name,limit_id,create_time) values (?,?,?,?,?,?,?,?,?,?)"
	//fmt.Println(query)
	err := Exec(query, registerData.Username, md5password, registerData.Nickname, salt, 0, registerData.Power, model.User_W_Avatar, registerData.Limit_Name, registerData.Limit_ID, time.Now().UnixNano())
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_User)
}
func ExecDeleteUser(uid int) error {
	query := "delete from " + userlist + " where uid=?"
	//fmt.Println(query)
	err := Exec(query, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_User)

}
