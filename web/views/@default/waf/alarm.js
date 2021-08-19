Tea.context(function () {
    // this.selectNode=1


    this.$delay(function () {
        let curSelectNode = localStorage.getItem("nfwSelectNodeId");
        if(curSelectNode){
            this.selectNode = curSelectNode
        }
    })


    this.getTime= function (time) {
        var d = new Date(time);
        return d.toLocaleDateString()+" "+d.toLocaleTimeString()
    }

    this.getEditName= function (act) {
        if(act == "drop"){
            return "警报"
        }
        return "丢弃";
    }

    this.getItemInfo = function (id) { 
        for (var i=0;i<this.tableData.length;i++){
            if(this.tableData[i].id ==id){
                return this.tableData[i]
            }
        }
        return null
    }

    //启用
    this.onOpenCig = function (id) {
        teaweb.confirm("确定启用规则？", function () {
            let that = this
            that.$post(".set")
                .params({
                    id: id,
                    nodeId: this.selectNode,
                })
                .refresh()
        })
    }
    //禁用
    this.onCloseCig = function (id) {
        teaweb.confirm("确定禁用规则？", function () {
            let that = this
            that.$post(".set")
                .params({
                    id: id,
                    nodeId: this.selectNode,
                })
                .refresh()
        })
    }
    //更换alert/drop
    this.onChangeCig = function (id,act) {
        // var itemData = this.getItemInfo(id)
        teaweb.confirm("确定修改规则？", function () {
            let that = this
            that.$post(".editAction")
                .params({
                    id: id,
                    nodeId: this.selectNode,
                    act:this.getEditName(act),
                })
                .refresh()
        })
     }

    // this.hostData = [
    //     {id:1,hostAddress:"成都-ddos-192.168.1.1",},
    //     {id:2,hostAddress:"成都-ddos-192.168.1.2",},
    //     {id:3,hostAddress:"成都-ddos-192.168.1.3",},
    //     {id:4,hostAddress:"成都-ddos-192.168.1.4",},
    // ]

    //
    // this.tableData = [
    //     {id:1,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:1,editType:1},
    //     {id:2,value1:"900505001",value2:"alert",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:2,editType:2},
    //     {id:3,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:1,editType:1},
    //     {id:4,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:2,editType:1},
    //     {id:5,value1:"900505001",value2:"drop",value3:"abuse.ch.feodotracker.rules",value4:"trojan-activity",value5:"Feodo Tracker: potential Dridex CnC Traf",status:1,editType:1},
    // ]
    //获取当前选中的节点
    this.GetSelectNode = function (event) {
        this.selectNode = event.target.value; //获取option对应的value值
        localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/nfw/ips?nodeId=' + node

    }
})