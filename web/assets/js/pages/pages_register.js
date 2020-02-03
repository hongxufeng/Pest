/*
 *  Document   : base_pages_register.js
 *  Author     : pixelcave
 *  Description: Custom JS code used in Register Page
 */

var BasePagesRegister = function() {
    // Init Register Form Validation, for more examples you can check out https://github.com/jzaefferer/jquery-validation
    var initValidationRegister = function() {
        jQuery('.js-validation-register').validate({
            errorClass: 'help-block text-right animated fadeInDown',
            errorElement: 'div',
            errorPlacement: function(error, e) {
                jQuery(e).parents('.form-group .form-material').append(error);
            },
            highlight: function(e) {
                jQuery(e).closest('.form-group').removeClass('has-error').addClass('has-error');
                jQuery(e).closest('.help-block').remove();
            },
            success: function(e) {
                jQuery(e).closest('.form-group').removeClass('has-error');
                jQuery(e).closest('.help-block').remove();
            },
            rules: {
                'register-username': {
                    required: true,
                    minlength: 3
                },
                'register-nickname': {
                    required: true,
                },
                'register-password': {
                    required: true,
                    minlength: 6
                },
                'register-password2': {
                    required: true,
                    equalTo: '#register-password'
                },
                'register-power': {
                    required: true,
                }
            },
            messages: {
                'register-username': {
                    required: '请输入用户名',
                    minlength: '用户名必须包含3个字符！'
                },
                'register-nickname': { required: '请输入昵称！' },
                'register-password': {
                    required: '请输入密码！',
                    minlength: '密码长度至少为6！'
                },
                'register-password2': {
                    required: '请确认密码！',
                    minlength: '密码长度至少为6！',
                    equalTo: '请输入和上面相同的密码！'
                },
                'register-power': { required: '请选择权限！' }
            }
        });
    };

    return {
        init: function() {
            // Init Register Form Validation
            initValidationRegister();
        }
    };
}();

// Initialize when page loads
jQuery(function() { BasePagesRegister.init(); });

$.validator.setDefaults({
    submitHandler: function() {
        var username = $("#register-username").val();
        var password = md5($("#register-password").val());
        var nickname = $("#register-nickname").val();
        var power = $("#register-power option:selected").val();
        alert(power)
        var failureAnimation = function(msg) {
            if (msg != undefined) {
                $("#tip").html(msg);
            }
            var failureBlock = $(".alert-dismissable");
            $(".login-title").css("display", "none");
            failureBlock.css("display", "block");
            if (failureBlock.hasClass("shake")) {
                failureBlock.removeClass("shake").addClass("wobble");
            } else {
                failureBlock.removeClass("wobble").addClass("shake");
            }
        };
        var postData = {
            method: "POST",
            url: "user/user/UserRegister",
            data: {
                username: username,
                password: password,
                nickname: nickname,
                power: power
            },
            success: function(data) {
                //已经是json对象无需解析
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    return false;
                }
                if (data.res.loginstatus === 0) {
                    alert("用户添加成功！");
                    window.location.href = "user.html";
                } else {
                    failureAnimation(data.res.faildata.msg);
                    return false;
                }
            },
            error: function() {
                failureAnimation("您需要先搭建服务器哦！");
                return false;
            }
        };
        $.ajax(postData);
        // failureAnimation();
        // return false
    }
});