Tea.context(function () {
    this.severity = 1

    this.nShowState = 1

    // this.id = 1             //操作的ID
    this.type = ""             //操作
    this.types = []            //操作
    this.interface = "wan"             //接口* ID
    this.interfaces = []             //接口选择下拉
    this.direction = "in"         //方向 *ID
    this.directions = []         //方向 下拉
    this.ipprotocol = ""    //TCP/IP版本*
    this.ipprotocols = []    //TCP/IP版本* 下拉
    this.protocol = ""       //协议*
    this.protocols = []      //协议* 下拉
    this.src = ""       //源 单机网络的id
    this.srcinput = ""       //源 单机网络的id
    this.srcs = []       //源 单机网络 下拉
    this.srcmask = "32"       //源 掩码
    this.dst = ""           //目标
    this.dstinput = ""           //目标
    this.dsts = []       //目标改界面需要下拉选择
    this.dstmask = 32       //目标 掩码
    this.descr = ""         //描述值

    this.srcbeginport = ""
    this.srcbeginportValue = ""
    this.srcbeginports = []
    this.srcbeginportinput = ""
    this.srcendport = ""
    this.srcendportValue = ""
    this.srcendports = []
    this.srcendportinput = ""
    this.dstbeginport = ""
    this.dstbeginportValue = ""
    this.dstbeginports = []
    this.dstbeginportinput = ""
    this.dstendport = ""
    this.dstendportValue = ""
    this.dstendports = []
    this.dstendportinput = ""
    // this.targerRangeStartId = 1 //修改中 目标范围的开始值
    // this.targerRangeEndId = 1 //修改中 目标范围的终止值

    this.$delay(function () {
        let curSelectNode = localStorage.getItem("nfwSelectNodeId");
        if(curSelectNode){
            this.selectNode = curSelectNode
        }
    })

    this.onChangeShowState = function (state) {
        if (this.nShowState != state) {
            this.nShowState = state
            if (state == 2) {
                // this.resetValue()
            }
        }
    }
    this.getInterfaceName = function (interface) {
        switch (interface) {
            case "wan":
                return "WAN";
            case "lan":
                return "LAN";
            case "lo0":
                return "Loopback";
        }
        return "";
    }
    this.getStatus = function (status) {
        return status == 1 ? "已开启" : "已停用";
    }

    this.getDirection = function (type) {
        return type == "in" ? "进" : "出";
    }


    this.getTactics = function (type) {
        switch (status) {
            case 1:
                return "通过"
            case 2:
                return "阻止"
            case 3:
                return "拒绝"
            default:
                return "拒绝"
        }
    }

    this.getItemInfo = function (id) {
        for (var i = 0; i < this.tableData.length; i++) {
            if (this.tableData[i].id == id) {
                return this.tableData[i]
            }
        }
        return null
    }

    //修改配置
    this.onOpenChangeView = function (id) {
        this.GetNatInfo(id)
    }

    //重置value
    this.resetValue = function () {
        this.id = 1
        this.interface = 1
        this.direction = 1
        this.ipprotocol = ""
        this.agreementId = 1
        this.sourceTypeId = 1
        this.sourceValue = "any"
        this.sourceId = 1
        this.targetTypeId = 1
        this.targetValue = "any"
        this.targetId = 1
        this.editSourceId = 1
        this.editTargetId = 1
        this.descValue = ""
        this.targerRangeStartId = 1
        this.targerRangeEndId = 1
    }

    //保存配置
    this.onSaveConfig = function () {
        console.log(this.srcbeginportValue)
        console.log(this.srcbeginport)
        console.log(this.srcbeginportinput)
        console.log(this.srcendport)
        console.log(this.srcendportinput)

        console.log(this.dstbeginport)
        console.log(this.dstbeginportinput)
        console.log(this.dstendport)
        console.log(this.dstendportinput)
        // return
        let srcbeginport = this.srcbeginportValue
        if(srcbeginport == ""){
            srcbeginport = this.srcbeginportinput
        }
        let srcendport = this.srcendportValue
        if(srcendport == ""){
            srcendport = this.srcendportinput
        }

        let dstbeginport = this.dstbeginportValue
        if(dstbeginport == ""){
            dstbeginport = this.dstbeginportinput
        }
        let dstendport = this.dstendportValue
        if(dstendport == ""){
            dstendport = this.dstendportinput
        }
        Tea.action("/nfw/acl/createPopup")
            .params({
                nodeId: this.selectNode,
                id: this.id,
                interface: this.interface,
                type: this.type,
                srcinput: this.srcinput,
                srcmask: this.srcmask,
                // dst: this.dst,
                dstinput: this.dstinput,
                dstmask: this.dstmask,
                descr: this.descr,
                direction: this.direction,
                ipprotocol: this.ipprotocol,
                protocol: this.protocol,
                srcbeginport: srcbeginport,
                srcendport: srcendport,
                dstbeginport: dstbeginport,
                dstendport: dstendport,

            })
            .post()
            .success(function (resp) {

            }).refresh()
    }

    //开启配置
    this.onOpenConfig = function (id, status, interface) {
        let tops = "禁用"
        let statusUp = "0"
        let that = this
        if (status == "0") {
            tops = "启用"
            statusUp = "1"
        }
        teaweb.confirm("确定要" + tops + "此项吗？", function () {
            that.$post(".set")
                .params({
                    id: id,
                    status: statusUp,
                    nodeId: this.selectNode,
                    interface: interface,
                })
                .refresh()
        })
    }

    //关闭配置
    this.onCloseConfig = function (id) {

    }
    //删除配置
    this.onDelConfig = function (id) {
        teaweb.confirm("确定要删除所选规则？", function () {
            let that = this
            that.$post(".delete")
                .params({
                    id: id,
                    nodeId: this.selectNode,
                })
                .refresh()
        })
    }


    //源和目标的位数
    this.masks = [
        {id: 1, value: 1}, {id: 2, value: 2}, {id: 3, value: 3}, {id: 4, value: 4},
        {id: 5, value: 5}, {id: 6, value: 6}, {id: 7, value: 7}, {id: 8, value: 8},
        {id: 9, value: 9}, {id: 10, value: 10}, {id: 11, value: 11}, {id: 12, value: 12},
        {id: 13, value: 13}, {id: 14, value: 14}, {id: 15, value: 15}, {id: 16, value: 16},
        {id: 17, value: 17}, {id: 18, value: 18}, {id: 19, value: 19}, {id: 20, value: 20},
        {id: 21, value: 21}, {id: 22, value: 22}, {id: 23, value: 23}, {id: 24, value: 24},
        {id: 25, value: 25}, {id: 26, value: 26}, {id: 27, value: 27}, {id: 28, value: 28},
        {id: 29, value: 29}, {id: 30, value: 30}, {id: 31, value: 31}, {id: 32, value: 32},
    ]

    //操作
    this.types = [
        {name: "通过", value: "通过"},
        {name: "阻止", value: "拒绝"},
        {name: "阻止", value: "阻止"}
    ]
    //方向
    this.dirctionData = [
        {name: "进", value: "in"},
        {name: "出", value: "out"},
    ]

    // this.hostData = [
    //     {id:1,hostAddress:"成都-ddos-192.168.1.1",},
    //     {id:2,hostAddress:"成都-ddos-192.168.1.2",},
    //     {id:3,hostAddress:"成都-ddos-192.168.1.3",},
    //     {id:4,hostAddress:"成都-ddos-192.168.1.4",},
    // ]


    //postIndex:接口* 选择的id  typeIndex 类型* 选择的id  sourceTypeIndex 源* 选择的id  targetIndex 目标* 选择的id
    // this.tableData = [
    //     {id:1,interface:"LAN",directionType:1,agreement:"IPv4 TCP",value1:"192.168.0.1",post1:"8877",value2:"192.168.1.1",post2:"3306",desc:"描述",tactics:1,status:1,postIndex:1,typeIndex:1,sourceTypeIndex:1,targetIndex:1,sourceValue:"源的值"},
    //     {id:2,interface:"WAN",directionType:2,agreement:"IPv4 TCP",value1:"192.168.0.1",post1:"8877",value2:"192.168.1.1",post2:"3306",desc:"描述",tactics:1,status:1,postIndex:1,typeIndex:1,sourceTypeIndex:1,targetIndex:1,sourceValue:"源的值"},
    //     {id:3,interface:"Loopback",directionType:1,agreement:"IPv4 TCP",value1:"192.168.0.1",post1:"8877",value2:"192.168.1.1",post2:"3306",desc:"描述",tactics:1,status:1,postIndex:1,typeIndex:1,sourceTypeIndex:1,targetIndex:1,sourceValue:"源的值"},
    // ]


    this.GetSelectNode = function (event) {
        this.selectNode = event.target.value; //获取option对应的value值
        localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/nfw/acl?nodeId=' + node

    }

    //通过ID 获取详细数据
    this.GetNatInfo = async function (id) {
        await Tea.action("/nfw/acl/detail")
            .params({
                nodeId: this.selectNode,
                id: id,
                "act": "get-info"
            })
            .get()
            .showDialog(true)
            .success(function (resp) {
                //操作
                if (resp.data.type.length > 0) {
                    this.types = resp.data.type
                    let typeList = resp.data.type
                    for (var i = 0; i < typeList.length; i++) {
                        if (typeList[i].selected == true) {
                            this.type = typeList[i].value
                        }

                    }
                    if (this.type == "") {
                        this.type = List[0].value
                    }

                }


                //接口
                if (resp.data.interface.length > 0) {
                    this.interface = ""
                    this.interfaces = resp.data.interface
                    for (var i = 0; i < resp.data.interface.length; i++) { //接口下拉
                        if (resp.data.interface[i].selected == true) {
                            this.interface = resp.data.interface[i].value
                        }
                    }
                    if (this.interface == "") {
                        this.interface = resp.data.interface[0].value
                    }
                }

                //方向
                if (resp.data.direction.length > 0) {
                    this.direction = ""
                    this.directions = resp.data.direction
                    for (var i = 0; i < resp.data.direction.length; i++) { //下拉
                        if (resp.data.direction[i].selected == true) {
                            this.direction = resp.data.direction[i].value
                        }
                    }
                    if (this.direction == "") {
                        this.direction = resp.data.direction[0].value
                    }
                }

                //tcp ip 版本
                if (resp.data.ipprotocol.length > 0) {
                    this.ipprotocol = ""
                    this.ipprotocols = resp.data.ipprotocol
                    // console.log(this.ipprotocols);
                    for (var i = 0; i < resp.data.ipprotocol.length; i++) { //下拉
                        if (resp.data.ipprotocol[i].selected == true) {
                            this.ipprotocol = resp.data.ipprotocol[i].value
                        }
                    }
                    if (this.ipprotocol == "") {
                        this.ipprotocol = resp.data.ipprotocol[0].value
                    }
                }

                //协议
                if (resp.data.protocol.length > 0) {
                    this.protocol = ""
                    this.protocols = resp.data.protocol
                    // console.log(this.protocols);
                    for (var i = 0; i < resp.data.protocol.length; i++) { //下拉
                        if (resp.data.protocol[i].selected == true) {
                            this.protocol = resp.data.protocol[i].value
                        }
                    }
                    if (this.protocol =="") {
                        this.protocol = resp.data.protocol[0].value
                    }
                }

                //源
                if (resp.data.src.length > 0) {
                    this.src = ""
                    this.srcinput = ""
                    this.srcs = resp.data.src

                    for (var i = 0; i < resp.data.src.length; i++) { //目标下拉
                        if (resp.data.src[i].selected == true) {
                            this.src = resp.data.src[i].value
                            this.srcinput = resp.data.src[i].value
                            if (resp.data.src[i].data_other == true && !this.checkDstSrcType(this.srcinput)) {
                                //单个主机或网络
                                resp.data.src[i].value = ""
                            }
                        }
                    }
                    if (this.src == "") {
                        this.src = resp.data.src[0].value
                        this.srcinput = resp.data.src[0].value
                    }

                }
                //源掩码
                this.srcmask = "32"
                if (resp.data.srcmask.length > 0) {
                    let srcmask = resp.data.srcmask
                    for (var i = 0; i < srcmask.length; i++) {
                        if (srcmask[i].selected == true) {
                            this.srcmask = srcmask[i].value
                        }
                    }

                }
                //源 开始端口
                if (resp.data.srcbeginport.length > 0) {
                    this.srcbeginport = ""
                    this.srcbeginportinput = ""
                    this.srcbeginports = resp.data.srcbeginport
                    let otherValue = ""
                    for (var i = 0; i < resp.data.srcbeginport.length; i++) { //目标下拉
                        if(resp.data.srcbeginport[i].data_other){
                            otherValue = resp.data.srcbeginport[i].value
                        }
                        if (resp.data.srcbeginport[i].selected == true) {
                            this.srcbeginport = resp.data.srcbeginport[i].value
                            this.srcbeginportinput = resp.data.srcbeginport[i].value

                        }
                        //下拉选择类型
                        if (resp.data.srcbeginport[i].data_other == true && this.checkPortSelect(resp.data.srcbeginport[i].value,resp.data.srcbeginport)) {
                            resp.data.srcbeginport[i].value = ""
                        }
                    }
                    //添加时  赋值默认数据
                    if(this.srcbeginport == "" && this.srcbeginportinput == ""){
                        if(otherValue == ""){
                            this.srcbeginportinput = "any"
                            this.srcbeginport = "any"
                        }else{
                            this.srcbeginportinput = otherValue
                            this.srcbeginport = otherValue
                        }

                    }
                    this.srcbeginports = resp.data.srcbeginport
                    this.srcbeginportValue = this.srcbeginportinput
                    console.log(this.srcbeginport)
                    console.log(this.srcbeginportinput)
                    console.log(this.srcbeginports)
                }

                //源 结束端口
                if (resp.data.srcendport.length > 0) {
                    this.srcendport = ""
                    this.srcendportinput = ""
                    this.srcendports = resp.data.srcendport
                    let otherValue = ""
                    for (var i = 0; i < resp.data.srcendport.length; i++) { //目标下拉
                        if(resp.data.srcendport[i].data_other){
                            otherValue = resp.data.srcendport[i].value
                        }
                        if (resp.data.srcendport[i].selected == true) {
                            this.srcendport = resp.data.srcendport[i].value
                            this.srcendportinput = resp.data.srcendport[i].value

                        }
                        //下拉选择类型
                        if (resp.data.srcendport[i].data_other == true && this.checkPortSelect(resp.data.srcendport[i].value,resp.data.srcendport)) {
                            resp.data.srcendport[i].value = ""
                        }
                    }
                    //添加时  赋值默认数据
                    if(this.srcendport == "" && this.srcendportinput == ""){
                        if(otherValue == ""){
                            this.srcendportinput = "any"
                            this.srcendport = "any"
                        }else{
                            this.srcendportinput = otherValue
                            this.srcendport = otherValue
                        }

                    }
                    this.srcendports = resp.data.srcendport
                    this.srcendportValue = this.srcendportinput


                }
                // 目标
                if (resp.data.dst.length > 0) {
                    this.dst = ""
                    this.dsts = resp.data.dst

                    for (var i = 0; i < resp.data.dst.length; i++) { //目标下拉
                        if (resp.data.dst[i].selected == true) {
                            this.dst = resp.data.dst[i].value
                            this.dstinput = resp.data.dst[i].value
                            if (resp.data.dst[i].data_other == true && !this.checkDstSrcType(this.dstinput)) {
                                //单个主机或网络
                                resp.data.dst[i].value = ""
                            }
                        }
                    }
                    if (this.dst == "") {
                        this.dst = resp.data.dst[0].value
                        this.dstinput = resp.data.dst[0].value
                    }

                }


                //目标掩码
                this.dstmask = "32"
                if (resp.data.dstmask.length > 0) {
                    let dstmask = resp.data.dstmask
                    for (var i = 0; i < dstmask.length; i++) {
                        if (dstmask[i].selected == true) {
                            this.dstmask = dstmask[i].value
                        }
                    }
                    if (this.dstmask == "") {
                        this.dstmask = dstmask[0].value

                    }
                }

                //目标 开始端口
                if (resp.data.dstbeginport.length > 0) {
                    this.dstbeginport = ""
                    this.dstbeginportinput = ""
                    this.dstbeginports = resp.data.dstbeginport
                    let otherValue = ""
                    for (var i = 0; i < resp.data.dstbeginport.length; i++) { //目标下拉
                        if(resp.data.dstbeginport[i].data_other){
                            otherValue = resp.data.dstbeginport[i].value
                        }
                        if (resp.data.dstbeginport[i].selected == true) {
                            this.dstbeginport = resp.data.dstbeginport[i].value
                            this.dstbeginportinput = resp.data.dstbeginport[i].value

                        }
                        //下拉类型
                        if (resp.data.dstbeginport[i].data_other == true && this.checkPortSelect(resp.data.dstbeginport[i].value,resp.data.dstbeginport)) {
                            resp.data.dstbeginport[i].value = ""
                        }
                    }
                    //添加时  赋值默认数据
                    if(this.dstbeginport == "" && this.dstbeginportinput == ""){
                        if(otherValue == ""){
                            this.dstbeginportinput = "any"
                            this.dstbeginport = "any"
                        }else{
                            this.dstbeginportinput = otherValue
                            this.dstbeginport = otherValue
                        }

                    }
                    this.dstbeginports = resp.data.dstbeginport
                    this.dstbeginportValue = this.dstbeginportinput


                }


                //目标 结束端口
                if (resp.data.dstendport.length > 0) {
                    this.dstendport = ""
                    this.dstendportinput = ""
                    this.dstendports = resp.data.dstendport
                    let otherValue = ""
                    for (var i = 0; i < resp.data.dstendport.length; i++) { //目标下拉
                        if(resp.data.dstendport[i].data_other){
                            otherValue = resp.data.dstendport[i].value
                        }
                        if (resp.data.dstendport[i].selected == true) {
                            this.dstendport = resp.data.dstendport[i].value
                            this.dstendportinput = resp.data.dstendport[i].value

                        }
                        //下拉类型
                        if (resp.data.dstendport[i].data_other == true && this.checkPortSelect(resp.data.dstendport[i].value,resp.data.dstendport)) {
                            resp.data.dstendport[i].value = ""
                        }
                    }
                    //添加时  赋值默认数据
                    if(this.dstendport == "" && this.dstendportinput == ""){
                        if(otherValue == ""){
                            this.dstendportinput = "any"
                            this.dstendport = "any"
                        }else{
                            this.dstendportinput = otherValue
                            this.dstendport = otherValue
                        }

                    }
                    this.dstendports = resp.data.dstendport
                    this.dstendportValue = this.dstendportinput


                }


                this.id = id


                // this.dstmask = this.tableDataList[i].dstmask //
                this.descr = resp.data.descr
                if (id != "") {
                    this.onChangeShowState(3)
                } else {
                    this.onChangeShowState(2)
                }


            })


    }
    //判断输入是否是输入类型
    this.checkDstSrcType = function(value){
        let vmap = ["bogons", "bogonsv6", "virusprot", "sshlockout", "any", "(self)", "lan", "lanip",
            "lo0", "wan", "wanip"]
        if (vmap.indexOf(value) == -1) {
            //输入类型
            return true
        }
        return false
    }

    //判断 是否是下拉选择
    this.checkPortSelect = function(port, srcportlist){
        for (var i = 0; i < srcportlist.length; i++) { //目标下拉
            if (!srcportlist[i].data_other ) {
                if(port == srcportlist[i].value){
                    return true
                }

            }
        }
        return false
    }

})