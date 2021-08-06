Tea.context(function () {

    this.curIndex = -1

    this.id = ""
    this.adminUser = ""
    this.platform = ""
    this.host = ""
    this.post = ""
    this.port = ""
    this.pubHost = ""
    this.maskStr = ""
    this.state = false
    this.protoData = [{value: "", proto: "ssh"}]

    this.bShowhAuth = false
    this.authValue = ""

    this.accountType = 1
    this.inputAuthUserName = ""
    this.inputAuthPassword = ""
    this.selectAuthCer=1

    this.pageState = 1

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
        }else if (id === 4){
            window.location = "/fortcloud/assets?pageState=4&asset="+this.id
        }

        if (this.pageState != id) {
            this.pageState = id
        }
    }

    this.onOpenDetail = function (id) {
        this.onChangeState(3)
        // this.id = id
        // this.$post(".details")
        //     .params({
        //         Id: id
        //     }).success(resp => {
        //     if (resp.code === 200) {
        //         let details = resp.data.details
        //         this.resData[0].value = details.id
        //         this.resData[1].value = details.hostname
        //         this.resData[2].value = details.ip
        //         this.resData[3].value = details.protocols
        //         this.resData[4].value = details.public_ip
        //         this.resData[5].value = details.admin_user_display
        //         this.resData[6].value = details.vendor
        //         this.resData[7].value = details.model
        //         this.resData[8].value = details.cpu_model
        //         this.resData[9].value = details.memory
        //         this.resData[10].value = details.disk_info
        //         this.resData[11].value = details.platform
        //         this.resData[12].value = details.os_arch
        //         this.resData[13].value = details.is_active ? '是' : '否'
        //         this.resData[14].value = details.sn
        //         this.resData[15].value = details.date_created
        //         this.resData[16].value = details.created_by
        //         this.resData[17].value = details.comment
        //         this.onChangeState(3)
        //     }
        // })

    }

    this.onConnect = function (id) {
        this.$post(".link")
            .params({
                Id: id
            }).success(resp => {
            if (resp.code === 200) {
                window.open(resp.data.url)
            }
        })
    }
    this.onOpenAuth = function (id) {
        //req
        this.id = id
        this.bShowhAuth = true
    }
    this.onCloseAuth = function () {
        this.bShowhAuth = false
        this.id = ""
    }
    this.onSaveAuth = function () {
        //req
        this.$post(".authorize")
            .params({
                Id: this.id,
                emails: this.authValue,
            }).success(resp =>{
                if(resp.code === 200){
                    this.authValue = ""
                    teaweb.success("授权成功")
                }
        })

        this.onCloseAuth()
    }
    this.onEdit = function (item) {

        this.id = item.id
        this.host = item.hostname
        this.post = item.ip
        this.platform = item.platform
        this.pubHost = item.public_ip
        this.maskStr = item.comment
        this.adminUser = item.admin_user
        this.state = item.is_active

        this.protoData = []
        let protoTemp = []
        for (let proto of item.protocols) {
            protoTemp = proto.split("/")
            this.protoData.push({value: protoTemp[1], proto: protoTemp[0]})
        }

        //赋值
        this.onChangeState(5)
    }
    this.onDelete = function (id) {

        teaweb.confirm("确定要删除该资产吗？", function () {
            this.$post(".delete")
                .params({
                    Id: id
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
                    platform: that.platform,
                    active: that.state,
                    publicIp: that.pubHost,
                    comment: that.maskStr,
                    adminUser: that.adminUser,
                    protocols: protocols,
                })
                .refresh()
        })
    }
    this.onSave = function () {
        let protocols = []
        for (let item of this.protoData) {
            protocols.push(item.proto + "/" + item.value)
        }
        this.$post(".")
            .params({
                hostName: this.host,
                ip: this.post,
                platform: this.platform,
                active: this.state,
                publicIp: this.pubHost,
                comment: this.maskStr,
                adminUser: this.adminUser,
                protocols: protocols,
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
        for(var index=0;index<radioList.length;index++){
            if(radioList[index].value==value){
                radioList[index].checked = true
                this.onListenHttpTypeChange(value)
                break
            }
        }
    }

    this.onListenHttpTypeChange=function (value) {
        var portMap={
            SSH:"22",
            RDP:"3389",
            VNC:"5900",
            TeInet:"23"
        }
        if(portMap[value]){
            this.port = portMap[value]
        }else{
            this.port = "22"
        }
        
    }

    this.onUpdateAuthPeople=function () {
        console.log("onUpdateAuthPeople")
    }

    this.noAuthPeopleListData = [
        {id:1,name:"未授权1"},
        {id:2,name:"未授权2"},
        {id:3,name:"未授权3"},
        {id:4,name:"未授权4"},
        {id:5,name:"未授权5"},
        {id:6,name:"未授权6"},
        {id:7,name:"未授权7"},
        {id:8,name:"未授权8"},
        {id:9,name:"未授权9"},
        {id:10,name:"未授权10"},
    ]
    this.authPeopleListData = [
        {id:51,name:"已授权51"},
        {id:52,name:"已授权52"},
        {id:53,name:"已授权53"},
        {id:54,name:"已授权54"},
        {id:55,name:"已授权55"},
        {id:56,name:"已授权56"},
        {id:57,name:"已授权57"},
        {id:58,name:"已授权58"},
        {id:59,name:"已授权59"},
        {id:60,name:"已授权60"},
    ]

    this.selectNoAuthPeopleListData = []
    this.selectAuthPeopleListData=[]

    this.onGetAuthPeopleItemInfo = function (id,table) {
        if(table && id && table.length>0 && id>0){
            for(var index=0;index<table.length;index++){
                if(table[index].id==id){
                    return table[index]
                }
            }
        }
        return null
    }
    this.onCheckHadValue = function (id,table) {
        if(table && id && table.length>0 && id>0){
            for(var index=0;index<table.length;index++){
                if(table[index].id==id){
                    return true
                }
            }
        }
        return false
    }
    //添加操作
    this.onCheckSelectAllNoAuth = function () {
        var tempElement = document.getElementById("noAuth-allSelect")
        for(var index=0;index<this.noAuthPeopleListData.length;index++){
            if(!this.onCheckHadValue(this.noAuthPeopleListData[index].id,this.selectNoAuthPeopleListData)){
                tempElement.checked = false
                return
            }
        }
        tempElement.checked = true
    }

    this.selectAllNoAuth = function () {
        var tempElement = document.getElementById("noAuth-allSelect")
        if(tempElement.checked){
            let noAuthList = document.getElementsByName("noAuthSelect")
            for(var index=0;index<noAuthList.length;index++){
                if(!noAuthList[index].checked){
                    noAuthList[index].checked = true
                    this.onAddSelectNoAuth(noAuthList[index].value,noAuthList[index].data)
                }
            }
        }else{
            let noAuthList = document.getElementsByName("noAuthSelect")
            for(var index=0;index<noAuthList.length;index++){
                if(noAuthList[index].checked){
                    noAuthList[index].checked = false
                    this.onRemoveSelectNoAuth(noAuthList[index].value)
                }
            }
        }
        
    }
    this.onListenClickNoAuthChange = function (id,name) {
        console.log(id)
        let noAuthList = document.getElementsByName("noAuthSelect")
        for(var index=0;index<noAuthList.length;index++){
            console.log(noAuthList[index].value)
            if(noAuthList[index].value==id){
                if(noAuthList[index].checked){
                    noAuthList[index].checked = false
                    this.onRemoveSelectNoAuth(id)
                }else{
                    noAuthList[index].checked = true
                    this.onAddSelectNoAuth(id,name)
                }
                break
            }
        }
        this.onCheckSelectAllNoAuth()
    }
    this.onListenSelectNoAuthChange = function (id,name) {
        var hadSelect = this.onCheckHadValue(id,this.selectNoAuthPeopleListData)
        if(hadSelect){
            this.onRemoveSelectNoAuth(id)
        }else{
            this.onAddSelectNoAuth(id,name)
        }
        this.onCheckSelectAllNoAuth()
    }
    this.onAddSelectNoAuth = function (id,name) {
        if(id && name){
            var tempData = {id:id,name:name}
            this.selectNoAuthPeopleListData.push(tempData)
        }
    }
    this.onRemoveSelectNoAuth = function (id) {
        this.selectNoAuthPeopleListData.splice(this.selectNoAuthPeopleListData.findIndex(i => i.id === id), 1);
    }

    this.onAddAuthPeople =function () {
        if(this.selectNoAuthPeopleListData.length>0){
            this.selectNoAuthPeopleListData.forEach(element => {
                this.authPeopleListData.push(element)
                this.noAuthPeopleListData.splice(this.noAuthPeopleListData.findIndex(i => i.id === element.id), 1);
            });
            this.selectNoAuthPeopleListData = []
        }
        let noAuthList = document.getElementsByName("noAuthSelect")
        for(var index=0;index<noAuthList.length;index++){
            noAuthList[index].checked = false
        }
    }


    //移除操作
    this.onCheckSelectAllAuth = function () {
        var tempElement = document.getElementById("auth-allSelect")
        for(var index=0;index<this.authPeopleListData.length;index++){
            if(!this.onCheckHadValue(this.authPeopleListData[index].id,this.selectAuthPeopleListData)){
                tempElement.checked = false
                return
            }
        }
        tempElement.checked = true
    }
    this.selectAllAuth = function () {
        let authList = document.getElementById("auth-allSelect")
        for(var index=0;index<authList.length;index++){
            if(!authList[index].checked){
                authList[index].checked = true
                this.onAddSelectAuth(authList[index].value,authList[index].data)
            }
        }
    }
    this.onListenClickAuthChange = function (id,name) {
        let authList = document.getElementsByName("authSelect")
        for(var index=0;index<authList.length;index++){
            if(authList[index].value==id){
                if(authList[index].checked){
                    authList[index].checked = false
                    this.onRemoveSelectAuth(id)
                }else{
                    authList[index].checked = true
                    this.onAddSelectAuth(id,name)
                }
                break
            }
        }
        this.onCheckSelectAllAuth()
    }
    this.onListenSelectAuthChange = function (id,name) {
        var hadSelect = this.onCheckHadValue(id,this.selectAuthPeopleListData)
        if(hadSelect){
            this.onRemoveSelectAuth(id)
        }else{
            this.onAddSelectAuth(id,name)
        }
        this.onCheckSelectAllAuth()
    }

    this.onAddSelectAuth = function (id,name) {
        if(id && name){
            var tempData = {id:id,name:name}
            this.selectAuthPeopleListData.push(tempData)
        }
    }
    this.onRemoveSelectAuth = function (id) {
        this.selectAuthPeopleListData.splice(this.selectAuthPeopleListData.findIndex(i => i.id === id), 1);
    }

    this.onRemoveAuthPeople =function () {
        if(this.selectAuthPeopleListData.length>0){
            this.selectAuthPeopleListData.forEach(element => {
                this.noAuthPeopleListData.push(element)
                this.authPeopleListData.splice(this.authPeopleListData.findIndex(i => i.id === element.id), 1);
            });
            this.selectAuthPeopleListData = []
        }
        let authList = document.getElementsByName("authSelect")
        for(var index=0;index<authList.length;index++){
            authList[index].checked = false
        }
    }

    this.resData = [
        {key: "ID:", value: "42f167c2-d91a-4f20-99b1-3d56dabd896a"},
        {key: "主机名:", value: "智安-安全审计系统服务器"},
        {key: "IP:", value: "182.150.0.104"},
        {key: "协议组:", value: "ssh/22"},
        {key: "公网IP:", value: "182.150.0.104"},
        {key: "管理账号:", value: "智安-安全审计服务器"},
        {key: "制造商:", value: "Red Hat"},
        {key: "型号:", value: "KVM"},
        {key: "CPU:", value: "Unknown"},
        {key: "内存:", value: "7.82 G"},
        {key: "硬盘:", value: '{"vda": "49.00 GB"}'},
        {key: "系统平台:", value: "Linux"},
        {key: "操作系统:", value: "x86_64"},
        {key: "激活:", value: "是"},
        {key: "序列号:", value: "24b64f7c-c262-4a76-965a-cc147db465d6"},
        {key: "资产编号:", value: "123646"},
        {key: "创建日期:", value: "2021/6/4 18:08:47"},
        {key: "创建者:", value: "Administrator"},
        {key: "备注:", value: "这是备注"},
    ]

    this.assets = [//authState 1 自己的资产 0 被授权的
        {id:1,hostname:"我的资产",hardware_info:"SSH",is_active:1,authCount:20,connectivity:{datetime:"2021-12-12 12:15:36 +"},authState:1},
        {id:2,hostname:"别人的资产",hardware_info:"RDP",is_active:0,authCount:25,connectivity:{datetime:"2021-10-12 12:15:36 +"},authState:0},
    ]
    
    this.accountTypeData=[
        {id:1,name:"密码"},{id:2,name:"授权"},
    ]

    this.authCerData=[
        {id:1,name:"创建的授权1"},
        {id:2,name:"创建的授权2"},
        {id:3,name:"给予的授权1"},
        {id:4,name:"给予的授权2"},
    ]
})