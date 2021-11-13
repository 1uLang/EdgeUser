Tea.context(function () {
    // this.selectNode=1


    this.$delay(function () {
        this.reloadDetailTableChart()
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
        window.location.href = '/waf/alarm?nodeId=' + node

    }

    this.reloadDetailTableChart = function () {
		let chartBox = document.getElementById("detail-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            grid: {
                top: 30,  
                left: 40, 
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
				data: this.detailTableData.lineValue
			},
			yAxis: {
                // name: 'GB',
                min:0, // 设置y轴刻度的最小值
                splitNumber:4,  // 设置y轴刻度间隔个数
            },
			tooltip: {
				trigger: "axis",
			},
			series: [
				{
                    name:"漏洞总数",
					type: "line",
					data: this.detailTableData.lineData,
					itemStyle: {
						color: "#0085fa"
					},
					lineStyle: {
						color: "#0085fa"
					}
				},
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

    // this.detailTableData={
    //     lineValue:["05-08","05-10","05-12","05-14","05-16","05-18","05-20","05-22","05-24","05-26","05-28"],
    //     lineData:[2,5,10,1,11,5,13,17,5,6,9]
    // }

    this.report = function (n) {
        localStorage.setItem("nfwSelectNodeId", this.selectNode);
        let node = this.selectNode
        window.location.href = '/waf/alarm?nodeId=' + node+"&report="+n

    }
})