Tea.context(function () {
	this.splitFormat = function (format) {
		let result = format.match(/^([0-9.]+)([a-zA-Z]+)$/)
		return [result[1], result[2]]
	}

	{
		let pieces = this.splitFormat(this.dashboard.monthlyTrafficBytes)
		this.dashboard.monthlyTrafficBytes = pieces[0]
		this.dashboard.monthlyTrafficBytesUnit = pieces[1]
	}

	{
		let pieces = this.splitFormat(this.dashboard.monthlyPeekTrafficBytes)
		this.dashboard.monthlyPeekTrafficBytes = pieces[0]
		this.dashboard.monthlyPeekTrafficBytesUnit = pieces[1].toLowerCase()
	}

	{
		let pieces = this.splitFormat(this.dashboard.dailyTrafficBytes)
		this.dashboard.dailyTrafficBytes = pieces[0]
		this.dashboard.dailyTrafficBytesUnit = pieces[1]
	}

	{
		let pieces = this.splitFormat(this.dashboard.dailyPeekTrafficBytes)
		this.dashboard.dailyPeekTrafficBytes = pieces[0]
		this.dashboard.dailyPeekTrafficBytesUnit = pieces[1].toLowerCase()
	}

	this.$delay(function () {
		this.reloadDailyTrafficChart()
		this.reloadDailyPeekTrafficChart()
	})

	this.reloadDailyTrafficChart = function () {
		let chartBox = document.getElementById("daily-traffic-chart-box")
		let chart = echarts.init(chartBox)

		let option = {
			xAxis: {
				data: this.dailyTrafficStats.map(function (v) {
					return v.day;
				})
			},
			yAxis: {
			},
			tooltip: {
				show: true,
				trigger: "item"
			},
			grid: {
				left: 40,
				top: 10,
				right: 20
			},
			series: [
				{
					name: "流量",
					type: "line",
					data: this.dailyTrafficStats.map(function (v) {
						return v.count;
					}),
					itemStyle: {
						color: "#9DD3E8"
					},
					lineStyle: {
						color: "#9DD3E8"
					},
					areaStyle: {
						color: "#9DD3E8"
					}
				}
			],
			animation: false
		}
		chart.setOption(option)
	}

	this.reloadDailyPeekTrafficChart = function () {
		let chartBox = document.getElementById("daily-peek-traffic-chart-box")
		let chart = echarts.init(chartBox)

		let option = {
			xAxis: {
				data: this.dailyPeekTrafficStats.map(function (v) {
					return v.day;
				})
			},
			yAxis: {
			},
			tooltip: {
				show: true,
				trigger: "item"
			},
			grid: {
				left: 40,
				top: 10,
				right: 20
			},
			series: [
				{
					name: "峰值带宽",
					type: "line",
					data: this.dailyPeekTrafficStats.map(function (v) {
						return v.count;
					}),
					itemStyle: {
						color: "#9DD3E8"
					},
					lineStyle: {
						color: "#9DD3E8"
					},
					areaStyle: {
						color: "#9DD3E8"
					}
				}
			],
			animation: false
		}
		chart.setOption(option)
	}
})