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
    this.otpIsOn = false
    this.otpParams = ""
    this.features = []
    this.selectList = []

    this.moveIndex = ""

    let that = this

    this.onChangeShowState = function (state) {
        //权限列表
        if (state === 3) {
            this.$get(".features")
                .params({
                    userId: this.userid
                }).success(resp => {
                if (resp.code === 200) {
                    this.features = resp.data.features
                    this.selectList = resp.data.selectList
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
        this.otpIsOn = item.otpIsOn
        this.otpParams = item.otpParams
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

    this.onListenEditCheckBoxOTP = function () {
        let enabled = document.getElementById("otpIsOn").checked
        if (enabled) {
            this.otpIsOn = true
        } else {
            this.otpIsOn = false
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
                    otpIsOn: that.otpIsOn,
                    remark: that.editRemark,
                }).success(resp => {
                if (resp.code === 200) {
                    teaweb.success("修改成功")
                }
            }).refresh()
        })

    }

    this.onSaveAuth = function () {
        this.$post(".features")
            .params({
                userId: this.userid,
                codes: JSON.stringify(this.selectList)
            }).success(resp => {
            if (resp.code === 200)
                teaweb.success("修改成功")
        })
    }

    this.onRemoveTableItem = function (code, table) {
        if (code && table && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index] == code) {
                    table.splice(index, 1)
                }
            }
        }
        return table
    }

    this.onGetTableItemInfo = function (code, table) {
        if (table && code && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].code == code) {
                    return table[index]
                }
            }
        }
        return null
    }

    this.onCheckHadValue = function (code, table) {
        if (table && code && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index] == code) {
                    return true
                }
            }
        }
        return false
    }

    this.getShowImageName = function (code) {

        let bAllSelect = that.onCheckHadValue(code, that.selectList)
        if (bAllSelect) {
            return "/images/select_select.png"
        } else {
            let bSelect = false
            let tempItem = that.onGetTableItemInfo(code, that.features)
            if (tempItem) {
                for (var index = 0; index < tempItem.children.length; index++) {
                    if (that.onCheckHadValue(tempItem.children[index].code, that.selectList)) {
                        bSelect = true
                        break
                    }
                }
            }
            if (bSelect) {
                return "/images/select_half_select.png"
            } else {
                return "/images/select_box.png"
            }
        }
    }


    this.getItemShowImageName = function (id) {
        if (that.onCheckHadValue) {
            if (that.onCheckHadValue(id, that.selectList)) {
                return "/images/select_select.png"
            } else {
                return "/images/select_box.png"
            }
        }
        return "/images/select_box.png"
    }

    this.onShowChildItem = function (code) {
        var tempItem = this.onGetTableItemInfo(code, this.features)
        if (tempItem) {
            tempItem.bShowChild = !tempItem.bShowChild
        }
    }

    this.onMouseEnter = function (code) {
        this.moveIndex = code
    }

    this.onMouseLeave = function () {
        this.moveIndex = ""
    }

    this.onSelectValue = function (code) {
        if (this.onCheckHadValue(code, this.selectList)) {
            this.onRemoveTableItem(code, this.selectList)
            let tempItem = this.onGetTableItemInfo(code, this.features)
            if (tempItem) {
                for (var index = 0; index < tempItem.children.length; index++) {
                    this.onRemoveTableItem(tempItem.children[index].code, this.selectList)
                }
            }
        } else {
            this.selectList.push(code)
            let tempItem = this.onGetTableItemInfo(code, this.features)
            if (tempItem) {
                for (var index = 0; index < tempItem.children.length; index++) {
                    if (!this.onCheckHadValue(tempItem.children[index].code, this.selectList)) {
                        this.selectList.push(tempItem.children[index].code)
                    }
                }
            }
        }
    }

    this.onSelectChildValue = function (code, parentId) {
        if (this.onCheckHadValue(code, this.selectList)) {
            this.onRemoveTableItem(code, this.selectList)
        } else {
            this.selectList.push(code)
        }

        let bFindAll = true
        let tempItem = this.onGetTableItemInfo(parentId, this.features)
        if (tempItem) {
            for (var index = 0; index < tempItem.children.length; index++) {
                if (!this.onCheckHadValue(tempItem.children[index].code, this.selectList)) {
                    console.log(tempItem.children[index].code)
                    bFindAll = false
                    break
                }
            }
        }
        if (bFindAll) {
            this.selectList.push(parentId)
        } else {
            this.onRemoveTableItem(parentId, this.selectList)
        }
    }

})