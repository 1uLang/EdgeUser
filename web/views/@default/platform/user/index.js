Tea.context(function () {

    this.page = ""
    this.pageState = 1

    this.userid = ""
    this.editUserName = "zhangxl1"
    this.editPassword = ""
    this.editPasswordConfirm = ""
    this.editFullName = ""
    this.editPhone = ""
    this.editEmail = ""
    this.editRemark = ""
    this.editEnabled = 0
    this.features = []

    this.onChangeShowState = function (state) {
        //权限列表
        if (state === 3) {
            this.$get(".features")
                .params({
                    userId: this.userid
                }).success(resp => {
                if (resp.code === 200) {
                    this.features = resp.data.features
                }
            })
        }

        if (this.pageState != state) {
            this.pageState = state
        }
    }
    this.createUser = function () {
        teaweb.popup(Tea.url(".create"), {
            height: "680px",
            width: "650px",
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload()
                })
            }
        })
    }

    this.onOpenDetail = function (item) {
        this.userid = item.id
        this.editUserName = item.username
        this.editPassword = ''
        this.editPasswordConfirm = ''
        this.editFullName = item.fullname
        this.editPhone = item.mobile
        this.editEmail = item.email
        this.editEnabled = item.isOn
        this.editRemark = item.remark
        console.log(item.remark)
        this.onChangeShowState(2)
    }

    this.onDelete = function (userId) {
        let that = this
        teaweb.confirm("确定要删除这个用户吗？", function () {
            that.$post(".delete")
                .params({
                    userId: userId
                }).success(resp => {
                if (resp.code === 200)
                    teaweb.success("删除成功")
            })
                .refresh()
        })
    }

    this.getStatus = function (status) {
        switch (status) {
            case 1:
                return "已启用"
            case 0:
                return "已停用"
            default:
                return "已停用"
        }
    }

    this.onListenEditCheckBox = function () {
        let enabled = document.getElementById("editCheckBox").checked
        if (enabled) {
            this.editEnabled = 1
        } else {
            this.editEnabled = 0
        }
    }

    this.onSaveEdit = function () {

        let that = this
        teaweb.confirm("确定要修改该用户吗？", function () {
            that.$post(".update")
                .params({
                    userId: that.userid,
                    username: that.editUserName,
                    pass1: that.editPassword,
                    pass2: that.editPasswordConfirm,
                    fullname: that.editFullName,
                    mobile: that.editPhone,
                    email: that.editEmail,
                    isOn: that.editEnabled,
                    remark: that.editRemark,
                }).success(resp => {
                if (resp.code === 200) {
                    teaweb.success("修改成功")
                }
            }).refresh()
        })

    }

    this.onSaveAuth = function () {
        let selectAuthData = []
        let authCheckBoxList = document.getElementsByName("authCheckBox")
        for (var index = 0; index < authCheckBoxList.length; index++) {
            if (authCheckBoxList[index].checked) {
                selectAuthData.push(authCheckBoxList[index].value)
            }
        }
        console.log(selectAuthData)
        this.$post(".features")
            .params({
                userId: this.userid,
                codes: JSON.stringify(selectAuthData)
            }).success(resp => {
            if (resp.code === 200)
                teaweb.success("修改成功")
        })
    }
})