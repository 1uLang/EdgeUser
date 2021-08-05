Tea.context(function () {

    this.curIndex = -1

    this.pageState = 1

    this.id = ""
    this.name = ""
    this.username = ""
    this.password = ""
    this.maskStr = ""
    this.assetsList = []


    this.getLinkStatus = function (status) {
        if (status !=="failed") {
            return "可连接"
        } else {
            return "不可连接"
        }
    }
    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.onChangeState = function (id) {
        let that = this
        if (id === 4) {//资产列表
            this.assetsList = []
            this.$post(".assetsList")
                .params({
                    Id: that.id
                }).success(resp => {
                if (resp.code === 200) {
                    that.assetsList = resp.data.assetsList
                }
            })
        }
        if (this.pageState != id) {
            this.pageState = id
        }
    }

    this.onOpenDetail = function (item) {
        this.id = item.id
        this.accountData[0].value = item.id
        this.accountData[1].value = item.name
        this.accountData[2].value = item.username
        this.accountData[3].value = ""
        this.accountData[4].value = item.date_created
        this.accountData[5].value = item.created_by

        this.onChangeState(3)
    }

    this.onEdit = function (item) {
        this.id = item.id
        this.name = item.name
        this.username = item.username
        this.password = item.password
        this.maskStr = item.comment
        //赋值
        this.onChangeState(5)
    }
    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该管理用户吗？", function () {
            this.$post(".delete")
                .params({
                    Id: id
                })
                .refresh()
        })
    }

    this.onSave = function () {

        this.$post(".")
            .params({
                name: this.name,
                username: this.username,
                password: this.password,
                Comment: this.maskStr,
            })
            .refresh()
    }
    this.onUpdate = function () {

        let that = this
        teaweb.confirm("确定要修改该管理用户信息吗？", function () {
            this.$post(".update")
                .params({
                    id: that.id,
                    name: that.name,
                    username: that.username,
                    password: that.password,
                    Comment: that.maskStr,
                })
                .refresh()
        })
    }

    this.onDeleteAuthAccount = function (id) {

    }
    this.accountData = [
        {key: "ID:", value: "42f167c2-d91a-4f20-99b1-3d56dabd896a"},
        {key: "名称:", value: "智安-安全审计系统服务器"},
        {key: "用户名:", value: "root"},
        {key: "SSH指纹:", value: "ssh/22"},
        {key: "创建日期:", value: "2021/6/4 18:04:46"},
        {key: "创建者:", value: "Administrator"},
    ]
})