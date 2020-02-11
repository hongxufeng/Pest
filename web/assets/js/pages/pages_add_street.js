$(document).ready(function() {
    $("#station-select").select2();
    $("#community-select").select2();
});
var changestation = function() {
    var district_no = $('#district-select option:selected').data('code');
    var url = "base/report/GetTableJson?page=1&rows=100&district_no=" + district_no
    $.post(url, {
            table: "station_list",
            configFile: "area"
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                if (data.res.count > 0) {
                    $("#station-select").empty();
                    var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                    if (data.res.table) {
                        var res = JSON.parse(data.res.table);
                    } else {
                        var res = JSON.parse(data.res.tables);
                    }
                    $.each(res, function(index, value) {
                        html += '<option data-code="' + res[index].uid + '" data-text="' + res[index].station_name + '" value="' + res[index].uid + '">' + res[index].station_name + '</option>';
                    });
                    $("#station-select").append(html);
                }
            }
        }
    );
}
var changecommunity = function() {
    var station_no = $('#station-select option:selected').val();
    var url = "base/report/GetTableJson?page=1&rows=100&station_no=" + station_no
    $.post(url, {
            table: "community_station",
            configFile: "area"
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                if (data.res.count > 0) {
                    $("#community-select").empty();
                    var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                    if (data.res.table) {
                        var res = JSON.parse(data.res.table);
                    } else {
                        var res = JSON.parse(data.res.tables);
                    }
                    $.each(res, function(index, value) {
                        html += '<option data-code="' + res[index].uid + '" data-text="' + res[index].community_name + '" value="' + res[index].uid + '">' + res[index].community_name + '</option>';
                    });
                    $("#community-select").append(html);
                }
            }
        }
    );
}
var addstreet = function() {
    var street_name = $('#street_name').val();
    var community_no = $('#community-select').val();
    //alert(street_name + community_no)
    $.post("user/area/AddStreet", {
            Street_Name: street_name,
            Community_No: community_no
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                alert(data.res.msg);
                location.href = "area.html";
            }
        }
    );
};