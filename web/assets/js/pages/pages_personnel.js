$(document).ready(function() {
    $("#table").rt({
        configFile: "pest",
        table: "personnel_list",
        complete: function() {
            $(".rt-edit").click(editPersonnel);
            $(".rt-delete").click(deletePersonnel);
            $(".rt-report").click(viewReport);
        }
    });
});