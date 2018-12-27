var $;
layui.use(['jquery', 'layer', 'table', 'form'], function () {
    var layer = layui.layer;
    $ = layui.$;
    var layer = layui.layer;

    console.log('111')
    initUser(layui.form);

    var active = {
        addUser: function () {
            layer.confirm("确认提交吗？", {
                btn: ["确定", '取消'],
                yes: function () {
                    var user = {
                        LoginName: $("#LoginName").val(),
                        Password: $("#Password").val(),
                        Name: $("#Name").val(),
                        Mobile: $("#Mobile").val(),
                        Email: $("#Email").val()
                    };
                    $.ajax({
                        method: 'post',
                        url: '/adduser',
                        dataType: 'json',
                        data: JSON.stringify(user),
                        success: function (data) {
                            layer.msg('保存成功', {
                                icon: 1,
                                time: 1000
                            }, function () {
                                var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
                                parent.layer.close(index); //再执行关闭
                                parent.userListTale.reload({});
                            });
                        }
                    });
                }
            });
        },
        cancelAdd: function () {
            var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
            parent.layer.close(index); //再执行关闭
        }
    }

    $('.layui-btn').on('click', function () {
        var othis = $(this), method = othis.data('method');
        if (method) {
            active[method] ? active[method].call(this, othis) : '';
        }
    });

});

function initUser(form) {
    $.ajax({
        url: '/queryOneUser/'+$("#userId").val(),
        type:'get',
        dataType: 'json',
        success: function (res) {
            if (res.code == 0) {
                form.val('editUserTable', res.data);
            } else {
                layer.alert("数据加载失败", {icon: 5});
            }
        }
    })
}
