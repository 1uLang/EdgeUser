Tea.context(function () {

    this.dayFrom = ""
    this.dayTo = ""
    this.keyword = ""
    //
    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")

        this.reloadLineTableChart()
        this.reloadBarTableChart()

        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }

    })
    this.onOpenVirusPage = function (url) {

        //向服务器请求获取需要显示的数据
        window.location = "/hids/invade/" + url
    }
    this.reloadLineTableChart = function () {
        let chartBox = document.getElementById("line-chart-box")
        let chart = echarts.init(chartBox)
        let option = {
            title:{
                text: '当前入侵威胁分布情况',
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
                top: 30,   // 等价于 y: '16%'
                left: 15, 
                right: 15,
                bottom: 30,
            },
            legend: {
                orient : 'horizontal',
                x : 'center',
                y : 'bottom',
                data:this.names,
            },
			tooltip : {
                trigger: 'item',
                formatter: "{a} <br/>{b} : {c} ({d}%)"
            },
            color:['#2698fb','#26c4c3','#26c46f','#ffce51','#f95c74','#8f59dd','#474fc5','#3d4c7d'],
			series: [
                {
                    name:"占比情况",
                    type: 'pie',
                    radius: "55%",
                    center: ['50%', '50%'],
                    data: this.dashboard,
                    itemStyle:{ 
                        normal:{ 
                            label:{ 
                                show: true, 
                                formatter: '{b} : {c} ({d}%)' 
                            }, 
                            labelLine :{show:true} 
                        } 
                    } 
                }
            ],
			animation: false
		}
        chart.setOption(option)
        chart.resize()
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
                right: 30,
                bottom: 30,
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
                    data: this.dashboard,
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
  