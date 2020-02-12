$(document).ready(function() {
    $("#table").rt({
        configFile: "area",
        table: "area_list"
    });
    $(".rt-qrcode").click(viewQr);
    $(".rt-create").click(addOne);
    $(".rt-edit").click(editThis);
    $(".rt-delete").click(deleteThis);
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
    var update_name = $('[name=street_name]').val();
    $.post("user/area/UpdateArea", {
        Cmd_Update: "street",
        Uid: uid,
        Update_Name: update_name
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
    var community_name = $(this).parent().siblings('[name=community_name]').data('value');
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
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在社区<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="community_name" type="text" class="form-control rt-form-control" placeholder="所在社区" value="' + community_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">街道/小区名称<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="street_name" type="text" class="form-control rt-form-control" placeholder="街道/小区名称" value="' + street_name + '"></div></div>';
    html += '<div class="form-group"><div class="col-sm-offset-3  col-sm-6"><input type="button" onclick="updatestreet()" class="btn btn-primary form-control rt-submit" value="提&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;交"></div></div>';
    $("#table").html(html);
}