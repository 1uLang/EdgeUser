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
        if (status !== "failed") {
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
        this.$post(".details")
            .params({
                id: item.id,
            }).success(resp => {
            if (resp.code === 200) {
                let cert = resp.data.cert
                this.name = cert.name
                this.username = cert.username
                this.password = cert.password
            }
        })

        //赋值
        this.onChangeState(5)
    }
    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该授权凭证吗？", function () {
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
            })
            .refresh()
    }
    this.onUpdate = function () {

        let that = this
        teaweb.confirm("确定要修改该授权凭证信息吗？", function () {
            this.$post(".update")
                .params({
                    id: that.id,
                    name: that.name,
                    username: that.username,
                    password: that.password,
                })
                .refresh()
        })
    }

    this.onDeleteAuthAccount = function (id) {

    }
})