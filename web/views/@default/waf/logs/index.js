Tea.context(function () {
	this.$delay(function () {
		let that = this
		teaweb.datepicker("day-input", function (day) {
			that.day = day
		})
        this.reloadDetailTableChart()
	})

    let that = this
    this.accessLogs.forEach(function (accessLog) {
        if (typeof (that.regions[accessLog.remoteAddr]) == "string") {
            accessLog.region = that.regions[accessLog.remoteAddr]
        } else {
            accessLog.region = ""
        }
    })

    this.reloadDetailTableChart = function () {
		console.log(this.detailTableData)
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
                    name:"日志数",
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

	this.report = function (n){
		window.location.href = '/waf/logs?report=' + n

	}
})