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

    this.getAppName = function (type) {
        switch (type) {
            case 1:
                return "iis"
            case 0:
                return "nginx"
            default:
                return "未知"
        }
    }

    this.onAddHost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
			height: "300px",
            width:"460px",
			callback: function () {
                this.onSearch()

            }
		})
    }

    this.onSearch = function () {
        this.$get(".").params({
            name: this.searchSQLName,
            ip: this.searchSQLIp,
            json: true,
        }).success(resp => {
            if (resp.code === 200) {
                if (resp.data.appList)
                    this.appList = resp.data.appList
                else
                    this.appList = []
                // this.level = resp.data.level
            }
        })
    }

    this.onOpenAuth = function (id) {
        teaweb.popup(Tea.url(".auth"), {
			height: "270px",
            width: "460px",
			callback: function () {
                this.onSearch()
			}
		})
    }

    this.onEdit = function (item) {
        teaweb.popup(Tea.url(".createPopup",{
            id:item.id,
            name:item.name,
            type:item.app_type,
            ip:item.ip,
            status:item.status,
            edit:true,
        }), {
			height: "300px",
            width:"460px",
			callback: function () {
                this.onSearch()
			}
		})
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该应用？", function () {
            this.$post(".delete").params({
                Opt: "delete",
                id: id,
            }).refresh()
        })
    }

    this.tableData = [
        {id:1,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:1},
        {id:2,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:0},
        {id:3,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:1},
        {id:4,value1:"robin_mysql",value2:"47.108.234.195",value3:"nginx",value4:0},
    ]
})