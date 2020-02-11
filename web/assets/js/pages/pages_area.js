$(document).ready(function() {
    $("#table").rt({
        configFile: "area",
        table: "area_list"
    });
    $(".rt-qrcode").click(viewQr);
    $(".rt-create").click(add)
});
var viewQr = function() {
    var key = $(this).data('args');
    var value = $(this).parent().siblings('[name=' + key + ']').data('value');
    location.href = value;
};
var add = function() {
    location.href = "add_street.html";
}