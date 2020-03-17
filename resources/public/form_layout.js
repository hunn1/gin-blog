/* ============================================================
 * Form Layouts
 * Form layout options available in Pages
 * For DEMO purposes only. Extract what you need.
 * ============================================================ */
(function ($) {

    'use strict';
    jQuery.extend(jQuery.validator.messages, {
        required: "必选字段",
        remote: "请修正该字段",
        email: "请输入正确格式的电子邮件",
        url: "请输入合法的网址",
        date: "请输入合法的日期",
        dateISO: "请输入合法的日期 (ISO).",
        number: "请输入合法的数字",
        digits: "只能输入整数",
        creditcard: "请输入合法的信用卡号",
        equalTo: "请再次输入相同的值",
        accept: "请输入拥有合法后缀名的字符串",
        maxlength: jQuery.validator.format("请输入一个 长度最多是 {0} 的字符串"),
        minlength: jQuery.validator.format("请输入一个 长度最少是 {0} 的字符串"),
        rangelength: jQuery.validator.format("请输入 一个长度介于 {0} 和 {1} 之间的字符串"),
        range: jQuery.validator.format("请输入一个介于 {0} 和 {1} 之间的值"),
        max: jQuery.validator.format("请输入一个最大为{0} 的值"),
        min: jQuery.validator.format("请输入一个最小为{0} 的值")
    });
    $(document).ready(function () {

        // Validation method for budget, profit, revenue fields
        $.validator.addMethod("usd", function (value, element) {
            return this.optional(element) || /^(\$?)(\d{1,3}(\,\d{3})*|(\d+))(\.\d{2})?$/.test(value);
        }, "Please specify a valid dollar amount");

        $('#start-date, #end-date').datepicker();

        $('#form-personal').validate();
        $("#form-project").validate();
        $(".form-ajax").validate();

        $('#form-personal').submit(function (e) {
            e.preventDefault()
        })
        $('#form-project').submit(function (e) {
            e.preventDefault()
        });
        
        $('.form-ajax').submit(function (e) {
            e.preventDefault();

            var url, data, method, options, enctype;
            var that = this;
            var form = $(this);
            if ($(this).hasClass('confirm')) {
                e.stopPropagation();
                if (!confirm($(this).attr('title') || '确认要执行该操作吗?')) {
                    return false;
                }
            }
            if (form.get(0).nodeName == 'A') {
                method = 'get';
                url = form.attr('href');
            } else {
                url = form.get(0).action;
                data = form.serialize();
                method = form.get(0).method.toLowerCase();
                enctype = form.get(0).enctype;
                if (enctype == 'multipart/form-data') {
                    options = {
                        processData: false,
                        contentType: false
                    };
                    data = new FormData(form.get(0));
                }
            }
            $.ajax($.extend({
                type: method,
                url: url,
                data: data,
                ifModified: false,
                traditional: false,
                timeout: 20000,
                cache: false,
                async: true,
                success: function success(json) {
                    if (json.code == 200) {
                        $('body').pgNotification({
                            title: '消息通知',
                            message: json.message,
                            style: 'bar',
                            timeout: 0,
                            type: 'success',
                            position: 'top',
                        }).show();
                        setTimeout(function () {
                            if (json.url) location.href = json.url; else location.reload();
                        }, 1500);
                    } else {
                        $('body').pgNotification({
                            title: '错误消息',
                            message: json.message,
                            style: 'bar',
                            timeout: 4000,
                            type: 'warning',
                            position: 'top',
                        }).show();

                        // window.bootbox.alert(json.message);
                    }
                },
                error: function error() {

                }
            }, options || {})).done(function (data) {
                console.log('done');
            }).fail(function (data) {
                var json = data.responseJSON;
                $.each(json.errors, function (i, item) {
                    $.each(item, function (k, val) {
                        $('body').pgNotification({
                            title: '错误消息' + json.message,
                            message: val,
                            style: 'bar',
                            timeout: 3000,
                            type: 'warning',
                            position: 'top',
                        }).show();
                    });
                });
            }).always(function (data) {
                console.log('always');
            });
        });
        $('button').each(function (key, item) {
            if ($(item).data('click') == 'clear_form') {
                $(item).on({
                    click: function () {
                        $(':input').not(':button,:submit,:reset,:hidden,:radio').val('');
                    }
                })
            }
        })

    });

})(window.jQuery);

function notice(message, val, type, position, timeout) {
    $('body').pgNotification({
        title: message,
        message: val || message,
        style: 'bar',
        timeout: timeout || 4000,
        type: type || 'warning',
        position: position || 'top',
    }).show();
}
