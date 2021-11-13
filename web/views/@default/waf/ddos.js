Tea.context(function () {

    this.address = ''
    this.nShowState = 1
    this.endTime = ""
    this.startTime = ""
    this.attackType = "0"
    this.status = "0"

    this.$delay(function () {

        this.reloadDetailTableChart()

        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")

        let curSelectNode = localStorage.getItem("ddosSelectNodeId");
        if (curSelectNode) {
            this.nodeId = curSelectNode
        }

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
        // this.getTraffic(this.nShowState)
    })
    this.getAttacks = function (state) {

        this.$get(".attacks").params({NodeId: this.nodeId, level: this.level}).success(resp => {
            if (resp.code === 200) {
                if (resp.data.attacks)
                    this.attacks = resp.data.attacks
                else
                    this.attacks = []
                this.nodeId = resp.data.nodeId
                this.nShowState = state
            }
        })
    }
    this.showHost = function () { //重新加载该页面
        let node = this.nodeId
        localStorage.setItem("ddosSelectNodeId", node);
        window.location.href = '/waf/ddos?nodeId=' + node
    }

    this.onChangeShowState = function (state) {
        this.level = 1
        if (this.nShowState != state) {
            if (state === 2) {
                this.getAttacks(state)
            } else {
                this.$get(".link").params({NodeId: this.nodeId, level: this.level}).success(resp => {
                    if (resp.code === 200) {
                        if (resp.data.links)
                            this.links = resp.data.links
                        else
                            this.links = []
                        this.level = resp.data.level
                        this.nShowState = state
                    }
                })
            }

        }
    }
    this.search = function () {
        let start = document.getElementById("day-from-picker").value
        let end = document.getElementById("day-to-picker").value
        window.location.href = '/waf/ddos?nodeId=' +  this.nodeId +"&startTime="+start+"&endTime="+end+"&attackType="+this.attackType+"&status="+this.status+"&address="+this.address
    }
    this.toShowStatus = function (st) {
        if (st === "2") {
            return "结束"
        } else {
            return "保护中"
        }
    };

    this.onDownLoad = function (id) {

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
    this.report = function (n) { //重新加载该页面
        let node = this.nodeId
        localStorage.setItem("ddosSelectNodeId", node);
        window.location.href = '/waf/ddos?nodeId=' + node+"&report="+n
    }
})
