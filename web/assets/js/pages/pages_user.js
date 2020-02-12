$(document).ready(function() {
    $("#table").rt({
        configFile: "user",
        table: "user"
    });
    $(".rt-delete").click(deleteUser);
    $(".rt-create").click(createUser);
});
var createUser = function() {
    location.href = "register.html";
}
var deleteUser = function() {
    var value = $(this).parent().siblings('[name=uid]').data('value');
    $.post("user/user/UserDelete", {
        uid: value
    }, function(data) {
        // var jsonObject = JSON.parse(data);
        if (data.status === "fail") {
            alert(data.msg);
            return false
        }
        if (data.status === "ok") {
            if (data.res.deletestatus === 1) {
                alert(data.res.msg);
            }
            location.href = location.href;
        }

    });
}