var addstation = function() {
    var station_name = $('#station_name').val();
    var station_head = $('#station_head').val();
    var station_phone = $('#station_phone').val();
    var province_no = $('#province-select option:selected').data('code');
    var province_name = $('#province-select').val();
    var city_no = $('#city-select option:selected').data('code');
    var city_name = $('#city-select').val();
    var district_no = $('#district-select option:selected').data('code');
    var district_name = $('#district-select').val();
    //alert(province_no)
    $.post("user/area/AddStation", {
            Station_Name: station_name,
            Station_Head: station_head,
            Station_Phone: station_phone,
            Province_No: province_no,
            Province_Name: province_name,
            City_No: city_no,
            City_Name: city_name,
            District_No: district_no,
            District_Name: district_name
        },
        function(data) {
            if (data.status === "fail") {
                alert(data.msg);
                return false;
            }
            if (data.status === "ok") {
                alert(data.res.msg);
                location.href = "station.html";
            }
        }
    );
};