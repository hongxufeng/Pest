$(document).ready(function() {
    $("#table").rt({
        configFile: "pest",
        table: "house_unit",
        complete: function() {
            $(".rt-edit").click(editUnit);
            $(".rt-delete").click(deleteUnit);
            $(".rt-personnel").click(viewUnitPersonnel);
        }
    });
});