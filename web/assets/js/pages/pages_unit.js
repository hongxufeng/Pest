$(document).ready(function() {
    $("#table").rt({
        configFile: "pest",
        table: "house_unit"
    });
    $(".rt-edit").click(editUnit);
    $(".rt-delete").click(deleteUnit);
    $(".rt-personnel").click(viewUnitPersonnel);
});