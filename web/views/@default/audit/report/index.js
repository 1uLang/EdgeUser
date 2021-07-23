Tea.context(function () {

    this.onAddHost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
			height: "400px",
            width:"520px",
			callback: function () {
				
			}
		})
    }


    this.onEdit = function (id) {
        teaweb.popup(Tea.url(".createPopup"), {
			height: "400px",
            width:"520px",
			callback: function () {
				
			}
		})
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该订阅任务？", function () {

        })
    }

    this.tableData = [
        {id:1,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
        {id:2,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
        {id:3,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
        {id:4,value1:"每日安全审计报表获取",value2:"每天 1:00",value3:"PDF",value4:"数据库",value5:"luobing_mysql",value6:"449588335@qq.com,luobing@zhiannet.com"},
    ]
})