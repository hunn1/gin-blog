$(document).ready(function () {
    $('.delete').each(function (i) {
        $(this).click(function () {
            var url = $(this).data('href');
            $.ajax({
                type: "GET",
                url: url,
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
        })
    })
});