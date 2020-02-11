$(document).ready(function() {
    $("#table").rt({
        configFile: "area",
        table: "community_station"
    });
    $(".rt-create").click(add)
});
var add = function() {
    location.href = "add_community.html";
}