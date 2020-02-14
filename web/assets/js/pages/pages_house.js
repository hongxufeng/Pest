$(document).ready(function() {
    $("#table").rt({
        configFile: "pest",
        table: "house_area"
    });
    $(".rt-edit").click(editHouse);
    $(".rt-delete").click(deleteHouse);
    $(".rt-personnel").click(viewHousePersonnel);
    $(".rt-unit").click(viewHouseUnit);
});