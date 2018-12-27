var $;
var layer;
var userListTale;
layui.use(['jquery', 'table', 'layer'], function () {
    var table = layui.table;
    layer = layui.layer;
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
                area: ['600px', '400px'],
                content: ['/addUserTpl', 'no']
            });
        },
        deleteUserMuilti: function () {
            var checkedData = table.checkStatus('userListTable');
            var userIds = [];
            for (var i = 0, len = checkedData.data.length; i < len; i++) {
                userIds.push(checkedData.data[i].Id);
            }
            deleteUser({UserIds: userIds});
        },
        disableUserMulti: function () {
            var checkedData = table.checkStatus('userListTable');
            var userIds = [];
            for (var i = 0, len = checkedData.data.length; i < len; i++) {
                userIds.push(checkedData.data[i].Id);
            }
            updateUserStatusMulti({UserIds: userIds, TargetStatus: 2});
        },
        enableUserMulti: function () {
            var checkedData = table.checkStatus('userListTable');
            var userIds = [];
            for (var i = 0, len = checkedData.data.length; i < len; i++) {
                userIds.push(checkedData.data[i].Id);
            }
            updateUserStatusMulti({UserIds: userIds, TargetStatus: 1});
        }
    };

    $('.layui-btn').on('click', function () {
        var othis = $(this), method = othis.data('method');
        if (method) {
            active[method] ? active[method].call(this, othis) : '';
        }
    });

    // 行工具按钮事件
    table.on('tool(userListTable)', function (obj) {
        var data = obj.data;
        var layEvent = obj.event;
        if (layEvent === 'deleteUserOne') {
            // 删除用户
            let userId = [];
            userId.push(data.Id);
            var user = {UserIds: userId};
            deleteUser(user);
        } else if (layEvent == 'updateUserStatus') {
            // 修改用户状态
            var userIds = [];
            userIds.push(data.Id);
            var targetStatus = 0;
            if (data.Status == 1) {
                targetStatus = 2;
            } else {
                targetStatus = 1;
            }
            newUser = {UserIds: userIds, TargetStatus: targetStatus}
            updateUserStatusMulti(newUser)
        } else if (layEvent == "ediUser") {
            layer.open({
                title: '修改用户',
                type: 2,
                area: ['600px', '400px'],
                content: ['/editUserTpl/'+data.Id, 'no']
            });
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
                fixed: 'right', width: 260, align: 'center',
                toolbar: '#userToolBar'
            }
        ]]
    });
}

/**
 * 删除用户
 * @param user
 */
function deleteUser(user) {
    if (!user || !user.UserIds || user.UserIds.length == 0) {
        layer.alert("请先选择用户！", {icon: 2});
        return;
    }
    layer.confirm('确认删除吗', {icon: 3, title: '提示'}, function (index) {
        $.post({
            url: '/deleteuser',
            data: JSON.stringify(user),
            dataType: 'json',
            success: function (res) {
                if (res.code == 0) {
                    layer.alert("操作成功", {icon: 6});
                    userListTale.reload({});
                } else {
                    layer.alert("操作失败", {icon: 5});
                }
            }
        });
        layer.close(index);
    });
}

/**
 * 禁用/启用
 * @param user
 */
function updateUserStatusMulti(user) {
    if (!user || !user.UserIds || user.UserIds.length == 0) {
        layer.alert("请先选择用户！", {icon: 2});
        return;
    }
    var message = '';
    if (user.TargetStatus == 1) {
        message = "确认启用用户吗？";
    } else {
        message = "确认禁用用户吗？";
    }
    layer.confirm(message, {icon: 3, title: '提示'}, function (index) {
        $.post({
            url: '/updateUserStatusMulti',
            data: JSON.stringify(user),
            dataType: 'json',
            success: function (res) {
                if (res.code == 0) {
                    layer.alert("操作成功", {icon: 6});
                    userListTale.reload({});
                } else {
                    layer.alert("操作失败", {icon: 5});
                }
            }
        });
        layer.close(index);
    });
}