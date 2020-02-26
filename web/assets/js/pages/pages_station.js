$(document).ready(function() {
    $("#table").rt({
        configFile: "area",
        table: "station_list",
        complete: function() {
            $(".rt-create").click(addOne);
            $(".rt-edit").click(editThis);
            $(".rt-delete").click(deleteThis);
        }
    });
});
var addOne = function() {
    location.href = "add_station.html";
}
var deleteThis = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("user/area/DeleteArea", {
        Cmd_Delete: "station",
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
var updatestation = function() {
    var uid = $('[name=uid]').val();
    var station_name = $('[name=station_name]').val();
    var station_head = $('[name=station_head]').val();
    var station_phone = $('[name=station_phone]').val();
    $.post("user/area/UpdateStation", {
        Uid: uid,
        Station_Name: station_name,
        Station_Head: station_head,
        Station_Phone: station_phone,
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
    var station_head = $(this).parent().siblings('[name=station_head]').data('value');
    var station_phone = $(this).parent().siblings('[name=station_phone]').data('value');
    var station_name = $(this).parent().siblings('[name=station_name]').data('value');
    var district_name = $(this).parent().siblings('[name=district_name]').data('value');
    var city_name = $(this).parent().siblings('[name=city_name]').data('value');
    var province_name = $(this).parent().siblings('[name=province_name]').data('value');
    var html = '<form class="form-horizontal" autocomplete="off">'
    html += '<div class="form-group"><label class="col-sm-3 control-label">UID<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="uid" type="text" class="form-control rt-form-control" placeholder="UID" value="' + uid + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在省<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="province_name" type="text" class="form-control rt-form-control" placeholder="所在省" value="' + province_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在市<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="city_name" type="text" class="form-control rt-form-control" placeholder="所在市" value="' + city_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">所在区<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input readonly name="district_name" type="text" class="form-control rt-form-control" placeholder="所在区" value="' + district_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">派出所名称<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="station_name" type="text" class="form-control rt-form-control" placeholder="派出所名称" value="' + station_name + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">派出所负责人<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="station_head" type="text" class="form-control rt-form-control" placeholder="派出所负责人" value="' + station_head + '"></div></div>';
    html += '<div class="form-group"><label class="col-sm-3 control-label">派出所电话<span class="rt-glyphicon-color">:</span></label><div class="col-sm-6"><input name="station_phone" type="text" class="form-control rt-form-control" placeholder="派出所电话" value="' + station_phone + '"></div></div>';
    html += '<div class="form-group"><div class="col-sm-offset-3  col-sm-6"><input type="button" onclick="updatestation()" class="btn btn-primary form-control rt-submit" value="提&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;交"></div></div>';
    $("#table").html(html);
}