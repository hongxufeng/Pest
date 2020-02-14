var DeleteHouse = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("user/pest/DeleteHouse", {
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
    $.post("user/pest/DeleteUnit", {
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
    $.post("user/pest/DeletePersonnel", {
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
        complete: function() {
            $(".rt-edit").click(editUnit);
            $(".rt-delete").click(deleteUnit);
            $(".rt-personnel").click(viewUnitPersonnel);
        }
    });
}
var viewHouseUnit = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    var query = "?house_id=" + value;
    $("#table").rt({
        configFile: "pest",
        table: "house_unit",
        query: query
    });
}
var viewUnitPersonnel = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    var query = "?unit_id=" + value;
    $("#table").rt({
        configFile: "pest",
        table: "unit_personnel",
        query: query,
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
        complete: function() {
            $(".rt-edit").click(editReport);
            $(".rt-delete").click(deleteReport);
        }
    });
}