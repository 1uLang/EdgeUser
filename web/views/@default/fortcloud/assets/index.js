Tea.context(function () {
    
    this.curIndex = -1

    this.pageState = 1

    this.host = ""
    this.post = ""
    this.pubHost = ""
    this.maskStr = ""
    this.state = false
    this.protoData = [{inputValue:"",protoIndex:-1}]

    this.bShowhAuth = false
    this.authValue = ""

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
        switch (status) {
            case 1:
                return "已启用"
            case 0:
                return "已停用"
            default:
                return "已停用"
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
        if( this.pageState!=id) {
            this.pageState = id
        }
    }

    
    this.onOpenDetail = function (item) { 
        this.onChangeState(3)
     }

    
    this.onConnect = function (id) {  
        
    }
    this.onOpenAuth = function (id) { 
        //req
       this.bShowhAuth = true
    }
    this.onCloseAuth = function () { 
        this.bShowhAuth = false
    }
    this.onSaveAuth = function () { 
        //req 
        this.onCloseAuth()
    }
    this.onEdit = function (id) { 
        //赋值
        this.onChangeState(2)
    }
    this.onDelete = function (id) { 
        
     }

    this.setState = function () { 
        this.state = !this.state
        // document.getElementById('btn-switch-state').checked = !this.state
    }

    this.onAddProtoData = function () { 
        let curData = {inputValue:"",protoIndex:-1}
        this.protoData.push(curData)
    }

    this.onRemoveProtoData = function (index) { 
        if(this.protoData.length > 1){
            this.protoData.splice(index, 1);
        }
    }

    this.onSave = function () { 

    }

    this.onDeleteAuthAccount=function (id) { 

    }

    this.onResRefresh = function () { 

    }

    this.onResTest = function () {

    }
 
    this.tableData = [
        {id:1,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:1,value5:1,value6:"2021-06-16T09:29:23.140",bAuth:1},
        {id:2,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:2,value5:0,value6:"2021-06-16T09:29:23.140",bAuth:0},
        {id:3,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:0,value5:1,value6:"2021-06-16T09:29:23.140",bAuth:1},
        {id:4,value1:"智安安全审计系统服务器",value2:"47.108.234.195",value3:"8 Core 7.82 G 49.0 G",value4:1,value5:0,value6:"2021-06-16T09:29:23.140",bAuth:0},
    ]
    this.userData = [
        {id:1,value1:"luobing",value2:"罗兵",value3:"uobing@zhiannet.com"},
        {id:2,value1:"luobing",value2:"罗兵",value3:"uobing@zhiannet.com"},
        {id:3,value1:"luobing",value2:"罗兵",value3:"uobing@zhiannet.com"},
    ]
    this.resData = [
        {key:"ID:",value:"42f167c2-d91a-4f20-99b1-3d56dabd896a"},
        {key:"主机名:",value:"智安-安全审计系统服务器"},
        {key:"IP:",value:"182.150.0.104"},
        {key:"协议组:",value:"ssh/22"},
        {key:"公网IP:",value:"182.150.0.104"},
        {key:"管理账号:",value:"智安-安全审计服务器"},
        {key:"制造商:",value:"Red Hat"},
        {key:"型号:",value:"KVM"},
        {key:"CPU:",value:"Unknown"},
        {key:"内存:",value:"7.82 G"},
        {key:"硬盘:",value:'{"vda": "49.00 GB"}'},
        {key:"系统平台:",value:"Linux"},
        {key:"操作系统:",value:"x86_64"},
        {key:"激活:",value:"是"},
        {key:"序列号:",value:"24b64f7c-c262-4a76-965a-cc147db465d6"},
        {key:"资产编号:",value:"123646"},
        {key:"创建日期:",value:"2021/6/4 18:08:47"},
        {key:"创建者:",value:"Administrator"},
        {key:"备注:",value:"这是备注"},
    ]
})