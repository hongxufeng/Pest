package report

import (
	"bytes"
	"database/sql"
	"datahelper/db"
	"model"
	"utils/function"
	"utils/service"
)

func GetCURD(req *service.HttpRequest, param *Param, settings *model.CRUDSettings, bodybuf *bytes.Buffer) (err error) {
	bodybuf.WriteString("<form class=\"form-horizontal\">")
	if settings.Cmd == model.Cmd_Create {
		err = BuildCBody(param, settings, bodybuf)
		if err != nil {
			return err
		}
	} else if settings.Cmd == model.Cmd_View || settings.Cmd == model.Cmd_Edit {
		err = BuildVEBody(req, param, settings, bodybuf)
		if err != nil {
			return err
		}
	}
	bodybuf.WriteString("</form>")
	return
}
func QueryThis(req *service.HttpRequest, param *Param, settings *model.CRUDSettings) (result *sql.Rows, size int, err error) {

	return
}
func BuildVEBody(req *service.HttpRequest, param *Param, settings *model.CRUDSettings, bodybuf *bytes.Buffer) (err error) {
	query, err := BuildQueryThisSQL(req, param, settings)
	//fmt.Println(query)
	if err != nil {
		return
	}
	row, err := db.Query(query)
	if err != nil {
		return
	}
	defer row.Close()
	columns, _ := row.Columns()
	size := len(columns)
	var s []interface{}
	for i := 0; i < size; i++ {
		var white = ""
		var p *string
		p = &white
		s = append(s, p)
	}
	for row.Next() {
		err = row.Scan(s...)
		if err != nil {
			return
		}
		//fmt.Println(function.PArrayToSArray(s))
		var current = 0
		for _, colConfig := range param.ColConfigDict {
			if colConfig.Tag == "buttons" || colConfig.Tag == "pagerbuttons" || colConfig.Visibility == "table-none" {
				continue
			}
			bodybuf.WriteString("<div class=\"form-group\">")
			bodybuf.WriteString("<label class=\"col-sm-3 control-label\">")
			bodybuf.WriteString(colConfig.Text)
			bodybuf.WriteString("&nbsp;&nbsp;<span class=\"rt-glyphicon-color\">:</span></label>")
			bodybuf.WriteString("<div class=\"col-sm-6\">")
			bodybuf.WriteString("<input ")
			if settings.Cmd == model.Cmd_View {
				bodybuf.WriteString("readonly ")
			}
			bodybuf.WriteString("name=\"")
			bodybuf.WriteString(colConfig.Tag)
			bodybuf.WriteString("\" type=\"text\" class=\"form-control rt-form-control\" placeholder=\"")
			bodybuf.WriteString(colConfig.Text)
			bodybuf.WriteString("\" value=\"")
			cell := function.ToString(s[current])
			cell, _ = Format(&param.ColConfigDict[current], cell)
			bodybuf.WriteString(cell)
			bodybuf.WriteString("\">")
			bodybuf.WriteString("</div>")
			bodybuf.WriteString("</div>")
			current++
		}
		bodybuf.WriteString("<div class=\"form-group\">")
		bodybuf.WriteString("<div class=\"col-sm-offset-3  col-sm-6\">")
		if settings.Cmd == model.Cmd_Create {
			bodybuf.WriteString("<input type=\"submit\" class=\"btn btn-primary form-control\" value=\"")
			if param.TableConfig.HasBtnCreateText {
				bodybuf.WriteString(param.TableConfig.BtnCreateText)
			} else {
				bodybuf.WriteString("新&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;增")
			}
			bodybuf.WriteString("\" />")
		} else if settings.Cmd == model.Cmd_View {
			bodybuf.WriteString("<input type=\"text\" class=\"btn btn-primary form-control rt-view-return\" value=\"")
			bodybuf.WriteString("浏&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;览")
			bodybuf.WriteString("\" />")
		} else if settings.Cmd == model.Cmd_Edit {
			bodybuf.WriteString("<input type=\"submit\" class=\"btn btn-primary form-control\" value=\"")
			bodybuf.WriteString("编&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;辑")
			bodybuf.WriteString("\" />")

		}
	}
	return
}
func BuildCBody(param *Param, settings *model.CRUDSettings, bodybuf *bytes.Buffer) (err error) {
	for _, colConfig := range param.ColConfigDict {
		if colConfig.Tag == "buttons" || colConfig.Tag == "pagerbuttons" || colConfig.Visibility == "table-none" {
			continue
		}
		bodybuf.WriteString("<div class=\"form-group\">")
		bodybuf.WriteString("<label class=\"col-sm-3 control-label\">")
		bodybuf.WriteString(colConfig.Text)
		bodybuf.WriteString("&nbsp;&nbsp;<span class=\"rt-glyphicon-color\">:</span></label>")
		bodybuf.WriteString("<div class=\"col-sm-6\">")
		bodybuf.WriteString("<input name=\"")
		bodybuf.WriteString(colConfig.Tag)
		bodybuf.WriteString("\" type=\"text\" class=\"form-control rt-form-control\" placeholder=\"")
		bodybuf.WriteString(colConfig.Text)
		bodybuf.WriteString("\">")
		bodybuf.WriteString("</div>")
		bodybuf.WriteString("</div>")
	}
	bodybuf.WriteString("<div class=\"form-group\">")
	bodybuf.WriteString("<div class=\"col-sm-offset-3  col-sm-6\">")
	if settings.Cmd == model.Cmd_Create {
		bodybuf.WriteString("<input type=\"submit\" class=\"btn btn-primary form-control\" value=\"")
		if param.TableConfig.HasBtnCreateText {
			bodybuf.WriteString(param.TableConfig.BtnCreateText)
		} else {
			bodybuf.WriteString("新&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;增")
		}
		bodybuf.WriteString("\" />")
	}
	return
}
