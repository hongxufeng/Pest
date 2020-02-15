var deleteHouse = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("base/pest/DeleteHouse", {
        Uid: value
    }, function(data) {
        // var jsonObject = JSON.parse(data);
        if (data.status === "fail") {
            alert(data.msg);
            return false
        }
        if (data.status === "ok") {
            if (data.res.deletestatus === 1) {
                alert(data.res.msg);
            }
            location.href = location.href;
        }
    });
}
var deleteUnit = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("base/pest/DeleteUnit", {
        Uid: value
    }, function(data) {
        // var jsonObject = JSON.parse(data);
        if (data.status === "fail") {
            alert(data.msg);
            return false
        }
        if (data.status === "ok") {
            if (data.res.deletestatus === 1) {
                alert(data.res.msg);
            }
            location.href = location.href;
        }
    });
}
var deletePersonnel = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("base/pest/DeletePersonnel", {
        Uid: value
    }, function(data) {
        // var jsonObject = JSON.parse(data);
        if (data.status === "fail") {
            alert(data.msg);
            return false
        }
        if (data.status === "ok") {
            if (data.res.deletestatus === 1) {
                alert(data.res.msg);
            }
            location.href = location.href;
        }
    });
}
var deleteReport = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("base/pest/DeleteDailyReport", {
        Uid: value
    }, function(data) {
        // var jsonObject = JSON.parse(data);
        if (data.status === "fail") {
            alert(data.msg);
            return false
        }
        if (data.status === "ok") {
            if (data.res.deletestatus === 1) {
                alert(data.res.msg);
            }
            location.href = location.href;
        }
    });
}
var viewHousePersonnel = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    var query = "?house_id=" + value;
    $("#table").rt({
        configFile: "pest",
        table: "house_personnel",
        query: query,
        selector: false,
        complete: function() {
            $(".rt-edit").click(editPersonnel);
            $(".rt-delete").click(deletePersonnel);
            $(".rt-report").click(viewReport);
        }
    });
}
var viewHouseUnit = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    var query = "?house_id=" + value;
    $("#table").rt({
        configFile: "pest",
        table: "house_unit",
        query: query,
        selector: false,
        complete: function() {
            $(".rt-edit").click(editUnit);
            $(".rt-delete").click(deleteUnit);
            $(".rt-personnel").click(viewUnitPersonnel);
        }
    });
}
var viewUnitPersonnel = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    var query = "?unit_id=" + value;
    $("#table").rt({
        configFile: "pest",
        table: "unit_personnel",
        query: query,
        selector: false,
        complete: function() {
            $(".rt-edit").click(editPersonnel);
            $(".rt-delete").click(deletePersonnel);
            $(".rt-report").click(viewReport);
        }
    });
}
var viewReport = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    var query = "?personnel_id=" + value;
    $("#table").rt({
        configFile: "dailyreport",
        table: "daily_report_list",
        query: query,
        selector: false,
        complete: function() {
            $(".rt-delete").click(deleteReport);
        }
    });
}
var editHouse = function() {
    var nature = $(this).parent().siblings('[name=nature]').data('value').split(',');
    var html = '<form class="form-horizontal" autocomplete="off">'
    html += '<div class="form-group"><label class="col-sm-3 control-label">UID<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="uid" type="text" class="form-control rt-form-control" placeholder="UID" value="' + $(this).parent().siblings('[name=uid]').data('value') + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所属街道/小区编号<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="street_no" type="text" class="form-control rt-form-control" placeholder="所属街道/小区编号" value="' + $(this).parent().siblings('[name=street_no]').data('value') + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所属街道/小区<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="street" type="text" class="form-control rt-form-control" placeholder="所属街道/小区" value="' + $(this).parent().siblings('[name=street]').data('value') + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">房屋性质<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><div class="checkbox"><label for="nature"><input type="checkbox" name="nature" value="自住">自住</label><label for="nature"><input type="checkbox" name="nature" value="出租">出租</label><label for="nature"><input type="checkbox" name="nature" value="单位">单位</label><label for="nature"><input type="checkbox" name="nature" value="闲置">闲置</label><label for="nature"><input type="checkbox" name="nature" value="废弃">废弃</label></div></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">详细地址<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input  name="address" type="text" class="form-control rt-form-control" placeholder="详细地址" value="' + $(this).parent().siblings('[name=address]').data('value') + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">常驻人口<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input  name="number" type="text" class="form-control rt-form-control" placeholder="常驻人口" value="' + $(this).parent().siblings('[name=number]').data('value') + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">备注<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="comment" type="text" class="form-control rt-form-control" placeholder="备注" value="' + $(this).parent().siblings('[name=comment]').data('value') + '"></div></div>';
    html += '<div class="form-group"><div class="col-sm-offset-3  col-sm-6"><input type="button" onclick="updatehouse()" class="btn btn-primary form-control rt-submit" value="提&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;交"></div></div>';
    $("#table").html(html);
    $.each($('input[name=nature]'), function() {
        var v = $(this).val();
        if ($.inArray(v, nature) > -1) {
            $(this).attr('checked', 'checked');
        }
    });
    $("input[name='nature']").each(function() {
        var currentEle = $(this);
        currentEle.click(function() {
            var elem = $(this);
            if ((elem.val() == '闲置' || elem.val() == '废弃') && elem.prop("checked") == true) { //如果选中废弃或者闲置，则清空其他选中状态
                $("input[name='nature']").each(function() {
                    $(this).removeAttr("checked");
                    elem.attr('checked', true);
                });
            } else {
                if (elem.prop("checked") == true) { //如果其他选项被选中，则清空废弃或闲置选中状态
                    $("input[name='nature']").eq(3).removeAttr("checked");
                    $("input[name='nature']").eq(4).removeAttr("checked");
                }
            }
        });
    });
}
var editUnit = function() {
    var uid = $(this).parent().siblings('[name=uid]').data('value');
    var name = $(this).parent().siblings('[name=name]').data('value');
    var license_number = $(this).parent().siblings('[name=license_number]').data('value');
    var identification_number = $(this).parent().siblings('[name=identification_number]').data('value');
    var picture = $(this).parent().siblings('[name=picture]').data('value');
    var kind = $(this).parent().siblings('[name=kind]').data('value');
    var scale = $(this).parent().siblings('[name=scale]').data('value');
    var tel = $(this).parent().siblings('[name=tel]').data('value');
    var bank_name = $(this).parent().siblings('[name=bank_name]').data('value');
    var bank_account = $(this).parent().siblings('[name=bank_account]').data('value');
    var comment = $(this).parent().siblings('[name=unit_comment]').data('value');
    var html = '<form class="form-horizontal" autocomplete="off">'
    html += '<div class="form-group"><label class="col-sm-3 control-label">UID<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Uid" type="text" class="form-control rt-form-control" placeholder="UID" value="' + uid + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">单位名称<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Name" type="text" class="form-control rt-form-control" placeholder="单位名称" value="' + name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">证件号码<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="License_Number" type="text" class="form-control rt-form-control" placeholder="营业执照号码或组织机构代码" value="' + license_number + '"></div></div>';
    html += '<div class="form-group hidden"><label class="col-sm-3 control-label">企业机构代码<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Identification_Number" type="text" class="form-control rt-form-control" placeholder="企业机构代码" value="' + identification_number + '"></div></div>';
    html += '<div class="form-group hidden"><label class="col-sm-3 control-label">证件照片<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Picture" type="text" class="form-control rt-form-control" placeholder="证件照片" value="' + picture + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">单位类别<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><label class="radio-inline"><input type="radio" name="Kind"  value="1" /> 政府机关</label><label class="radio-inline"><input type="radio" name="Kind"  value="2" /> 事业单位</label><label class="radio-inline"><input type="radio" name="Kind"  value="3" /> 企业单位</label><label class="radio-inline"><input type="radio" name="Kind"  value="4" /> 个体商户</label><label class="radio-inline"><input type="radio" name="Kind"  value="5" /> 民间机构</label></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">单位规模<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><label class="radio-inline"><input type="radio" name="Scale"  value="1" /> 10人以下</label><label class="radio-inline"><input type="radio" name="Scale"  value="2" /> 10-50人</label><label class="radio-inline"><input type="radio" name="Scale"  value="3" /> 50-200人</label><label class="radio-inline"><input type="radio" name="Scale"  value="4" /> 200人以上</label></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">单位电话<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Tel" type="text" class="form-control rt-form-control" placeholder="单位电话" value="' + tel + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">开户行名称<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Bank_Name" type="text" class="form-control rt-form-control" placeholder="开户行名称" value="' + bank_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">开户行帐号<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Bank_Account" type="text" class="form-control rt-form-control" placeholder="开户行帐号" value="' + bank_account + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">备注<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Comment" type="text" class="form-control rt-form-control" placeholder="备注" value="' + comment + '"></div></div>';
    html += '<div class="form-group"><div class="col-sm-offset-3  col-sm-6"><input type="button" onclick="updateunit()" class="btn btn-primary form-control rt-submit" value="提&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;交"></div></div>';
    $("#table").html(html);
    $("input:radio[name=Kind][value='" + kind + "']").prop("checked", "checked");
    $("input:radio[name=Scale][value='" + scale + "']").prop("checked", "checked");
}
var editPersonnel = function() {
    var uid = $(this).parent().siblings('[name=uid]').data('value');
    var name = $(this).parent().siblings('[name=name]').data('value');
    var card_no = $(this).parent().siblings('[name=card_no]').data('value');
    var phone = $(this).parent().siblings('[name=phone]').data('value');
    var occupation = $(this).parent().siblings('[name=occupation]').data('value');
    var sex = $(this).parent().siblings('[name=sex]').data('value');
    var nation = $(this).parent().siblings('[name=nation]').data('value');
    var birthday = $(this).parent().siblings('[name=birthday]').data('value');
    var address = $(this).parent().siblings('[name=address]').data('value');
    var sign_organization = $(this).parent().siblings('[name=sign_organization]').data('value');
    var limited_date = $(this).parent().siblings('[name=limited_date]').data('value');
    var history = $(this).parent().siblings('[name=history]').data('value');
    var remark = $(this).parent().siblings('[name=remark]').data('value');
    var face_picture = $(this).parent().siblings('[name=face_picture]').data('value');
    var card_picture_front = $(this).parent().siblings('[name=card_picture_front]').data('value');
    var card_picture_back = $(this).parent().siblings('[name=card_picture_back]').data('value');
    var html = '<form class="form-horizontal" autocomplete="off">'
    html += '<div class="form-group"><label class="col-sm-3 control-label">UID<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Uid" type="text" class="form-control rt-form-control" placeholder="UID" value="' + uid + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">人员姓名<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Name" type="text" class="form-control rt-form-control" placeholder="人员姓名" value="' + name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">身份证号<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Card_No" type="text" class="form-control rt-form-control" placeholder="身份证号" value="' + card_no + '"></div></div>';
    html += '<div class="form-group hidden"><label class="col-sm-3 control-label">手机号<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Phone" type="text" class="form-control rt-form-control" placeholder="手机号" value="' + phone + '"></div></div>';
    html += '<div class="form-group hidden"><label class="col-sm-3 control-label">职业<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Occupation" type="text" class="form-control rt-form-control" placeholder="职业" value="' + occupation + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">性别<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><label class="radio-inline"><input type="radio" name="Sex"  value="1" /> 男</label><label class="radio-inline"><input type="radio" name="Sex"  value="2" /> 女</label></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">民族<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Nation" type="text" class="form-control rt-form-control" placeholder="民族" value="' + nation + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">出生年月日<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Birthday" type="text" class="form-control rt-form-control" placeholder="出生年月日" value="' + birthday + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">身份证地址<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Address" type="text" class="form-control rt-form-control" placeholder="身份证地址" value="' + address + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">签发机构<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Sign_Organization" type="text" class="form-control rt-form-control" placeholder="签发机构" value="' + sign_organization + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">有效期<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Limited_Date" type="text" class="form-control rt-form-control" placeholder="有效期" value="' + limited_date + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">接触史<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="History" type="text" class="form-control rt-form-control" placeholder="接触史" value="' + history + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">备注<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="Remark" type="text" class="form-control rt-form-control" placeholder="备注" value="' + remark + '"></div></div>';
    html += '<div class="form-group hidden"><label class="col-sm-3 control-label">人脸照片<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Face_Picture" type="text" class="form-control rt-form-control" placeholder="人脸照片" value="' + face_picture + '"></div></div>';
    html += '<div class="form-group hidden"><label class="col-sm-3 control-label">身份证正面<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Card_Picture_Front" type="text" class="form-control rt-form-control" placeholder="身份证正面" value="' + card_picture_front + '"></div></div>';
    html += '<div class="form-group hidden"><label class="col-sm-3 control-label">身份证背面<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="Card_Picture_Back" type="text" class="form-control rt-form-control" placeholder="身份证背面" value="' + card_picture_back + '"></div></div>';
    html += '<div class="form-group"><div class="col-sm-offset-3  col-sm-6"><input type="button" onclick="updatepersonnel()" class="btn btn-primary form-control rt-submit" value="提&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;交"></div></div>';
    $("#table").html(html);
    $("input:radio[name=Sex][value='" + sex + "']").prop("checked", "checked");
}

var updatehouse = function() {
    var arr = [];
    $.each($('input[name=nature]:checked'), function() {
        arr.push($(this).val());
    });
    var nature = arr.join(",");
    var uid = $('[name=uid]').val();
    var street_no = $('[name=street_no]').val();
    var street = $('[name=street]').val();
    var address = $('[name=address]').val();
    var number = $('[name=number]').val();
    var comment = $('[name=comment]').val();
    $.post("base/pest/UpdateHouse", {
        Uid: uid,
        Nature: nature,
        Street: street,
        Street_No: street_no,
        Address: address,
        Number: number,
        Comment: comment
    }, function(data) {
        // var jsonObject = JSON.parse(data);
        if (data.status === "fail") {
            alert(data.msg);
            return false
        }
        if (data.status === "ok") {
            if (data.res.updatestatus === 1) {
                alert(data.res.msg);
            }
            location.href = location.href;
        }
    });
}
var updateunit = function() {
    $.post("base/pest/UpdateUnit",
        $("form").serialize(),
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                if (data.res.updatestatus === 1) {
                    alert(data.res.msg);
                }
                location.href = location.href;
            }
        }
    );
};
var updatepersonnel = function() {
    $.post("base/pest/UpdatePersonnel",
        $("form").serialize(),
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                if (data.res.updatestatus === 1) {
                    alert(data.res.msg);
                }
                location.href = location.href;
            }
        }
    );
};