$(document).ready(function() {
    $("#table").rt({
        configFile: "pest",
        table: "personnel_list"
    });
    $(".rt-edit").click(editPersonnel);
    $(".rt-delete").click(deletePersonnel);
    $(".rt-report").click(viewReport);
});