$(document).ready(function() {
    $("#station-select").select2();
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
                var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
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
var addcommunity = function() {
    var office_name = $('#office_name').val();
    var office_head = $('#office_head').val();
    var office_phone = $('#office_phone').val();
    var station_no = $('#station-select').val();
    //alert(community_name + station_no)
    $.post("user/area/AddOffice", {
            Office_Name: office_name,
            Office_Head: office_head,
            Office_Phone: office_phone,
            Station_No: station_no
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                alert(data.res.msg);
                location.href = "office.html";
            }
        }
    );
};