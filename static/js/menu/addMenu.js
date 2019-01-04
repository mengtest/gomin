var $;
layui.use(['jquery', 'layer', 'table'], function () {
    var layer = layui.layer;
    $ = layui.$;
    var layer = layui.layer;

    var active = {
        addMenu: function () {
            layer.confirm("确认提交吗？", {
                btn: ["确定", '取消'],
                yes: function () {
                    var menu = {
                        MenuName: $("#MenuName").val(),
                        MenuUrl: $("#MenuUrl").val(),
                        Level: $("#Level").val(),
                        ParentId: $("#ParentId").val(),
                        Order: $("#Order").val(),
                        Remark: $("#Remark").val()
                    };
                    $.ajax({
                        method: 'post',
                        url: '/addmenu',
                        dataType: 'json',
                        data: JSON.stringify(menu),
                        success: function (data) {
                            layer.msg('保存成功', {
                                icon: 1,
                                time: 1000
                            }, function () {
                                var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
                                parent.layer.close(index); //再执行关闭
                                parent.menuListTable.reload({});
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

