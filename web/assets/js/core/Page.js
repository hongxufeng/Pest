var BasePagesInit = function() {
    var initPage = function() {
        if (Cookies.get('avatar') == undefined) {
            alert("您未进行登录，并不能获取数据哦！");
            window.location.href = "login.html";
        } else {
            // alert(decodeURI(Cookies.get("avatar")))
            $("[alt=\"Avatar\"]").attr("src", decodeURI(Cookies.get("avatar")));
        }
    };

    return {
        init: function() {
            // Init
            initPage();
        }
    };
}();

// Initialize when page loaded
$(document).ready(function() {
    BasePagesInit.init();
});