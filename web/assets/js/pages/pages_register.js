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
                },
                'register-limit-name': {
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
                'register-power': { required: '请选择角色！' },
                'register-limit-name': { required: '请选择权限！' }
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
jQuery(function() {
    BasePagesRegister.init();
    $("#station-select").select2();
    $("#community-select").select2();
    $("#street-select").select2();
    $('#province').addClass("hidden");
    $('#city').addClass("hidden");
    $('#district').addClass("hidden");
    $('#station').addClass("hidden");
    $('#office').addClass("hidden");
    $('#community').addClass("hidden");
    $('#street').addClass("hidden");
    $("#register-limit-name").change(function() {
        var limit_name = $("#register-limit-name option:selected").val();
        switch (limit_name) {
            case "province_no":
                $('#province').removeClass("hidden");
                $('#city').addClass("hidden");
                $('#district').addClass("hidden");
                $('#station').addClass("hidden");
                $('#office').addClass("hidden");
                $('#community').addClass("hidden");
                $('#street').addClass("hidden");
                break;
            case "city_no":
                $('#province').removeClass("hidden");
                $('#city').removeClass("hidden");
                $('#district').addClass("hidden");
                $('#station').addClass("hidden");
                $('#office').addClass("hidden");
                $('#community').addClass("hidden");
                $('#street').addClass("hidden");
                break;
            case "district_no":
                $('#province').removeClass("hidden");
                $('#city').removeClass("hidden");
                $('#district').removeClass("hidden");
                $('#station').addClass("hidden");
                $('#office').addClass("hidden");
                $('#community').addClass("hidden");
                $('#street').addClass("hidden");
                break;
            case "station_no":
                $('#province').removeClass("hidden");
                $('#city').removeClass("hidden");
                $('#district').removeClass("hidden");
                $('#station').removeClass("hidden");
                $('#office').addClass("hidden");
                $('#community').addClass("hidden");
                $('#street').addClass("hidden");
                break;
            case "office_no":
                $('#province').removeClass("hidden");
                $('#city').removeClass("hidden");
                $('#district').removeClass("hidden");
                $('#station').removeClass("hidden");
                $('#office').removeClass("hidden");
                $('#community').addClass("hidden");
                $('#street').addClass("hidden");
                break;
            case "community_no":
                $('#province').removeClass("hidden");
                $('#city').removeClass("hidden");
                $('#district').removeClass("hidden");
                $('#station').removeClass("hidden");
                $('#office').removeClass("hidden");
                $('#community').removeClass("hidden");
                $('#street').addClass("hidden");
                break;
            case "street_no":
                $('#province').removeClass("hidden");
                $('#city').removeClass("hidden");
                $('#district').removeClass("hidden");
                $('#station').removeClass("hidden");
                $('#office').removeClass("hidden");
                $('#community').removeClass("hidden");
                $('#street').removeClass("hidden");
                break;
            default:
                $('#province').addClass("hidden");
                $('#city').addClass("hidden");
                $('#district').addClass("hidden");
                $('#station').addClass("hidden");
                $('#office').removeClass("hidden");
                $('#community').addClass("hidden");
                $('#street').addClass("hidden");
        }
    })
});

$.validator.setDefaults({
    submitHandler: function() {
        var username = $("#register-username").val();
        var password = md5($("#register-password").val());
        var nickname = $("#register-nickname").val();
        var power = $("#register-power option:selected").val();
        var limit_name = $("#register-limit-name option:selected").val();
        var limit_id;
        switch (limit_name) {
            case "all":
                limit_id = 0;
                break;
            case "province_no":
                limit_id = $('#province-select option:selected').data('code');
                break;
            case "city_no":
                limit_id = $('#city-select option:selected').data('code');
                break;
            case "district_no":
                limit_id = $('#district-select option:selected').data('code');
                break;
            case "station_no":
                limit_id = $('#station-select option:selected').data('code');
                break;
            case "office_no":
                limit_id = $('#office-select option:selected').data('code');
                break;
            case "community_no":
                limit_id = $('#community-select option:selected').data('code');
                break;
            case "street_no":
                limit_id = $('#street-select option:selected').data('code');
                break;
            default:
                alert("权限错误!");
        }
        if (limit_id.length == 0) {
            alert("请选择详细权限！");
            return false;
        }
        var postData = {
            method: "POST",
            url: "user/user/UserRegister",
            data: {
                username: username,
                password: password,
                nickname: nickname,
                power: power,
                limit_name: limit_name,
                limit_id: limit_id
            },
            success: function(data) {
                //已经是json对象无需解析
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    return false;
                }
                if (data.res.registerstatus === 1) {
                    alert(data.res.msg);

                }
                location.href = "user.html";
            },
            error: function() {
                alert("您需要先搭建服务器哦！");
                return false;
            }
        };
        $.ajax(postData);
        // failureAnimation();
        // return false
    }
});