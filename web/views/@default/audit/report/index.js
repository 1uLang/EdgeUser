Tea.context(function () {

    this.onAddHost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
			height: "400px",
            width:"520px",
			callback: function () {
				   window.location = "/audit/report"

			}
		})
    }


    this.getSendTime = function (cycle,cycleType,send_time) {
        switch (cycle) {
            case 1:
                return "每天 "+send_time
            case 2:
                return "每周 "+cycleType+" "+send_time
            case 3:
                return "每月 "+cycleType+" "+send_time
            default:
                return "未知"
        }
    }
    this.getFormat = function (format) {
        switch (format) {
            case 1:
                return "pdf"
            case 2:
                return "html"
            default:
                return "未知"
        }
    }

    this.getAssets = function (type) {
        switch (type) {
            case 1:
                return "数据库"
            case 2:
                return "主机"
            case 3:
                return "应用"
            default:
                return "全部"
        }
    }
    this.onEdit = function (id) {
        teaweb.popup(Tea.url(".createPopup",{id:id}), {
			height: "400px",
            width:"520px",
			callback: function () {
                window.location = "/audit/report"
			}
		})
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该订阅任务？", function () {
            this.$post("/audit/report/delete").params({
                Opt: "delete",
                id: id,
            }).refresh()
        })
    }

    this.tableData = [
        {id:1,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
        {id:2,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
        {id:3,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
        {id:4,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
    ]
})