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
    this.getSystemName = function (system) {
        switch (system) {
            case 0:
                return "windows"
            case 1:
                return "linux"
            default:
                return "未知"
        }
    }
    this.onAddHost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
            height: "300px",
            width: "460px",
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
                if (resp.data.hostList)
                    this.hostList = resp.data.hostList
                else
                    this.hostList = []
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
        teaweb.popup(Tea.url(".createPopup", {
            id: item.id,
            name: item.name,
            ip: item.ip,
            system: item.system,
            status: item.status,
            edit: true,
        }), {
            height: "300px",
            width: "460px",
            callback: function () {
                this.onSearch()
            }
        })
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该主机？", function () {
            this.$post("/audit/host/delete").params({
                Opt: "delete",
                id: id,
            }).refresh()
        })
    }

    this.tableData = [
        {id: 1, value1: "robin_mysql", value2: "47.108.234.195", value3: "windows", value4: 1},
        {id: 2, value1: "robin_mysql", value2: "47.108.234.195", value3: "linux", value4: 0},
        {id: 3, value1: "robin_mysql", value2: "47.108.234.195", value3: "windows", value4: 1},
        {id: 4, value1: "robin_mysql", value2: "47.108.234.195", value3: "linux", value4: 0},
    ]
})