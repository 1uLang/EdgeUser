Tea.context(function () {

    this.curIndex = -1

    this.pageState = 1

    this.id = ""
    this.name = ""
    this.username = ""
    this.password = ""
    this.maskStr = ""
    this.assetsList = []
    this.bShowhAuth = false
    this.cert_name = ""
    this.allUsers = []
    this.authUsers = []
    this.selectNoAuthPeopleListData = []
    this.selectAuthPeopleListData = []
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

    this.mouseLeave = function () {
        this.curIndex = -1
    }

    this.mouseEnter = function (index) {
        this.curIndex = index
    }
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

    this.onRefreshAuth = function () {
        this.$get(".authorize")
            .params({
                Id: this.id,
            }).success(resp => {
            if (resp.code === 200) {
                this.allUsers = resp.data.allUsers
                this.authUsers = resp.data.authUsers
                this.onResetAuthView()
            }
        })
    }


    //添加操作
    this.onCheckSelectAllNoAuth = function () {
        var tempElement = document.getElementById("noAuth-allSelect")
        for (var index = 0; index < this.allUsers.length; index++) {
            if (!this.allUsers[index].my &&!this.onCheckHadValue(this.allUsers[index].id, this.selectNoAuthPeopleListData)) {
                tempElement.checked = false
                return
            }
        }
        tempElement.checked = true
    }

    this.selectAllNoAuth = function () {
        var tempElement = document.getElementById("noAuth-allSelect")
        let noAuthList = document.getElementsByName("noAuthSelect")
        if (tempElement.checked) {
            for (var index = 0; index < noAuthList.length; index++) {
                if (!noAuthList[index].checked && !noAuthList[index].disabled) {
                    noAuthList[index].checked = true
                    this.onAddSelectNoAuth(noAuthList[index].value, noAuthList[index].getAttribute("data"))
                }
            }
        } else {

            for (var index = 0; index < noAuthList.length; index++) {
                if (noAuthList[index].checked && !noAuthList[index].disabled) {
                    noAuthList[index].checked = false
                    this.onRemoveSelectNoAuth(noAuthList[index].value)
                }
            }
        }
        
        

    }
    this.onListenClickNoAuthChange = function (item) {
        if(item.my){
            return
        }
        let noAuthList = document.getElementsByName("noAuthSelect")
        for (var index = 0; index < noAuthList.length; index++) {
            if (noAuthList[index].value == item.id) {
                if (noAuthList[index].checked) {
                    noAuthList[index].checked = false
                    this.onRemoveSelectNoAuth(item.id)
                } else {
                    noAuthList[index].checked = true
                    this.onAddSelectNoAuth(item.id, item.name)
                }
                break
            }
        }
        this.onCheckSelectAllNoAuth()
    }
    this.onListenSelectNoAuthChange = function (item) {
        var hadSelect = this.onCheckHadValue(item.id, this.selectNoAuthPeopleListData)
        if (hadSelect) {
            this.onRemoveSelectNoAuth(item.id)
        } else {
            this.onAddSelectNoAuth(item.id, item.name)
        }
        this.onCheckSelectAllNoAuth()
    }
    this.onAddSelectNoAuth = function (id, name) {
        if (id && name) {
            var tempData = {id: id, name: name,my:false}
            this.selectNoAuthPeopleListData.push(tempData)
        }
    }
    this.onRemoveSelectNoAuth = function (id) {
        this.selectNoAuthPeopleListData.splice(this.selectNoAuthPeopleListData.findIndex(i => i.id === id), 1);
    }

    this.onAddAuthPeople = function () {
        if (this.selectNoAuthPeopleListData.length > 0) {
            this.selectNoAuthPeopleListData.forEach(element => {
                this.authUsers.push(element)
                this.allUsers.splice(this.allUsers.findIndex(i => i.id === element.id), 1);
            });
            this.selectNoAuthPeopleListData = []
        }
        var tempElement = document.getElementById("noAuth-allSelect")
        tempElement.checked = false
        let noAuthList = document.getElementsByName("noAuthSelect")
        for (var index = 0; index < noAuthList.length; index++) {
            if(!noAuthList[index].disabled){
                noAuthList[index].checked = false
            }

        }
    }

    this.onGetAuthPeopleItemInfo = function (id, table) {
        if (table && id && table.length > 0 && id > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].id == id) {
                    return table[index]
                }
            }
        }
        return null
    }
    this.onCheckHadValue = function (id, table) {
        if (table && id && table.length > 0 && id > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].id == id) {
                    return true
                }
            }
        }
        return false
    }


    //移除操作
    this.onCheckSelectAllAuth = function () {
        var tempElement = document.getElementById("auth-allSelect")
        for (var index = 0; index < this.authUsers.length; index++) {
            if (!this.onCheckHadValue(this.authUsers[index].id, this.selectAuthPeopleListData)) {
                tempElement.checked = false
                return
            }
        }
        tempElement.checked = true
    }
    this.selectAllAuth = function () {
        var tempElement = document.getElementById("auth-allSelect")
        let authList = document.getElementsByName("authSelect")
        if (tempElement.checked) {
            for (var index = 0; index < authList.length; index++) {
                if (!authList[index].checked) {
                    authList[index].checked = true
                    this.onAddSelectAuth(authList[index].value, authList[index].getAttribute("data"))
                }
            }
        } else {
            for (var index = 0; index < authList.length; index++) {
                if (authList[index].checked) {
                    authList[index].checked = false
                    this.onRemoveSelectAuth(authList[index].value)
                }
            }
        }
    }
    this.onListenClickAuthChange = function (item) {
        let authList = document.getElementsByName("authSelect")
        for (var index = 0; index < authList.length; index++) {
            if (authList[index].value == item.id) {
                if (authList[index].checked) {
                    authList[index].checked = false
                    this.onRemoveSelectAuth(item.id)
                } else {
                    authList[index].checked = true
                    this.onAddSelectAuth(item.id, item.name)
                }
                break
            }
        }
        this.onCheckSelectAllAuth()
    }
    this.onListenSelectAuthChange = function (item) {
        var hadSelect = this.onCheckHadValue(item.id, this.selectAuthPeopleListData)
        if (hadSelect) {
            this.onRemoveSelectAuth(item.id)
        } else {
            this.onAddSelectAuth(item.id, name)
        }
        this.onCheckSelectAllAuth()
    }

    this.onAddSelectAuth = function (id, name) {
        if (id && name) {
            var tempData = {id: id, name: name,my:false}
            this.selectAuthPeopleListData.push(tempData)
        }
    }
    this.onRemoveSelectAuth = function (id) {
        this.selectAuthPeopleListData.splice(this.selectAuthPeopleListData.findIndex(i => i.id === id), 1);
    }

    this.onRemoveAuthPeople = function () {
        if (this.selectAuthPeopleListData.length > 0) {
            this.selectAuthPeopleListData.forEach(element => {
                this.allUsers.push(element)
                this.authUsers.splice(this.authUsers.findIndex(i => i.id === element.id), 1);
            });

            this.selectAuthPeopleListData = []
        }
        var tempElement = document.getElementById("auth-allSelect")
        tempElement.checked = false
        let authList = document.getElementsByName("authSelect")
        for (var index = 0; index < authList.length; index++) {
            authList[index].checked = false
        }
    }


    this.onCloseAuth = function () {
        this.bShowhAuth = false
        this.id = ""
    }
    this.onResetAuthView = function (params) {
        var tempElement1 = document.getElementById("noAuth-allSelect")
        tempElement1.checked = false
        let noAuthList = document.getElementsByName("noAuthSelect")
        for (var index = 0; index < noAuthList.length; index++) {
            if (!noAuthList[index].disabled) {
                noAuthList[index].checked = false
            }
        }
        var tempElement2 = document.getElementById("auth-allSelect")
        tempElement2.checked = false
        let authList = document.getElementsByName("authSelect")
        for (var index = 0; index < authList.length; index++) {
            authList[index].checked = false
        }
        this.selectNoAuthPeopleListData = []
        this.selectAuthPeopleListData = []

    }
    this.onOpenAuth = function (item) {
        //req
        this.cert_name = item.name
        this.id = item.id

        this.$get(".authorize")
            .params({
                Id: this.id,
            }).success(resp => {
            if (resp.code === 200) {
                this.allUsers = resp.data.allUsers
                this.authUsers = resp.data.authUsers
                this.bShowhAuth = true
            }
        })
    }

    this.onGetIdList = function (array){
        var tempList = []
        if(array && array.length>0){
            for(var index=0;index<array.length;index++){
                if(!this.onCheckHadValue(array[index].id,tempList)){
                    tempList.push(array[index].id)
                }
            }
        }
        return tempList
    }
    this.onSaveAuth = function () {
        //req
        this.$post(".authorize")
            .params({
                Id: this.id,
                Users: this.onGetIdList(this.authUsers),
            }).success(resp => {
            if (resp.code === 200) {
                teaweb.success("授权成功")
            }
        }).refresh()

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
})