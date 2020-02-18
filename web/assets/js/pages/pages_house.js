$(document).ready(function() {
    var limit_name = Cookies.get('limit_name');
    var limit_id = Cookies.get('limit_id');
    if (limit_name == undefined || limit_id == undefined) {
        window.location.href = "login.html";
    } else {
        var limit;
        switch (limit_name) {
            case "all":
                limit = "";
                break;
            case "province_no":
                limit = "?" + limit_name + "=" + limit_id;
                break;
            case "city_no":
                limit = "?" + limit_name + "=" + limit_id;
                break;
            case "district_no":
                limit = "?" + limit_name + "=" + limit_id;
                break;
            case "station_no":
                limit = "?" + limit_name + "=" + limit_id;
                break;
            case "community_no":
                limit = "?" + limit_name + "=" + limit_id;
                break;
            case "street_no":
                limit = "?" + limit_name + "=" + limit_id;
                break;
            default:
                alert("用户权限错误,请重新登录!");
                window.location.href = "login.html";
        }

        var search = '<div id="distpicker"><div><span class="rt-search-heading">所在省：</span><select class="rt-search-txt form-control" id="province-select" name="province_name" data-sign="%7e%7e"></select></div><div><span class="rt-search-heading">所在市：</span><select class="rt-search-txt form-control" id="city-select" name="city_name" data-sign="%7e%7e"></select></div><div><span class="rt-search-heading">所在县/区：</span><select class="rt-search-txt form-control" id="district-select" name="district_name" data-sign="%7e%7e" onchange="changestation()"></select></div></div><div><div><span class="rt-search-heading">所属派出所：</span><select class="rt-search-txt form-control" id="station-select" name="station_no" data-sign="%7e%7e" data-station="---- 请选择 ----" onchange="changecommunity()"><option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option></select></div><div><span class="rt-search-heading">所属社区：</span><select class="rt-search-txt form-control" id="community-select" name="community_no" data-sign="%7e%7e" data-community="---- 请选择 ----" onchange="changestreet()"><option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option></select></div><div><span class="rt-search-heading">所属小区：</span><select class="rt-search-txt form-control" id="street-select" name="street_no" data-sign="%7e%7e" data-street="---- 请选择 ----"><option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option></select></div></div>'
        $("#table").rt({
            configFile: "pest",
            table: "house_area",
            query: limit,
            searchBar: search,
            complete: function() {
                $(".rt-edit").click(editHouse);
                $(".rt-delete").click(deleteHouse);
                $(".rt-personnel").click(viewHousePersonnel);
                $(".rt-unit").click(viewHouseUnit);
                $("#station-select").select2();
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
                        case "community_no":
                            var url = "base/report/GetTableJson?page=1&rows=100&uid=" + limit_id
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
                                            $('#station-select').select2('val', res[0].station_no);
                                            $('#station-select').attr("disabled", true).trigger("change");
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
        });
    }
});