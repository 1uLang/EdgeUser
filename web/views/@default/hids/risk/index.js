Tea.context(function () {
    this.seriousPer = "0%"
    this.heightPer = "0%"
    this.middlePer = "0%"
    this.lowPer = "0%"

    this.$delay(function () {
        this.reloadBarTableChart()
        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })

    this.onOpenDetail = function(type){
        switch(type){
            case 1:
                window.location.href = "/hids/risk/systemRisk"
                break
            case 2:
                window.location.href = "/hids/risk/weak"
                break
            case 3:
                window.location.href = "/hids/risk/dangerAccount"
                break
            case 4:
                window.location.href = "/hids/risk/configDefect"
                break
            default:
                break
        }
    }

    this.reloadBarTableChart = function () {
        let chartBox = document.getElementById("bar-chart-box")
        let chart = echarts.init(chartBox)
        let option = {
            title:{
                text: '当前漏洞风险分布情况',
                x:'left',
                y: 'top',
                textStyle: { 
                    fontSize: 16,
                    color: '#333',
                    fontWeight:"normal"
                },
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 60,   // 等价于 y: '16%'
                left: 30,
                right: 30,
                bottom: 15,
                containLabel: true
            },
            xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                data: this.names,
                axisLabel: {
                    rotate: 0, // 旋转角度
                    interval: 0  //设置X轴数据间隔几个显示一个，为0表示都显示
                },
            },
            yAxis: {
                // name: 'GB',
                min: 0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber: 5,  // 设置y轴刻度间隔个数
                // axisLine: {
                //     lineStyle: {
                //         // 设置y轴颜色
                //         color: '#fff'
                //     }
                // },
            },
            tooltip: {
                trigger: "item",
            },
            series: [
                {
                    type: "bar",
                    data: this.datas,
                    barWidth: "70px",
                    color: "#2698fb"
                },

            ],
            animation: false
        }
        chart.setOption(option)
        chart.resize()
    }
});
  