$(document).ready(function() {
    $("#station-select").select2();
    $("#office-select").select2();
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
                $("#station-select").empty();
                $("#office-select").empty();
                var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                $("#office-select").append(html);
                if (data.res.count > 0) {
                    if (data.res.table) {
                        var res = JSON.parse(data.res.table);
                    } else {
                        var res = JSON.parse(data.res.tables);
                    }
                    $.each(res, function(index, value) {
                        html += '<option data-code="' + res[index].uid + '" data-text="' + res[index].station_name + '" value="' + res[index].uid + '">' + res[index].station_name + '</option>';
                    });
                }
                $("#station-select").append(html);
            }
        }
    );
}
var changeoffice = function() {
    var station_no = $('#station-select option:selected').val();
    var url = "base/report/GetTableJson?page=1&rows=100&station_no=" + station_no
    $.post(url, {
            table: "office_station",
            configFile: "area"
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                $("#office-select").empty();
                var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                if (data.res.count > 0) {
                    if (data.res.table) {
                        var res = JSON.parse(data.res.table);
                    } else {
                        var res = JSON.parse(data.res.tables);
                    }
                    $.each(res, function(index, value) {
                        html += '<option data-code="' + res[index].uid + '" data-text="' + res[index].office_name + '" value="' + res[index].uid + '">' + res[index].office_name + '</option>';
                    });
                }
                $("#office-select").append(html);
            }
        }
    );
}
var addcommunity = function() {
    var community_name = $('#community_name').val();
    var community_head = $('#community_head').val();
    var community_phone = $('#community_phone').val();
    var office_no = $('#office-select').val();
    //alert(community_name + station_no)
    $.post("user/area/AddCommunity", {
            Community_Name: community_name,
            Community_Head: community_head,
            Community_Phone: community_phone,
            Office_No: office_no
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                alert(data.res.msg);
                location.href = "community.html";
            }
        }
    );
};