package user

import (
	"datahelper/user"
	"fmt"
	"model"
	"utils/config"
	"utils/service"

	"github.com/hongxufeng/fileLogger"
)

var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger

type UserModule struct {
	level service.LEVEL
}

func (module *UserModule) Init(conf *config.Config) error {
	module.level = service.SetEnvironment(conf.Environment)
	Info = fileLogger.NewDefaultLogger(conf.LogDir, "User_Info.log")
	Error = fileLogger.NewDefaultLogger(conf.LogDir, "User_Error.log")
	Info.SetPrefix("[USER] ")
	Error.SetPrefix("[USER] ")
	return nil
}

func (module *UserModule) Base_UserLogin(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var loginData model.LoginData
	err = req.ParseEncodeUrl(false, "username", &loginData.Username, "password", &loginData.Password)
	if err != nil {
		return
	}
	if module.level >= service.DEV {
		fmt.Println(loginData.Username)
	}
	uid := user.GetUidbyName(loginData.Username)
	if module.level >= service.DEV {
		fmt.Println(uid)
	}
	if uid == 0 {
		result["res"] = user.CreateFailResp(service.ERR_USER_NOT_FOUND, "貌似您输入的帐号不存在！")
		return
	}
	if state := user.CheckUserState(uid); state {
		result["res"] = user.CreateFailResp(service.RR_STATUS_DENIED, "现如今用户登录状态关闭呢！")
		return
	}
	//判断登录频繁，防止暴力破解
	if forbid := user.CheckUserForbid(uid); forbid {
		result["res"] = user.CreateFailResp(service.ERR_LOGIN_FREQUENT, "您登录有点频繁，请稍事休息！")
		return
	}
	ud, e := user.CheckAuth(uid, loginData.Password)
	if e != nil {
		Info.Info("%d  auth failed", uid)
		//这里增加判断登录频繁次数
		if forbid := user.CheckUserLoginErr(uid); forbid {
			Info.Info("forbid user :%d", uid)
			if module.level >= service.DEV {
				fmt.Printf("forbid user :%d", uid)
			}
		}
		result["res"] = user.CreateFailResp(service.ERR_INVALID_USER, "少侠，您输入的密码有误啊！")
		return
	}
	result["res"] = user.CreateSuccessResp(ud)
	return
}
func (module *UserModule) User_UserRegister(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var registerData model.RegisterData
	err = req.ParseEncodeUrl(false, "username", &registerData.Username, "password", &registerData.Password, "nickname", &registerData.Nickname, "power", &registerData.Power, "limit_name", &registerData.Limit_Name, "limit_id", &registerData.Limit_ID)
	if err != nil {
		return
	}
	if registerData.Username == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "用户名不能为空哦！")
		return
	} else {
		isexist, e := user.CheckUser("username", registerData.Username)
		if e != nil {
			return
		}
		if isexist {
			err = service.NewError(service.ERR_USER_USERNAME_EXIST, "用户名已经存在了哦！")
			return
		}
	}

	if registerData.Password == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "密码不能为空哦！")
		return
	}
	if registerData.Nickname == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "昵称不能为空哦！")
		return
	}
	if registerData.Power < 0 {
		err = service.NewError(service.ERR_INVALID_PARAM, "角色输入错误哦！")
		return
	}
	if registerData.Limit_Name != model.Limit_All || registerData.Limit_Name != model.Limit_City || registerData.Limit_Name != model.Limit_Community || registerData.Limit_Name != model.Limit_District || registerData.Limit_Name != model.Limit_Province || registerData.Limit_Name != model.Limit_Station || registerData.Limit_Name != model.Limit_Street {
		err = service.NewError(service.ERR_INVALID_PARAM, "权限输入错误哦！")
		return
	}
	userreg, err := user.CreateUser(&registerData)
	if err != nil {
		return
	}
	result["res"] = userreg
	return
}
func (module *UserModule) User_UserDelete(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var uid int
	err = req.ParseEncodeUrl(false, "uid", &uid)
	if err != nil {
		return
	}
	res, err := user.DeleteUser(uid)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
