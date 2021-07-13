Tea.context(function () {

    this.nShowState = 1
    this.nTimeSelect = 1
    this.dayFrom = ""
    this.dayTo = ""

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
        this.reloadFindTableChart()
        // this.reloadDetailTableChart()
    })

    this.onChangeShowState = function (state){
        if(this.nShowState!=state){
            this.nShowState = state
  
            if(this.nShowState ==1){
                this.$delay(function () {
                    this.reloadFindTableChart()
                })
            }else{
                this.$delay(function () {
                    this.reloadDetailTableChart()
                })
            }
            
        }
    }

    this.onChangeTimeSelect = function(index){
        if(this.nTimeSelect!=index){
            this.nTimeSelect = index
        }
    }

    this.reloadFindTableChart = function () {
		let chartBox = document.getElementById("find-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 40, 
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                // boundaryGap: false,
				data: this.findTableData.lineValue
			},
			yAxis: {
                // name: 'GB',
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
                    name:"漏洞总数",
					type: "line",
					data: this.findTableData.lineData,
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

    this.reloadDetailTableChart = function () {
		let chartBox = document.getElementById("detail-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 40, 
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                // boundaryGap: false,
				data: this.detailTableData.lineValue
			},
			yAxis: {
                // name: 'GB',
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

    this.data = {
        hostCount:5,
        hostOnline:1,
        detailAttCount:10,
        detailAttTodayCount:1,
        detailloopholeCount:120,
        detailloopholeTodayCount:12,
    }

    this.findTableData={
        lineValue:["05-08","05-10","05-12","05-14","05-16","05-18","05-20","05-22","05-24","05-26","05-28"],
        lineData:[20,21,1,21,31,25,15,12,13,16,9]
    }

    this.detailTableData={
        lineValue:["05-08","05-10","05-12","05-14","05-16","05-18","05-20","05-22","05-24","05-26","05-28"],
        lineData:[2,5,10,1,11,5,13,17,5,6,9]
    }



});
  