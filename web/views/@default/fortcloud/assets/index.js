Tea.context(function () {

    this.curIndex = -1

    this.id = ""
    this.adminUser = ""
    this.platform = ""
    this.host = ""
    this.post = ""
    this.port = "22"
    this.pubHost = ""
    this.maskStr = ""
    this.httpType = "ssh"

    this.protoData = [{value: "", proto: "ssh"}]

    this.bShowhAuth = false
    this.authValue = ""

    this.accountType = 'custom'
    this.inputAuthUserName = ""
    this.inputAuthPassword = ""
    this.selectAuthCer = 1

    this.pageState = 1
    this.allUsers = []
    this.authUsers = []
    this.getLinkStatus = function (status) {
        switch (status) {
            case 1:
                return "可连接"
            case 0:
                return "不可连接"
            default:
                return "未知"
        }
    }
    this.getAuthCount = function (tags) {
        let tmp = tags.split(",")
        return tmp.length - 1
    }
    this.checkAuth = function (item) {

    }
    this.getStatus = function (status) {
        if (status) {
            return "运行中"
        } else {
            return "不可用"
        }
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            resultTime = time.substring(0, time.indexOf(" +"));
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

        if (id === 2) {
            //初始化变量
            this.host = ""
            this.post = ""
            this.platform = ""
            this.pubHost = ""
            this.protoData = [{value: "", proto: "ssh"}]
            this.adminUser = ""
            this.state = false
            this.maskStr = ""
        } else if (id === 4) {
            window.location = "/fortcloud/assets?pageState=4&asset=" + this.id
        }

        if (this.pageState != id) {
            this.pageState = id
        }
    }

    this.onOpenDetail = function (id) {
        this.onChangeState(3)
    }

    this.onConnect = function (id) {
        this.$post(".connect")
            .params({
                Id: id
            }).success(resp => {
            if (resp.code === 200) {
                window.open(resp.data.url)
            }
        })
    }
    this.onOpenAuth = function (item) {
        //req
        this.asset_name = item.name
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

    this.onCloseAuth = function () {
        this.bShowhAuth = false
        this.id = ""
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
        })
            .refresh()

    }
    this.onEdit = function (id) {

        this.id = id
        let that = this
        this.$post(".details")
            .params({
                id: id,
            }).success(resp => {
            if (resp.code === 200) {
                let asset = resp.data.asset
                that.host = asset.name
                that.post = asset.ip
                that.port = asset.port
                that.accountType = asset.accountType
                that.httpType = asset.protocol
                that.inputAuthPassword = asset.password
                that.inputAuthUserName = asset.username
                that.selectAuthCer = asset.credentialId
                that.maskStr = asset.description
            }
            this.onChangeState(5)
        })

        //赋值
    }
    this.onDelete = function (id) {

        teaweb.confirm("确定要删除该资产吗？", function () {
            this.$post(".delete")
                .params({
                    Id: id
                }).success(resp=>{
                if(resp.code===200)
                    teaweb.success("删除成功")
            })
                .refresh()
        })
    }

    this.setState = function () {
        this.state = !this.state
        // document.getElementById('btn-switch-state').checked = !this.state
    }

    this.onAddProtoData = function () {
        let curData = {value: "", proto: "ssh"}
        this.protoData.push(curData)
    }

    this.onRemoveProtoData = function (index) {
        if (this.protoData.length > 1) {
            this.protoData.splice(index, 1);
        }
    }

    this.onUpdate = function () {

        let that = this
        teaweb.confirm("确定要修改该资产信息吗？", function () {
            let protocols = []
            for (let item of that.protoData) {
                protocols.push(item.proto + "/" + item.value)
            }
            this.$post(".update")
                .params({
                    Id: that.id,
                    hostName: that.host,
                    ip: that.post,
                    type: that.accountType,
                    protocol: that.httpType,
                    password: that.inputAuthPassword,
                    username: that.inputAuthUserName,
                    certId: that.selectAuthCer,
                    description: that.maskStr,
                    port: that.port,
                }).success(resp => {
                if (resp.code === 200) {
                    teaweb.success("修改成功")
                }
            })
                .refresh()
        })
    }
    this.onSave = function () {
        this.$post(".")
            .params({
                hostName: this.host,
                ip: this.post,
                type: this.accountType,
                protocol: this.httpType,
                password: this.inputAuthPassword,
                username: this.inputAuthUserName,
                certId: this.selectAuthCer,
                description: this.maskStr,
                port: this.port,
            }).success(resp => {
            if (resp.code === 200) {
                teaweb.success("创建成功")
            }
        })
            .refresh()
    }

    this.onDeleteAuthAccount = function (id) {
        teaweb.confirm("确定要删除该资产授权吗？", function () {
            this.$post(".delAuthorize")
                .params({
                    Id: id
                })
                .refresh()
        })
    }

    this.onResRefresh = function () {
        this.$post(".refresh")
            .params({
                Id: this.id,
            }).success(resp => {
            if (resp.code === 200) {
                console.log(resp.data.url)
                window.open(resp.data.url)
            }
        })
    }

    this.onResTest = function () {

        this.$post(".checkLink")
            .params({
                Id: this.id,
            }).success(resp => {
            if (resp.code === 200) {
                console.log(resp.data.url)
                window.open(resp.data.url)
            }
        })
    }

    this.onSelectHttpType = function (value) {
        let radioList = document.getElementsByName("httpType")
        for (var index = 0; index < radioList.length; index++) {
            if (radioList[index].value == value) {
                radioList[index].checked = true
                this.onListenHttpTypeChange(value)
                break
            }
        }
    }

    this.onListenHttpTypeChange = function (value) {
        var portMap = {
            ssh: "22",
            rdp: "3389",
            vnc: "5900",
            telnet: "23"
        }
        if (portMap[value]) {
            this.port = portMap[value]
        } else {
            this.port = "22"
        }

    }

    this.selectNoAuthPeopleListData = []
    this.selectAuthPeopleListData = []

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
    this.accountTypeData = [
        {type: "custom", name: "密码"}, {type: "credential", name: "授权"},
    ]

})