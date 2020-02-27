$(document).ready(function() {
    $("#table").rt({
        configFile: "area",
        table: "area_list",
        complete: function() {
            $(".rt-qrcode").click(viewQr);
            $(".rt-create").click(addOne);
            $(".rt-edit").click(editThis);
            $(".rt-delete").click(deleteThis);
        }
    });
});
var viewQr = function() {
    var key = $(this).data('args');
    var value = $(this).parent().siblings('[name=' + key + ']').data('value');
    location.href = value;
};
var addOne = function() {
    location.href = "add_street.html";
}
var deleteThis = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("user/area/DeleteArea", {
        Cmd_Delete: "street",
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
var updatestreet = function() {
    var uid = $('[name=uid]').val();
    var street_name = $('[name=street_name]').val();
    var street_head = $('[name=street_head]').val();
    var street_phone = $('[name=street_phone]').val();
    var street_property_name = $('[name=street_property_name]').val();
    var street_property_phone = $('[name=street_property_phone]').val();
    $.post("user/area/UpdateStreet", {
        Uid: uid,
        Street_Name: street_name,
        Street_Head: street_head,
        Street_Phone: street_phone,
        Street_Property_Name: street_property_name,
        Street_Property_Phone: street_property_phone,
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
var editThis = function() {
    var uid = $(this).parent().siblings('[name=uid]').data('value');
    var street_name = $(this).parent().siblings('[name=street_name]').data('value');
    var street_head = $(this).parent().siblings('[name=street_head]').data('value');
    var street_phone = $(this).parent().siblings('[name=street_phone]').data('value');
    var street_property_name = $(this).parent().siblings('[name=street_property_name]').data('value');
    var street_property_phone = $(this).parent().siblings('[name=street_property_phone]').data('value');
    var community_name = $(this).parent().siblings('[name=community_name]').data('value');
    var office_name = $(this).parent().siblings('[name=office_name]').data('value');
    var station_name = $(this).parent().siblings('[name=station_name]').data('value');
    var district_name = $(this).parent().siblings('[name=district_name]').data('value');
    var city_name = $(this).parent().siblings('[name=city_name]').data('value');
    var province_name = $(this).parent().siblings('[name=province_name]').data('value');
    var html = '<form class="form-horizontal" autocomplete="off">'
    html += '<div class="form-group"><label class="col-sm-3 control-label">UID<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="uid" type="text" class="form-control rt-form-control" placeholder="UID" value="' + uid + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在省<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="province_name" type="text" class="form-control rt-form-control" placeholder="所在省" value="' + province_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在市<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="city_name" type="text" class="form-control rt-form-control" placeholder="所在市" value="' + city_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在区<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="district_name" type="text" class="form-control rt-form-control" placeholder="所在区" value="' + district_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在派出所<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="station_name" type="text" class="form-control rt-form-control" placeholder="所在派出所" value="' + station_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在办事处<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="office_name" type="text" class="form-control rt-form-control" placeholder="所在办事处" value="' + office_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在社区<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="community_name" type="text" class="form-control rt-form-control" placeholder="所在社区" value="' + community_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">小区名称<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="street_name" type="text" class="form-control rt-form-control" placeholder="小区名称" value="' + street_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">小区负责人<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="street_head" type="text" class="form-control rt-form-control" placeholder="小区负责人" value="' + street_head + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">小区电话<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="street_phone" type="text" class="form-control rt-form-control" placeholder="小区电话" value="' + street_phone + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">物业名称<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="street_property_name" type="text" class="form-control rt-form-control" placeholder="物业名称" value="' + street_property_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">物业电话<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="street_property_phone" type="text" class="form-control rt-form-control" placeholder="物业电话" value="' + street_property_phone + '"></div></div>';
    html += '<div class="form-group"><div class="col-sm-offset-3  col-sm-6"><input type="button" onclick="updatestreet()" class="btn btn-primary form-control rt-submit" value="提&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;交"></div></div>';
    $("#table").html(html);
}