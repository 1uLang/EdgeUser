Tea.context(function () {
    this.localProgressData = localStorage.getItem("examinProgressData")
    this.progressListData =  this.localProgressData ? JSON.parse(this.localProgressData) : [];//{id:1,curPer:1,disabled:1}

    this.Items = this.examineItems !== ""? this.examineItems.split(","): []
    this.curIndex = -1

    this.bShowCheckDetail = false
    this.pCheckDetailData = null
    this.sSelectCheckValue = ["01","02","03","04","13","14","15"]

    this.webSearchKey = ""  //网页后门
    this.searchPath = ""    //病毒木马
    this.MacCode = ""
    this.serverIp = ""
    this.bTimeOutTip = false
    this.bShowScanPath = false

    let that = this

    this.sTopSelectItem = [
        {id: "01", value: "系统漏洞"},
        {id: "02", value: "弱口令"},
        {id: "03", value: "风险账号"},
        {id: "04", value: "配置缺陷"},
        {id: "11", value: "病毒木马"},
        {id: "12", value: "网页后门"}
    ]
    this.sBottomSelectItem = [
        {id: "13", value: "反弹shell"},
        {id: "14", value: "异常账号"},
        {id: "15", value: "系统命令篡改"},
        {id: "16", value: "异常进程"},
        {id: "17", value: "日志异常删除"},
    ]
 
    this.$delay(function () {

        
        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }

        // that.onCreateLoopTimeOut()
        window.addEventListener('beforeunload', function () {
            that.onReleaseUpdateTimeOut()
            // that.onReleaseTimeOut()
        })
    })
    

    this.onCallBack = function () {
        if (this.checkScans()) {
            this.$post(".").success(resp => {
                if (resp.code === 200) {
                    this.datas = resp.data.datas
                    this.state = resp.data.state
                    this.Type = resp.data.Type
                    this.score = resp.data.score
                    this.examineItems = resp.data.examineItems
                    this.startTime = resp.data.startTime
                    this.endTime = resp.data.endTime
                }
            })
        }
    }
    this.onCreateLoopTimeOut = function () {
        this.onReleaseTimeOut()
        this.checkTimer = createTimer(this.onCallBack, {timeout: 60000});
        this.checkTimer.start();
    }
    this.onReleaseTimeOut = function () {
        if (this.checkTimer) {
            this.checkTimer.stop()
            this.checkTimer = null
        }
    }
    this.checkScans = function () {
        for (item of this.datas) {
            if (item.state === 1 ||
                item.state === 4) {
                return true
            }
        }
        return false
    }
    //判断是否有主机正在体检...
    this.checkExamine = function (){
        for(item of this.datas){
            console.log(item)
        }
        return false
    }
    this.onChangeCheckState = function (state) {
        if (this.state != state) {
            this.state = state
        }
        this.refreshPage()
    }

    this.onChangeHealthNumState = function (state) {
        if (this.score != state) {
            this.score = state
        }
        this.refreshPage()
    }

    this.onChangeResultState = function (state) {
        if (this.Type != state) {
            this.Type = state
        }
        this.refreshPage()
    }
    this.refreshPage = function () {
        let url = "/hids/examine?state=" + this.state + "&score=" + this.score + "&Type=" + this.Type
        if (this.Items.length > 0) {
            url += "&examineItems=" + this.Items.toString()
        }
        window.location.href= url

    }

    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }


    //检测是否包含元素
    this.checkSelectValue = function (index, selectValue) {
        if (selectValue) {
            for (var i = 0; i < selectValue.length; i++) {
                if (selectValue[i] == index) {
                    return true;
                }
            }
        }

        return false;
    }
    //添加/删除元素
    this.onAddSelectValue = function (index) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(index, this.Items);
        }
        if (bValue) {
            this.Items = this.Items.filter((itemIndex) => {
                return itemIndex != index;
            });
        } else {
            this.Items.push(index);
        }

        this.refreshPage()
    }


    this.getShowSelectImage = function (id) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(id, this.Items);
        }
        if (bValue) {
            return "/images/select_select.png";
        }
        return "/images/select_box.png";
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.getStatusName = function (status) {
        switch (status) {
            case 0:
                return "未体检"
            case 1:
                return "体检中"
            case 2:
                return "已完成"
            case 4:
                return "取消中"
            case 6:
                return "已取消"
            default:
                return "未知"
        }
    }

    this.getCheckNumName = function (num) {
        return (num && num > 0) ? num + "分" : "未得出"
    }

    this.getHealthName = function (score) {
        if (score =>0 && score <= 59){
            return '不健康'
        }else if (score >=60 && score <=89){
            return '亚健康'
        }else{
            return '健康'
        }
    }

    this.enters = function (index) {
        // this.curIndex = index;
    }

    this.leaver = function (index) {
        this.curIndex = -1;
    }

    this.onOpenDetail = function (item) {
        window.location.href = "/hids/examine/detail?macCode="+item.macCode
    }

    this.onOpenCheck = function (item) {
        this.MacCode = item.macCode
        this.serverIp = item.serverIp
        this.sSelectCheckValue = ["01","02","03","04","13","14","15"]
        this.pCheckDetailData = [
            {
                checkName: "漏洞风险检查项：",
                checkValue: [
                    {id: "01", value: "系统漏洞"},
                    {id: "02", value: "弱口令"},
                    {id: "03", value: "风险账号"},
                    {id: "04", value: "配置缺陷"},
                ]
            },
            {
                checkName: "入侵威胁检查项：",
                checkValue: this.sBottomSelectItem
            }
        ]
        if (this.pCheckDetailData) {
            this.bShowCheckDetail = true
        } else {
            this.bShowCheckDetail = false
        }
    }

    this.onSelectCheckValue = function (index) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(index, this.sSelectCheckValue);
        }
        if (bValue) {
            this.sSelectCheckValue = this.sSelectCheckValue.filter((itemIndex) => {
                return itemIndex != index;
            });
        } else {
            this.sSelectCheckValue.push(index);
        }
    }

    this.getShowSelectValueImage = function (id) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(id, this.sSelectCheckValue);
        }
        if (bValue) {
            return "/images/select_select.png";
        }
        return "/images/select_box.png";
    }

    this.onCheckSelectItem = function (id) {
        let bValue = false;
        if (this.checkSelectValue) {
            bValue = this.checkSelectValue(id, this.sSelectCheckValue);
        }
        return bValue
    }

    this.onCloseCheck = function () {
        this.sSelectCheckValue = []
        this.bShowCheckDetail = false
    }

    this.onStartCheck = function (item) {
        if(this.sSelectCheckValue.length == 0){
            teaweb.warn("请选择体检项目")
            return
        }
        teaweb.confirm("确定立即体检吗？", function () {
            this.$post(".scans").params({
                Opt:'now',
                serverIp:this.serverIp,
                VirusPath:this.searchPath,
                WebShellPath:this.webSearchKey,
                MacCode: [this.MacCode],
                ScanItems: this.sSelectCheckValue.join(","),
            }).success(function (){
                teaweb.closePopup()
                window.location.reload()
            }).error(function (){
                teaweb.warn("失败：该主机agent已暂停服务，命令无法执行！")
            })
        })

        this.bShowCheckDetail = false

        //
    }
    this.onStopCheck = function (item) {
        teaweb.confirm("确定取消体检吗？", function () {
            this.$post(".scans").params({
                Opt:'cancel',
                serverIp:item.os.serverIp,
                MacCode: [item.macCode],
            }).success(function (){
                window.location.reload()
            }).error(function (){
                teaweb.warn("失败：该主机agent已暂停服务，命令无法执行", function () {})
            })
        })
    }
    //检测是否显示扫描路径的输入框和提示框
    this.onCheckSelectValue = function () {

        var selextBox = document.getElementsByName("customScan")
        if (selextBox) {
            for (var item of selextBox) {
                if (item.checked) {
                    if (item.value == 2) {
                        this.bTimeOutTip = true
                        this.bShowScanPath = false
                    } else if (item.value == 3) {
                        this.bTimeOutTip = false
                        this.bShowScanPath = true
                    } else {
                        this.bTimeOutTip = false
                        this.bShowScanPath = false
                    }
                }
            }
        }
    }


    this.checkShowColor = function (curValue, maxValue) {
        if(curValue && maxValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            return tempValue >= 100
        }
        return false
    }

    this.getProgressItemInfo = function (id) {
        if(id){
            for(var index=0;index<that.progressListData.length;index++){
                if(that.progressListData[index].id == id){
                    return that.progressListData[index]
                }
            }
        }
        return null
    }
    
    this.getProgressPerStr = function (curValue, maxValue,id,state) {
        if(!that.getProgressItemInfo){return "0%"}

        if(curValue == 0 ){
            if(state==1){
                that.onCreateUpdateTimeOut()
                let curData = that.getProgressItemInfo(id)
                if(curData){
                    if(curData.curPer == 0){
                        return "1%"
                    }
                    return curData.curPer+"%"
                }
                that.onCreateProgressItemInfo(id)
                return "1%"
            }else{
                that.onChangeProgressDataState(id,state)
                return ""
            }
        }else if(curValue == 100){
            that.onChangeProgressDataState(id,state)
        }

        if(curValue && maxValue && maxValue>0 && maxValue >= curValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            if(tempValue>=100){
                return "已完成"
            }else if(tempValue<1 && state && state==1){
                return "1%"
            }

            return tempValue + "%"
        }
        return "0%"
    }

    // let that = this
    // this.testFunc = function(curValue, maxValue,id,state){
    //     console.log(that.getProgressPerStr)
    //     console.log(that.getProgressPerStr(curValue, maxValue,id,state))
    //     return "0%"
    // }
    // this.getInitProgressPerString = function(curValue, maxValue,id,state){
    //     let that = this
    //     this.testFunc = function(){
    //         console.log(that.getProgressPerStr)
    //         console.log(that.getProgressPerStr(curValue, maxValue,id,state))
    //     }
    //     this.testFunc()
    // }

    this.getProgressPer = function (curValue, maxValue,id,state) {
        if(!that.getProgressItemInfo){return "0%"}
        
        if(curValue == 0 ){
            if(state && state==1){
                let curData = that.getProgressItemInfo(id)
                if(curData){
                    return curData.curPer+"%"
                }
                that.onCreateProgressItemInfo(id)
                return "1%"
            }
        }

        if(curValue && maxValue && maxValue>0 && maxValue >= curValue){
            var tempValue = ((curValue / maxValue) * 100).toFixed(1)
            if(tempValue<1 && state && state==1 ){
                return "1%"
            }
            return tempValue + "%"
        }
        return "0%"
    }
    
    //选择时间之后的回调
    this.onTimeChange = function () {
        let startTime = document.getElementById("day-from-picker").value
        if(startTime.length > 0){
            let tempCharCount=startTime.split(":").length-1;
            if(tempCharCount<=1){
                document.getElementById("day-from-picker").value = ""
            }else{
                startTime = startTime.replace("T", " ");
            }
        }else{
            document.getElementById("day-from-picker").value = ""
        }

        let endTime = document.getElementById("day-to-picker").value
        if(endTime.length > 0){
            let tempCharCount=endTime.split(":").length-1;
            if(tempCharCount<=1){
                document.getElementById("day-to-picker").value = ""
            }else{
                endTime = endTime.replace("T", " ");
            }
        }else{
            document.getElementById("day-to-picker").value = ""
        }
        //todo req
    }


    this.updateTimeId = null

    this.onCreateProgressItemInfo = function (id) {
        let curData = {id:id,curPer:0,state:1,disabled:0}
        that.progressListData.push(curData)
        that.onSaveProgressData()
    }
    this.onChangeProgressDataState = function (id,state) {

        for(var index=0;index<that.progressListData.length;index++){
            if(that.progressListData[index].id==id){
                that.progressListData[index].state = state
                break
            }
        }
        if(that.progressListData.length>0){
            that.progressListData = that.progressListData.filter((item) => {
                return item.state == 1;
            });
        }
        that.onSaveProgressData()
    }
    // 进度的缓存数据
    this.onUpdateProgressData = function () {
        for(var index=0;index<that.progressListData.length;index++){
            if(that.progressListData[index].state==1){
                that.progressListData[index].curPer = that.progressListData[index].curPer+5
                if(that.progressListData[index].curPer>=95){
                    that.progressListData[index].curPer = 95
                }
            }
        }
        that.onSaveProgressData()
    }
    this.onSaveProgressData = function () {
        localStorage.setItem("examinProgressData", JSON.stringify(that.progressListData));
    }
    

    //计时器
    this.onCreateUpdateTimeOut = function () {
        if(!that.updateTimeId){
            that.updateTimeId = createTimer(that.onUpdateProgressData, {timeout: 5000});
            that.updateTimeId.start();
        }
    }

    this.onReleaseUpdateTimeOut = function () {

        if (that.updateTimeId) {
            that.updateTimeId.stop()
            that.updateTimeId = null
        }
    }

})