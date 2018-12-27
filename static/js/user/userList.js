var $;
var userListTale
layui.use(['jquery', 'table', 'layer'], function () {
    // var userListTale = layui.table;
    var layer = layui.layer;
    $ = layui.$;

    // 初始化表格
    userListTale = initTable(layui.table);

    var active = {
        openAddUserModal: function () {
            var that = this;
            // 添加用户
            layer.open({
                title: '添加用户',
                type: 2,
                area:['600px', '400px'],
                content: ['addUserTpl', 'no']
            });
        }
    }

    $('.layui-btn').on('click', function(){
        var othis = $(this), method = othis.data('method');
        if (method) {
            active[method] ? active[method].call(this, othis) : '';
        }
    });

});

/**
 * 初始化表格
 * @param table
 */
function initTable(table) {
    return table.render({
        elem: '#userListTable',
        url: '/userlist',
        page: true,
        cellMinWidth: 80 //全局定义常规单元格的最小宽度，layui 2.2.1 新增
        , cols: [[
            {type: "checkbox", title: "序号", sort: false},
            {type: "numbers", title: "序号", sort: false},
            {field: 'CreateTime', width: 220, title: '创建时间'},
            {field: 'UpdateTime', width: 220, title: '更新时间'},
            {field: 'LoginName', title: '登录名'},
            {field: 'Password', title: '密码'}, //minWidth：局部定义当前单元格的最小宽度，layui 2.2.1 新增,
            {field: 'Name', title: '姓名'},
            {field: 'Mobile', title: '手机号'},
            {field: 'Email', width: 220, title: '邮箱'},
            {
                field: 'Status', title: '状态', templet: function (row) {
                    let userState = {"1": "正常", "2": "停用"};
                    return userState[row.Status]
                }
            },
            {
                fixed: 'right', width: 250, align: 'center',
                toolbar: '#userToolBar'
            }
        ]]
    });
}