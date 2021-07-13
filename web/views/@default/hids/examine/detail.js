Tea.context(function () {

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
        this.reloadBarTableChart()
        this.reloadCircularTableChart()

        // let that = this
        // window.addEventListener("resize", function () {
        //     that.resizeBarTableChart()
        //     that.resizeCircularTableChart()
        // })
    })


    this.resizeBarTableChart = function () {
        let chartBox = document.getElementById("bar-chart-box")
        let chart = echarts.init(chartBox)
        chart.resize()
    }

    this.resizeCircularTableChart = function () {
        let chartBox = document.getElementById("circular-chart-box")
        let chart = echarts.init(chartBox)
        chart.resize()
    }

    this.onGoBack = function () {
        window.location.href = "/hids/examine"
    }

    this.getHealthName = function (score) {
        if (score < 60){
            return '不健康'
        }else if (score < 90){
            return '亚健康'
        }else{
            return '健康'
        }
    }
    this.parseServerLocalIp = function (ip) {
        let ips = ip.split(";")
        return ips.slice(-1)[0]
    }

    this.reloadBarTableChart = function () {
        let chartBox = document.getElementById("bar-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 15, 
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
				data: this.tableData.titleValue,
                axisLabel: {
                    // rotate: -30, // 旋转角度
                    interval: 0  //设置X轴数据间隔几个显示一个，为0表示都显示
                },
			},
			yAxis: {
                name: '体检概况',
                min:0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber:4,  // 设置y轴刻度间隔个数
                // axisLine: {
                //     lineStyle: {
                //         // 设置y轴颜色
                //         color: '#fff'
                //     }
                // },
            },
			tooltip: {
				trigger: "axis",
			},
			series: [
				{
					type: "bar",
					data: this.tableData.itemValue,
					barWidth:"120px",
                    color: "#2698fb"
				},
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
    }

    this.reloadCircularTableChart = function () {
        let chartBox = document.getElementById("circular-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 15, 
                right: 60,
                bottom: 30,
                containLabel: true
            },
            legend: {
                top: '5%',
                left: 'center'
            },
			tooltip: {
				trigger: "item",
			},
            color:['#FD0112','#FF8F40','#00A2EC','#78deff'],
			series: [
                {
                    type: 'pie',
                    radius: ['40%', '70%'],
                    avoidLabelOverlap: false,
                    label: {
                        show: false,
                        position: 'center'
                    },
                    emphasis: {
                        label: {
                            show: true,
                            fontSize: '40',
                            fontWeight: 'bold'
                        }
                    },
                    labelLine: {
                        show: false
                    },
                    data: this.tableData.circularValue
                }
            ],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
    }

    this.tableData = {
        titleValue:["系统漏洞","弱口令","危险账户","配置缺陷"],
        itemValue:[ this.details.risk,this.details.weak,this.details.danger_account,this.details.config_defect],
        circularValue:[
            {value: this.details.risk, name: '系统漏洞'},
            {value: this.details.weak, name: '弱口令'},
            {value: this.details.danger_account, name: '危险账户'},
            {value: this.details.config_defect, name: '配置缺陷'},
        ]
    }
})