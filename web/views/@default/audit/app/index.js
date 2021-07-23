Tea.context(function () {
    
    this.searchSQLName = ""
    this.searchSQLIp = ""

    this.getStatus = function (status) {
        switch (status) {
            case 1:
                return "已启用"
            case 0:
                return "已停用"
            default:
                return "已停用"
        }
    }

    this.onAddHost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
			height: "300px",
            width:"460px",
			callback: function () {
				
			}
		})
    }

    this.onSearch = function () {
        
    }

    this.onOpenAuth = function (id) {
        teaweb.popup(Tea.url(".auth"), {
			height: "270px",
            width: "460px",
			callback: function () {
				
			}
		})
    }

    this.onEdit = function (id) {
        teaweb.popup(Tea.url(".createPopup"), {
			height: "300px",
            width:"460px",
			callback: function () {
				
			}
		})
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该应用？", function () {

        })
    }

    this.tableData = [
        {id:1,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:1},
        {id:2,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:0},
        {id:3,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:1},
        {id:4,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:0},
    ]
})