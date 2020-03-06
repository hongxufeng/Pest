var changestation = function() {
    var district_no = $('#district-select option:selected').data('code');
    var url = "base/report/GetTableJson?page=1&rows=100&district_no=" + district_no;
    $.ajaxSettings.async = false;
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
                $("#community-select").empty();
                $("#street-select").empty();
                var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                $("#office-select").append(html);
                $("#community-select").append(html);
                $("#street-select").append(html);
                if (data.res.count > 0 && $('#district-select').val()) {
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
    var url = "base/report/GetTableJson?page=1&rows=100&station_no=" + station_no;
    $.ajaxSettings.async = false;
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
                $("#community-select").empty();
                $("#street-select").empty();
                var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                $("#community-select").append(html);
                $("#street-select").append(html);
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
var changecommunity = function() {
    var office_no = $('#office-select option:selected').val();
    var url = "base/report/GetTableJson?page=1&rows=100&office_no=" + office_no;
    $.ajaxSettings.async = false;
    $.post(url, {
            table: "community_office",
            configFile: "area"
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                $("#community-select").empty();
                $("#street-select").empty();
                var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                $("#street-select").append(html);
                if (data.res.count > 0) {
                    if (data.res.table) {
                        var res = JSON.parse(data.res.table);
                    } else {
                        var res = JSON.parse(data.res.tables);
                    }
                    $.each(res, function(index, value) {
                        html += '<option data-code="' + res[index].uid + '" data-text="' + res[index].community_name + '" value="' + res[index].uid + '">' + res[index].community_name + '</option>';
                    });
                }
                $("#community-select").append(html);
            }
        }
    );
}
var changestreet = function() {
    var community_no = $('#community-select option:selected').val();
    var url = "base/report/GetTableJson?page=1&rows=100&community_no=" + community_no;
    $.ajaxSettings.async = false;
    $.post(url, {
            table: "area_list",
            configFile: "area"
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                $("#street-select").empty();
                var html = '<option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option>';
                if (data.res.count > 0) {
                    if (data.res.table) {
                        var res = JSON.parse(data.res.table);
                    } else {
                        var res = JSON.parse(data.res.tables);
                    }
                    $.each(res, function(index, value) {
                        html += '<option data-code="' + res[index].uid + '" data-text="' + res[index].street_name + '" value="' + res[index].uid + '">' + res[index].street_name + '</option>';
                    });
                }
                $("#street-select").append(html);
            }
        }
    );
}
var initarea = function() {
    $("#station-select").select2();
    $("#office-select").select2();
    $("#community-select").select2();
    $("#street-select").select2();
    var limit_name = Cookies.get('limit_name');
    var limit_id = Cookies.get('limit_id');
    if (limit_name == undefined || limit_id == undefined) {
        window.location.href = "login.html";
    } else {
        switch (limit_name) {
            case "all":
                $('#distpicker').distpicker();
                break;
            case "province_no":
                $('[name=province_name]').attr("disabled", true);
                $('#distpicker').distpicker({                                
                    province: limit_id,
                    city: '---- 所在市 ----',
                    district: '---- 所在区 ----'                            
                });
                break;
            case "city_no":
                $('[name=province_name]').attr("disabled", true);
                $('[name=city_name]').attr("disabled", true);
                var province_name = limit_id.substring(0, 3) + '000';                            
                $('#distpicker').distpicker({                                
                    province: province_name,
                    city: limit_id,
                    district: '---- 所在区 ----'
                });
                break;
            case "district_no":
                var city_name = limit_id.substring(0, 4) + '00';
                var province_name = limit_id.substring(0, 3) + '000';
                var  district_name  = limit_id;                            
                $('#distpicker').distpicker({                                
                    province: province_name,
                    city: city_name,
                    district: district_name                            
                });
                $('[name=province_name]').attr("disabled", true);
                $('[name=city_name]').attr("disabled", true);
                $('[name=district_name]').attr("disabled", true);
                break;
            case "station_no":
                var url = "base/report/GetTableJson?page=1&rows=100&uid=" + limit_id
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
                                if (data.res.table) {
                                    var res = JSON.parse(data.res.table);
                                } else {
                                    var res = JSON.parse(data.res.tables);
                                }
                                $('#distpicker').distpicker({
                                    province: res[0].province_name,
                                    city: res[0].city_name,
                                    district: res[0].district_name
                                });
                                $('[name=province_name]').attr("disabled", true);
                                $('[name=city_name]').attr("disabled", true);
                                $('[name=district_name]').attr("disabled", true).trigger("change");
                                $('#station-select').select2('val', res[0].uid);
                                $('#station-select').attr("disabled", true);
                            } else {
                                alert("用户权限错误,请尝试登录!");
                                window.location.href = "login.html";
                            }
                        }
                    }
                );
                break;
            case "office_no":
                var url = "base/report/GetTableJson?page=1&rows=100&uid=" + limit_id
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
                            if (data.res.count > 0) {
                                if (data.res.table) {
                                    var res = JSON.parse(data.res.table);
                                } else {
                                    var res = JSON.parse(data.res.tables);
                                }
                                $('#distpicker').distpicker({
                                    province: res[0].province_name,
                                    city: res[0].city_name,
                                    district: res[0].district_name
                                });
                                $('[name=province_name]').attr("disabled", true);
                                $('[name=city_name]').attr("disabled", true);
                                $('[name=district_name]').attr("disabled", true).trigger("change");
                                $('#station-select').select2('val', res[0].station_no);
                                $('#station-select').attr("disabled", true).trigger("change");
                                $('#office-select').select2('val', res[0].uid);
                                $('#office-select').attr("disabled", true);
                            } else {
                                alert("用户权限错误,请尝试登录!");
                                window.location.href = "login.html";
                            }
                        }
                    }
                );
                break;
            case "community_no":
                var url = "base/report/GetTableJson?page=1&rows=100&uid=" + limit_id
                $.post(url, {
                        table: "community_office",
                        configFile: "area"
                    },
                    function(data) {
                        if (data.status === "fail") {
                            alert(data.msg);
                            return false;
                        }
                        if (data.status === "ok") {
                            if (data.res.count > 0) {
                                if (data.res.table) {
                                    var res = JSON.parse(data.res.table);
                                } else {
                                    var res = JSON.parse(data.res.tables);
                                }
                                $('#distpicker').distpicker({
                                    province: res[0].province_name,
                                    city: res[0].city_name,
                                    district: res[0].district_name
                                });
                                $('[name=province_name]').attr("disabled", true);
                                $('[name=city_name]').attr("disabled", true);
                                $('[name=district_name]').attr("disabled", true).trigger("change");
                                $('#station-select').select2('val', res[0].station_no);
                                $('#station-select').attr("disabled", true).trigger("change");
                                $('#office-select').select2('val', res[0].office_no);
                                $('#office-select').attr("disabled", true).trigger("change");
                                $('#community-select').select2('val', res[0].uid);
                                $('#community-select').attr("disabled", true);
                            } else {
                                alert("用户权限错误,请尝试登录!");
                                window.location.href = "login.html";
                            }
                        }
                    }
                );
                break;
            case "street_no":
                var url = "base/report/GetTableJson?page=1&rows=100&uid=" + limit_id
                $.post(url, {
                        table: "area_list",
                        configFile: "area"
                    },
                    function(data) {
                        if (data.status === "fail") {
                            alert(data.msg);
                            return false;
                        }
                        if (data.status === "ok") {
                            if (data.res.count > 0) {
                                if (data.res.table) {
                                    var res = JSON.parse(data.res.table);
                                } else {
                                    var res = JSON.parse(data.res.tables);
                                }
                                $('#distpicker').distpicker({
                                    province: res[0].province_name,
                                    city: res[0].city_name,
                                    district: res[0].district_name
                                });
                                $('[name=province_name]').attr("disabled", true);
                                $('[name=city_name]').attr("disabled", true);
                                $('[name=district_name]').attr("disabled", true).trigger("change");
                                $('#station-select').select2('val', res[0].station_no);
                                $('#station-select').attr("disabled", true).trigger("change");
                                $('#office-select').select2('val', res[0].office_no);
                                $('#office-select').attr("disabled", true).trigger("change");
                                $('#community-select').select2('val', res[0].community_no);
                                $('#community-select').attr("disabled", true).trigger("change");
                                $('#street-select').select2('val', res[0].uid);
                                $('#street-select').attr("disabled", true);
                            } else {
                                alert("用户权限错误,请尝试登录!");
                                window.location.href = "login.html";
                            }
                        }
                    }
                );
                break;
            default:
                alert("用户权限错误,请尝试重新登录!");
                window.location.href = "login.html";
        }
    }
}