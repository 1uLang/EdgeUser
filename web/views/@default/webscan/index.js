Tea.context(function () {

	this.$delay(function () {
		if(this.data.nodeErr != ""){
			return
		}
		// this.reloadHighVulnerabilitiesChart()
		// this.reloadMedVulnerabilitiesChart()
		// this.reloadLowVulnerabilitiesChart()

		let that = this
		window.addEventListener("resize", function () {
			that.resizeHighVulnerabilitiesChart()
			that.resizeMedVulnerabilitiesChart()
			that.resizeLowVulnerabilitiesChart()
		})
	})

	this.resizeHighVulnerabilitiesChart = function () {
		let chartBox = document.getElementById("high-vulnerabilities-chart-box")
		let chart = echarts.init(chartBox)
		chart.resize()
	}

	this.reloadHighVulnerabilitiesChart = function () {
		let chartBox = document.getElementById("high-vulnerabilities-chart-box")
		let chart = echarts.init(chartBox)
		let option ={
			tooltip: {
				trigger: 'item',
				formatter: '{a} <br/>{b} : {c} ({d}%)'
			},
			legend: {
				left: 'center',
				bottom: '10',
				data: ['高危漏洞']
			},
			series: [
				{
					name: '统计占比',
					type: 'pie',
					roseType: 'radius',
					radius: [15, 95],
					center: ['50%', '48%'],
					data: [
						{ value: this.data.vuln_count.high, name: '高危漏洞',itemStyle:{normal:{color:'#ec808d'}} },
					],
					animationEasing: 'cubicInOut',
					animationDuration: 2600
				}
			]
		}
		chart.setOption(option)
		chart.resize()
	}

	this.resizeMedVulnerabilitiesChart = function () {
		let chartBox = document.getElementById("med-vulnerabilities-chart-box")
		let chart = echarts.init(chartBox)
		chart.resize()
	}

	this.reloadMedVulnerabilitiesChart = function () {
		let chartBox = document.getElementById("med-vulnerabilities-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
			tooltip: {
				trigger: 'item',
				formatter: '{a} <br/>{b} : {c} ({d}%)'
			},
			legend: {
				left: 'center',
				bottom: '10',
				data: ['中危漏洞']
			},
			series: [
				{
					name: '统计占比',
					type: 'pie',
					roseType: 'radius',
					radius: [15, 95],
					center: ['50%', '48%'],
					data: [
						{ value: this.data.vuln_count.med, name: '中危漏洞',itemStyle:{normal:{color:'#fcc77d'}} }
					],
					animationEasing: 'cubicInOut',
					animationDuration: 2600
				}
			]
		}
		chart.setOption(option)
		chart.resize()
	}

	this.resizeLowVulnerabilitiesChart = function () {
		let chartBox = document.getElementById("low-vulnerabilities-chart-box")
		let chart = echarts.init(chartBox)
		chart.resize()
	}

	this.reloadLowVulnerabilitiesChart = function () {
		let chartBox = document.getElementById("low-vulnerabilities-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
			tooltip: {
				trigger: 'item',
				formatter: '{a} <br/>{b} : {c} ({d}%)'
			},
			legend: {
				left: 'center',
				bottom: '10',
				data: ['低危漏洞']
			},
			series: [
				{
					name: '统计占比',
					type: 'pie',
					roseType: 'radius',
					radius: [15, 95],
					center: ['50%', '48%'],
					data: [
						{ value: this.data.vuln_count.low, name: '低危漏洞',itemStyle:{normal:{color:'#3abee8'}} }
					],
					animationEasing: 'cubicInOut',
					animationDuration: 2600
				}
			]
		}
		chart.setOption(option)
		chart.resize()
	}
})
