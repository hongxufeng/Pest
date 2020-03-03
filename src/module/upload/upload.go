package upload

import (
	"utils/config"
	"utils/service"
	"utils/upload"

	"github.com/hongxufeng/fileLogger"
)

var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger

type UploadModule struct {
	level service.LEVEL
}

func (module *UploadModule) Init(conf *config.Config) error {
	module.level = service.SetEnvironment(conf.Environment)
	Info = fileLogger.NewDefaultLogger(conf.LogDir, "Upload_Info.log")
	Error = fileLogger.NewDefaultLogger(conf.LogDir, "Upload_Error.log")
	Info.SetPrefix("[Upload] ")
	Error.SetPrefix("[Upload] ")
	return nil
}
func (r *UploadModule) Base_UploadFile(req *service.HttpRequest, res map[string]interface{}) (e error) {
	bytes, filename, e := req.PostFileInfo("file")
	stype := req.PostParam("type")
	switch stype {
	case "avatar", "picture", "voice", "video":
	default:
		stype = "other"
	}
	//get extname from fname
	if e != nil {
		e = service.NewError(service.ERR_INVALID_PARAM, e.Error(), "上传错误")
		return
	}
	remoteName, e := upload.PutFileBytes(bytes, filename, stype)
	if e != nil {
		e = service.NewError(service.ERR_INVALID_PARAM, e.Error(), "上传文件到服务器错误")
		return
	}

	rs := make(map[string]interface{})
	rs["url"] = remoteName
	res["res"] = rs
	return
}
