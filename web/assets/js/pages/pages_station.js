$(document).ready(function() {
    $("#table").rt({
        configFile: "area",
        table: "station_list"
    });
    $(".rt-create").click(add)
});
var add = function() {
    location.href = "add_station.html";
}