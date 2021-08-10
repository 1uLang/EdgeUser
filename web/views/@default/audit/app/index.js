Tea.context(function () {
    
    this.searchSQLName = ""
    this.searchSQLIp = ""

    this.bShowhAuth = false
    this.asset_name = ""

    this.allUsers = []
    this.authUsers = []
    this.selectNoAuthPeopleListData = []
    this.selectAuthPeopleListData = []


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

    this.getAppName = function (type) {
        switch (type) {
            case 1:
                return "iis"
            case 0:
                return "nginx"
            default:
                return "未知"
        }
    }

    this.onAddHost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
			height: "300px",
            width:"460px",
			callback: function () {
                this.onSearch()

            }
		})
    }

    this.onSearch = function () {
        this.$get(".").params({
            name: this.searchSQLName,
            ip: this.searchSQLIp,
            json: true,
        }).success(resp => {
            if (resp.code === 200) {
                if (resp.data.appList)
                    this.appList = resp.data.appList
                else
                    this.appList = []
                // this.level = resp.data.level
            }
        })
    }

    // this.onOpenAuth = function (id) {
    //     teaweb.popup(Tea.url(".auth"), {
	// 		height: "270px",
    //         width: "460px",
	// 		callback: function () {
    //             this.onSearch()
	// 		}
	// 	})
    // }

    this.onEdit = function (item) {
        teaweb.popup(Tea.url(".createPopup",{
            id:item.id,
            name:item.name,
            type:item.app_type,
            ip:item.ip,
            status:item.status,
            edit:true,
        }), {
			height: "300px",
            width:"460px",
			callback: function () {
                this.onSearch()
			}
		})
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该应用？", function () {
            this.$post(".delete").params({
                Opt: "delete",
                id: id,
            }).refresh()
        })
    }

    this.tableData = [
        {id:1,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:1},
        {id:2,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:0},
        {id:3,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:1},
        {id:4,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:0},
    ]


    this.onOpenAuth = function (item) {
        //req
        this.asset_name = item.name
        this.bShowhAuth = true
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
        console.log(this.authUsers)
        // return
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

    this.onRemoveTableItem=function (id,table) {
        if(id && table && table.length>0 ){
            for(var index=0;index<table.length;index++){
                if(table[index].id==id){
                    table.splice(index,1)
                }
            }
        }
        return table
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
        this.onRemoveTableItem(id,this.selectNoAuthPeopleListData)
    }

    this.onAddAuthPeople = function () {
        if (this.selectNoAuthPeopleListData.length > 0) {
            this.selectNoAuthPeopleListData.forEach(element => {
                this.authUsers.push(element)
                this.onRemoveTableItem(element.id,this.allUsers)
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
            this.onAddSelectAuth(item.id, item.name)
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
        this.onRemoveTableItem(id,this.selectAuthPeopleListData)
    }

    this.onRemoveAuthPeople = function () {
        if (this.selectAuthPeopleListData.length > 0) {
            this.selectAuthPeopleListData.forEach(element => {
                this.allUsers.push(element)
                this.onRemoveTableItem(element.id,this.authUsers)
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
})