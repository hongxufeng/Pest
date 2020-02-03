package report

import (
	"datahelper/report"
	"fmt"
	"model"
	"utils/config"
	"utils/service"

	"github.com/hongxufeng/fileLogger"
)

var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger

type ReportModule struct {
	level service.LEVEL
}

func (module *ReportModule) Init(conf *config.Config) error {
	module.level = service.SetEnvironment(conf.Environment)
	Info = fileLogger.NewDefaultLogger(conf.LogDir, "Report_Info.log")
	Error = fileLogger.NewDefaultLogger(conf.LogDir, "Report_Error.log")
	Info.SetPrefix("[REPORT] ")
	Error.SetPrefix("[REPORT] ")
	return nil
}

func (module *ReportModule) User_GetTable(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var settings model.Settings
	err = req.GetParams(false, "page", &settings.Page, "rows", &settings.Rows, "colpage", &settings.ColPage)
	if err != nil {
		return
	}
	_ = req.GetParams(true, "sort", &settings.Order)
	fmt.Println(settings.TableID)
	err = req.ParseEncodeUrl(false, "table", &settings.TableID, "configFile", &settings.ConfigFile, "hasCheckbox", &settings.HasCheckbox, "style", &settings.Style, "rowList", &settings.RowList)
	if err != nil {
		return
	}
	if module.level >= service.DEV {
		fmt.Println(settings)
	}
	param, err := report.New(req.Uid, settings.ConfigFile, settings.TableID)
	if err != nil {
		return
	} else {
		result["res"], err = param.Table(req, &settings)
	}
	return
}

func (module *ReportModule) User_GetTableJson(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var settings model.Settings
	err = req.GetParams(false, "page", &settings.Page, "rows", &settings.Rows)
	if err != nil {
		return
	}
	_ = req.GetParams(true, "sort", &settings.Order)
	fmt.Println(settings.TableID)
	err = req.ParseEncodeUrl(false, "table", &settings.TableID, "configFile", &settings.ConfigFile)
	if err != nil {
		return
	}
	if module.level >= service.DEV {
		fmt.Println(settings)
	}
	param, err := report.New(req.Uid, settings.ConfigFile, settings.TableID)
	if err != nil {
		return
	} else {
		result["res"], err = param.TableJson(req, &settings)
	}
	return
}
func (module *ReportModule) Base_GetTableJson(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var settings model.Settings
	err = req.GetParams(false, "page", &settings.Page, "rows", &settings.Rows)
	if err != nil {
		return
	}
	_ = req.GetParams(true, "sort", &settings.Order)
	fmt.Println(settings.TableID)
	err = req.ParseEncodeUrl(false, "table", &settings.TableID, "configFile", &settings.ConfigFile)
	if err != nil {
		return
	}
	if module.level >= service.DEV {
		fmt.Println(settings)
	}
	param, err := report.New(req.Uid, settings.ConfigFile, settings.TableID)
	if err != nil {
		return
	} else {
		result["res"], err = param.TableJson(req, &settings)
	}
	return
}
func (module *ReportModule) User_SearchTree(req *service.HttpRequest, result map[string]interface{}) (err error) {
	return
}

func (module *ReportModule) User_LocateNode(req *service.HttpRequest, result map[string]interface{}) (err error) {
	return
}

func (module *ReportModule) User_GetPageCURD(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var settings model.CRUDSettings
	err = req.ParseEncodeUrl(false, "table", &settings.TableID, "configFile", &settings.ConfigFile, "cmd", &settings.Cmd)
	if err != nil {
		return
	}
	if settings.Cmd != model.Cmd_Create {
		err = req.ParseEncodeUrl(true, "queryKey", &settings.QueryKey, "queryValue", &settings.QueryValue)
		if err != nil {
			return
		}
	}
	if module.level >= service.DEV {
		fmt.Println(settings)
	}
	//fmt.Println("req.Uid" + string(req.Uid) + "settings.ConfigFile" + settings.ConfigFile + "settings.TableID" + settings.TableID)
	param, err := report.New(req.Uid, settings.ConfigFile, settings.TableID)
	if err != nil {
		return
	} else {
		result["res"], err = param.CURD(req, &settings)
	}
	return
}
