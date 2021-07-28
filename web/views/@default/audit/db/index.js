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
    this.getTypeName = function (type) {
        switch (type) {
            case 0:
                return "mariadb"
            case 1:
                return "mysql"
            case 2:
                return "sqlServer"
            default:
                return "未知"
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
            height: "390px",
            width: "600px",
            callback: function () {
                this.onSearch()

            }
        })
    }

    this.onSearch = function () {
        console.log(this.searchSQLName)
        // window.location.href = '/audit/db?name=' + this.searchSQLName+"&ip="+this.searchSQLIp
        // this.$post("/audit/db")
        //     .params({
        //         Name: this.searchSQLName,
        //         Ip: this.searchSQLIp,
        //         ShowJson:true,
        //     })
        //     .refresh()

        this.$get(".").params({
            name: this.searchSQLName,
            ip: this.searchSQLIp,
            json: true,
        }).success(resp => {
            if (resp.code === 200) {
                if (resp.data.dbList)
                    this.dbList = resp.data.dbList
                else
                    this.dbList = []
                // this.level = resp.data.level
            }
        })
    }

    this.onOpenAuth = function (id) {
        teaweb.popup(Tea.url("/audit/db/auth",{id:id}), {
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
            type:item.type,
            version:item.version,
            ip:item.ip,
            port:item.port,
            system:item.system,
            status:item.status,
            edit:true,
        }), {
            height: "390px",
            width: "600px",
            callback: function () {
                this.onSearch()

            }
        })
    }

    this.onDelete = function (id) {
        teaweb.confirm("确定要删除该数据库？", function () {
            this.$post("/audit/db/delete").params({
                Opt: "delete",
                id: id,
            }).refresh()
        })
    }

    this.tableData = [
        {
            id: 1,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "3306",
            value4: "mysql",
            value5: "5.8",
            value6: "linux",
            value7: 1
        },
        {
            id: 2,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "3306",
            value4: "mysql",
            value5: "5.8",
            value6: "linux",
            value7: 0
        },
        {
            id: 3,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "3306",
            value4: "mysql",
            value5: "5.8",
            value6: "linux",
            value7: 1
        },
        {
            id: 4,
            value1: "robin_mysql",
            value2: "47.108.234.195",
            value3: "3306",
            value4: "mysql",
            value5: "5.8",
            value6: "linux",
            value7: 0
        },
    ]
})