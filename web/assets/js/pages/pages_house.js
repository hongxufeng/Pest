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
            case "office_no":
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

        var search = '<div id="distpicker"><div><span class="rt-search-heading">所在省：</span><select class="rt-search-txt form-control" id="province-select" name="province_name" data-sign="%7e%7e"></select></div><div><span class="rt-search-heading">所在市：</span><select class="rt-search-txt form-control" id="city-select" name="city_name" data-sign="%7e%7e"></select></div><div><span class="rt-search-heading">所在县/区：</span><select class="rt-search-txt form-control" id="district-select" name="district_name" data-sign="%7e%7e" onchange="changestation()"></select></div></div><div><div><span class="rt-search-heading">所属派出所：</span><select class="rt-search-txt form-control" id="station-select" name="station_no" data-sign="%7e%7e" data-station="---- 请选择 ----" onchange="changeoffice()"><option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option></select></div><div><span class="rt-search-heading">所属办事处：</span><select class="rt-search-txt form-control" id="office-select" name="office_no" data-sign="%7e%7e" data-office="---- 请选择 ----" onchange="changecommunity()"><option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option></select></div><div><span class="rt-search-heading">所属社区：</span><select class="rt-search-txt form-control" id="community-select" name="community_no" data-sign="%7e%7e" data-community="---- 请选择 ----" onchange="changestreet()"><option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option></select></div></div><div><span class="rt-search-heading">所属小区：</span><select class="rt-search-txt form-control" id="street-select" name="street_no" data-sign="%7e%7e" data-street="---- 请选择 ----"><option data-code data-text="---- 请选择 ----" value="">---- 请选择 ----</option></select></div>'
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
                initarea();
            }
        });
    }
});